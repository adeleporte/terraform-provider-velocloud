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
}

output "my_group" {
    value = data.velocloud_port_group.my_group.logicalid
}

```

## Argument Reference

* `name` - (Required) Name of the port group.
* `enterpriseid` - (Optional) Enterprise ID (Customer). To be specified only if logged as an operator


## Attribute Reference

* `logicalid` - ID to be used as a reference by other resources.