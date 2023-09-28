package api

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"vmware-rest-proxy/internal"
)

// VMResponseValue is the value in response from the VM endpoint
type VMResponseValue struct {
	VM   string `json:"vm"`
	Name string `json:"name"`
}

// VMSResponse holds a list of VMResponseValue as returned from the VM endpoint
type VMSResponse struct {
	Value []VMResponseValue `json:"value"`
}

// GetVMs returns all VMs from the VM endpoint
func GetVMs(c internal.Config, username string, password string) ([]VMResponseValue, error) {
	if s, err := GetSession(c, username, password); err != nil {
		return []VMResponseValue{}, err
	} else {
		logrus.Debugf("Fetching all VMs from %s for %s", c.Resty.BaseURL, username)
		var vmsResponse VMSResponse
		if r, err := c.Resty.
			R().
			SetHeader("vmware-api-session-id", s).
			SetResult(&vmsResponse).
			Get("/rest/vcenter/vm"); err != nil {
			logrus.Errorf("Error fetching VMs: %s", err)
			return []VMResponseValue{}, err
		} else {
			if r.IsError() {
				err := fmt.Errorf("error getting vms (%s): %s", r.Status(), r.Body())
				logrus.Error(err)
				return []VMResponseValue{}, err
			}
			return vmsResponse.Value, nil
		}
	}
}

// AttachedTagsResponse holds a response from the tags endpoint with the list of attached tag ids
type AttachedTagsResponse struct {
	Value []string `json:"value"`
}

// TagResponseValue Holds the value of TagResponse describing a tag from the tag endpoint
type TagResponseValue struct {
	CategoryID string `json:"category_id"`
	Name       string `json:"name"`
}

// TagResponse holds the response from the tag endpoint
type TagResponse struct {
	Value TagResponseValue `json:"value"`
}

// CategoryResponseValue holds the value of CategoryResponse describing a category from the category endpoint
type CategoryResponseValue struct {
	Name string `json:"name"`
}

// CategoryResponse holds the response of the category endpoint
type CategoryResponse struct {
	Value CategoryResponseValue `json:"value"`
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
		var attachedTagsResponse AttachedTagsResponse
		if r, err := c.Resty.
			R().
			SetHeader("vmware-api-session-id", s).
			SetResult(&attachedTagsResponse).
			SetBody(IDBody{ObjectID: VMIDBody{
				Type: "VirtualMachine",
				ID:   VMID,
			}}).
			SetQueryParam("~action", "list-attached-tags").
			Post("/rest/com/vmware/cis/tagging/tag-association"); err != nil {
			logrus.Error(err)
			return tags, err
		} else {
			if r.IsError() {
				err := fmt.Errorf("error getting tags (%s): %s", r.Status(), r.Body())
				logrus.Error(err)
				return tags, err
			}
			for _, tagID := range attachedTagsResponse.Value {
				logrus.Debugf("Loading tag information for tag id %s from vm %s", tagID, VMID)
				var tagResponse TagResponse
				if r, err := c.Resty.
					R().
					SetHeader("vmware-api-session-id", s).
					SetResult(&tagResponse).
					SetPathParam("tagID", tagID).
					Get("/rest/com/vmware/cis/tagging/tag/id:{tagID}"); err != nil {
					logrus.Error(err)
					return tags, err
				} else {
					if r.IsError() {
						err := fmt.Errorf("error getting tags (%s): %s", r.Status(), r.Body())
						logrus.Error(err)
						return tags, err
					}
					logrus.Debugf("Loading category information for tag %s from vm %s", tagID, VMID)
					var categoryResponse CategoryResponse
					if r, err := c.Resty.
						R().
						SetHeader("vmware-api-session-id", s).
						SetResult(&categoryResponse).
						SetPathParam("categoryID", tagResponse.Value.CategoryID).
						Get("/rest/com/vmware/cis/tagging/category/id:{categoryID}"); err != nil {
						logrus.Error(err)
						return tags, err
					} else {
						if r.IsError() {
							err := fmt.Errorf("error getting category (%s): %s", r.Status(), r.Body())
							logrus.Error(err)
							return tags, err
						}

						tags = append(tags, VMTag{
							Value:    tagResponse.Value.Name,
							Category: categoryResponse.Value.Name,
						})
					}
				}
			}
			return tags, nil
		}
	}
}
