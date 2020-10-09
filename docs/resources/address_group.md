---
page_title: "Manage address groups"
---

# Resource: address_group

An address group consists of a range of IP addresses. When you create business policies and firewall rules, you can define the rules for a range of IP addresses, by including the address groups in the rule definitions.

You can create Address groups to save the range of valid IP addresses. You can simplify the policy management by creating object groups of specific types and reusing them in policies and rules.

Using Object Groups, you can:

Manage policies easily
Modularize and reuse the policy components
Update all referenced business and firewall policies easily
Reduce the number of policies
Improve the policy debugging and readability

## Example Usage

```hcl
resource "velocloud_address_group" "my_group" {
  name  =   "my_group"
  description = "my_group_desc"

  range {
    ip          = "1.1.1.1"
  }

  range {
    ip          = "2.2.2.0"
    mask        = "255.255.255.0"
    rule_type   = "netmask"
  }
}

output "my_group" {
    value = velocloud_address_group.my_group.logicalid
}

```

## Argument Reference

* `name` - (Required) Name of the address group.
* `description` - (Optional) Description of the address group.
* `range` - (Optional) List of address group ranges.

### Range argument Reference
* `ip` - (Required) IP or Subnet to match (ie 1.1.1.1 or 1.1.0.0).
* `mask` - (Optional) Mask to match (ie 255.255.255.0).
* `rule_type` - (Optional) Type of rule. Valid values: `exact`, `subnet`, `prefix`, `wildcard`.

## Attribute Reference

* `logicalid` - ID to be used as a reference for other resources.