# edge

This datasource is used to fetch status information about a specific edge

## Example Usage

```hcl
data "velocloud_edge" "edge2" {
    name = "Antoine-HomeOffice"
}

output "my_group" {
    value = data.velocloud_edge.edge2
}

```

## Argument Reference

* `name` - (Required) Name of the port group.
* `enterpriseid` - (Optional) Enterprise ID (Customer). To be specified only if logged as an operator


## Attribute Reference

* `activationkey` - The activation key to be used to activate this edge
* `activationstate` - The activation state of this edge. Possible values are: `UNASSIGNED`, `PENDING`, `ACTIVATED`, `REACTIVATION_PENDING`.
* `edgestate` - The state of this edge. Possible values are: `NEVER_ACTIVATED`, `DEGRADED`, `OFFLINE`, `DISABLED`, `EXPIRED`, `CONNECTED`.
* `hastate` - The HA state of this edge. Possible values are: `UNCONFIGURED`, `PENDING_INIT`, `PENDING_CONFIRMATION`, `PENDING_CONFIRMED`, `PENDING_DISSOCIATION`, `READY`, `FAILED`.
* `servicestate` - The service state of this edge. Possible values are: `IN_SERVICE`, `OUT_OF_SERVICE`, `PENDING_SERVICE`.

