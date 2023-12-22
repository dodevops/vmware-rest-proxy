package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/thoas/go-funk"
)

// GetVMs returns all VMs from the VM endpoint
func (d DefaultVSphereProxyApi) GetVMs(username string, password string) ([]VM, error) {
	if s, err := d.GetSession(username, password); err != nil {
		return []VM{}, err
	} else {
		logrus.Debugf("Fetching all VMs from %s for %s", d.Resty.BaseURL, username)
		var vmsResponse []VM
		if r, err := d.Resty.
			R().
			SetHeader("vmware-api-session-id", s).
			SetResult(&vmsResponse).
			Get("/api/vcenter/vm"); err != nil {
			logrus.Errorf("Error fetching VMs: %s", err)
			return []VM{}, err
		} else {
			if r.IsError() {
				err := fmt.Errorf("error getting vms (%s): %s", r.Status(), r.Body())
				logrus.Error(err)
				return []VM{}, err
			}
			return vmsResponse, nil
		}
	}
}

// GetVMTags retrieves a list of tags associated with the given vm
func (d DefaultVSphereProxyApi) GetVMTags(username string, password string, VMID string) ([]VMTag, error) {
	var tags []VMTag
	if s, err := d.GetSession(username, password); err != nil {
		return tags, err
	} else {
		logrus.Debugf("Loading the attached tags for vm %s from %s for %s", VMID, d.Resty.BaseURL, username)

		var attachedTagsResponse []string
		if r, err := d.Resty.
			R().
			SetHeader("vmware-api-session-id", s).
			SetResult(&attachedTagsResponse).
			SetBody(gin.H{"object_id": gin.H{
				"type": "VirtualMachine",
				"id":   VMID,
			}}).
			SetQueryParam("action", "list-attached-tags").
			Post("/api/cis/tagging/tag-association"); err != nil {
			logrus.Error(err)
			return tags, err
		} else {
			if r.IsError() {
				err := fmt.Errorf("error getting tags (%s): %s", r.Status(), r.Body())
				logrus.Error(err)
				return tags, err
			}
			for _, tagID := range attachedTagsResponse {
				logrus.Debugf("Loading tag information for tag id %s from vm %s", tagID, VMID)
				var tagResponse struct {
					CategoryID string `json:"category_id"`
					Name       string `json:"name"`
				}
				if r, err := d.Resty.
					R().
					SetHeader("vmware-api-session-id", s).
					SetResult(&tagResponse).
					SetPathParam("tagID", tagID).
					Get("/api/cis/tagging/tag/{tagID}"); err != nil {
					logrus.Error(err)
					return tags, err
				} else {
					if r.IsError() && r.StatusCode() != 404 {
						err := fmt.Errorf("error getting tag information for tag %s (%s): %s", tagID, r.Status(), r.Body())
						logrus.Error(err)
						return tags, err
					} else if r.StatusCode() == 404 || tagResponse.CategoryID == "" {
						logrus.Warnf("Invalid tag %s. Either not found or has no category", tagID)
						continue
					}
					logrus.Debugf("Loading category information for tag %s from vm %s", tagID, VMID)
					var categoryResponse struct {
						Name string `json:"name"`
					}
					if r, err := d.Resty.
						R().
						SetHeader("vmware-api-session-id", s).
						SetResult(&categoryResponse).
						SetPathParam("categoryID", tagResponse.CategoryID).
						Get("/api/cis/tagging/category/{categoryID}"); err != nil {
						logrus.Error(err)
						return tags, err
					} else {
						if r.IsError() {
							err := fmt.Errorf("error getting category (%s): %s", r.Status(), r.Body())
							logrus.Error(err)
							return tags, err
						}

						tags = append(tags, VMTag{
							Value:    tagResponse.Name,
							Category: categoryResponse.Name,
						})
					}
				}
			}
			return tags, nil
		}
	}
}

// GetFQDN uses the VMware guest tools to get the fqdn of a VM (if possible)
func (d DefaultVSphereProxyApi) GetFQDN(username string, password string, VMID string) (string, error) {
	if s, err := d.GetSession(username, password); err != nil {
		return "", err
	} else {
		logrus.Debugf("Trying to figure out the fqdn for vm %s from %s for %s", VMID, d.Resty.BaseURL, username)

		var gR struct {
			DNSValues struct {
				DomainName string `json:"domain_name"`
				HostName   string `json:"host_name"`
			} `json:"dns_values"`
		}
		if r, err := d.Resty.
			R().
			SetHeader("vmware-api-session-id", s).
			SetResult(&gR).
			SetPathParam("vm", VMID).
			Get("/api/vcenter/vm/{vm}/guest/networking"); err != nil {
			logrus.Error(err)
			return "", err
		} else {
			if r.IsError() {
				return "", fmt.Errorf("can not get FQDN (%s): %s", r.Status(), r.Body())
			}
			return fmt.Sprintf(
				"%s.%s",
				gR.DNSValues.HostName,
				gR.DNSValues.DomainName,
			), nil
		}
	}
}

func (d DefaultVSphereProxyApi) GetVMInfo(username string, password string, VMID string) (VMInfo, error) {
	v := VMInfo{}
	if s, err := d.GetSession(username, password); err != nil {
		return v, err
	} else {
		logrus.Debugf("Getting basic information for vm %s from %s for %s", VMID, d.Resty.BaseURL, username)
		var vR struct {
			Name string `json:"name"`
			CPU  struct {
				CoresPerSocket int `json:"cores_per_socket"`
				Count          int `json:"count"`
			} `json:"cpu"`
			Memory struct {
				SizeMiB int `json:"size_MiB"`
			} `json:"memory"`
		}
		if r, err := d.Resty.
			R().
			SetHeader("vmware-api-session-id", s).
			SetResult(&vR).
			SetPathParam("vm", VMID).
			Get("/api/vcenter/vm/{vm}"); err != nil {
			logrus.Error(err)
			return v, err
		} else {
			if r.IsError() {
				return v, fmt.Errorf("can not get vm information (%s): %s", r.Status(), r.Body())
			}
			v.Name = vR.Name
			v.CPUCores = vR.CPU.CoresPerSocket * vR.CPU.Count
			v.ProvisionedRAM = vR.Memory.SizeMiB

			logrus.Debugf("Getting local filesystem information for vm %s from %s for %s", VMID, d.Resty.BaseURL, username)
			type fs struct {
				FreeSpace int `json:"free_space"`
				Capacity  int `json:"capacity"`
			}
			var fR map[string]fs
			if r, err := d.Resty.
				R().
				SetHeader("vmware-api-session-id", s).
				SetResult(&fR).
				SetPathParam("vm", VMID).
				Get("/api/vcenter/vm/{vm}/guest/local-filesystem"); err != nil {
				logrus.Error(err)
				return v, err
			} else {
				if r.IsError() {
					return v, fmt.Errorf("can not get info about local filesystems (%s): %s", r.Status(), r.Body())
				}
				v.ProvisionedStorage = funk.Reduce(funk.Values(fR), func(acc int, f fs) int { return acc + f.Capacity }, 0).(int)
				v.UsedStorage = funk.Reduce(funk.Values(fR), func(acc int, f fs) int { return acc + f.Capacity - f.FreeSpace }, 0).(int)
				return v, nil
			}
		}
	}
}
