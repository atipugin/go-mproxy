package mproxy

import (
	"net"
	"net/http"
	"net/http/httputil"
)

type ReverseProxy struct {
	httputil.ReverseProxy
	*Registry
}

func NewReverseProxy(r *Registry) *ReverseProxy {
	p := &ReverseProxy{httputil.ReverseProxy{}, r}

	p.Director = func(r *http.Request) {
		r.URL.Scheme = "http"
		r.URL.Host = "proxyhost" // ...
	}

	p.Transport = &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		Dial: func(network, addr string) (net.Conn, error) {
			d := &net.Dialer{}

			for {
				e, err := p.Registry.RandomEndpoint()
				if err != nil {
					break
				}

				conn, err := d.Dial(network, e.URL.Host)
				if err != nil {
					e.MarkAsUnavailable()
					continue
				}

				return conn, err
			}

			return nil, ErrNoEndpointsAvailable
		},
	}

	return p
}
