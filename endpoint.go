package mproxy

import (
	"net/url"
)

type Endpoint struct {
	*url.URL
	Available bool
}

func NewEndpoint(u *url.URL) *Endpoint {
	return &Endpoint{u, true}
}

func (e *Endpoint) MarkAsUnavailable() {
	e.Available = false
}
