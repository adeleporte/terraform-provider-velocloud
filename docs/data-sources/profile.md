---
page_title: "Get profiles"
---

# Datasource: profile

Profiles provide a composite of the configurations created in Networks and Network Services. It also adds configuration for Business Policy and Firewall rules. Use this datasource to get the ID of the profile, to be used as a reference by other resources (i.e business_policies)

## Example Usage

```hcl
data "velocloud_profile" "newtf" {
    name = "newtf"
}

output "my_application" {
    value = data.velocloud_profile.newtf.id
}

```

## Argument Reference

* `name` - (Required) Name of the profile.

## Attribute Reference

* `id` - ID to be used as a reference by other resources.