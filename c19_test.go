package c19

import (
	"net/http"
	"net/http/httptest"
)

var (
	mux *http.ServeMux

	client *Client

	server *httptest.Server
)

func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	client = NewClient()

	client.BaseURL = server.URL
}

func teardown() {
	server.Close()
}
