# Installation

**note: this is for edgeOS v1.10.x. If they are using Systemd in v2.x.x, you may be able to use the Systemd units :)**

1. Compile hosts2dnsmasq for the proper architecture (see "build" in README at the project's root).
2. Send the hosts2dnsmasq binary to your EdgeRouter.
To make it persistent after upgrades, put it in `/config`. For example, in `/config/scripts/hosts2dnsmasq`
2. Send the `wrapper_hosts2dnsmasq.sh` to `/config/scripts/wrapper_hosts2dnsmasq.sh`
3. Update the configuration. See the `config` file
