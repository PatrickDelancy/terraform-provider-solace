# terraform-provider-solace

[![Build Status](https://travis-ci.org/ExalDraen/terraform-provider-solace.svg?branch=master)](https://travis-ci.org/ExalDraen/terraform-provider-solace)
[![Go Report Card](https://goreportcard.com/badge/github.com/ExalDraen/terraform-provider-solace)](https://goreportcard.com/report/github.com/ExalDraen/terraform-provider-solace)
[![GoDoc](https://godoc.org/github.com/ExalDraen/terraform-provider-solace/solace?status.svg)](https://godoc.org/github.com/ExalDraen/terraform-provider-solace/solace)

<img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" width="300px" />
<img src="https://3yecy51kdipx3blyi37oute1-wpengine.netdna-ssl.com/wp-content/uploads/2018/05/Solace_Logo_Green_360x100-1.png" width="300px" />

## Table of Contents

- [terraform-provider-solace](#terraform-provider-solace)
  - [Using A Binary Package](#using-a-binary-package)
  - [Building The Provider From Source](#building-the-provider-from-source)
  - [Developing the Provider](#developing-the-provider)
  - [Documentation](#documentation)

[Terraform](https://www.terraform.io) provider plugin for managing [Solace](https://solace.com/) appliances using the [Solace SEMP v2 API](https://docs.solace.com/SEMP/Using-SEMP.htm).

## Using the Binary Packages

Download the [release binary](https://github.com/ExalDraen/terraform-provider-solace/releases) from and copy it to the `$HOME/terraform.d/plugins/<os>_<arch>/terraform-provider-solace` as explained in the [terraform docs](https://www.terraform.io/docs/configuration/providers.html#third-party-plugins).

For example, in a Linux environment you would copy it to `/home/youruser/terraform.d/plugins/linux_amd64/terraform-provider-solace`, for a MacOS environment to `/Users/youruser/terraform.d/plugins/darwin_amd64/terraform-provider-solace`.

## Building The Provider From Source

Clone repository to: `$GOPATH/src/github.com/ExalDraen/terraform-provider-solace`

```sh
mkdir -p $GOPATH/src/github.com/ExalDraen; cd $GOPATH/src/github.com/ExalDraen
git clone git@github.com:ExalDraen/terraform-provider-solace
```

Enter the provider directory and build the provider

```sh
cd $GOPATH/src/github.com/ExalDraen/terraform-provider-solace
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

## Documentation

The provider very closely mirrors the SEMPv2 API, including resource and attribute names. In most cases, the name of attributes and resources is the same as in the SEMP API, with `snake_case` instead of `CamelCase`. For example `max_connection_count` instead of `MaxConnectionCount`.

See [examples/main.tf](examples/main.tf) for some usage examples.