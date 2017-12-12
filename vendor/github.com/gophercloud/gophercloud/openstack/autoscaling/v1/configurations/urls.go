package configurations

import (
	"github.com/gophercloud/gophercloud"
)

func createURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("scaling_configuration")
}

func getURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL("scaling_configuration", id)
}

func deleteURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL("scaling_configuration", id)
}

func listURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("scaling_configuration")
}
