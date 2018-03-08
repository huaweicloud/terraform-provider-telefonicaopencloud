package telefonicaopencloud

import (
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/huaweicloud/golangsdk/openstack/networking/v2/extensions/elb/listeners"
)

const nameELBListener = "ELB-Listener"

func resourceELBListener() *schema.Resource {
	return &schema.Resource{
		Create: resourceELBListenerCreate,
		Read:   resourceELBListenerRead,
		Update: resourceELBListenerUpdate,
		Delete: resourceELBListenerDelete,

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

			"loadbalancer_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},

			"backend_protocol": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"backend_port": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},

			"lb_algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"session_sticky": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			"sticky_session_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"cookie_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			"tcp_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			"tcp_draining": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			"tcp_draining_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			"certificate_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"certificates": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"udp_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			"ssl_protocols": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "TLSv1.2",
			},

			"ssl_ciphers": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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

			"admin_state_up": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},

			"member_number": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},

			"healthcheck_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceELBListenerCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkingClient, err := chooseELBClient(d, config)
	if err != nil {
		return fmt.Errorf("Error creating TelefonicaOpenCloud networking client: %s", err)
	}

	var opts listeners.CreateOpts
	err, not_pass_params := buildCreateParam(&opts, d)
	if err != nil {
		return fmt.Errorf("Error creating %s: building parameter failed:%s", nameELBListener, err)
	}
	log.Printf("[DEBUG] Create %s Options: %#v", nameELBListener, opts)

	switch {
	case opts.Protocol == "HTTPS" || opts.Protocol == "SSL" && !hasFilledParam(d, "certificate_id"):
		return fmt.Errorf("certificate_id is mandatory when protocol is set to HTTPS or SSL")
	}
	l, err := listeners.Create(networkingClient, opts, not_pass_params).Extract()
	if err != nil {
		return fmt.Errorf("Error creating %s: %s", nameELBListener, err)
	}
	log.Printf("[DEBUG] Create %s: %#v", nameELBListener, *l)

	// Wait for Listener to become active before continuing
	timeout := d.Timeout(schema.TimeoutCreate)
	err = waitForELBListenerActive(networkingClient, l.ID, timeout)
	if err != nil {
		return err
	}

	// If all has been successful, set the ID on the resource
	d.SetId(l.ID)

	return resourceELBListenerRead(d, meta)
}

func resourceELBListenerRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkingClient, err := chooseELBClient(d, config)
	if err != nil {
		return fmt.Errorf("Error creating TelefonicaOpenCloud networking client: %s", err)
	}

	l, err := listeners.Get(networkingClient, d.Id()).Extract()
	if err != nil {
		return CheckDeleted(d, err, "listener")
	}
	log.Printf("[DEBUG] Retrieved %s %s: %#v", nameELBListener, d.Id(), l)

	sp := d.Get("ssl_protocols")
	if l.SslProtocols == "" && sp != nil {
		l.SslProtocols = sp.(string)
	}
	return refreshResourceData(l, d)
}

func resourceELBListenerUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkingClient, err := chooseELBClient(d, config)
	if err != nil {
		return fmt.Errorf("Error creating TelefonicaOpenCloud networking client: %s", err)
	}

	lId := d.Id()

	var opts listeners.UpdateOpts
	err, not_pass_params := buildUpdateParam(&opts, d)
	if err != nil {
		return fmt.Errorf("Error updating %s %s: building parameter failed:%s", nameELBListener, lId, err)
	}

	protocol := d.Get("protocol").(string)
	switch {
	case protocol == "HTTPS" || protocol == "SSL" && !hasFilledParam(d, "certificate_id"):
		return fmt.Errorf("certificate_id is mandatory when protocol is set to HTTPS or SSL")
	}
	// Wait for Listener to become active before continuing
	timeout := d.Timeout(schema.TimeoutUpdate)
	err = waitForELBListenerActive(networkingClient, lId, timeout)
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating %s %s with options: %#v", nameELBListener, lId, opts)
	err = resource.Retry(timeout, func() *resource.RetryError {
		_, err := listeners.Update(networkingClient, lId, opts, not_pass_params).Extract()
		if err != nil {
			return checkForRetryableError(err)
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("Error updating %s %s: %s", nameELBListener, lId, err)
	}

	return resourceELBListenerRead(d, meta)
}

func resourceELBListenerDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkingClient, err := chooseELBClient(d, config)
	if err != nil {
		return fmt.Errorf("Error creating TelefonicaOpenCloud networking client: %s", err)
	}

	lId := d.Id()
	log.Printf("[DEBUG] Deleting %s %s", nameELBListener, lId)

	timeout := d.Timeout(schema.TimeoutDelete)
	err = resource.Retry(timeout, func() *resource.RetryError {
		err := listeners.Delete(networkingClient, lId).ExtractErr()
		if err != nil {
			return checkForRetryableError(err)
		}
		return nil
	})
	if err != nil {
		if isResourceNotFound(err) {
			log.Printf("[INFO] deleting an unavailable %s: %s", nameELBListener, lId)
			return nil
		}
		return fmt.Errorf("Error deleting %s %s: %s", nameELBListener, lId, err)
	}

	return nil
}
