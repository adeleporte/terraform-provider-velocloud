# Terraform Velocloud Provider

This is the repository for the Terraform Velocloud Provider, which one can use with
Terraform to work with [VMware Velocloud SDWAN][vmware-nsxt].

[vmware-nsxt]: https://www.vmware.com/fr/products/sd-wan-by-velocloud.html

For general information about Terraform, visit the [official
website][tf-website] and the [GitHub project page][tf-github].

[tf-website]: https://terraform.io/
[tf-github]: https://github.com/hashicorp/terraform


Documentation on the Velocloud platform can be found at the [VMware SD-WAN Documentation page](https://www.vmware.com/fr/products/sd-wan-by-velocloud.html)


# Using the Provider

The latest version of this provider requires Terraform v0.12 or higher to run.

The VMware supported version of the provider requires Velocloud version 3.4 onwards and Terraform 0.12 onwards.


Note that you need to run `terraform init` to fetch the provider before
deploying. Read about the provider split and other changes to TF v0.10.0 in the
official release announcement found [here][tf-0.10-announce].

[tf-0.10-announce]: https://www.hashicorp.com/blog/hashicorp-terraform-0-10/

### Controlling the provider version

Note that you can also control the provider version. This requires the use of a
`provider` block in your Terraform configuration if you have not added one
already.

The syntax is as follows:

```hcl
provider "velocloud" {
  version = "~> 1.0"
  ...
}
```


Version locking uses a pessimistic operator, so this version lock would mean
anything within the 1.x namespace, including or after 1.0.0. [Read
more][provider-vc] on provider version control.

[provider-vc]: https://www.terraform.io/docs/configuration/providers.html#provider-versions

# Installation (manual)


**NOTE:** Recommended way to compile the provider is using [Go Modules](https://blog.golang.org/using-go-modules), however vendored dependencies are still supported.

**NOTE:** For terraform 0.13, please refer to [provider installation configuration][install-013] in order to use custom provider.
[install-013]: https://www.terraform.io/docs/commands/cli-config.html#provider-installation


## Cloning the Project

First, you will want to clone the repository to
`$GOPATH/src/github.com/adeleporte/terraform-provider-velocloud`:

```sh
mkdir -p $GOPATH/src/github.com/adeleporte
cd $GOPATH/src/github.com/adeleporte
git clone https://github.com/adeleporte/terraform-provider-velocloud.git
```

## Building and Installing the Provider

After the clone has been completed, you can enter the provider directory and build the provider.

```sh
cd $GOPATH/src/github.com/adeleporte/terraform-provider-velocloud
make
```

After the build is complete, if your terraform running folder does not match your GOPATH environment, you need to copy the `terraform-provider-velocloud` executable to your running folder and re-run `terraform init` to make terraform aware of your local provider executable.

After this, your project-local `.terraform/plugins/ARCH/lock.json` (where `ARCH`
matches the architecture of your machine) file should contain a SHA256 sum that
matches the local plugin. Run `shasum -a 256` on the binary to verify the values
match.

# Usage

In order to use the Velocloud Terraform provider you must first configure the provider to communicate with the Velocloud Orchestrator. The Velocloud Orchestrator is the system which serves the VMware SDWAN REST API and provides a way to configure the desired state of the Velocloud system. The configuration of the Velocloud provider requires the IP address, hostname, or FQDN of the Velocloud Orchestrator (VCO).

The Velocloud provider offers several ways to authenticate to the VCO. Credentials can be provided statically or provided as environment variables.

Setting the `operator` parameter to true will enable the interaction with the VCO as an operator. This authentication level is mandatory to create some resources (i.e customer, gateways...)


## Example of Provider Configuration


```hcl
provider "velocloud" {
  host     = "vco.vcn.net"
  username = "admin"
  password = "default"
  operator = false
}
```

```hcl
provider "velocloud" {
  host     = "vco.vcn.net"
  token    = "jhkjlhjkhjkhjkhjkhhjkhkjhkjhjkhkjhkjhjkhjkh"
}
```

## Example of Setting Environment Variables

```sh
export VCO_HOST     ="vco.vcn.cloud"
export VCO_USERNAME ="admin"
export VCO_PASSWORD ="default"
```


# Velocloud Provider Resources

## Enterprise Resource
An Enterprise is a Velocloud tenant and is the first resource to be managed.


### Usage example

```hcl
resource "velocloud_enterprise" "tf_customer1" {
  name      = "My-First-Customer"
  username  = "adeleporte"
  password  = "my-password"
  email     = "adeleporte@vmware.com"
}
```

### Argument reference

Name | Required | Description
------------ | ------------- | -------------
`name` | yes | The enterprise name
`username` | yes | The username used to manage this enterprise
`password` | yes | The password
`email` | yes | The email of the username managing this enterprise


### Attribute reference
In addition to arguments listed above, the following attributes are exported:

Name | Description
------------ | ------------- 
`id` | ID of the entreprise


## Edge Resource
An Edge is the Velocloud Appliance (virtual or physical) connecting the network to the SDWAN


### Usage example

```hcl
resource "velocloud_edge" "tf_edge_melbourne" {
  name              = "Edge-Melbourne"
  model             = "virtual"
  contact           = "adeleporte"
  email             = "adeleporte@vmware.com"
  lat               = 20.1234
  lon               = 30.1234
  enterpriseid      = data.tf_customer1.id
  configurationid   = data.tf_businessprofile1.id
}
```

### Argument reference

Name | Required | Description
------------ | ------------- | -------------
`name` | yes | The enterprise name
`model` | yes | The username used to manage this enterprise
`contact` | yes | The password
`email` | yes | The email of the username managing this enterprise
`lat` | no | The geographical latitude of the edge
`lon` | no | The geographical longitude of the edge
`enterpriseid` | yes | The enterprise associated to the edge
`configurationid` | yes | The business profile to associate the edge with


### Attribute reference
In addition to arguments listed above, the following attributes are exported:

Name | Description
------------ | ------------- 
`id` | ID of the edge
`activationkey` | The activation key needed to activate the edge onsite
`devicemodule` | ID of the Device Settings module of the edge
`qosmodule` | ID of the Business Profile module of the edge
`firewallmodule` | ID of the Firewall module of the edge
`wanmodule` | ID of the Overlay module of the edge

## Business Policy Resource
VMware SD-WAN provides an enhanced Quality of Service feature called Business Policy. This feature is defined using the Business Policy tab in a Profile or at the Edge override level.
Based on the business policy configuration, VMware SD-WAN examines the traffic being used, identifies the Application behavior, the business service objective required for a given app (High, Med, or Low), and the Edge WAN Link conditions. Based on this, the Business Policy optimizes Application behavior driving queuing, bandwidth utilization, link steering, and the mitigation of network errors.

### Usage example

```hcl
resource "velocloud_qos_rule" "tf_business_policy" {
  enterpriseid      = data.tf_customer1.id
  configurationid   = data.velocloud_configuration.tf_config.id
  qosmodule         = data.velocloud_configuration.tf_config.qosmodule

  rule {
    name            = "rule1"
    dip             = "1.1.1.1"
    dmask           = "255.255.255.255"
    dport           = 80
    proto           = 6
    bandwidthpct    = 20
    class           = "transactional"
    servicegroup    = "ALL"
  }

  rule {
    name            = "rule2"
    dip             = "2.2.2.0"
    dmask           = "255.255.255.0"
    dport           = 443
    proto           = 6
    bandwidthpct    = 80
    class           = "realtime"
    servicegroup    = "PRIVATE_WIRED"
  }
```

### Argument reference

Name | Required | Description
------------ | ------------- | -------------
`enterpriseid` | yes | The enterprise ID
`configurationid` | yes | The profile ID
`qosmodule` | yes | The Business Policy module ID

#### Rule reference

Name | Required | Description | Default
------------ | ------------- | ------------- | -------------
`name` | yes | The name of the rule | n/a
`dip` | no | The destination subnet | "any"
`dmask` | no | The destination mask | "255.255.255.255" (match IP)
`dport` | no | The destination port | "-1" (All ports)
`proto` | no | The destination protocol | "-1" (All protocol) (tcp=6)
`bandwidthpct` | no | The bandwidth rate limiting | "-1" (No limit) (percent, ingress and egress)
`routeaction` | no | Use overlay or direct | "edge2Cloud" (edge2Cloud or gateway)
`class` | no | QoS service class | "bulk" (bulk, transactional or realtime)
`servicegroup` | no | Link steering policy | "ALL" (ALL, PUBLIC_WIRED or PRIVATE_WIRED)


## Firewall Rule Resource
SD-WAN Orchestrator allows you to configure Firewall rules at the Profile and Edge levels to allow, drop, reject, or skip inbound and outbound traffic. The firewall uses the parameters such as source IP address/port, destination IP address/port, applications, application categories, and DSCP tags to create firewall rules.

### Usage example

```hcl
resource "velocloud_firewall_rule" "tf_firewall_policy" {
  enterpriseid      = data.tf_customer1.id
  configurationid    = data.velocloud_configuration.tf_config.id
  fwmodule          = data.velocloud_configuration.tf_config.fwmodule

  enabled           = true
  logging           = true

  rule {
    name            = "rule1"
    sip             = "1.1.1.1"
    smask           = "255.255.255.255"
    dip             = "2.2.2.2"
    dmask           = "255.255.255.255"
    sport           = -1
    dport           = 80
    proto           = 6
    action          = "allow"
  }

  rule {
    name            = "rule2"
    sip             = "3.3.3.3"
    smask           = "255.255.255.255"
    dip             = "4.4.4.4"
    dmask           = "255.255.255.255"
    sport           = -1
    dport           = 443
    proto           = 6
    action          = "deny"
  }
```

### Argument reference

Name | Required | Description
------------ | ------------- | -------------
`enterpriseid` | yes | The enterprise ID
`configurationid` | yes | The profile ID
`fwmodule` | yes | The Firewall Policy module ID

#### Rule reference

Name | Required | Description | Default
------------ | ------------- | ------------- | -------------
`name` | yes | The name of the rule | n/a
`sip` | no | The source subnet | "any"
`smask` | no | The source mask | "255.255.255.255" (match IP)
`dip` | no | The destination subnet | "any"
`dmask` | no | The destination mask | "255.255.255.255" (match IP)
`sport` | no | The source port | "-1" (All ports)
`dport` | no | The destination port | "-1" (All ports)
`proto` | no | The destination protocol | "-1" (All protocol) (tcp=6)
`action` | yes | The firewall action | n/a (allow or deny)

# Velocloud Provider Datasources

## Enterprise Datasource
An Enterprise is a Velocloud tenant.


### Usage example

```hcl
data "velocloud_enterprise" "tf_customer1" {
  name      = "My-First-Customer"
  username  = "adeleporte"
  password  = "my-password"
  email     = "adeleporte@vmware.com"
}
```

### Argument reference

Name | Required | Description
------------ | ------------- | -------------
`name` | yes | The enterprise name
`username` | yes | The username used to manage this enterprise
`password` | yes | The password
`email` | yes | The email of the username managing this enterprise


### Attribute reference
In addition to arguments listed above, the following attributes are exported:

Name | Description
------------ | ------------- 
`id` | ID of the entreprise


## Profile datasource
The velocloud profile contains Device Settings, Business Policy, Firewall and Overlay modules configuration, and is applied to multiples Edges.


### Usage example

```hcl
data "velocloud_configuration" "tf_configuration1" {
  enterpriseid      = data.velocloud_enterprise.tf_customer1.id
  name              = "Quick Start Profile"
}
```

### Argument reference

Name | Required | Description
------------ | ------------- | -------------
`enterpriseid` | yes | The enterprise ID
`name` | yes | The profile we want to fetch



### Attribute reference
In addition to arguments listed above, the following attributes are exported:

Name | Description
------------ | ------------- 
`id` | ID of the profile
`devicesettings` | ID of the device settings module
`firewall` | ID of the firewall module
`qosmodule` | ID of the Business Policy module
`wan` | ID of the Overlay module