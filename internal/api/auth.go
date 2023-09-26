package api

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

// SessionResponse holds the value returned by the VMware session controller
type SessionResponse struct {
	Value string `json:"value"`
}

// GetSession returns the vmware session id to be used by other requests
func GetSession(url string, username string, password string) (string, error) {
	logrus.Debugf("Creating VMware session for user %s at %s", username, url)
	var sessionRespone SessionResponse
	if r, err := resty.
		New().
		SetBasicAuth(username, password).
		SetBaseURL(url).
		R().
		SetResult(&sessionRespone).
		Post("/rest/com/vmware/cis/session"); err != nil {
		return "", err
	} else {
		if r.IsError() {
			return "", fmt.Errorf("error getting session (%s): %s", r.Status(), r.Body())
		}
		return sessionRespone.Value, nil
	}
}
