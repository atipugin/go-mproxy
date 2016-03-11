package mproxy

import (
	"net/url"
	"testing"
)

func TestNewEndpoint(t *testing.T) {
	urls := []string{
		"https://127.0.0.1:800",
		"http://127.0.0.1:800",
		"http://127.0.0.1",
		"127.0.0.1",
	}

	for _, v := range urls {
		u, _ := url.Parse(v)
		e := NewEndpoint(u)
		if u.Path != e.Path {
			t.Fatalf("Bad: %#v", e)
		}
	}
}
