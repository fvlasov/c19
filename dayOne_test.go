package c19

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"
)

var dayOneJSON = `
[
	{
		"Country": "Russian Federation",
		"CountryCode": "RU",
		"Province": "",
		"City": "",
		"CityCode": "",
		"Lat": "61.52",
		"Lon": "105.32",
		"Cases": 2,
		"Status": "confirmed",
		"Date": "2020-01-31T00:00:00Z"
	},
	{
		"Country": "Russian Federation",
		"CountryCode": "RU",
		"Province": "",
		"City": "",
		"CityCode": "",
		"Lat": "61.52",
		"Lon": "105.32",
		"Cases": 2,
		"Status": "confirmed",
		"Date": "2020-02-01T00:00:00Z"
	}
]
`

func TestGetDayOne(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/country/russia/status/confirmed", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, dayOneJSON)
	})

	res, err := client.GetDayOne(context.Background(), "russia", "confirmed")
	if err != nil {
		t.Fatal("err", err)
	}

	n := []DayOne{{"Russian Federation", "RU", "", "", "", "61.52", "105.32", 2, "confirmed", time.Date(2020, 01, 31, 00, 00, 00, 00, time.UTC)},
		{"Russian Federation", "RU", "", "", "", "61.52", "105.32", 2, "confirmed", time.Date(2020, 02, 01, 00, 00, 00, 00, time.UTC)}}

	if !reflect.DeepEqual(n, res) {
		t.Errorf("have %+v want %+v", n, res)
	}
}

var GetDayOneAllStatusJSON = `
[
	{
		"Country": "Russian Federation",
		"CountryCode": "RU",
		"Province": "",
		"City": "",
		"CityCode": "",
		"Lat": "61.52",
		"Lon": "105.32",
		"Confirmed": 2,
		"Deaths": 0,
		"Recovered": 0,
		"Active": 2,
		"Date": "2020-01-31T00:00:00Z"
	},
	{
		"Country": "Russian Federation",
		"CountryCode": "RU",
		"Province": "",
		"City": "",
		"CityCode": "",
		"Lat": "61.52",
		"Lon": "105.32",
		"Confirmed": 2,
		"Deaths": 0,
		"Recovered": 0,
		"Active": 2,
		"Date": "2020-02-01T00:00:00Z"
	}
]
`

func TestGetDayOneAllStatus(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/dayone/country/russia", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, GetByCountryAllStatusJSON)
	})

	res, err := client.GetDayOneAllStatus(context.Background(), "russia")
	if err != nil {
		t.Fatal("err", err)
	}

	if len(res) == 0 {
		t.Fatal("res should not be empty", res)
	}
}

func TestGetDayOneLive(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/dayone/country/russia/status/confirmed/live", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, dayOneJSON)
	})

	res, err := client.GetDayOneLive(context.Background(), "russia", "confirmed")
	if err != nil {
		t.Fatal("err", err)
	}

	if len(res) == 0 {
		t.Fatal("res should not be empty", res)
	}
}

func TestGetDayOneTotal(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/total/dayone/country/russia/status/confirmed", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, GetDayOneAllStatusJSON)
	})

	res, err := client.GetDayOneTotal(context.Background(), "russia", "confirmed")
	if err != nil {
		t.Fatal("err", err)
	}

	if len(res) == 0 {
		t.Fatal("res should not be empty", res)
	}
}
