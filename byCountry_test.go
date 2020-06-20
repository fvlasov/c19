package c19

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"
)

var TestGetByCountryJSON = `
[
	{"Country":"South Africa",
	"CountryCode":"ZA",
	"Province":"",
	"City":"",
	"CityCode":"",
	"Lat":"-30.56",
	"Lon":"22.94",
	"Cases":0,
	"Status":"confirmed",
	"Date":"2020-03-01T00:00:00Z"},

	{"Country":"South Africa",
	"CountryCode":"ZA",
	"Province":"",
	"City":"",
	"CityCode":"",
	"Lat":"-30.56",
	"Lon":"22.94",
	"Cases":0,
	"Status":"confirmed",
	"Date":"2020-03-02T00:00:00Z"}
]
`

func TestGetByCountry(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/dayone/country/south-africa/status/confirmed", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, TestGetByCountryJSON)
	})
	
	res, err := client.GetByCountry(context.Background(), "south-africa", "confirmed", time.Time{}, time.Time{})
	if err != nil {
		t.Fatal("err", err)
	}

	n := []ByCountry{{Country: "South Africa",
		CountryCode: "ZA",
		Province:    "",
		City:        "",
		CityCode:    "",
		Lat:         "-30.56",
		Lon:         "22.94",
		Cases:       0,
		Status:      "confirmed",
		Date:        time.Date(2020, 3, 1, 0, 0, 0, 0, time.UTC)},
		{Country: "South Africa",
			CountryCode: "ZA",
			Province:    "",
			City:        "",
			CityCode:    "",
			Lat:         "-30.56",
			Lon:         "22.94",
			Cases:       0,
			Status:      "confirmed",
			Date:        time.Date(2020, 3, 2, 0, 0, 0, 0, time.UTC)},
	}

	if !reflect.DeepEqual(n, res) {
		t.Errorf("have %+v want %+v", n, res)
	}
}

var GetByCountryAllStatusJSON = `
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
		"Recovered": 2,
		"Active": 0,
		"Date": "2020-03-01T00:00:00Z"
	},
	{
		"Country": "Russian Federation",
		"CountryCode": "RU",
		"Province": "",
		"City": "",
		"CityCode": "",
		"Lat": "61.52",
		"Lon": "105.32",
		"Confirmed": 3,
		"Deaths": 0,
		"Recovered": 2,
		"Active": 1,
		"Date": "2020-03-02T00:00:00Z"
	}
]
`

func TestGetByCountryAllStatus(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/dayone/country/russia", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, GetByCountryAllStatusJSON)
	})

	res, err := client.GetByCountryAllStatus(context.Background(), "russia", time.Time{}, time.Time{})
	if err != nil {
		t.Fatal("err:", err)
	}

	n := []ByCountryAllStatus{{"Russian Federation", "RU", "", "", "", "61.52", "105.32", 2, 0, 2, 0, time.Date(2020, 03, 01, 00, 00, 00, 00, time.UTC)}, {"Russian Federation", "RU", "", "", "", "61.52", "105.32", 3, 0, 2, 1, time.Date(2020, 03, 02, 00, 00, 00, 00, time.UTC)}}

	if !reflect.DeepEqual(n, res) {
		t.Errorf("have %+v want %+v", n, res)
	}
}

func TestGetByCountryLive(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/dayone/country/south-africa/status/confirmed/live", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, TestGetByCountryJSON)
	})
	res, err := client.GetByCountryLive(context.Background(), "south-africa", "confirmed", time.Time{}, time.Time{})
	if err != nil {
		t.Fatal("err", err)
	}

	n := []ByCountry{{Country: "South Africa",
		CountryCode: "ZA",
		Province:    "",
		City:        "",
		CityCode:    "",
		Lat:         "-30.56",
		Lon:         "22.94",
		Cases:       0,
		Status:      "confirmed",
		Date:        time.Date(2020, 3, 1, 0, 0, 0, 0, time.UTC)},
		{Country: "South Africa",
			CountryCode: "ZA",
			Province:    "",
			City:        "",
			CityCode:    "",
			Lat:         "-30.56",
			Lon:         "22.94",
			Cases:       0,
			Status:      "confirmed",
			Date:        time.Date(2020, 3, 2, 0, 0, 0, 0, time.UTC)},
	}

	if !reflect.DeepEqual(n, res) {
		t.Errorf("have %+v want %+v", n, res)
	}
}
