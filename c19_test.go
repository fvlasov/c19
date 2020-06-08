package c19

import (
	"net/http"
	"net/http/httptest"
)

var (
	mux *http.ServeMux

	server *httptest.Server
)

func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
}

func teardown() {
	server.Close()
}
