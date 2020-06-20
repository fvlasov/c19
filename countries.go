package c19

import (
	"context"
	"encoding/json"
	"net/http"
)

type Countries struct {
	Country string `json:"Country"`
	Slug    string `json:"Slug"`
	ISO2    string `json:"ISO2"`
}

var countriesURL = "/countries"

//GetCountries returns all the available countries and provinces, as well as the country slug for per country requests
//For more details, see https://documenter.getpostman.com/view/10808728/SzS8rjbc?version=latest#7934d316-f751-4914-9909-39f1901caeb8
func (c Client) GetCountries(ctx context.Context) ([]Countries, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", c.BaseURL+countriesURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var res []Countries
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
