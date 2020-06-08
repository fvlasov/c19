package c19

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type DayOne struct {
	Country     string    `json:"Country"`
	CountryCode string    `json:"CountryCode"`
	Province    string    `json:"Province,omitempty"`
	City        string    `json:"City,omitempty"`
	CityCode    string    `json:"cityCode,omitempty"`
	Lat         string    `json:"Lat"`
	Lon         string    `json:"Lon"`
	Cases       int       `json:"Cases"`
	Status      string    `json:"Status,omitempty"`
	Date        time.Time `json:"Date"`
}

func isAvelableStatus(status string) bool {
	if strings.Contains(status, "confirmed") && strings.Contains(status, "recovered") && strings.Contains(status, "deaths") {
		return true
	}
	return false
}

var baseDayOneURL = baseURL + "/dayone"

//GetDayOne returns all cases by case type for a country from the first recorded case
//country must be the Slug from /countries or /summary. Cases must be one of: confirmed, recovered, deaths
//For more details, see https://documenter.getpostman.com/view/10808728/SzS8rjbc?version=latest#4b88f773-be9b-484f-b521-bb58dda0315c
func GetDayOne(ctx context.Context, country string, status string) ([]DayOne, error) {
	if isAvelableStatus(status) {
		return nil, errStatus
	}

	dayOneURL := fmt.Sprintf("/country/%s/status/%s", country, status)

	req, err := http.NewRequestWithContext(ctx, "GET", baseDayOneURL+dayOneURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var res []DayOne
	json.NewDecoder(resp.Body).Decode(&res)

	return res, nil
}

//GetDayOneAllStatus returns all cases by case type for a country from the first recorded case
//country must be the Slug from /countries or /summary. Cases must be one of: confirmed, recovered, deaths
//For more details, see https://documenter.getpostman.com/view/10808728/SzS8rjbc?version=latest#d0ca988a-ac5f-4d30-ab64-b188e45149e4
func GetDayOneAllStatus(ctx context.Context, country string) ([]DayOne, error) {
	if isAvelableStatus(status) {
		return nil, errStatus
	}

	dayOneAllURL := fmt.Sprintf("dayone/country/%s", country)

	req, err := http.NewRequestWithContext(ctx, "GET", baseDayOneURL+dayOneAllURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var res []DayOne
	json.NewDecoder(resp.Body).Decode(&res)

	return res, nil
}

//GetDayOneLive returns all cases by case type for a country from the first recorded case with the latest record being the live count
//Country must be the Slug from /countries or /summary. Cases must be one of: confirmed, recovered, deaths
//this function is like GetDateOne function but with live count
//For more details, see https://documenter.getpostman.com/view/10808728/SzS8rjbc?version=latest#81447902-b68a-4e79-9df9-1b371905e9fa
func GetDayOneLive(ctx context.Context, country string, status string) ([]DayOne, error) {
	if isAvelableStatus(status) {
		return nil, errStatus
	}

	dayOneURL := fmt.Sprintf("dayone/country/%s/status/%s/live", country, status)

	req, err := http.NewRequestWithContext(ctx, "GET", baseURL+dayOneURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var res []DayOne
	json.NewDecoder(resp.Body).Decode(&res)

	return res, nil
}

//GetDayOneTotal Returns all cases by case type for a country from the first recorded case
//Country must be the slug from /countries or /summary. Cases must be one of: confirmed, recovered, deaths
//For more details, see https://api.covid19api.com/total/dayone/country/south-africa/status/confirmed
func GetDayOneTotal(ctx context.Context, country string, status string) ([]DayOne, error) {
	if !isAvelableStatus(status) {
		return nil, errStatus
	}

	if !IsValidCountry(country) {
		return nil, errCountry
	}

	dayOneURL := fmt.Sprintf("/total/dayone/country/%s/status/%s", country, status)

	req, err := http.NewRequestWithContext(ctx, "GET", baseURL+dayOneURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var res []DayOne
	json.NewDecoder(resp.Body).Decode(&res)

	return res, nil
}
