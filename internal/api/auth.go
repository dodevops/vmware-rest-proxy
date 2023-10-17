package api

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"vmware-rest-proxy/internal"
)

type GetSessionResponse struct {
	User string `json:"user"`
}

var sessionCache map[string]string

// GetSession returns the vmware session id to be used by other requests
func GetSession(c internal.Config, username string, password string) (string, error) {
	if sessionCache == nil {
		sessionCache = make(map[string]string)
	}
	if s, ok := sessionCache[username]; ok {
		logrus.Debugf("Checking cached session for user %s", username)
		var getSessionResponse GetSessionResponse
		if r, err := c.Resty.
			R().
			SetHeader("vmware-api-session-id", s).
			SetBasicAuth(username, password).
			SetResult(&getSessionResponse).
			Get("/api/session"); err != nil {
			return "", err
		} else {
			if r.StatusCode() == 401 {
				delete(sessionCache, username)
			} else if r.IsError() {
				return "", fmt.Errorf("error checking session for user %s (%s): %s", username, r.Status(), r.Body())
			} else {
				logrus.Debugf("Cached session still valid for user %s", username)
				return s, nil
			}
		}
	}
	logrus.Debugf("Creating VMware session for user %s at %s", username, c.Resty.BaseURL)
	var sessionToken string
	if r, err := c.Resty.
		R().
		SetBasicAuth(username, password).
		SetResult(&sessionToken).
		Post("/api/session"); err != nil {
		return "", err
	} else {
		if r.IsError() {
			return "", fmt.Errorf("error getting session (%s): %s", r.Status(), r.Body())
		}
		sessionCache[username] = sessionToken
		return sessionToken, nil
	}
}
