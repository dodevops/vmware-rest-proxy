package api

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

func (d DefaultVSphereProxyApi) GetHosts(username string, password string) ([]Host, error) {
	if s, err := d.GetSession(username, password); err != nil {
		return []Host{}, err
	} else {
		logrus.Debugf("Fetching all hosts from %s for %s", d.Resty.BaseURL, username)
		var hostsResponse []Host
		if r, err := d.Resty.
			R().
			SetHeader("vmware-api-session-id", s).
			SetResult(&hostsResponse).
			Get("/api/vcenter/host"); err != nil {
			logrus.Errorf("Error fetching hosts: %s", err)
			return []Host{}, err
		} else {
			if r.IsError() {
				err := fmt.Errorf("error getting hosts (%s): %s", r.Status(), r.Body())
				logrus.Error(err)
				return []Host{}, err
			}
			return hostsResponse, nil
		}
	}
}
