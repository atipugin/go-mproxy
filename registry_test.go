package mproxy

import "testing"

func TestNewRegistry(t *testing.T) {
	okUrls := []string{
		"https://127.0.0.1:8001",
		"http://127.0.0.1:8002",
		"http://127.0.0.1",
	}

	_, err := NewRegistry(okUrls)
	if err != nil {
		t.Fatalf("Error with parse: %s", err)
	}

	badUrls := []string{
		"127.0.0.1:8001",
		"127.0.0.1",
	}

	for _, u := range badUrls {
		_, err := NewRegistry([]string{u})
		if err == nil {
			t.Fatalf("Error with parse: %s", u)
		}
	}
}
