package mproxy

import (
	"errors"
	"math/rand"
	"net/url"
)

var (
	ErrNoEndpointsAvailable = errors.New("no endpoints available")
	ErrNoUrls               = errors.New("no urls provided")
)

type Registry struct {
	Endpoints map[string]*Endpoint
}

func NewRegistry(urls []string) (*Registry, error) {
	if len(urls) == 0 {
		return nil, ErrNoUrls
	}

	e := map[string]*Endpoint{}
	for _, v := range urls {
		u, err := url.Parse(v)
		if err != nil {
			return nil, err
		}

		e[u.Host] = NewEndpoint(u)
	}

	return &Registry{e}, nil
}

func (r *Registry) RandomEndpoint() (*Endpoint, error) {
	a := r.AvailableEndpoints()
	if len(a) == 0 {
		return nil, ErrNoEndpointsAvailable
	}

	return a[rand.Intn(len(a))], nil
}

func (r *Registry) AvailableEndpoints() []*Endpoint {
	var e []*Endpoint
	for _, v := range r.Endpoints {
		if v.Available {
			e = append(e, v)
		}
	}

	return e
}
