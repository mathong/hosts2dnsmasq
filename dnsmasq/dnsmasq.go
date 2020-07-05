package dnsmasq

import (
	"fmt"
	"io"
)

const (
	// DefaultConfigFormat represents the configuration file content. '%s' is the host name.
	DefaultConfigFormat = "address=/%s/0.0.0.0\n"
)

// WriteConfig writes dnsmasq configuration to a writer, following a given format.
func WriteConfig(writer io.Writer, format string, hosts []string) error {
	for i := range hosts {
		_, err := fmt.Fprintf(writer, format, hosts[i])
		if err != nil {
			return err
		}
	}
	return nil
}
