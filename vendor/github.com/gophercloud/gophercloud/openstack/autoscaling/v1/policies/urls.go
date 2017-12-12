package policies

import (
	"github.com/gophercloud/gophercloud"
)

//createURL will build the rest query url of creation
//the create url is endpoint/scaling_policy
func createURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("scaling_policy")
}

//deleteURL will build the url of deletion
//its pattern is endpoint/scaling_policy/<policy-id>
func deleteURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL("scaling_policy", id)
}

//getURL will build the get url of get function
func getURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL("scaling_policy", id)
}

func updateURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("scaling_policy", id)
}
