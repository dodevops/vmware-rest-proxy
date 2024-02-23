package api

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"time"
)

// SessionLifetime sets the maximum lifetime of a session as the vCenter API for handling sessions isn't quite
// up for that
const SessionLifetime = 5 * time.Minute

type session struct {
	token   string
	created time.Time
}

var sessionCache map[string]session

// GetSession returns the vmware session id to be used by other requests
func (d DefaultVSphereProxyApi) GetSession(username string, password string) (string, error) {
	if sessionCache == nil {
		sessionCache = make(map[string]session)
	}
	if s, ok := sessionCache[username]; ok {
		logrus.Debugf("Checking cached session for user %s", username)

		if time.Now().Sub(sessionCache[username].created) <= SessionLifetime {
			return sessionCache[username].token, nil
		}

		logrus.Debugf("Session of user %s has been expired. Recreating it", username)

		delete(sessionCache, username)

		if r, err := d.Resty.
			R().
			SetHeader("vmware-api-session-id", s.token).
			SetBasicAuth(username, password).
			Delete("/api/session"); err != nil {
			return "", err
		} else {
			if r.IsError() {
				return "", fmt.Errorf("error deleting session for user %s (%s): %s", username, r.Status(), r.Body())
			}
		}
	}
	logrus.Debugf("Creating VMware session for user %s at %s", username, d.Resty.BaseURL)
	var sessionToken string
	if r, err := d.Resty.
		R().
		SetBasicAuth(username, password).
		SetResult(&sessionToken).
		Post("/api/session"); err != nil {
		return "", err
	} else {
		if r.IsError() {
			return "", fmt.Errorf("error getting session (%s): %s", r.Status(), r.Body())
		}
		sessionCache[username] = session{
			token:   sessionToken,
			created: time.Now(),
		}
		return sessionToken, nil
	}
}
