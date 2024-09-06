package net_ops

import (
	"errors"
	"net"
)

// DNSLookup returns the first IP address of a domain name.
func DNSLookup(domain string) (string, error) {
	// Perform the DNS lookup
	ips, err := net.LookupIP(domain)
	if err != nil {
		return "", err
	}

	// Check if any IP addresses were returned
	if len(ips) == 0 {
		return "", errors.New("no IP addresses found for domain")
	}

	// Return the first IP address as a string
	return ips[0].String(), nil
}

func DNSLookupNS(tldomain string) ([]string, error) {
	// Perform NS lookup
	nsRecords, err := net.LookupNS(tldomain)
	if err != nil {
		return nil, err
	}

	// Collect the NS hostnames in a string slice
	var nsHosts []string
	for _, ns := range nsRecords {
		nsHosts = append(nsHosts, ns.Host)
	}

	return nsHosts, nil
}
