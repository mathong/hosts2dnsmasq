# Installation

1. Put `adblock-dnsmasq.service` and `adblock-dnsmasq.timer` in `/etc/systemd/system`
2. `systemctl daemon-reload`
3. Try to start the service `systemctl start adblock-dnsmasq`
4. Check logs: `journalctl -u adblock-dnsmasq.service`
5. Enable the service and timer: `systemctl enable adblock-dnsmasq`
6. Check that the timer started: `systemctl list-timers`