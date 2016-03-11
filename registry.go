package mproxy

import (
	"errors"
	"net/url"
	"sync"
)

var (
	ErrNoUrls   = errors.New("no urls provided")
	ErrNoScheme = errors.New("invalid url scheme provided")
	ErrNoHost   = errors.New("invalid url host provided")
)

type Registry struct {
	Endpoints       []*Endpoint
	currEndpointIdx int
	mtx             *sync.Mutex
}

func NewRegistry(urls []string) (*Registry, error) {
	if len(urls) == 0 {
		return nil, ErrNoUrls
	}

	var e []*Endpoint
	for _, v := range urls {
		u, err := url.Parse(v)
		if err != nil {
			return nil, err
		}
		if len(u.Scheme) == 0 {
			return nil, ErrNoScheme
		}
		if len(u.Host) == 0 {
			return nil, ErrNoScheme
		}

		e = append(e, NewEndpoint(u))
	}

	return &Registry{e, 0, &sync.Mutex{}}, nil
}

func (r *Registry) Endpoint() *Endpoint {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	e := r.Endpoints[r.currEndpointIdx]

	r.currEndpointIdx++
	if r.currEndpointIdx >= len(r.Endpoints) {
		r.currEndpointIdx = 0
	}

	return e
}
