package c19

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"
)

var TestgetWorrldWIPJSON = `
[
	{
		"NewConfirmed": 88833,
		"TotalConfirmed": 5581829,
		"NewDeaths": 3214,
		"TotalDeaths": 354790,
		"NewRecovered": 63212,
		"TotalRecovered": 2231124
	},
	{
		"NewConfirmed": 90813,
		"TotalConfirmed": 3578992,
		"NewDeaths": 4736,
		"TotalDeaths": 251397,
		"NewRecovered": 84595,
		"TotalRecovered": 1159713
	}
]
`

func TestGetWorldWIP(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/world", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, TestgetWorrldWIPJSON)
	})

	res, err := client.GetWorldWIP(context.Background(), time.Time{}, time.Time{})
	if err != nil {
		t.Fatal("err", err)
	}

	n := []WorldTotal{{88833, 5581829, 3214, 354790, 63212, 2231124}, {90813, 3578992, 4736, 251397, 84595, 1159713}}

	if !reflect.DeepEqual(n, res) {
		t.Errorf("have %+v want %+v", n, res)
	}
}

var WorldTotalWIPJSON = `
{
	"TotalConfirmed": 8274306,
	"TotalDeaths": 451939,
	"TotalRecovered": 3954518
}
`

func TestGetWorldTotalWIP(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/world/total", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, WorldTotalWIPJSON)
	})

	res, err := client.GetWorldTotalWIP(context.Background())
	if err != nil {
		t.Fatal("err", err)
	}

	n := &WorldTotalWIP{8274306, 451939, 3954518}

	if !reflect.DeepEqual(n, res) {
		t.Errorf("have %+v want %+v", n, res)
	}
}
