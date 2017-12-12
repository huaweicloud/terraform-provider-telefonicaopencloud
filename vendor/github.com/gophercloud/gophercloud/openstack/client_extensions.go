package openstack

import (
	"github.com/gophercloud/gophercloud"
)

//NewRdsServiceV1 creates the a ServiceClient that may be used to access the v1
//rds service which is a service of db instances management of huawei cloud
func NewRdsServiceV1(client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "rds")
	//sc.ResourceBase = sc.Endpoint + "rds/v1" //TODO make sure if it's neccessary to set the base uri
	return sc, err
}

//NewAutoScalingService creates a ServiceClient that may be used to access the
//auto-scaling service of huawei public cloud
func NewAutoScalingService(client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "autoscaling-api")
	return sc, err
}
