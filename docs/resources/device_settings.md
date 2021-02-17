# device_settings

Use this resource to configure device settings of edges

## Usage example

```hcl
resource "velocloud_device_settings" "dv1" {
  profile         = velocloud_edge.edge1.edgeprofileid

  vlan {
    cidr_ip         = "1.1.1.1"
    cidr_prefix     = 24
    advertise       = true
    override        = false
    dhcp_enabled    = false
  }

  routed_interface {
    name            = "GE3"
    cidr_ip         = "3.3.3.3"
    cidr_prefix     = 24
    netmask         = "255.255.255.0"
    type            = "STATIC"
  }

  routed_interface {
    name            = "GE4"
    cidr_ip         = "4.4.4.4"
    cidr_prefix     = 24
    gateway         = "4.4.4.254"
    netmask         = "255.255.255.0"
    type            = "STATIC"
    nat_direct      = true
    wan_overlay     = false
  }

  static_route {
    subnet_cidr_ip     = "8.8.8.8"
    subnet_cidr_prefix = 32
    next_hop           = "1.2.3.4"
    interface          = "GE3"
  }

  static_route {
    subnet_cidr_ip     = "1.1.1.1"
    subnet_cidr_prefix = 24
    next_hop           = "4.4.4.4"
    interface          = "GE4"
    cost               = 10
    advertise          = false
  }

}
```

## Argument Reference

- `profile` - (Required) ID of the edge specific profile.
- `vlan` - (Optional) Configuration of VLAN.
- `routed_interface` - (Optional) List of configuration for Routed Interfaces.
- `static_route` - (Optional) List of configuration for static routes.

### vlan reference

- `cidr_ip` - (Required) CIDR of the VLAN.
- `cidr_prefix` - (Optional) Prefix Length. Default is `24`
- `advertise` - (Optional) Whether to advertise the VLAN. Default is `true`
- `override` - (Optional) Whether to use Edge-specific override for VLAN. Default is `false`
- `dhcp_enabled` - (Optional) Whether DHCP is enabled on VLAN. Default is `true`

### routed_interface reference

- `name` - (Required) Name of the routed interface.
- `cidr_ip` - (Required) CIDR of the interface.
- `cidr_prefix` - (Optional) Prefix Length. Default is `24`
- `gateway` - (Optional) Gateway of the interface.
- `netmask` - (Required) Netmask of the interface.
- `type` - (Optional) Type of addressing. Valid values are `STATIC` and `DHCP`.
- `override` - (Optional) Override status of the interface. Default is `true`
- `nat_direct` - (Optional) Send traffic via WAN interface, bypassing SD-WAN gateway. Default is `true`
- `wan_overlay` - (Optional) Use auto-detected WAN overlay. Default is `true`

### static_route reference

- `subnet_cidr_ip` - (Required) CIDR IP of subnet to route.
- `subnet_cidr_prefix` - (Required) CIDR prefix of subnet to route.
- `next_hop` - (Required) IP address of next hop.
- `interface` - (Required) Name of interface to route to.
- `cost` - (Optional) Metric cost of route. Default is `0`
- `preferred` - (Optional) Set route to preferred. Default is `true`
- `advertise` - (Optional) Advertise the route. Default is `true`

## Attribute Reference

n/a
