package c19

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type ByCountry struct {
	Country     string    `json:"Country"`
	CountryCode string    `json:"CountryCode"`
	Province    string    `json:"Province"`
	City        string    `json:"City"`
	CityCode    string    `json:"CityCode"`
	Lat         string    `json:"Lat"`
	Lon         string    `json:"Lon"`
	Cases       int       `json:"Cases"`
	Status      string    `json:"Status"`
	Date        time.Time `json:"Date"`
}

//GetByCountry returns all cases by case type for a country. Country must be the slug from /countries or /summary
//cases must be one of: confirmed, recovered, deaths
//For more details, see  https://documenter.getpostman.com/view/10808728/SzS8rjbc?version=latest#b07f97ba-24f4-4ebe-ad71-97fa35f3b683
func (c Client) GetByCountry(ctx context.Context, country string, status string, from time.Time, to time.Time) ([]ByCountry, error) {
	if !isAvailableStatus(status) {
		return nil, errStatus
	}

	if !IsValidCountry(country) {
		return nil, errCountry
	}

	byCountryURL := fmt.Sprintf("/dayone/country/%s/status/%s?from=%s&to=%s", country, status, from.Format(time.RFC3339Nano), to.Format(time.RFC3339Nano))

	req, err := http.NewRequestWithContext(ctx, "GET", c.BaseURL+byCountryURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var res []ByCountry
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type ByCountryAllStatus struct {
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

//GetByCountryAllStatus returns all cases by case type for a country
//country must be the slug from /countries or /summary. Cases must be one of: confirmed, recovered, deaths
//For more details, see  https://documenter.getpostman.com/view/10808728/SzS8rjbc?version=latest#071be6ab-ebcc-40dc-be8b-9209ab7caca5
func (c Client) GetByCountryAllStatus(ctx context.Context, country string, from time.Time, to time.Time) ([]ByCountryAllStatus, error) {
	if !IsValidCountry(country) {
		return nil, errCountry
	}

	byCountryURL := fmt.Sprintf("/dayone/country/%s?from=%s&to=%s", country, from.Format(time.RFC3339Nano), to.Format(time.RFC3339Nano))

	req, err := http.NewRequestWithContext(ctx, "GET", c.BaseURL+byCountryURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var res []ByCountryAllStatus
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

//GetByCountryLive returns all cases by case type for a country, the latest record being the live count
//Country must be the slug from /countries or /summary. Cases must be one of: confirmed, recovered, deaths
//For more details, see  https://documenter.getpostman.com/view/10808728/SzS8rjbc?version=latest#c34162be-7c20-418e-9866-a24dca632b3c
func (c Client) GetByCountryLive(ctx context.Context, country string, status string, from time.Time, to time.Time) ([]ByCountry, error) {
	if !isAvailableStatus(status) {
		return nil, errStatus
	}

	if !IsValidCountry(country) {
		return nil, errCountry
	}

	byCountryURL := fmt.Sprintf("/dayone/country/%s/status/%s/live?from=%s&to=%s", country, status, from.Format(time.RFC3339Nano), to.Format(time.RFC3339Nano))

	req, err := http.NewRequestWithContext(ctx, "GET", c.BaseURL+byCountryURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var res []ByCountry
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
