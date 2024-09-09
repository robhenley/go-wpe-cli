package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"slices"
	"strings"
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

	if res.StatusCode != http.StatusOK {
		return installResponse{}, fmt.Errorf("%s", res.Status)
	}

	ir := installResponse{}
	err = json.NewDecoder(res.Body).Decode(&ir)
	if err != nil {
		return installResponse{}, err
	}

	return ir, nil

}

func (a *API) InstallsGet(installID string) (install, error) {
	install := install{}

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/installs/%s", a.Config.BaseURL, installID), nil)
	if err != nil {
		return install, err
	}
	req.Header.Set("Authorization", "Basic "+a.Config.AuthToken)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return install, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return install, fmt.Errorf("%s", res.Status)
	}

	err = json.NewDecoder(res.Body).Decode(&install)
	if err != nil {
		return install, err
	}

	return install, nil
}

func (a *API) InstallsCreate(name, accountID, siteID, environment string) (install, error) {
	install := install{}

	ir := installCreateRequest{
		Name:        name,
		AccountID:   accountID,
		SiteID:      siteID,
		Environment: environment,
	}

	j, err := json.Marshal(ir)
	if err != nil {
		return install, err
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/installs", a.Config.BaseURL), bytes.NewReader(j))
	if err != nil {
		return install, err
	}
	req.Header.Set("Authorization", "Basic "+a.Config.AuthToken)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return install, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusCreated {
		return install, fmt.Errorf("%s", res.Status)
	}

	err = json.NewDecoder(res.Body).Decode(&install)
	if err != nil {
		return install, err
	}

	return install, nil
}

func (a *API) InstallsPurgeCache(installID, cacheType string) (installPurgeCacheResponse, error) {
	pr := installPurgeCacheResponse{
		CacheType: cacheType,
		IsPurged:  false,
	}

	if !isValidCacheType(cacheType) {
		return pr, fmt.Errorf("invalid cache type: %s", cacheType)
	}

	purgeReq := installPurgeCacheRequest{
		CacheType: cacheType,
	}

	j, err := json.Marshal(purgeReq)
	if err != nil {
		return pr, err
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/installs/%s/purge_cache", a.Config.BaseURL, installID), bytes.NewBuffer(j))
	if err != nil {
		return pr, err
	}
	req.Header.Set("Authorization", "Basic "+a.Config.AuthToken)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return pr, err
	}

	if res.StatusCode != http.StatusAccepted {
		return pr, fmt.Errorf("%s", res.Status)
	}

	pr.IsPurged = true

	return pr, nil
}

func isValidCacheType(cacheType string) bool {
	valid := []string{"object", "page", "cdn"}

	return slices.Contains(valid, strings.ToLower(cacheType))
}
