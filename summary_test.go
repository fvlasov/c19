package c19

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

var SummaryJSON = `
{
	"Global": {
	  "NewConfirmed": 139803,
	  "TotalConfirmed": 8274306,
	  "NewDeaths": 6829,
	  "TotalDeaths": 451939,
	  "NewRecovered": 97831,
	  "TotalRecovered": 3954518
	},
	"Countries": [
	  {
		"Country": "Afghanistan",
		"CountryCode": "AF",
		"Slug": "afghanistan",
		"NewConfirmed": 783,
		"TotalConfirmed": 26310,
		"NewDeaths": 13,
		"TotalDeaths": 491,
		"NewRecovered": 344,
		"TotalRecovered": 5508,
		"Date": "2020-06-17T17:30:02Z"
	  },
	  {
		"Country": "Zimbabwe",
		"CountryCode": "ZW",
		"Slug": "zimbabwe",
		"NewConfirmed": 4,
		"TotalConfirmed": 391,
		"NewDeaths": 0,
		"TotalDeaths": 4,
		"NewRecovered": 8,
		"TotalRecovered": 62,
		"Date": "2020-06-17T17:30:02Z"
	  }
	],
	"Date": "2020-06-17T17:30:02Z"
  }
`

func TestGetSummary(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc(summaryURL, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, SummaryJSON)
	})

	res, err := client.GetSummary(context.Background())
	if err != nil {
		t.Fatal("err", err)
	}

	if res == nil {
		t.Fatal("res should not be empty", res)
	}
}
