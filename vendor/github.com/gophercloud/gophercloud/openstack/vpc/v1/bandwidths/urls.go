package bandwidths

import "github.com/gophercloud/gophercloud"

const resourcePath = "bandwidths"

func resourceURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL(resourcePath, id)
}
