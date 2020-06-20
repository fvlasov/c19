package c19

import (
	"context"
	"io/ioutil"
	"net/http"
)

//GetVersion returns api version
//For more details information see https://documenter.getpostman.com/view/10808728/SzS8rjbc?version=latest#efea2a6d-af2e-433b-9a1f-b8a266a0077f
func (c Client) GetVersion(ctx context.Context) (string, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", "https://api.covid19api.com/version", nil)
	if err != nil {
		return "", err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", nil
	}

	return string(b), nil
}
