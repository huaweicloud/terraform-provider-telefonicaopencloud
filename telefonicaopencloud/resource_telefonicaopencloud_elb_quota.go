package telefonicaopencloud

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/elb/quotas"
)

func resourceELBQuota() *schema.Resource {
	return &schema.Resource{
		Read: resourceELBQuotaRead,

		Schema: map[string]*schema.Schema{
			"quotas": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"used": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"quota": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"max": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"min": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceELBQuotaRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkingClient, err := chooseELBClient(d, config)
	if err != nil {
		return fmt.Errorf("Error creating OpenStack networking client: %s", err)
	}

	q, err := quotas.Get(networkingClient).Extract()
	if err != nil {
		return CheckDeleted(d, err, "listener")
	}
	log.Printf("[DEBUG] Retrieved ELB-Quotas %#v", q)

	return refreshResourceData(q, d)
}
