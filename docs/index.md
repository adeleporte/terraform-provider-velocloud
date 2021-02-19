# Velocloud Provider

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

# Installation (automatic)

To install this provider, copy and paste this code into your Terraform configuration. Then, run terraform init.

```hcl
terraform {
  required_providers {
    velocloud = {
      source = "adeleporte/velocloud"
    }
  }
}

provider "velocloud" {
  vco       = "https://vco.vcn.cloud/portal/rest"
  token     = "my-token"
}
```

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

In order to use the Velocloud Terraform provider you must first configure the provider to communicate with the Velocloud Orchestrator. The Velocloud Orchestrator is the system which serves the VMware SDWAN REST API and provides a way to configure the desired state of the Velocloud system.

## API Key
The configuration of the Velocloud provider requires the hostname and a API token of the Velocloud Orchestrator (VCO).

## Username/Password
The configuration of the Velocloud provider requires the hostname of the Velocloud Orchestrator (VCO) + a login/password

## Operator level
The configuration of the Velocloud provider requires the hostname of the Velocloud Orchestrator (VCO) + a login/password and operator=true


## Example of Provider Configuration

```hcl
provider "velocloud" {
  vco       = "https://vco.vcn.cloud/portal/rest"
  token     = "my-token"
}
```

## Example of Provider Configuration with Username / Password Access

```hcl
provider "velocloud" {
  vco       = "https://vco.vcn.cloud/portal/rest"

  username                = "supertest@vcn.cloud"
  password                = "changeme!"
}
```

## Example of Provider Configuration with Operator Level Access

```hcl
provider "velocloud" {
  vco       = "https://vco.vcn.cloud/portal/rest"

  username                = "supertest@vcn.cloud"
  password                = "changeme!"

  skip_ssl_verification   = true
  operator                = true
}
```

## Example of Setting Environment Variables

```sh
export VCO_URL      = "https://vco.vcn.cloud/portal/rest"
export VCO_TOKEN    = "my-token"
```
