package mproxy

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"testing"
)

type TestHandler struct {
	sync.Mutex
	count int
}

func (h *TestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func TestReverseProxyInit(t *testing.T) {
	r, err := NewRegistry([]string{
		"http://localhost:8001",
		"http://localhost:8002",
		"http://localhost:8003",
	})

	if err != nil {
		t.Fatalf("Error: %s", err)
	}

	p := NewReverseProxy(r)

	if len(p.Endpoints) != 3 {
		t.Fatalf("Missed endpoints")
	}
}

func TestReverseProxy(t *testing.T) {
	go func() {
		mux := http.NewServeMux()
		handler := &TestHandler{}
		mux.Handle("/", handler)
		server := http.Server{Handler: mux, Addr: ":8001"}
		log.Fatal(server.ListenAndServe())
	}()

	go func() {
		r, err := NewRegistry([]string{
			"http://localhost:8001",
		})

		if err != nil {
			log.Fatalf("Error: %s", err)
		}

		p := NewReverseProxy(r)
		http.Handle("/", p)
		http.ListenAndServe(":8080", nil)
	}()

	resp, err := http.Get("http://localhost:8080")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 200 {
		t.Fatalf("Received non-200 response: %d\n", resp.StatusCode)
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if "Hello" != string(content) {
		t.Fatal("Expected the message Hello")
	}
}
