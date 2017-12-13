package telefonicaopencloud

import (
	"fmt"
	"log"
	"regexp"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/elb/certificate"
)

const nameELBCert = "ELB-Certificate"

func resourceELBCertificate() *schema.Resource {
	return &schema.Resource{
		Create: resourceELBCertificateCreate,
		Read:   resourceELBCertificateRead,
		Update: resourceELBCertificateUpdate,
		Delete: resourceELBCertificateDelete,

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
			},

			"domain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: func(v interface{}, k string) (ws []string, errors []error) {
					value := v.(string)
					vv := regexp.MustCompile("^[A-Z0-9]+[a-zA-Z0-9-.]{0,253}$")
					if !vv.MatchString(value) {
						errors = append(errors, fmt.Errorf("%s is a string that contains a maximum of 254 characters, can only contain letters, digits, hyphens (-), and periods (.), and must start with uppercase letters or digits", k))
					}
					return
				},
			},

			"certificate": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"private_key": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"update_time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			"create_time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}

}

func resourceELBCertificateCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkingClient, err := chooseELBClient(d, config)
	if err != nil {
		return fmt.Errorf("Error creating OpenStack networking client: %s", err)
	}

	var createOpts certificate.CreateOpts
	err = buildELBCreateParam(&createOpts, d)
	if err != nil {
		return fmt.Errorf("Error creating %s: building parameter failed:%s", nameELBCert, err)
	}
	log.Printf("[DEBUG] Create %s Options: %#v", nameELBCert, createOpts)

	c, err := certificate.Create(networkingClient, createOpts).Extract()
	if err != nil {
		return fmt.Errorf("Error creating %s: %s", nameELBCert, err)
	}
	log.Printf("[DEBUG] Create %s: %#v", nameELBCert, *c)

	// If all has been successful, set the ID on the resource
	d.SetId(c.ID)

	return resourceELBCertificateRead(d, meta)
}

func resourceELBCertificateRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkingClient, err := chooseELBClient(d, config)
	if err != nil {
		return fmt.Errorf("Error creating OpenStack networking client: %s", err)
	}

	c, err := certificate.Get(networkingClient, d.Id()).Extract()
	if err != nil {
		return CheckDeleted(d, err, "certificate")
	}
	log.Printf("[DEBUG] Retrieved %s(%s): %#v", nameELBCert, d.Id(), *c)

	return refreshResourceData(c, d)
}

func resourceELBCertificateUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkingClient, err := chooseELBClient(d, config)
	if err != nil {
		return fmt.Errorf("Error creating OpenStack networking client: %s", err)
	}

	cId := d.Id()
	var updateOpts certificate.UpdateOpts
	err = buildELBUpdateParam(&updateOpts, d)
	if err != nil {
		return fmt.Errorf("Error updating %s(%s): building parameter failed:%s", nameELBCert, cId, err)
	}
	b, err := updateOpts.IsNeedUpdate()
	if err != nil {
		return err
	}
	if !b {
		log.Printf("[INFO] Updating %s %s with no changes", nameELBCert, cId)
		return nil
	}
	log.Printf("[DEBUG] Updating %s(%s) with options: %#v", nameELBCert, cId, updateOpts)

	timeout := d.Timeout(schema.TimeoutUpdate)
	err = resource.Retry(timeout, func() *resource.RetryError {
		_, err := certificate.Update(networkingClient, cId, updateOpts).Extract()
		if err != nil {
			return checkForRetryableError(err)
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("Error updating %s(%s): %s", nameELBCert, cId, err)
	}

	return resourceELBCertificateRead(d, meta)
}

func resourceELBCertificateDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	networkingClient, err := chooseELBClient(d, config)
	if err != nil {
		return fmt.Errorf("Error creating OpenStack networking client: %s", err)
	}

	cId := d.Id()
	log.Printf("[DEBUG] Deleting %s %s", nameELBCert, cId)

	timeout := d.Timeout(schema.TimeoutDelete)
	err = resource.Retry(timeout, func() *resource.RetryError {
		err := certificate.Delete(networkingClient, cId).ExtractErr()
		if err != nil {
			return checkForRetryableError(err)
		}
		return nil
	})
	if err != nil {
		if isELBResourceNotFound(err) {
			log.Printf("[INFO] deleting an unavailable %s: %s", nameELBCert, cId)
			return nil
		}
		return fmt.Errorf("Error deleting %s(%s): %s", nameELBCert, cId, err)
	}

	return nil
}
