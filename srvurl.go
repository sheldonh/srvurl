package srvul

import (
	"fmt"
	"net"
	URL "net/url"
	"strings"
)

func ResolveUrl(url string) string {
	u, err := URL.Parse(url)
	if err != nil {
		return url
	}
	components := strings.Split(u.Host, ":")
	if len(components) == 1 {
		_, srvs, err := net.LookupSRV("", "", components[0])
		if err == nil {
			host, port := srvs[0].Target, srvs[0].Port
			u.Host = fmt.Sprintf("%s:%d", host[:len(host)-1], port)
		}
	}

	return u.String()
}
