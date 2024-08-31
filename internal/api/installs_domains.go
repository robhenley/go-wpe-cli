package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (a *API) InstallDomainsList(installID string, page int) ([]domain, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/installs/%s/domains", a.Config.BaseURL, installID), nil)
	if err != nil {
		return []domain{}, err
	}
	req.Header.Set("Authorization", "Basic "+a.Config.AuthToken)

	q := req.URL.Query()
	offset := (page - 1) * a.Config.Limit
	if offset > 0 {
		q.Add("offset", fmt.Sprintf("%d", offset))
	}

	if a.Config.Limit > 0 {
		q.Add("limit", fmt.Sprintf("%d", a.Config.Limit))
	}
	req.URL.RawQuery = q.Encode()

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return []domain{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return []domain{}, fmt.Errorf("%s", res.Status)
	}

	dr := installDomainsListResponse{}
	err = json.NewDecoder(res.Body).Decode(&dr)
	if err != nil {
		return []domain{}, err
	}

	return dr.Results, nil
}
