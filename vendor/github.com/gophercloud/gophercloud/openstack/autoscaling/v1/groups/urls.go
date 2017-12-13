package groups

import (
	"github.com/gophercloud/gophercloud"
)

const resourcePath = "scaling_group"
const apiVersion = "v1"

func createURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL(apiVersion, c.ProjectID, resourcePath)
}

func deleteURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL(apiVersion, c.ProjectID, resourcePath, id)
}

func getURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL(apiVersion, c.ProjectID, resourcePath, id)
}

func listURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL(apiVersion, c.ProjectID, resourcePath)
}

func enableURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL(apiVersion, c.ProjectID, resourcePath, id, "action")
}

func updateURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL(apiVersion, c.ProjectID, resourcePath, id)
}
