package telefonicaopencloud

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/elb/quotas"
)

func dataResourceELBQuota() *schema.Resource {
	return &schema.Resource{
		Read: dataResourceELBQuotaRead,

		Schema: map[string]*schema.Schema{
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
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
	}
}

func dataResourceELBQuotaRead(d *schema.ResourceData, meta interface{}) error {
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

	t := d.Get("type").(string)
	for _, v := range q {
		if v.Type == t {
			return refreshResourceData(v, d)
		}
	}
	return fmt.Errorf("Error does not find the quota for %s", t)
}
