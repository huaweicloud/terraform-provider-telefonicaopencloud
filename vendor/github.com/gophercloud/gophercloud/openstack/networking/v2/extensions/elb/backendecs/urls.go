package backendecs

import "github.com/gophercloud/gophercloud"

const (
	rootPath     = "elbaas"
	resourcePath = "listeners"
	memberPath   = "members"
)

func rootURL(c *gophercloud.ServiceClient, lId string) string {
	return c.ServiceURL(c.ProjectID, rootPath, resourcePath, lId, memberPath)
}

func actionURL(c *gophercloud.ServiceClient, lId string) string {
	return c.ServiceURL(c.ProjectID, rootPath, resourcePath, lId, memberPath, "action")
}
