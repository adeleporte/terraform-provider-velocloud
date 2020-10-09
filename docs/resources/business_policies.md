---
page_title: "Manage business policies"
---

# Resource: business_policies
VMware SD-WAN provides an enhanced Quality of Service feature called Business Policy. This feature is defined using the Business Policy tab in a Profile or at the Edge override level.
Based on the business policy configuration, VMware SD-WAN examines the traffic being used, identifies the Application behavior, the business service objective required for a given app (High, Med, or Low), and the Edge WAN Link conditions. Based on this, the Business Policy optimizes Application behavior driving queuing, bandwidth utilization, link steering, and the mitigation of network errors.

## Usage example

```hcl
resource "velocloud_business_policies" "newtfbp" {

  profile = data.velocloud_profile.newtf.id

  rule {
    name            = "app1"
    dip             = "1.1.1.1"
    linksteering    = "PRIVATE_WIRED"
    serviceclass    = "transactional"
    networkservice  = "fixed"
    priority        = "high"
    rxbandwidthpct  = 50
    txbandwidthpct  = 75
  }

  rule {
    name            = "app2"
    saddressgroup   = data.velocloud_address_group.silver.logicalid
    daddressgroup   = data.velocloud_address_group.gold.logicalid
    sportgroup      = data.velocloud_port_group.tcp22.id
    dportgroup      = data.velocloud_port_group.tcp22.id
    linksteering    = "ALL" // or PUBLIC_WIRED, PRIVATE_WIRED
    serviceclass    = "realtime" // or bulk, transactional
    networkservice  = "auto" // or fixed
    priority        = "low" // or high, normal
  }

    rule {
    name            = "app3"
    appid           = data.velocloud_application.bittorrent.id
    sportgroup      = velocloud_port_group.test.logicalid
    dportgroup      = velocloud_port_group.test.logicalid
    linksteering    = "PRIVATE_WIRED"
    serviceclass    = "transactional"
    networkservice  = "auto"
    priority        = "high"
    rxbandwidthpct  = 50
    txbandwidthpct  = 75
  }

}
```

## Argument Reference

* `profile` - (Required) ID of the profile.
* `rule` - (Optional) List of business policies.

### Rule reference
* `name` - (Required) Name of the rule.
* `appid` - (Optional) Application to be matched. Use application datasource to get the ID of a specifig application
* `dip` - (Optional) Destination IP of the rule.
* `dsm` - (Optional) Destination Mask of the rule.
* `dport_low` - (Optional) Destination Low Port of the rule.
* `dport_high` - (Optional) Destination High Port of the rule.
* `proto` - (Optional) Destination Protocol of the rule.
* `saddressgroup` - (Optional) Source Address Group of the rule. Use address group datasource or resource to get the logicalid
* `daddressgroup` - (Optional) Destination Address Group of the rule. Use address group datasource or resource to get the logicalid
* `sportgroup` - (Optional) Source Port Group of the rule. Use port group datasource or resource to get the logicalid
* `dportgroup` - (Optional) Destination Port Group of the rule. Use port group datasource or resource to get the logicalid
* `priority` - (Optional) Priority of the rule. Valid values: `high`, `normal`, `high`.
* `serviceclass` - (Optional) Service Class of the rule. Valid values: `realtime`, `transactional`, `bulk`.
* `networkservice` - (Optional) Network service policy of the rule. Valid values: `auto` (DMPO), `fixed` (Direct).
* `transportgroup` - (Optional) Transport Group policy of the rule. Valid values: `ALL`, `PUBLIC_WIRED`, `PRIVATE_WIRED`, `PUBLIC_WIRELESS`.
* `rxbandwidthpct` - (Optional) Inbound Rate limit of the rule. Valid values: `0-100`.
* `txbandwidthpct` - (Optional) Outbound Rate limit of the rule. Valid values: `0-100`.

## Attribute Reference

n/a