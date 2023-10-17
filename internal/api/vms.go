package api

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"vmware-rest-proxy/internal"
)

// VMResponse is the value in response from the VM endpoint
type VMResponse struct {
	VM   string `json:"vm"`
	Name string `json:"name"`
}

// GetVMs returns all VMs from the VM endpoint
func GetVMs(c internal.Config, username string, password string) ([]VMResponse, error) {
	if s, err := GetSession(c, username, password); err != nil {
		return []VMResponse{}, err
	} else {
		logrus.Debugf("Fetching all VMs from %s for %s", c.Resty.BaseURL, username)
		var vmsResponse []VMResponse
		if r, err := c.Resty.
			R().
			SetHeader("vmware-api-session-id", s).
			SetResult(&vmsResponse).
			Get("/api/vcenter/vm"); err != nil {
			logrus.Errorf("Error fetching VMs: %s", err)
			return []VMResponse{}, err
		} else {
			if r.IsError() {
				err := fmt.Errorf("error getting vms (%s): %s", r.Status(), r.Body())
				logrus.Error(err)
				return []VMResponse{}, err
			}
			return vmsResponse, nil
		}
	}
}

// TagResponse Holds the value of TagResponse describing a tag from the tag endpoint
type TagResponse struct {
	CategoryID string `json:"category_id"`
	Name       string `json:"name"`
}

// CategoryResponse holds the value of CategoryResponse describing a category from the category endpoint
type CategoryResponse struct {
	Name string `json:"name"`
}

// VMIDBody holds the object id from IDBody pointing to a vm
type VMIDBody struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

// IDBody holds the information required to search for a specific VM
type IDBody struct {
	ObjectID VMIDBody `json:"object_id"`
}

// VMTag holds a tag from vSphere
type VMTag struct {
	Value    string `json:"value"`
	Category string `json:"category"`
}

// GetVMTags retrieves a list of tags associated with the given vm
func GetVMTags(c internal.Config, username string, password string, VMID string) ([]VMTag, error) {
	var tags []VMTag
	if s, err := GetSession(c, username, password); err != nil {
		return tags, err
	} else {
		logrus.Debugf("Loading the attached tags for vm %s from %s for %s", VMID, c.Resty.BaseURL, username)
		var attachedTagsResponse []string
		if r, err := c.Resty.
			R().
			SetHeader("vmware-api-session-id", s).
			SetResult(&attachedTagsResponse).
			SetBody(IDBody{ObjectID: VMIDBody{
				Type: "VirtualMachine",
				ID:   VMID,
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
				var tagResponse TagResponse
				if r, err := c.Resty.
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
						logrus.Warn("Invalid tag %s. Either not found or has no category", tagID)
						continue
					}
					logrus.Debugf("Loading category information for tag %s from vm %s", tagID, VMID)
					var categoryResponse CategoryResponse
					if r, err := c.Resty.
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

type GuestNetworkingResponseDNSValues struct {
	DomainName string `json:"domain_name"`
	HostName   string `json:"host_name"`
}

type GuestNetworkingResponse struct {
	DNSValues GuestNetworkingResponseDNSValues `json:"dns_values"`
}

// GetFQDN uses the VMware guest tools to get the fqdn of a VM (if possible)
func GetFQDN(c internal.Config, username string, password string, VMID string) (string, error) {
	if s, err := GetSession(c, username, password); err != nil {
		return "", err
	} else {
		logrus.Debugf("Trying to figure out the fqdn for vm %s from %s for %s", VMID, c.Resty.BaseURL, username)

		var guestNetworkingResponse GuestNetworkingResponse
		if r, err := c.Resty.
			R().
			SetHeader("vmware-api-session-id", s).
			SetResult(&guestNetworkingResponse).
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
				guestNetworkingResponse.DNSValues.HostName,
				guestNetworkingResponse.DNSValues.DomainName,
			), nil
		}
	}
}
