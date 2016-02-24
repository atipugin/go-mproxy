package mproxy

import (
	"net/url"
)

type Endpoint struct {
	*url.URL
}

func NewEndpoint(u *url.URL) *Endpoint {
	return &Endpoint{u}
}
