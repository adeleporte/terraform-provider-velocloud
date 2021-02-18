# edge

This resource is used to manage edges

## Example Usage

```hcl
resource "velocloud_edge" "myedge" {

  configurationid               = data.velocloud_profile.newtf.id
  modelnumber                   = "virtual"

  name                          = "edge-test"

  site {
    name                        = "site1"
    contactname                 = "Antoine DELEPORTE"
    contactphone                = "+331010101010"
    contactmobile               = "+336010101010"
    contactemail                = "adeleporte@vmware.com"
    streetaddress               = "Terrasse Boildieu"
    city                        = "Paris"
    country                     = "France"

    shippingstreetaddress       = "Terrasse Boildieu"
    shippingcity                = "Paris"
    shippingcountry             = "France"
  
    lat                         = 10.4567
    lon                         = 20.23

    shippingsameaslocation      = true
  }

output "my_group" {
    value = velocloud_edge.myedge
}

```

## Argument Reference

* `name` - (Required) Name of the edge.
* `configurationid` - (Required) Profile to associate the edge with.
* `modelnumber` - (Required) Model of the edge. Valid values are: `edge500`, `edge5X0`, `edge510`, `edge510lte`, `edge6X0`, `edge840`, `edge1000`, `edge1000qat`,`edge3X00`, `edge1000qat`, `virtual`.
* `enterpriseid` - (Optional) Enterprise ID (Customer). To be specified only if logged as an operator

### Site Reference

* `name` - (Required) Name of the site.
* `contactname` - (Optional) Name of the contact.
* `contactphone` - (Optional) Phone number of the contact.
* `contactmobile` - (Optional) Mobile number of the contact.
* `contactemail` - (Optional) Email of the contact.
* `streetaddress` - (Optional) Street Address.
* `streetaddress2` - (Optional) Street Address2.
* `city` - (Optional) City.
* `state` - (Optional) State.
* `country` - (Optional) Country.
* `postalcode` - (Optional) Postal Code.
* `lat` - (Optional) Latitude of the edge (float64).
* `lon` - (Optional) Longitude of the edge (float64).
* `timezone` - (Optional) Timezone.
* `locale` - (Optional) Locale.
* `shippingsameaslocation` - (Optional) Same address for shipping (Boolean).
* `shippingcontactname` - (Optional) Contact name for shipping.
* `shippingstreetaddress` - (Optional) Street Address of shipping.
* `shippingstreetaddress2` - (Optional) Street Address2 of shipping.
* `shippingcity` - (Optional) Shipping City.
* `shippingcountry` - (Optional) Shipping Country.
* `shippingpostalcode` - (Optional) Shipping Postal Code.


## Attribute Reference

* `activationkey` - The activation key to be used to activate this edge
* `activationstate` - The activation state of this edge. Possible values are: `UNASSIGNED`, `PENDING`, `ACTIVATED`, `REACTIVATION_PENDING`.
* `edgestate` - The state of this edge. Possible values are: `NEVER_ACTIVATED`, `DEGRADED`, `OFFLINE`, `DISABLED`, `EXPIRED`, `CONNECTED`.
* `hastate` - The HA state of this edge. Possible values are: `UNCONFIGURED`, `PENDING_INIT`, `PENDING_CONFIRMATION`, `PENDING_CONFIRMED`, `PENDING_DISSOCIATION`, `READY`, `FAILED`.
* `servicestate` - The service state of this edge. Possible values are: `IN_SERVICE`, `OUT_OF_SERVICE`, `PENDING_SERVICE`.


