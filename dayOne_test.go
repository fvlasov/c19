package c19

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"
)

var dayOneRes = `[
	{
	  "Country": "South Africa",
	  "CountryCode": "ZA",
	  "Province": "",
	  "City": "",
	  "CityCode": "",
	  "Lat": "-30.56",
	  "Lon": "22.94",
	  "Cases": 1,
	  "Status": "confirmed",
	  "Date": "2020-03-05T00:00:00Z"
	}
	]`

var dayOneTotalRes = `[
	{
	  "Country": "Switzerland",
	  "CountryCode": "CH",
	  "Province": "",
	  "City": "",
	  "CityCode": "",
	  "Lat": "46.82",
	  "Lon": "8.23",
	  "Confirmed": 1,
	  "Deaths": 0,
	  "Recovered": 0,
	  "Active": 1,
	  "Date": "2020-02-25T00:00:00Z"
	}
	]`

func TestGetDayOne(t *testing.T) {
	setup()
	defer teardown()
	baseDayOneURL = server.URL + "/dayone"

	mux.HandleFunc("/dayone/country/south-africa/status/confirmed", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, dayOneRes)
	})

	r, err := GetDayOne(context.Background(), "south-africa", "confirmed")
	if err != nil {
		t.Errorf("Error %v", err)
	}

	pt, err := time.Parse(time.RFC3339Nano, "2020-03-05T00:00:00Z")
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	sl := []DayOne{{"South Africa", "ZA", "", "", "", "-30.56", "22.94", 1, "confirmed", pt}}

	if !reflect.DeepEqual(r, sl) {
		t.Errorf("returned keys %+v, expected %+v", r, sl)
	}
}


