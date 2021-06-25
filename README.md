# terraform-provider-solace

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/PatrickDelancy/terraform-provider-solace)
[![GoDoc](https://godoc.org/github.com/PatrickDelancy/terraform-provider-solace/solace?status.svg)](https://godoc.org/github.com/PatrickDelancy/terraform-provider-solace/solace)
[![Go Report Card](https://goreportcard.com/badge/github.com/PatrickDelancy/terraform-provider-solace)](https://goreportcard.com/report/github.com/PatrickDelancy/terraform-provider-solace)
<a href="https://codeclimate.com/github/PatrickDelancy/terraform-provider-solace/maintainability"><img src="https://api.codeclimate.com/v1/badges/e9f9e33f3a7160cc431a/maintainability" /></a>

![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/PatrickDelancy/terraform-provider-solace)
[![Create Release](https://github.com/PatrickDelancy/terraform-provider-solace/actions/workflows/release.yml/badge.svg)](https://github.com/PatrickDelancy/terraform-provider-solace/actions/workflows/release.yml)

![GitHub](https://img.shields.io/github/license/PatrickDelancy/terraform-provider-solace)
[![PRs Welcome](https://img.shields.io/badge/PRs-Welcome-brightgreen)](http://makeapullrequest.com)

## Table of Contents

- [terraform-provider-solace](#terraform-provider-solace)
  - [Using A Binary Package](#using-a-binary-package)
  - [Building The Provider From Source](#building-the-provider-from-source)
  - [Developing the Provider](#developing-the-provider)
  - [Documentation](#documentation)

[Terraform](https://www.terraform.io) provider plugin for managing [Solace](https://solace.com/) appliances using the [Solace SEMP v2 API](https://docs.solace.com/SEMP/Using-SEMP.htm).

## Using the Binary Packages

Download the [release binary](https://github.com/PatrickDelancy/terraform-provider-solace/releases) from and copy it to `$HOME/terraform.d/plugins/<os>_<arch>/terraform-provider-solace` as explained in the [terraform docs](https://www.terraform.io/docs/configuration/providers.html#third-party-plugins).

For example, in a Linux environment you would copy it to `/home/youruser/terraform.d/plugins/linux_amd64/terraform-provider-solace`, for a MacOS environment to `/Users/youruser/terraform.d/plugins/darwin_amd64/terraform-provider-solace`.

## Building The Provider From Source

Clone repository to: `$GOPATH/src/github.com/PatrickDelancy/terraform-provider-solace`

```sh
mkdir -p $GOPATH/src/github.com/PatrickDelancy; cd $GOPATH/src/github.com/PatrickDelancy
git clone git@github.com:PatrickDelancy/terraform-provider-solace
```

Enter the provider directory and build the provider

```sh
cd $GOPATH/src/github.com/PatrickDelancy/terraform-provider-solace
make
```

## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.8+ is *required*). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ make build
...
$ $GOPATH/bin/terraform-provider-solace
...
```

### Testing the Provider

To run acceptance tests you will need to run a local solace instance. Once the instance is running you will need to export the following environment variables.

```sh
export SOLACE_USER="admin"
export SOLACE_PASSWORD="alex"
export SOLACE_HOST="localhost:8080"
```

You can then run the acceptance tests with

```sh
TF_ACC=1 make test
```

## Documentation

The provider very closely mirrors the SEMPv2 API, including resource and attribute names. In most cases, the name of attributes and resources is the same as in the SEMP API, with `snake_case` instead of `CamelCase`. For example `max_connection_count` instead of `MaxConnectionCount`.

See [examples/main.tf](examples/main.tf) for some usage examples.

## Configuration

The provider may be configured using terraform as normal, e.g.

```hcl
provider "solace" {
    host = "localhost:8080"
    base_path = "/SEMP/v2/config"
    user = "admin"
    password = "alex"
    msg_vpn = "go-tf-1"
}
```

or using the corresponding `SOLACE_*` environment variables. Check `provider.go` to see which variables are available.