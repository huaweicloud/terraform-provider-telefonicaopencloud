package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/terraform-providers/terraform-provider-telefonicaopencloud/telefonicaopencloud"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: telefonicaopencloud.Provider})
}
