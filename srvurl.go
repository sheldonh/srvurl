package srvul

import (
	"fmt"
	"net"
	URL "net/url"
	"strings"
)

type Resolver struct {
}

func NewResolver() *Resolver {
	return new(Resolver)
}

func (r *Resolver) ResolveUrl(url string) string {
	u, err := URL.Parse(url)
	if err != nil {
		return url
	}
	components := strings.Split(u.Host, ":")
	if len(components) == 1 {
		_, srvs, err := net.LookupSRV("", "", components[0])
		if err == nil {
			u.Host = srvAddress(r.selectSrv(srvs))
		}
	}

	return u.String()
}

func srvAddress(srv *net.SRV) string {
	host, port := srv.Target, srv.Port
	return fmt.Sprintf("%s:%d", host[:len(host)-1], port)
}

func (r *Resolver) selectSrv(srvs []*net.SRV) *net.SRV {
	return srvs[0]
}
