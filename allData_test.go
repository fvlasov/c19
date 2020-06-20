package c19

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"
)

var allDataJson = `
[
	{
	 "Country": "Afghanistan",
	 "CountryCode": "AF",
	 "Province": "",
	 "City": "",
	 "CityCode": "",
	 "Lat": "33.94",
	 "Lon": "67.71",
	 "Confirmed": 0,
	 "Deaths": 0,
	 "Recovered": 0,
	 "Active": 0,
	 "Date": "2020-01-22T00:00:00Z"
	},
	{
	 "Country": "Afghanistan",
	 "CountryCode": "AF",
	 "Province": "",
	 "City": "",
	 "CityCode": "",
	 "Lat": "33.94",
	 "Lon": "67.71",
	 "Confirmed": 0,
	 "Deaths": 0,
	 "Recovered": 0,
	 "Active": 0,
	 "Date": "2020-01-23T00:00:00Z"
	}
]
`

func TestGetAllData(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/all", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, allDataJson)
	})

	res, err := client.GetAllData(context.Background())
	if err != nil {
		t.Fatal("error:", err)
	}
	n := []AllData{{"Afghanistan", "AF", "", "", "", "33.94", "67.71", 0, 0, 0, 0, time.Date(2020, 1, 22, 0, 0, 0, 0, time.UTC)},
		{"Afghanistan", "AF", "", "", "", "33.94", "67.71", 0, 0, 0, 0, time.Date(2020, 1, 23, 0, 0, 0, 0, time.UTC)}}

	if !reflect.DeepEqual(n, res) {
		t.Errorf("have %+v want %+v", n, res)
	}
}
