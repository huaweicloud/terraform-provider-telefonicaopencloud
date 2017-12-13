package policies

import (
	"github.com/gophercloud/gophercloud"
)

const resourcePath = "scaling_policy"
const apiVersion = "v1"

//createURL will build the rest query url of creation
//the create url is endpoint/scaling_policy
func createURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL(apiVersion, client.ProjectID, resourcePath)
}

//deleteURL will build the url of deletion
//its pattern is endpoint/scaling_policy/<policy-id>
func deleteURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL(apiVersion, client.ProjectID, resourcePath, id)
}

//getURL will build the get url of get function
func getURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL(apiVersion, client.ProjectID, resourcePath, id)
}

func updateURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL(apiVersion, client.ProjectID, resourcePath, id)
}
