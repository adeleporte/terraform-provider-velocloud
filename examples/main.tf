terraform {
  required_providers {
    velocloud = {
      versions = ["0.1"]
      source = "vcn.cloud/edu/velocloud"
    }
  }
}

provider velocloud {
}

data "velocloud_profile" "newtf" {
    name = "newtf"
}

/*
data "velocloud_address_group" "gold" {
    name = "Gold"
}

data "velocloud_address_group" "silver" {
    name = "Silver"
}

data "velocloud_port_group" "tcp22" {
    name = "tcp22"
}

data "velocloud_application" "bittorrent" {
    name = "bittorrent"
}

data "velocloud_edge" "edge1" {
    name = "Antoine-HomeOffice"
}

resource "velocloud_address_group" "test" {
  name  =   "test"
  description = "test"

  range {
    ip = "1.1.1.1"
  }

  range {
    ip = "2.2.2.2"
  }
}

resource "velocloud_port_group" "test" {
  name  =   "test"
  description = "test"

  range {
    proto       = 17
    port_low    = 443
    port_high   = 443
  }

  range {
    proto       = 6
    port_low    = 443
    port_high   = 9443
  }

    range {
    proto       = 6
    port_low    = 80
    port_high   = 80
  }
}


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


output "newtf" {
    value = data.velocloud_profile.newtf
}

output "newtfbp" {
    value = velocloud_business_policies.newtfbp
}

output "edge1" {
    value = data.velocloud_edge.edge1
}

*/

data "velocloud_edge" "edge2" {
    name = "Antoine-HomeOffice"
}


resource "velocloud_edge" "edge1" {

  configurationid               = data.velocloud_profile.newtf.id
  modelnumber                   = "virtual"

  name                          = "edge-test3"

  site {
    name                        = "site1"
    contactname                 = "Antoine DELEPORTE2"
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

  
}
