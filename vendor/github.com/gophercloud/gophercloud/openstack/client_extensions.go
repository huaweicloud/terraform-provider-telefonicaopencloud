package openstack

import (
	"github.com/gophercloud/gophercloud"
)


//NewAutoScalingService creates a ServiceClient that may be used to access the
//auto-scaling service of huawei public cloud
func NewAutoScalingService(client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "as")
	return sc, err
}

//NewVpcServiceV1 creates the a ServiceClient that may be used to access the v1
//vpc service which is a service of public ip management of huawei cloud
func NewVpcServiceV1(client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	sc, err := initClientOpts(client, eo, "network")
	return sc, err
}
