package api

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"vmware-rest-proxy/internal"
)

// SessionResponse holds the value returned by the VMware session controller
type SessionResponse struct {
	Value string `json:"value"`
}

// GetSession returns the vmware session id to be used by other requests
func GetSession(c internal.Config, username string, password string) (string, error) {
	logrus.Debugf("Creating VMware session for user %s at %s", username, c.Resty.BaseURL)
	var sessionRespone SessionResponse
	if r, err := c.Resty.
		R().
		SetBasicAuth(username, password).
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
