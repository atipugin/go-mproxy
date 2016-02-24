package mproxy

import (
	"errors"
	"math/rand"
	"net/url"
)

type Registry struct {
	Endpoints map[string]*Endpoint
}

func NewRegistry(urls []string) (*Registry, error) {
	if len(urls) == 0 {
		return nil, errors.New("urls not provided")
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

func (r *Registry) RandomEndpoint() *Endpoint {
	a := r.AvailableEndpoints()

	return a[rand.Intn(len(a))]
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
