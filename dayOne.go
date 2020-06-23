package c19

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type DayOne struct {
	Country     string    `json:"Country"`
	CountryCode string    `json:"CountryCode"`
	Province    string    `json:"Province"`
	City        string    `json:"City"`
	CityCode    string    `json:"cityCode"`
	Lat         string    `json:"Lat"`
	Lon         string    `json:"Lon"`
	Cases       int       `json:"Cases"`
	Status      string    `json:"Status"`
	Date        time.Time `json:"Date"`
}

//GetDayOne returns all cases by case type for a country from the first recorded case
//country must be the Slug from /countries or /summary. Cases must be one of: confirmed, recovered, deaths
//For more details, see https://documenter.getpostman.com/view/10808728/SzS8rjbc?version=latest#4b88f773-be9b-484f-b521-bb58dda0315c
func (c Client) GetDayOne(ctx context.Context, country string, status string) ([]DayOne, error) {
	if !isAvailableStatus(status) {
		return nil, errStatus
	}

	if !IsValidCountry(country) {
		return nil, errCountry
	}

	dayOneURL := fmt.Sprintf("/country/%s/status/%s", country, status)

	req, err := http.NewRequestWithContext(ctx, "GET", c.BaseURL+dayOneURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var res []DayOne
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type DayOneAllStatus struct {
	Country     string    `json:"Country"`
	CountryCode string    `json:"CountryCode"`
	Province    string    `json:"Province"`
	City        string    `json:"City"`
	CityCode    string    `json:"CityCode"`
	Lat         string    `json:"Lat"`
	Lon         string    `json:"Lon"`
	Confirmed   int       `json:"Confirmed"`
	Deaths      int       `json:"Deaths"`
	Recovered   int       `json:"Recovered"`
	Active      int       `json:"Active"`
	Date        time.Time `json:"Date"`
}

//GetDayOneAllStatus returns all cases by case type for a country from the first recorded case
//country must be the Slug from /countries or /summary. Cases must be one of: confirmed, recovered, deaths
//For more details, see https://documenter.getpostman.com/view/10808728/SzS8rjbc?version=latest#d0ca988a-ac5f-4d30-ab64-b188e45149e4
func (c Client) GetDayOneAllStatus(ctx context.Context, country string) ([]DayOneAllStatus, error) {
	if !IsValidCountry(country) {
		return nil, errCountry
	}

	dayOneAllURL := fmt.Sprintf("/dayone/country/%s", country)

	req, err := http.NewRequestWithContext(ctx, "GET", c.BaseURL+dayOneAllURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var res []DayOneAllStatus
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

//GetDayOneLive returns all cases by case type for a country from the first recorded case with the latest record being the live count
//Country must be the Slug from /countries or /summary. Cases must be one of: confirmed, recovered, deaths
//this function is like GetDateOne function but with live count
//For more details, see https://documenter.getpostman.com/view/10808728/SzS8rjbc?version=latest#81447902-b68a-4e79-9df9-1b371905e9fa
func (c Client) GetDayOneLive(ctx context.Context, country string, status string) ([]DayOne, error) {
	if !IsValidCountry(country) {
		return nil, errCountry
	}

	if !isAvailableStatus(status) {
		return nil, errStatus
	}

	dayOneURL := fmt.Sprintf("/dayone/country/%s/status/%s/live", country, status)

	req, err := http.NewRequestWithContext(ctx, "GET", c.BaseURL+dayOneURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var res []DayOne
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

//GetDayOneTotal Returns all cases by case type for a country from the first recorded case
//Country must be the slug from /countries or /summary. Cases must be one of: confirmed, recovered, deaths
//For more details, see https://api.covid19api.com/total/dayone/country/south-africa/status/confirmed
func (c Client) GetDayOneTotal(ctx context.Context, country string, status string) ([]DayOneAllStatus, error) {
	if !isAvailableStatus(status) {
		return nil, errStatus
	}

	if !IsValidCountry(country) {
		return nil, errCountry
	}

	dayOneURL := fmt.Sprintf("/total/dayone/country/%s/status/%s", country, status)

	req, err := http.NewRequestWithContext(ctx, "GET", c.BaseURL+dayOneURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var res []DayOneAllStatus
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
