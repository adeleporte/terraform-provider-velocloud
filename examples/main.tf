terraform {
  required_providers {
    velocloud = {
      versions = ["0.1"]
      source = "vcn.cloud/edu/velocloud"
    }
    aws = {
      source  = "hashicorp/aws"
    }
  }
}

provider velocloud {
  vco                     = "https://172.17.9.254/portal/rest"
  //username                = "supertest@velocloud.net"
  //password                = "VMware1!"
  skip_ssl_verification   = true
  //operator                = true
  token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0b2tlblV1aWQiOiJlYzJiYzQxMC0wNjIxLTRiMWQtYTk4My00YmZjMWYwY2M3OTIiLCJleHAiOjE2MzgyOTE1ODgwMDAsInV1aWQiOiI1YzgzMTMzNC0xZjZlLTExZWItODc3Ni0wMDUwNTZhZjg2MjYiLCJpYXQiOjE2MDY3NTU1OTR9.JbQqp5xWvlpABsOWu5FCrn0MGd9AAnFcBgst2oulKsY"
}
/*
data "velocloud_enterprise" "lab" {
    name = "lab"
}
*/
data "velocloud_profile" "default" {
    //enterpriseid = data.velocloud_enterprise.lab.id
    name = "Quick Start Profile"
}

resource "velocloud_edge" "edge1" {

  //enterpriseid = data.velocloud_enterprise.lab.id
  configurationid               = data.velocloud_profile.default.id
  modelnumber                   = "virtual"

  name                          = "edge-test3-bis"

  site {
    name                        = "site2"
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
/*

resource "velocloud_address_group" "test" {
  enterpriseid = data.velocloud_enterprise.lab.id
  name  =   "test2"
  description = "test2"
  

  range {
    ip = "1.1.1.1"
  }

  range {
    ip = "2.2.2.2"
  }

}

resource "velocloud_port_group" "test" {
  enterpriseid = data.velocloud_enterprise.lab.id
  name  =   "test2"
  description = "test2"
  

//  range {
//    proto       = 17
//    port_low    = 443
//    port_high   = 443
//  }

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
*/
/*
data "velocloud_enterprise" "ent" {
    name = "test"
}



data "velocloud_profile" "default" {
    enterpriseid = velocloud_enterprise.lab.id
    name = "Quick Start Profile"
}


output "out1" {
    value = data.velocloud_enterprise.ent
}

output "out2" {
    value = data.velocloud_profile.newtf
}


resource "velocloud_enterprise" "lab" {
  name        = "lab"

  user {
    username  = "lab"
    password  = "VMware1!"
    password2 = "VMware1!"
    email     = "test@vcn.cloud"
  }

}

resource "velocloud_address_group" "test" {
  name  =   "test"
  description = "test"
  enterpriseid = velocloud_enterprise.lab.id

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
  enterpriseid = velocloud_enterprise.lab.id

//  range {
//    proto       = 17
//    port_low    = 443
//    port_high   = 443
//  }

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

resource "velocloud_edge" "edge1" {

  enterpriseid                  = velocloud_enterprise.lab.id
  configurationid               = data.velocloud_profile.default.id
  modelnumber                   = "virtual"

  name                          = "edge-test3"

  site {
    name                        = "site2"
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

resource "velocloud_device_settings" "dv1" {
  enterpriseid     = velocloud_enterprise.lab.id
  profile          = velocloud_edge.edge1.edgeprofileid


  vlan {
    cidr_ip         = "1.1.1.1"
    cidr_prefix     = 24
  }


  routed_interface {
    name            = "GE3"
    cidr_ip         = "3.3.3.3"
    cidr_prefix     = 24
    gateway         = "3.3.3.254"
    netmask         = "255.255.255.0"
    type            = "STATIC"
  }

  routed_interface {
    name            = "GE4"
    cidr_ip         = "5.5.5.5"
    cidr_prefix     = 24
    gateway         = "5.5.5.254"
    netmask         = "255.255.255.0"
    type            = "STATIC"
  }

}

resource "velocloud_firewall_rules" "newtffw" {

  enterpriseid     = velocloud_enterprise.lab.id
  profile          = data.velocloud_profile.default.id

  firewall_status = true
  firewall_logging = true
  firewall_stateful = true
  firewall_syslog = true

  rule {
    name            = "rule1"
    dip             = "1.1.1.1"
    action          = "allow"
  }

  rule {
    name            = "rule2"
    sip             = "4.4.4.5"
    dip             = "3.3.3.3"
    action          = "deny"
  }


}

resource "velocloud_business_policies" "newtfbp" {

  enterpriseid     = velocloud_enterprise.lab.id
  profile          = data.velocloud_profile.default.id

  rule {
    name            = "app1"
    dip             = "1.1.1.2"
    linksteering    = "PRIVATE_WIRED"
    serviceclass    = "transactional"
    networkservice  = "fixed"
    priority        = "high"
    rxbandwidthpct  = 50
    txbandwidthpct  = 75
  }

}

/*
data "velocloud_profile" "newtf" {
    name = "newtf"
}
*/
/*
data "velocloud_profile" "tf2" {
    name = "tf2"
}
*/
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


data "velocloud_edge" "edge2" {
    name = "Antoine-HomeOffice"
}

*/

/*

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

output "edge1" {
    value = velocloud_edge.edge1
}



resource "velocloud_device_settings" "dv1" {
  #profile         = data.velocloud_profile.tf2.id
  profile          = velocloud_edge.edge1.edgeprofileid


  vlan {
    cidr_ip         = "1.1.1.1"
    cidr_prefix     = 24
  }


  routed_interface {
    name            = "GE3"
    cidr_ip         = "3.3.3.3"
    cidr_prefix     = 24
    gateway         = "3.3.3.254"
    netmask         = "255.255.255.0"
    type            = "STATIC"
  }

  routed_interface {
    name            = "GE4"
    cidr_ip         = "4.4.4.4"
    cidr_prefix     = 24
    gateway         = "4.4.4.254"
    netmask         = "255.255.255.0"
    type            = "STATIC"
  }

}

*/
/*
resource "velocloud_firewall_rules" "newtffw" {

  profile = data.velocloud_profile.newtf.id

  firewall_status = true
  firewall_logging = true
  firewall_stateful = true
  firewall_syslog = true

  rule {
    name            = "rule1"
    dip             = "1.1.1.1"
    action          = "allow"
  }

  rule {
    name            = "rule2"
    sip             = "4.4.4.4"
    dip             = "3.3.3.3"
    action          = "deny"
  }


}
*/