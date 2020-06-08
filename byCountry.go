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
	Lat         string    `json:"Lat"`
	Lon         string    `json:"Lon"`
	Cases       int       `json:"Cases"`
	Status      string    `json:"Status"`
	Date        time.Time `json:"Date"`
}

//GetByCountry returns all cases by case type for a country. Country must be the slug from /countries or /summary
//cases must be one of: confirmed, recovered, deaths
//For more details, see  https://documenter.getpostman.com/view/10808728/SzS8rjbc?version=latest#b07f97ba-24f4-4ebe-ad71-97fa35f3b683
func GetByCountry(ctx context.Context, country string, status string, from time.Time, to time.Time) ([]ByCountry, error) {
	if !IsValidCountry(country) {
		return nil, errCountry
	}

	if !isAvelableStatus(status) {
		return nil, errStatus
	}

	byCountryURL := fmt.Sprintf("dayone/country/%s/status/%s?from=%s&to=%s", country, status, from.Format(time.RFC3339Nano), to.Format(time.RFC3339Nano))
	fmt.Println(byCountryURL)

	req, err := http.NewRequestWithContext(ctx, "GET", baseURL+byCountryURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var res []ByCountry
	json.NewDecoder(resp.Body).Decode(&res)

	return res, nil
}

//GetByCountryAllStatus returns all cases by case type for a country
//country must be the slug from /countries or /summary. Cases must be one of: confirmed, recovered, deaths
//For more details, see  https://documenter.getpostman.com/view/10808728/SzS8rjbc?version=latest#071be6ab-ebcc-40dc-be8b-9209ab7caca5
func GetByCountryAllStatus(ctx context.Context, country string, from time.Time, to time.Time) ([]ByCountry, error) {
	if !IsValidCountry(country) {
		return nil, errCountry
	}

	byCountryURL := fmt.Sprintf("dayone/country/%s?from=%s&to=%s", country, from.Format(time.RFC3339Nano), to.Format(time.RFC3339Nano))
	fmt.Println(byCountryURL)

	req, err := http.NewRequestWithContext(ctx, "GET", baseURL+byCountryURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var res []ByCountry
	json.NewDecoder(resp.Body).Decode(&res)

	return res, nil
}

//GetByCountryLive returns all cases by case type for a country, the latest record being the live count
//Country must be the slug from /countries or /summary. Cases must be one of: confirmed, recovered, deaths
//For more details, see  https://documenter.getpostman.com/view/10808728/SzS8rjbc?version=latest#c34162be-7c20-418e-9866-a24dca632b3c
func GetByCountryLive(ctx context.Context, country string, status string, from time.Time, to time.Time) ([]ByCountry, error) {
	if !IsValidCountry(country) {
		return nil, errCountry
	}

	if !isAvelableStatus(status) {
		return nil, errStatus
	}

	byCountryURL := fmt.Sprintf("dayone/country/%s/status/%s/live?from=%s&to=%s", country, status, from.Format(time.RFC3339Nano), to.Format(time.RFC3339Nano))
	fmt.Println(byCountryURL)

	req, err := http.NewRequestWithContext(ctx, "GET", baseURL+byCountryURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var res []ByCountry
	json.NewDecoder(resp.Body).Decode(&res)

	return res, nil
}
