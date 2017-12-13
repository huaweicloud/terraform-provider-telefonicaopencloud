package eips

import "github.com/gophercloud/gophercloud"

const resourcePath = "publicips"

func rootURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL(resourcePath)
}

func resourceURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL(resourcePath, id)
}
