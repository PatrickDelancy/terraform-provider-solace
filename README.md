<img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" width="400px">

# terraform-provider-solace
[![Build Status](https://travis-ci.org/ExalDraen/terraform-provider-solace.svg?branch=master)](https://travis-ci.org/ExalDraen/terraform-provider-solace)

<!-- markdown-toc start - Don't edit this section. Run M-x markdown-toc-refresh-toc -->
**Table of Contents**

- [terraform-provider-solace](#terraform-provider-solace)
    - [Using A Binary Package](#using-a-binary-package)
    - [Building The Provider From Source](#building-the-provider-from-source)
    - [Developing the Provider](#developing-the-provider)
    - [Documentation](#documentation)

<!-- markdown-toc end -->


[Terraform](https://www.terraform.io) provider plugin for managing Solace appliances using the [Solace SEMP v2 API](https://docs.solace.com/SEMP/Using-SEMP.htm).

**Work in Progress !**

## Using A Binary Package

Download the release binary and copy it to the `$HOME/terraform.d/plugins/<os>_<arch>/terraform-provider-solace` as explain in the [terraform docs](https://www.terraform.io/docs/configuration/providers.html#third-party-plugins).
For example `/home/youruser/terraform.d/plugins/linux_amd64/terraform-provider-solace` for a Linux environment or `/Users/youruser/terraform.d/plugins/darwin_amd64/terraform-provider-solace` for a MacOS environment.

## Building The Provider From Source

Clone repository to: `$GOPATH/src/github.com/ExalDraen/terraform-provider-solace`

```sh
$ mkdir -p $GOPATH/src/github.com/ExalDraen; cd $GOPATH/src/github.com/ExalDraen
$ git clone git@github.com:ExalDraen/terraform-provider-solace
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/ExalDraen/terraform-provider-solace
$ make
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
TODO
