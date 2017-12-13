package configurations

import (
	"github.com/gophercloud/gophercloud"
)

const resourcePath = "scaling_configuration"
const apiVersion = "v1"

func createURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL(apiVersion, client.ProjectID, resourcePath)
}

func getURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL(apiVersion, client.ProjectID, resourcePath, id)
}

func deleteURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL(apiVersion, client.ProjectID, resourcePath, id)
}

func listURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL(apiVersion, client.ProjectID, resourcePath)
}
