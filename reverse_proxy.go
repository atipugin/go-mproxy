package mproxy

import (
	"errors"
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
		e := p.Registry.RandomEndpoint()

		r.URL.Scheme = e.URL.Scheme
		r.URL.Host = e.URL.Host
	}

	p.Transport = &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		Dial: func(network, addr string) (net.Conn, error) {
			d := &net.Dialer{}
			conn, err := d.Dial(network, addr)
			if err != nil {
				p.Registry.Endpoints[addr].MarkAsUnavailable()
				e := p.Registry.AvailableEndpoints()
				for _, v := range e {
					conn, err = d.Dial(network, v.URL.Host)
					if err != nil {
						continue
					}

					return conn, err
				}

				return nil, errors.New("no endpoints available")
			}

			return conn, err
		},
	}

	return p
}
