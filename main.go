package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/terraform-providers/terraform-provider-telefornicaopencloud/telefornicaopencloud"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: telefornicaopencloud.Provider})
}
