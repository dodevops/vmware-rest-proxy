package api

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

func (d DefaultVSphereProxyApi) GetDatastores(username string, password string) ([]Datastore, error) {
	if s, err := d.GetSession(username, password); err != nil {
		return []Datastore{}, err
	} else {
		logrus.Debugf("Fetching all datastores from %s for %s", d.Resty.BaseURL, username)
		var datastores []Datastore
		if r, err := d.Resty.
			R().
			SetHeader("vmware-api-session-id", s).
			SetResult(&datastores).
			Get("/api/vcenter/datastore"); err != nil {
			logrus.Errorf("Error fetching datastores: %s", err)
			return []Datastore{}, err
		} else {
			if r.IsError() {
				err := fmt.Errorf("error getting datastores (%s): %s", r.Status(), r.Body())
				logrus.Error(err)
				return []Datastore{}, err
			}
			return datastores, nil
		}
	}
}
