package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (a *API) APIStatus() {
}

// InstallDomainCDNStatus submits a request to check the status of the domain
func (a *API) InstallDomainCDNStatus(install, domainID string) (InstallDomainCDNStatusResponse, error) {
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/installs/%s/domains/%s/check_status", a.Config.BaseURL, install, domainID), nil)
	if err != nil {
		// fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return InstallDomainCDNStatusResponse{}, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", a.Config.AuthToken))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		// fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return InstallDomainCDNStatusResponse{}, err
	}
	defer res.Body.Close()

	ir := InstallDomainCDNStatusResponse{}
	err = json.NewDecoder(res.Body).Decode(&ir)
	if err != nil {
		// fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return InstallDomainCDNStatusResponse{}, err
	}

	return ir, nil

}

func (a *API) InstallsList(page int, accountID string) (installResponse, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/installs", a.Config.BaseURL), nil)
	if err != nil {
		return installResponse{}, err
	}
	req.Header.Set("Authorization", "Basic "+a.Config.AuthToken)

	q := req.URL.Query()
	if accountID != "" {
		q.Add("account_id", accountID)
	}

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
		return installResponse{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return installResponse{}, fmt.Errorf("%s", res.Status)
	}

	ir := installResponse{}
	err = json.NewDecoder(res.Body).Decode(&ir)
	if err != nil {
		return installResponse{}, err
	}

	return ir, nil

}
