package quotas

import (
	"github.com/gophercloud/gophercloud"
)

// Get retrieves a particular Loadbalancer based on its unique ID.
func Get(c *gophercloud.ServiceClient) (r GetResult) {
	_, r.Err = c.Get(rootURL(c), &r.Body, nil)
	return
}
