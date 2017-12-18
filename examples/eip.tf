resource "telefonicaopencloud_vpc_eip_v1" "eip" {
    publicip {
        type = "5_bgp"
	ip_address = "170.84.209.207"
    }
    bandwidth {
        name = "test4terra"
        size = 8
        share_type = "PER"
        charge_mode = "traffic"
    }
}
