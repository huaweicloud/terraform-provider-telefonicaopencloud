package telefonicaopencloud

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/gophercloud/gophercloud/openstack/autoscaling/v1/groups"
	"log"
)

func TestAccASV1Group_basic(t *testing.T) {
	var asGroup groups.Group

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckASV1GroupDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccASV1Group_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckASV1GroupExists("telefonicaopencloud_as_group_v1.my_group", &asGroup),
				),
			},
		},
	})
}

func testAccCheckASV1GroupDestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)
	asClient, err := config.autoscalingV1Client(OS_REGION_NAME)
	if err != nil {
		return fmt.Errorf("Error creating telefonicaopencloud autoscaling client: %s", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "telefonicaopencloud_as_group_v1" {
			continue
		}

		_, err := groups.Get(asClient, rs.Primary.ID).Extract()
		if err == nil {
			return fmt.Errorf("AS group still exists")
		}
	}

	return nil
}

func testAccCheckASV1GroupExists(n string, group *groups.Group) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		config := testAccProvider.Meta().(*Config)
		asClient, err := config.autoscalingV1Client(OS_REGION_NAME)
		if err != nil {
			return fmt.Errorf("Error creating telefonicaopencloud autoscaling client: %s", err)
		}

		found, err := groups.Get(asClient, rs.Primary.ID).Extract()
		if err != nil {
			return err
		}

		if found.ID != rs.Primary.ID {
			return fmt.Errorf("Autoscaling Group not found")
		}
		log.Printf("[DEBUG] test found is: %#v", found)
		group = &found

		return nil
	}
}

var testAccASV1Group_basic = fmt.Sprintf(`
resource "telefonicaopencloud_compute_keypair_v2" "hth_key" {
  name = "hth_key"
}

resource "telefonicaopencloud_networking_router_v2" "hth_router" {
  name             = "hth-asg-router"
  admin_state_up   = "true"
}

resource "telefonicaopencloud_networking_network_v2" "hth_network" {
  name           = "hth-asg-network"
  admin_state_up = "true"
}

resource "telefonicaopencloud_networking_subnet_v2" "hth_subnet" {
  name            = "hth-asg-subnet"
  network_id      = "${telefonicaopencloud_networking_network_v2.hth_network.id}"
  cidr            = "192.168.111.0/24"
  ip_version      = 4
  dns_nameservers = ["8.8.8.8", "8.8.4.4"]
}

resource "telefonicaopencloud_networking_secgroup_v2" "hth_secgroup" {
  name = "hth_secgroup"
  description = "terraform asg acceptance test"
}

resource "telefonicaopencloud_as_group_v1" "hth_asg"{
  scaling_group_name = "hth_asg"
  scaling_configuration_id = "${telefonicaopencloud_as_configuration_v1.hth_as_config.id}"
  desire_instance_number = 1
  min_instance_number = 0
  max_instance_number = 3
  networks = [
    {id = "${telefonicaopencloud_networking_network_v2.hth_network.id}"},
  ]
  security_groups = [
    {id = "${telefonicaopencloud_networking_secgroup_v2.hth_secgroup.id}"},
  ]
  vpc_id = "${telefonicaopencloud_networking_router_v2.hth_router.id}"
  delete_publicip = true
  delete_instances = "yes"
}

resource "telefonicaopencloud_as_configuration_v1" "hth_as_config"{
  scaling_configuration_name = "hth_as_config"
  instance_config = {
    flavor = "%s"
    image = "%s"
    disk = [
      {size = 40
      volume_type = "SATA"
      disk_type = "SYS"}
    ]
    key_name = "${telefonicaopencloud_compute_keypair_v2.hth_key.id}"
  }
}
`, OS_FLAVOR_NAME, OS_IMAGE_ID)
