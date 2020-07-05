# hosts2dnsmasq

Dumb tool that transforms a hosts file into a dnsmasq configuration file.
The goal is to automate DNS blacklisting like [PiHole](https://github.com/pi-hole/pi-hole), but a lot simpler. You just need a host that can run dnsmasq and a Go binary.

Concept:

* Setup dnsmasq on a lacal machine, available 24/7
* Set this host as your DNS resolver
* Forward DNS queries from this host to your favorite provider, using dnsmasq
* Block domains based on a block list
* Update the block list automatically

Default block list comes from (https://github.com/StevenBlack/hosts) ([adware + malware](https://raw.githubusercontent.com/StevenBlack/hosts/master/hosts)).

# Requirements

Dnsmasq, OS/Architecture supported by Go, and access to the internet.

Tested on Ubuntu and EdgeOS v1.10.x.

# Build

See the [official Go installation instructions](https://golang.org/doc/install) to setup your build environment.
Requirement: Go >= 1.11 (since the project uses Go modules). 

## Build for you local architecture

`make build`

## Build for ERX

`make build-mipsle`

# Usage

See `hosts2dnsmasq -h`
	
## Installation and automation

Files are provided in `dist`, to automate the block list update.