# port_group

An port group consists of a range of TCP/UDP ports. When you create business policies and firewall rules, you can define the rules for a range of TCP/UDP ports, by including the port groups in the rule definitions.

You can create Port groups to save the range of valid TCP/UDP ports. You can simplify the policy management by creating object groups of specific types and reusing them in policies and rules.

Using Object Groups, you can:

Manage policies easily
Modularize and reuse the policy components
Update all referenced business and firewall policies easily
Reduce the number of policies
Improve the policy debugging and readability

## Example Usage

```hcl
resource "velocloud_port_group" "my_group" {
  name  =   "my_group"
  description = "my_group_desc"

  range {
    proto           = 6
    port_low        = 80
    port_high       = 80
  }

  range {
    proto           = 17
    port_low        = 443
    port_high       = 443
  }
}

output "my_group" {
    value = velocloud_port_group.my_group.logicalid
}

```

## Argument Reference

* `name` - (Required) Name of the port group.
* `description` - (Optional) Description of the port group.
* `range` - (Optional) List of port group ranges.
* `enterpriseid` - (Optional) Enterprise ID (Customer). To be specified only if logged as an operator

### Range argument Reference
* `proto` - (Required) IP Protocol to match. Valid values: `6` (TCP), `17` (UDP).
* `port_low` - (Required) Low port match. Valid values: `1-65535`
* `port_high` - (Required) High port. Valid values: `1-65535`.

## Attribute Reference

* `logicalid` - ID to be used as a reference for other resources.