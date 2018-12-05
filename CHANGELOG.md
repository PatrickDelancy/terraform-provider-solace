# Change Log
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/)
and this project adheres to [Semantic Versioning](http://semver.org/).


## [Unreleased]
### Known Issues
* The models used for parameters by the `semp-client` library are tagged with the label `omitempty`. Per [the json docs](https://golang.org/pkg/encoding/json/#Marshal) this means that these fields will *not* be included in a request if have an "empty value". 
This is a problem when trying to update an attribute from a *non-empty* to an *empty* value, for example to change a boolean from `false` to `true`: such a value is recognized as an "empty value" and not included in any update (`PATCH`) requests.

### Added
* Basic scaffolding for terraform provider
* Very basic MSG VPN resource with a few select attributes and create/destroy/update capability

[Unreleased]: https://github.com/ExalDraen/terraform-provider-solace/compare/
