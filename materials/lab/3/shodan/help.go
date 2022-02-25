// Do not use this file directly, do not attemp to compile this source file directly
// Go To lab/3/shodan/main/main.go

package shodan

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HelpInfo struct {
	AvailableFacets []string
	AvailableFilters []string
}

// Returns HelpInfo with currently available facets and filters from the Shodan API.
func (s *Client) Helper() (*HelpInfo, error) {
	res1, err := http.Get(fmt.Sprintf("%s/shodan/host/search/facets?key=%s", BaseURL, s.apiKey))
	if err != nil {
		return nil, err
	}
	defer res1.Body.Close()

	var ret1 []string
	if err := json.NewDecoder(res1.Body).Decode(&ret1); err != nil {
		return nil, err
	}

	res2, err := http.Get(fmt.Sprintf("%s/shodan/host/search/filters?key=%s", BaseURL, s.apiKey))
	if err != nil {
		return nil, err
	}
	defer res2.Body.Close()

	var ret2 []string
	if err := json.NewDecoder(res2.Body).Decode(&ret2); err != nil {
		return nil, err
	}
	return &HelpInfo {
		AvailableFacets: ret1,
		AvailableFilters: ret2,
	}, nil
}