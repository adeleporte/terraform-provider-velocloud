---
page_title: "Deploy and configure an edge on AWS"
---

## Usage example

```hcl

provider "aws" {
  region = "us-west-1"
}

data "aws_ami" "velocloud" {
  most_recent = true

  filter {
    name   = "name"
    values = ["VeloCloud VCE 3.3.1-71-R331-20190925-GA-31ebe508-57ab-4a11-ad63-89bffdedead5-ami-0806d8f5d5558de00.4"]
  }

  owners = ["aws-marketplace"]

}

data template_file "userdata" {
  template = file("${path.module}/cloud-init.yaml")

  vars = {
    activationcode      = velocloud_edge.edge1.activationkey
  }
}

resource "aws_instance" "web" {
  ami           = data.aws_ami.velocloud.id
  instance_type = "c5.large"

  key_name = "aws-california"

  tags = {
    Name = "Velocloud"
  }

  network_interface {
      device_index = 0
      network_interface_id = "eni-09c2735c768c43583"
  }

  network_interface {
      device_index = 1
      network_interface_id = "eni-030c6ec61b0aa4b7b"
  }

  network_interface {
      device_index = 2
      network_interface_id = "eni-0962c0b9a50c1856b"
  }

  user_data = data.template_file.userdata.rendered
}

provider velocloud {
}

data "velocloud_profile" "newtf" {
    name = "newtf"
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

output "edge1" {
    value = velocloud_edge.edge1
}


```

