package c19

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

var contriesJSON = `
[{"Country":"Macao, SAR China","Slug":"macao-sar-china","ISO2":"MO"},{"Country":"Oman","Slug":"oman","ISO2":"OM"}]
`

func TestGetCountries(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(countriesURL, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, contriesJSON)
	})

	res, err := client.GetCountries(context.Background())
	if err != nil {
		t.Fatal("err", err)
	}

	n := []Countries{{"Macao, SAR China", "macao-sar-china", "MO"}, {"Oman", "oman", "OM"}}

	if !reflect.DeepEqual(n, res) {
		t.Errorf("have %+v want %+v", n, res)
	}
}
