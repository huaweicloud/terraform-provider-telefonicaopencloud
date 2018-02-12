package telefonicaopencloud

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// Suppress all changes?
func suppressDiffAll(k, old, new string, d *schema.ResourceData) bool {
	return true
}

// Suppress changes if we get a computed min_disk_gb if value is unspecified (default 0)
func suppressMinDisk(k, old, new string, d *schema.ResourceData) bool {
	return new == "0" || old == new
}

// Suppress changes if we get a fixed ip when not expecting one, if we have a floating ip (generates fixed ip).
func suppressComputedFixedWhenFloatingIp(k, old, new string, d *schema.ResourceData) bool {
	if v, ok := d.GetOk("floating_ip"); ok && v != "" {
		return new == "" || old == new
	}
	return false
}
