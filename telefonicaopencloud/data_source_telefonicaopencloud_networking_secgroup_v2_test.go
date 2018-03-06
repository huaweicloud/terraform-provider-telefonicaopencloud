package telefonicaopencloud

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccTelefonicaOpenCloudNetworkingSecGroupV2DataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccTelefonicaOpenCloudNetworkingSecGroupV2DataSource_group,
			},
			resource.TestStep{
				Config: testAccTelefonicaOpenCloudNetworkingSecGroupV2DataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkingSecGroupV2DataSourceID("data.telefonicaopencloud_networking_secgroup_v2.secgroup_1"),
					resource.TestCheckResourceAttr(
						"data.telefonicaopencloud_networking_secgroup_v2.secgroup_1", "name", "secgroup_1"),
				),
			},
		},
	})
}

func TestAccTelefonicaOpenCloudNetworkingSecGroupV2DataSource_secGroupID(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccTelefonicaOpenCloudNetworkingSecGroupV2DataSource_group,
			},
			resource.TestStep{
				Config: testAccTelefonicaOpenCloudNetworkingSecGroupV2DataSource_secGroupID,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkingSecGroupV2DataSourceID("data.telefonicaopencloud_networking_secgroup_v2.secgroup_1"),
					resource.TestCheckResourceAttr(
						"data.telefonicaopencloud_networking_secgroup_v2.secgroup_1", "name", "secgroup_1"),
				),
			},
		},
	})
}

func testAccCheckNetworkingSecGroupV2DataSourceID(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Can't find security group data source: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Security group data source ID not set")
		}

		return nil
	}
}

const testAccTelefonicaOpenCloudNetworkingSecGroupV2DataSource_group = `
resource "telefonicaopencloud_networking_secgroup_v2" "secgroup_1" {
        name        = "secgroup_1"
	description = "My neutron security group"
}
`

var testAccTelefonicaOpenCloudNetworkingSecGroupV2DataSource_basic = fmt.Sprintf(`
%s

data "telefonicaopencloud_networking_secgroup_v2" "secgroup_1" {
	name = "${telefonicaopencloud_networking_secgroup_v2.secgroup_1.name}"
}
`, testAccTelefonicaOpenCloudNetworkingSecGroupV2DataSource_group)

var testAccTelefonicaOpenCloudNetworkingSecGroupV2DataSource_secGroupID = fmt.Sprintf(`
%s

data "telefonicaopencloud_networking_secgroup_v2" "secgroup_1" {
	secgroup_id = "${telefonicaopencloud_networking_secgroup_v2.secgroup_1.id}"
}
`, testAccTelefonicaOpenCloudNetworkingSecGroupV2DataSource_group)
