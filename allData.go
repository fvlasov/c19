package c19

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

type AllData struct {
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

var allDataURL = "/all"

//GetAllData returns all daily data
//This call results in 10MB of data being returned and should be used infrequently
//see https://documenter.getpostman.com/view/10808728/SzS8rjbc?version=latest#81415d42-eb53-4a85-8484-42d2349debfe
func (c Client) GetAllData(ctx context.Context) ([]AllData, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", c.BaseURL+allDataURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var res []AllData
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
