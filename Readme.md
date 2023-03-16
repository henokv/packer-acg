[![goreleaser](https://github.com/henokv/docs-azurerm/actions/workflows/release.yml/badge.svg)](https://github.com/henokv/docs-azurerm/actions/workflows/release.yml)

# packer-acg

This is a tool to add azure compute gallery to the default packer build of github runners ound in the msft repo

## Installation
To install download the latest version from the [releases](https://github.com/henokv/packer-acg/releases) page or if you have go installed run the command
```shell
go install github.com/henokv/packer-acg@latest
```

## Usage
Example 1:
```shell
packer-acg -i ubuntu-22.04 -n gallery-name -g galler-rg -f source.hcl
```

Example 2:
```shell
packer-acg -i ubuntu-22.04 -n gallery-name -g galler-rg -f source.hcl -o destination.hcl
```
