package net_ops

import (
	"errors"
	"net"
)

// Check interface
func CheckInterface(iface string) (ipaddress string, err error) {
	// get the interface
	netInterface, err := net.InterfaceByName(iface)
	if err != nil {
		return "", errors.New("failed to get interface: " + err.Error())
	}

	// Add a return statement for the missing return
	addrs, err := netInterface.Addrs()
	if err != nil {
		return "", errors.New("failed to get interface address: " + err.Error())
	}

	if len(addrs) > 0 {
		ip, _, _ := net.ParseCIDR(addrs[0].String())
		return ip.String(), nil
	}

	return "", errors.New("no IP address found for the interface")
}
