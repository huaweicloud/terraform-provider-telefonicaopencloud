package telefonicaopencloud

import (
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/huaweicloud/golangsdk/openstack/networking/v2/extensions/elb"
	"github.com/huaweicloud/golangsdk/openstack/networking/v2/extensions/elb/loadbalancers"
)

const nameELBLB = "ELB-LoadBalancer"

func resourceELBLoadBalancer() *schema.Resource {
	return &schema.Resource{
		Create: resourceELBLoadBalancerCreate,
		Read:   resourceELBLoadBalancerRead,
		Update: resourceELBLoadBalancerUpdate,
		Delete: resourceELBLoadBalancerDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"vpc_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"bandwidth": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"admin_state_up": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},

			"vip_subnet_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"az": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			"charge_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "bandwidth",
			},

			"eip_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			"security_group_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"vip_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"tenantid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			"update_time": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},

			"create_time": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},

			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceELBLoadBalancerCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkingClient, err := chooseELBClient(d, config)
	if err != nil {
		return fmt.Errorf("Error creating TelefonicaOpenCloud networking client: %s", err)
	}

	var opts loadbalancers.CreateOpts
	_, err = buildCreateParam(&opts, d, nil)
	if err != nil {
		return fmt.Errorf("Error creating %s: building parameter failed:%s", nameELBLB, err)
	}
	log.Printf("[DEBUG] Create %s Options: %#v", nameELBLB, opts)

	switch {
	case opts.Type == "External" && !hasFilledOpt(d, "bandwidth"):
		return fmt.Errorf("bandwidth is mandatory when type is set to External")

	case opts.Type == "Internal" && !hasFilledOpt(d, "vip_subnet_id"):
		return fmt.Errorf("vip_subnet_id is mandatory when type is set to Internal")

	case opts.Type == "Internal" && !hasFilledOpt(d, "az"):
		return fmt.Errorf("az is mandatory when type is set to Internal")

	case opts.Type == "Internal" && !hasFilledOpt(d, "security_group_id"):
		return fmt.Errorf("security_group_id is mandatory when type is set to Internal")

	case opts.Type == "Internal" && !hasFilledOpt(d, "tenantid"):
		return fmt.Errorf("tenantid is mandatory when type is set to Internal")
	}

	j, err := loadbalancers.Create(networkingClient, opts).Extract()
	if err != nil {
		return fmt.Errorf("Error creating %s: %s", nameELBLB, err)
	}
	log.Printf("[DEBUG] Create %s, the job is: %#v", nameELBLB, *j)

	// Wait for LoadBalancer to become active before continuing
	timeout := d.Timeout(schema.TimeoutCreate)
	jobInfo, err := waitForELBJobSuccess(networkingClient, j, timeout)
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] Create %s, the job is: %#v", nameELBLB, jobInfo)

	e, ok := jobInfo.Entities["elb"]
	if !ok {
		return fmt.Errorf("Error creating %s: get the entity from job info failed", nameELBLB)
	}
	i, ok := e.(map[string]interface{})
	if !ok {
		return fmt.Errorf("Error creating %s: convert job entity to map failed", nameELBLB)
	}
	eid, ok := i["id"]
	if !ok {
		return fmt.Errorf("Error creating %s: get elb id from job entity failed", nameELBLB)
	}

	// If all has been successful, set the ID on the resource
	d.SetId(eid.(string))

	return resourceELBLoadBalancerRead(d, meta)
}

func resourceELBLoadBalancerRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkingClient, err := chooseELBClient(d, config)
	if err != nil {
		return fmt.Errorf("Error creating TelefonicaOpenCloud networking client: %s", err)
	}

	lb, err := loadbalancers.Get(networkingClient, d.Id()).Extract()
	if err != nil {
		return CheckDeleted(d, err, "loadbalancer")
	}
	log.Printf("[DEBUG] Retrieved %s %s: %#v", nameELBLB, d.Id(), lb)

	return refreshResourceData(lb, d, nil)
}

func resourceELBLoadBalancerUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkingClient, err := chooseELBClient(d, config)
	if err != nil {
		return fmt.Errorf("Error creating TelefonicaOpenCloud networking client: %s", err)
	}

	lbId := d.Id()

	var updateOpts loadbalancers.UpdateOpts
	not_pass_param, err := buildUpdateParam(&updateOpts, d, nil)
	if err != nil {
		return fmt.Errorf("Error updating %s %s: building parameter failed:%s", nameELBLB, lbId, err)
	}

	// Wait for LoadBalancer to become active before continuing
	timeout := d.Timeout(schema.TimeoutUpdate)
	err = waitForELBLoadBalancerActive(networkingClient, lbId, timeout)
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating %s %s with options: %#v", nameELBLB, lbId, updateOpts)
	var job *elb.Job
	err = resource.Retry(timeout, func() *resource.RetryError {
		j, err := loadbalancers.Update(networkingClient, lbId, updateOpts, not_pass_param).Extract()
		if err != nil {
			return checkForRetryableError(err)
		}
		job = j
		return nil
	})
	if err != nil {
		return fmt.Errorf("Error updating %s %s: %s", nameELBLB, lbId, err)
	}

	// Wait for LoadBalancer to become active before continuing
	_, err = waitForELBJobSuccess(networkingClient, job, timeout)
	if err != nil {
		return err
	}

	return resourceELBLoadBalancerRead(d, meta)
}

func resourceELBLoadBalancerDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkingClient, err := chooseELBClient(d, config)
	if err != nil {
		return fmt.Errorf("Error creating TelefonicaOpenCloud networking client: %s", err)
	}

	lbId := d.Id()
	log.Printf("[DEBUG] Deleting %s %s", nameELBLB, lbId)

	var job *elb.Job
	timeout := d.Timeout(schema.TimeoutDelete)
	err = resource.Retry(timeout, func() *resource.RetryError {
		j, err := loadbalancers.Delete(networkingClient, lbId).Extract()
		if err != nil {
			return checkForRetryableError(err)
		}
		job = j
		return nil
	})
	if err != nil {
		if isResourceNotFound(err) {
			log.Printf("[INFO] deleting an unavailable %s: %s", nameELBLB, lbId)
			return nil
		}
		return fmt.Errorf("Error deleting %s %s: %s", nameELBLB, lbId, err)
	}
	log.Printf("[DEBUG] Delete %s, the job is: %#v", nameELBLB, *job)

	_, err = waitForELBJobSuccess(networkingClient, job, timeout)
	if err != nil {
		return err
	}
	return nil
}
