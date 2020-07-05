package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"git.gautier.ovh/hosts2dnsmasq/dnsmasq"
	"git.gautier.ovh/hosts2dnsmasq/hosts"
)

const (
	version               = "v1.0.0-rc.1"
	dnsmasqConfigFileMode = 0600
)

func main() {
	hostFileURL := flag.String("host-url", hosts.DefaultURL, "Host file URL")
	confFilePath := flag.String("conf-path", "conf", "Config file path")
	confFileFormat := flag.String("conf-format", dnsmasq.DefaultConfigFormat, "Which format to use while writing the configuration")
	showVersion := flag.Bool("version", false, "Print binary version")

	flag.Parse()

	if *showVersion {
		fmt.Println(version)
		return
	}

	r, err := hosts.Download(&http.Client{}, *hostFileURL)
	if err != nil {
		log.Printf("Failed to download hosts file from '%s': %s", *hostFileURL, err)
		os.Exit(1)
	}
	hosts, err := hosts.Parse(r)
	if err != nil {
		log.Printf("Failed to parse hosts file: %s", err)
		os.Exit(1)
	}

	configFile, err := os.OpenFile(*confFilePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, dnsmasqConfigFileMode)
	if err != nil {
		log.Printf("Failed to open Dnsmasq configuration: %s", err)
	}
	defer configFile.Close()

	err = dnsmasq.WriteConfig(configFile, *confFileFormat, hosts)
	if err != nil {
		log.Printf("Failed to write Dnsmasq configuration: %s", err)
	}
}
