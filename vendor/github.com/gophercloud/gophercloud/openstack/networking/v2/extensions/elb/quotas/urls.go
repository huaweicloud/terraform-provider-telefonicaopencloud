package quotas

import "github.com/gophercloud/gophercloud"

const (
	rootPath     = "elbaas"
	resourcePath = "quotas"
)

func rootURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL(c.ProjectID, rootPath, resourcePath)
}
