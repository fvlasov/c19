package c19

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type WorldTotal struct {
	NewConfirmed   int `json:"NewConfirmed"`
	TotalConfirmed int `json:"TotalConfirmed"`
	NewDeaths      int `json:"NewDeaths"`
	TotalDeaths    int `json:"TotalDeaths"`
	NewRecovered   int `json:"NewRecovered"`
	TotalRecovered int `json:"TotalRecovered"`
}

type WorldTotalWIP struct {
	TotalConfirmed int `json:"TotalConfirmed"`
	TotalDeaths    int `json:"TotalDeaths"`
	TotalRecovered int `json:"TotalRecovered"`
}

//GetWorldWIP Returns all cases from all world
func (c Client) GetWorldWIP(ctx context.Context, from time.Time, to time.Time) ([]WorldTotal, error) {
	worldURL := fmt.Sprintf("/world?from=%s&to=%s", from.Format(time.RFC3339Nano), to.Format(time.RFC3339Nano))

	req, err := http.NewRequestWithContext(ctx, "GET", c.BaseURL+worldURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var res []WorldTotal
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

//GetWorldTotalWIP returnrs total recovered and total confirmed and total recovered from all world
func (c Client) GetWorldTotalWIP(ctx context.Context) (*WorldTotalWIP, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", c.BaseURL+"/world/total", nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var res *WorldTotalWIP
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
