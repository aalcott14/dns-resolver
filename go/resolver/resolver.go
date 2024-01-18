package resolver

import (
	"context"
	"net"
	"net/http"
	"time"
)

var (
	dnsResolverIP        = "8.8.8.8:53"
	dnsResolverProto     = "udp"
	dnsResolverTimeoutMs = 5000
)

func NewDnsClient() http.Client {
	dnsResolver := newDnsResolver()
	http.DefaultTransport.(*http.Transport).DialContext = dnsResolver.NewDialContext

	return http.Client{}
}

type dnsResolver struct {
	dialer *net.Dialer
}

func (r *dnsResolver) NewDialContext(ctx context.Context, network, addr string) (net.Conn, error) {
	return r.dialer.DialContext(ctx, network, addr)
}

func newDnsResolver() *dnsResolver {
	return &dnsResolver{
		dialer: &net.Dialer{
			Resolver: &net.Resolver{
				PreferGo: true,
				Dial:     dial,
			},
		},
	}
}

func dial(ctx context.Context, network, address string) (net.Conn, error) {
	d := net.Dialer{
		Timeout: time.Duration(dnsResolverTimeoutMs) * time.Millisecond,
	}
	return d.DialContext(ctx, dnsResolverProto, dnsResolverIP)
}
