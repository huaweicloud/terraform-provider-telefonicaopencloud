package telefonicaopencloud

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccTelefonicaopenCloudRdsFlavorV1DataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccTelefonicaopenCloudRdsFlavorV1DataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRdsFlavorV1DataSourceID("data.telefonicaopencloud_rds_flavors_v1.flavor"),
					resource.TestCheckResourceAttrSet(
						"data.telefonicaopencloud_rds_flavors_v1.flavor", "name"),
					resource.TestCheckResourceAttrSet(
						"data.telefonicaopencloud_rds_flavors_v1.flavor", "id"),
					resource.TestCheckResourceAttrSet(
						"data.telefonicaopencloud_rds_flavors_v1.flavor", "speccode"),
				),
			},
		},
	})
}

func TestAccTelefonicaopenCloudRdsFlavorV1DataSource_speccode(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccTelefonicaopenCloudRdsFlavorV1DataSource_speccode,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkingNetworkV2DataSourceID("data.telefonicaopencloud_rds_flavors_v1.flavor"),
					resource.TestCheckResourceAttr(
						"data.telefonicaopencloud_rds_flavors_v1.flavor", "name", "OTC_PGCM_XLARGE"),
					resource.TestCheckResourceAttr(
						"data.telefonicaopencloud_rds_flavors_v1.flavor", "speccode", "rds.pg.s1.xlarge"),
				),
			},
		},
	})
}

func testAccCheckRdsFlavorV1DataSourceID(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Can't find rds data source: %s ", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Rds data source ID not set ")
		}

		return nil
	}
}

var testAccTelefonicaopenCloudRdsFlavorV1DataSource_basic = `

data "telefonicaopencloud_rds_flavors_v1" "flavor" {
    region = "sa-chile-1"
	datastore_name = "PostgreSQL"
    datastore_version = "9.5.5"
}
`

var testAccTelefonicaopenCloudRdsFlavorV1DataSource_speccode = `

data "telefonicaopencloud_rds_flavors_v1" "flavor" {
    region = "sa-chile-1"
	datastore_name = "PostgreSQL"
    datastore_version = "9.5.5"
    speccode = "rds.pg.s1.xlarge"
}
`
