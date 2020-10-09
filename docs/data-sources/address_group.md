# address_group

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
data "velocloud_address_group" "my_group" {
  name  =   "my_group"
}

output "my_group" {
    value = data.velocloud_address_group.my_group.logicalid
}

```

## Argument Reference

* `name` - (Required) Name of the address group.

## Attribute Reference

* `logicalid` - ID to be used as a reference by other resources.