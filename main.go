package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"git.gautier.ovh/hosts2dnsmasq/dnsmasq"
	"git.gautier.ovh/hosts2dnsmasq/hosts"
)

const (
	version               = "v1.0.0-rc.1"
	dnsmasqConfigFileMode = 0600
	hostsGetTimeout       = 30 * time.Second
)

func main() {
	hostsFileURL := flag.String("host-url", hosts.DefaultURL, "Host file URL")
	confFilePath := flag.String("conf-path", "conf", "Config file path")
	confFileFormat := flag.String("conf-format", dnsmasq.DefaultConfigFormat, "Which format to use while writing the configuration")
	showVersion := flag.Bool("version", false, "Print binary version")

	flag.Parse()

	if *showVersion {
		fmt.Println(version)
		return
	}

	log.Printf("Downloading hosts file %s", *hostsFileURL)
	client := &http.Client{Timeout: hostsGetTimeout}
	r, err := hosts.Download(client, *hostsFileURL)
	if err != nil {
		log.Printf("Failed to download hosts file from '%s': %s", *hostsFileURL, err)
		os.Exit(1)
	}

	log.Print("Parsing hosts file")
	hosts, err := hosts.Parse(r)
	if err != nil {
		log.Printf("Failed to parse hosts file: %s", err)
		os.Exit(1)
	}
	log.Printf("Parsed %d hosts", len(hosts))

	configFile, err := os.OpenFile(*confFilePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, dnsmasqConfigFileMode)
	if err != nil {
		log.Printf("Failed to open Dnsmasq configuration: %s", err)
	}
	defer configFile.Close()

	log.Printf("Writing dnsmasq config to '%s'", *confFilePath)
	err = dnsmasq.WriteConfig(configFile, *confFileFormat, hosts)
	if err != nil {
		log.Printf("Failed to write Dnsmasq configuration: %s", err)
	}
}
