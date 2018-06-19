package telefonicaopencloud

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/huaweicloud/golangsdk/openstack/smn/v2/topics"
)

func TestAccSMNV2Topic_basic(t *testing.T) {
	var topic topics.TopicGet

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSMNTopicV2Destroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: TestAccSMNV2TopicConfig_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSMNV2TopicExists("telefonicaopencloud_smn_topic_v2.topic_1", &topic),
					resource.TestCheckResourceAttr(
						"telefonicaopencloud_smn_topic_v2.topic_1", "name", "topic_1"),
					resource.TestCheckResourceAttr(
						"telefonicaopencloud_smn_topic_v2.topic_1", "display_name",
						"The display name of topic_1"),
				),
			},
			resource.TestStep{
				Config: TestAccSMNV2TopicConfig_update,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"telefonicaopencloud_smn_topic_v2.topic_1", "display_name",
						"The update display name of topic_1"),
					resource.TestCheckResourceAttr(
						"telefonicaopencloud_smn_topic_v2.topic_1", "name", "topic_1"),
				),
			},
		},
	})
}

func testAccCheckSMNTopicV2Destroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)
	smnClient, err := config.SmnV2Client(OS_REGION_NAME)
	if err != nil {
		return fmt.Errorf("Error creating telefonicaopencloud smn: %s", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "telefonicaopencloud_smn_topic_v2" {
			continue
		}

		_, err := topics.Get(smnClient, rs.Primary.ID).Extract()
		if err == nil {
			return fmt.Errorf("Topic still exists")
		}
	}

	return nil
}

func testAccCheckSMNV2TopicExists(n string, topic *topics.TopicGet) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		config := testAccProvider.Meta().(*Config)
		smnClient, err := config.SmnV2Client(OS_REGION_NAME)
		if err != nil {
			return fmt.Errorf("Error creating telefonicaopencloud smn client: %s", err)
		}

		found, err := topics.Get(smnClient, rs.Primary.ID).ExtractGet()
		if err != nil {
			return err
		}

		if found.TopicUrn != rs.Primary.ID {
			return fmt.Errorf("Topic not found")
		}

		*topic = *found

		return nil
	}
}

var TestAccSMNV2TopicConfig_basic = `
resource "telefonicaopencloud_smn_topic_v2" "topic_1" {
  name		  = "topic_1"
  display_name    = "The display name of topic_1"
}
`

var TestAccSMNV2TopicConfig_update = `
resource "telefonicaopencloud_smn_topic_v2" "topic_1" {
  name		  = "topic_1"
  display_name    = "The update display name of topic_1"
}
`
