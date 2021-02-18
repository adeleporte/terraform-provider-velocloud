# firewall_rules
VMware SD-WAN orchestrator supports configuration of stateless and stateful firewalls for profiles and edges.
Use this resource to configure firewall rules and enable or disable firewall status and logs.

## Usage example

```hcl
resource "velocloud_firewall_rules" "newtffw" {

  profile = data.velocloud_profile.newtf.id

  firewall_status = true
  firewall_logging = true
  firewall_stateful = true
  firewall_syslog = true

  rule {
    name            = "rule1"
    dip             = "1.1.1.1"
    action          = "allow"
  }

  rule {
    name            = "rule2"
    sip             = "4.4.4.4"
    dip             = "3.3.3.3"
    action          = "deny"
  }
```

## Argument Reference

* `profile` - (Required) ID of the profile.
* `segment` - (Optional) ID of the segment.
* `rule` - (Required) List of business policies.
* `firewall_status` - (Optional) Status of the firewall. Default to `true`
* `firewall_stateful` - (Optional) Indicates if firewalling should be stateful. Default to `false`
* `firewall_logging` - (Optional) Status of firewall logging. Default to `false`
* `firewall_syslog` - (Optional) Indicates if firewall should send logs as syslog. Default to `false`
* `enterpriseid` - (Optional) Enterprise ID (Customer). To be specified only if logged as an operator

### Rule reference
* `name` - (Required) Name of the rule.

* `sip` - (Optional) Source IP of the rule.
* `ssm` - (Optional) Source Mask of the rule.
* `sport_low` - (Optional) Source Low Port of the rule.
* `sport_high` - (Optional) Source High Port of the rule.
* `s_address_group` - (Optional) Source Address Group of the rule. Use address group datasource or resource to get the logicalid
* `s_port_group` - (Optional) Source Port Group of the rule. Use port group datasource or resource to get the logicalid
* `svlan` - (Optional) Source VLAN to be matched.
* `smac` - (Optional) Source MAC address to be matched.
* `s_rule_type` - (Optional) Type of source object to be matched.

* `dip` - (Optional) Destination IP of the rule.
* `dsm` - (Optional) Destination Mask of the rule.
* `d_rule_type` - (Optional) Type of destination object to be matched.
* `dport_low` - (Optional) Destination Low Port of the rule.
* `dport_high` - (Optional) Destination High Port of the rule.
* `appid` - (Optional) Application to be matched. Use application datasource to get the ID of a specifig application
* `proto` - (Optional) Destination Protocol of the rule.
* `d_address_group` - (Optional) Destination Address Group of the rule. Use address group datasource or resource to get the logicalid
* `d_port_group` - (Optional) Destination Port Group of the rule. Use port group datasource or resource to get the logicalid
* `classid` - (Optional) Class to be matched.
* `dscp` - (Optional) Dscp to be matched.
* `dvlan` - (Optional) Destnation VLAN to be matched.
* `hostname` - (Optional) Hostname to be matched.
* `os_version` - (Optional) OS Version to be matched.



## Attribute Reference

n/a