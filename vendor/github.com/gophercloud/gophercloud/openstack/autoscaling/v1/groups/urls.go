package groups

import (
	"github.com/gophercloud/gophercloud"
)

func createURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("scaling_group")
}

func deleteURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("scaling_group", id)
}

func getURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("scaling_group", id)
}

func listURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("scaling_group")
}

func enableURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("scaling_group", id, "action")
}

func updateURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("scaling_group", id)
}
