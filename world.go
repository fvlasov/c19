package c19

import (
	"context"
	"encoding/json"
	"net/http"
)

type WorldTotal struct {
	TotalConfirmed int `json:"TotalConfirmed"`
	TotalDeaths    int `json:"TotalDeaths"`
	TotalRecovered int `json:"TotalRecovered"`
}

//GetWorldTotalWIP returnrs 
func GetWorldTotalWIP(ctx context.Context) (*WorldTotal, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", baseURL+summaryURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var res *WorldTotal
	json.NewDecoder(resp.Body).Decode(&res)

	return res, nil
}
