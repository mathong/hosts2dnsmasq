package hosts

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
)

const (
	// DefaultURL point to the main hosts file provided by https://github.com/StevenBlack/hosts
	DefaultURL = "https://raw.githubusercontent.com/StevenBlack/hosts/master/hosts"
)

// Parse parses hosts file's content.
func Parse(r io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(r)
	re, err := regexp.Compile(`^0\.0\.0\.0 ([^ ^\n]+)`)
	if err != nil {
		return nil, err
	}
	hosts := make([]string, 0)
	for scanner.Scan() {
		host := re.FindStringSubmatch(scanner.Text())
		if len(host) == 2 {
			if host[1] != "0.0.0.0" {
				hosts = append(hosts, host[1])
			}
		}
	}
	return hosts, scanner.Err()
}

// Download downloads a remote hosts file and returns an io.Reader that can be passed to Parse()
func Download(client Client, url string) (io.Reader, error) {
	log.Printf("Downloading host file %s", url)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 300 {
		return nil, fmt.Errorf("HTTP status code >300: '%d'", resp.StatusCode)
	}

	return resp.Body, nil
}
