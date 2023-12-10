package util

import (
	"fmt"
	"strings"

	"github.com/imzoloft/soundcheck/util/inout"
)

type ProxyInfo struct {
	IP       string
	Port     string
	Username string
	Password string
}

func ParseProxy(proxyString string) *ProxyInfo {
	parts := strings.Split(proxyString, ":")

	var username, password string

	switch len(parts) {
	case 2:
		return &ProxyInfo{
			IP:   parts[0],
			Port: parts[1],
		}
	case 4:
		username = parts[2]
		password = parts[3]
	default:
		fmt.Printf("INVALID PROXY FORMAT %s\n%d", proxyString, len(parts))
		inout.FatalError("INVALID PROXY FORMAT sdf", proxyString)
	}
	return &ProxyInfo{
		IP:       parts[0],
		Port:     parts[1],
		Username: username,
		Password: password,
	}
}
