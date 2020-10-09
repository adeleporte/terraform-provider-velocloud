---
page_title: "Get applications"
---

# Datasource: application

VMware SD-WAN provides an initial Application Map with possible applications.
Use this datasource to get the ID of a specific application. This ID can be used as a reference in other resources.

## Example Usage

```hcl
data "velocloud_application" "bittorrent" {
    name = "bittorrent"
}

output "my_application" {
    value = data.velocloud_application.bittorrent.id
}

```

## Argument Reference

* `name` - (Required) Name of the address group.

## Attribute Reference

* `id` - ID to be used as a reference by other resources.