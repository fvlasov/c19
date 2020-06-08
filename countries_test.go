package c19

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

var countriesRes = `
	[
  {
    "Country": "Barbados",
    "Slug": "barbados",
    "ISO2": "BB"
  }
]
`

func TestGetCountries(t *testing.T) {
	setup()
	defer teardown()
	countriesURL = server.URL + "/countries"
	baseURL = server.URL
	fmt.Println(baseURL, countriesURL)
	mux.HandleFunc("/countries", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, countriesRes)
	})

	r, err := GetCountries(context.Background())
	if err != nil {
		t.Errorf("Error %v", err)
	}

	sl := []Countries{{"Barbados", "barbados", "BB"}}

	if !reflect.DeepEqual(r, sl) {
		t.Errorf("returned keys %+v, expected %+v", r, sl)
	}
}
