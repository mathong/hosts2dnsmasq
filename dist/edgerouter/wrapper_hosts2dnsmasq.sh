#!/bin/bash

set -e
logger "updating dsnmasq configuration"
/config/scripts/hosts2dnsmasq -conf-path /etc/dnsmasq.d/block_adware_malware.conf
logger "reloading dnsmasq"
/etc/init.d/dnsmasq force-reload
