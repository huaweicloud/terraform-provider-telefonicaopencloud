# Configure an Auto-scaling Group with Alarm

As you may know, there are many kinds of configurations for auto-scaling, one of the most important kind is to configure one group of servers which is able to scale via their status. the scenario

> End user want a server group, this group will contain several servers which number can increase or decrease according to current workload.

This example will show you how to finish the configuration combining Auto-scaling service and Cloud Eye service. to simplify this example, we recommend you to read the [docs](https://github.com/huaweicloud/terraform-provider-telefonicaopencloud/tree/master/website/docs/r) first.

Three steps will guide you to achieve this configuration. first configuration an Auto-scaling group. because Auto-scaling group depends on Auto-scaling configuration resource, configuration creation need to finish before group. please refer to the docs [configuration](https://github.com/huaweicloud/terraform-provider-telefonicaopencloud/blob/master/website/docs/r/as_configuration_v1.html.markdown), [group](https://github.com/huaweicloud/terraform-provider-telefonicaopencloud/blob/master/website/docs/r/as_group_v1.html.markdown) for more detailed information

```
resource "telefonicaopencloud_as_configuration_v1" "as_configuration"{
  scaling_configuration_name = "terraform"
  instance_config = {
    flavor = "${var.flavor}"
    image = "${var.image}"
    disk = [
      {size = 40
      volume_type = "SATA"
      disk_type = "SYS"}
    ]
    key_name = "${var.keyname}"
    user_data = "${file("userdata.txt")}"
  }
}
```

```
resource "telefonicaopencloud_as_group_v1" "as_group"{
  scaling_group_name = "terraform"
  scaling_configuration_id = "${telefonicaopencloud_as_configuration_v1.as_configuration.id}"
  desire_instance_number = 2
  min_instance_number = 0
  max_instance_number = 3
  networks = [
    {id = "${telefonicaopencloud_networking_network_v2.network.id}"},
  ]
  security_groups = [
    {id = "${telefonicaopencloud_compute_secgroup_v2.secgroup.id}"},
  ]
  vpc_id = "${telefonicaopencloud_networking_router_v2.router.id}"
  delete_publicip = true
  delete_instances = "yes" 
  depends_on = ["telefonicaopencloud_networking_router_interface_v2.int_01"]
}
```

Second, setup the alarm rule upon this Auto-scaling group, for more detail [alarm](https://github.com/huaweicloud/terraform-provider-telefonicaopencloud/blob/master/website/docs/r/ces_alarm_rule.html.markdown)

```
resource "telefonicaopencloud_ces_alarmrule" "alarm_rule" {
  "alarm_action_enabled" = "false" 
  "alarm_name" = "terraform"
  "metric" {
    "namespace" = "SYS.AS"
    "metric_name" = "cpu_util"
    "dimensions" {
        "name" = "AutoScalingGroup"
        "value" = "${telefonicaopencloud_as_group_v1.as_group.id}"
    }
  }
  "condition"  {
    "period" = 300
    "filter" = "average"
    "comparison_operator" = ">"
    "value" = 80
    "unit" = "%"
    "count" = 2
  }
}
```

To classify this is a rule for Auto-scaling group, first we should configure the "namesapce" as SYS.AS and "dimensions.name" as AutoScalingGroup. second we should define which group you would like to set up via setting "value" parameter with Auto-scaling group id.

Lastly. Configure the Auto-scaling policy with alarm type. 

```
resource "telefonicaopencloud_as_policy_v1" "as_policy"{
  scaling_policy_name = "terraform"
  scaling_group_id = "${telefonicaopencloud_as_group_v1.as_group.id}"
  scaling_policy_type= "ALARM"
  scaling_policy_action = {
    operation = "ADD"
    instance_number = 1
  }
  alarm_id="${telefonicaopencloud_ces_alarmrule.alarm_rule.id}"
}
```

You can find the full example of how to configuration policy as alarm type in [doc](https://github.com/huaweicloud/terraform-provider-telefonicaopencloud/blob/master/website/docs/r/as_policy_v1.html.markdown)
