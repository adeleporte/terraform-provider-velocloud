# device_settings

Use this resource to configure device settings of edges

## Usage example

```hcl
resource "velocloud_device_settings" "dv1" {
  profile         = velocloud_edge.edge1.edgeprofileid

  vlan {
    cidr_ip         = "1.1.1.1"
    cidr_prefix     = 24
  }

  routed_interface {
    name            = "GE3"
    cidr_ip         = "3.3.3.3"
    cidr_prefix     = 24
    gateway         = "3.3.3.254"
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
  }

}
```

## Argument Reference

* `profile` - (Required) ID of the edge specific profile.
* `vlan` - (Optional) Configuration of VLAN.
* `routed_interface` - (Optional) List of configuration for Routed Interfaces.
* `enterpriseid` - (Optional) Enterprise ID (Customer). To be specified only if logged as an operator

### vlan reference
* `cidr_ip` - (Required) CIDR of the VLAN.
* `cidr_prefix` - (Optional) Prefix Length. Default is `24`

### routed_interface reference
* `name` - (Required) Name of the routed interface.
* `cidr_ip` - (Required) CIDR of the interface.
* `cidr_prefix` - (Optional) Prefix Length. Default is `24`
* `gateway` - (Required) Gateway of the interface.
* `netmask` - (Required) Netmask of the interface.
* `type` - (Optional) Type of addressing. Valid values are `STATIC` and `DHCP`.
* `override` - (Optional) Override status of the interface. Default is `true`


## Attribute Reference

n/a