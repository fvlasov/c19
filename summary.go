package c19

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

type Summary struct {
	Global    global      `json:"Global"`
	Countries []countries `json:"Countries"`
	Date      time.Time   `json:"Date"`
}

type global struct {
	NewConfirmed   int `json:"NewConfirmed"`
	TotalConfirmed int `json:"TotalConfirmed"`
	NewDeaths      int `json:"NewDeaths"`
	TotalDeaths    int `json:"TotalDeaths"`
	NewRecovered   int `json:"NewRecovered"`
	TotalRecovered int `json:"TotalRecovered"`
}

type countries struct {
	Country        string    `json:"Country"`
	CountryCode    string    `json:"CountryCode"`
	Slug           string    `json:"Slug"`
	NewConfirmed   int       `json:"NewConfirmed"`
	TotalConfirmed int       `json:"TotalConfirmed"`
	NewDeaths      int       `json:"NewDeaths"`
	TotalDeaths    int       `json:"TotalDeaths"`
	NewRecovered   int       `json:"NewRecovered"`
	TotalRecovered int       `json:"TotalRecovered"`
	Date           time.Time `json:"Date"`
}

var summaryURL = "summary"

//GetSummary returns summary of new and total cases per country.
//For more details information see https://documenter.getpostman.com/view/10808728/SzS8rjbc?version=latest#00030720-fae3-4c72-8aea-ad01ba17adf8
func (c Client) GetSummary(ctx context.Context) (*Summary, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", c.BaseURL+summaryURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var res *Summary
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
