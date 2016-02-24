package mproxy

import (
	"errors"
	"net"
	"net/http"
	"net/http/httputil"
)

var (
	ErrNoEndpointsAvailable = errors.New("no endpoints available")
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

			for i := 0; i < len(p.Registry.Endpoints); i++ {
				e := p.Registry.Endpoint()
				conn, err := d.Dial(network, e.URL.Host)
				if err != nil {
					continue
				}

				return conn, err
			}

			return nil, ErrNoEndpointsAvailable
		},
	}

	return p
}
