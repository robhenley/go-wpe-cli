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
func (a *API) InstallDomainCDNStatus(install, domainID string) (objAccepted, error) {
	oa := objAccepted{
		ID:         domainID,
		IsAccepted: false,
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/installs/%s/domains/%s/check_status", a.Config.BaseURL, install, domainID), nil)
	if err != nil {
		return oa, err
	}
	req.Header.Set("Authorization", "Basic "+a.Config.AuthToken)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return oa, err
	}
	defer res.Body.Close()

	err = a.checkErrorResponse(res)
	if err != nil {
		return oa, err
	}

	// NOTE: For some reason I'm only ever getting a status code of 202.
	if res.StatusCode == http.StatusAccepted {
		oa.IsAccepted = true
		return oa, nil
	}

	return oa, nil
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

	err = a.checkErrorResponse(res)
	if err != nil {
		return installResponse{}, err
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

	err = a.checkErrorResponse(res)
	if err != nil {
		return install, err
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
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Basic "+a.Config.AuthToken)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return install, err
	}
	defer res.Body.Close()

	err = a.checkErrorResponse(res)
	if err != nil {
		return install, err
	}

	err = json.NewDecoder(res.Body).Decode(&install)
	if err != nil {
		return install, err
	}

	return install, nil
}

func (a *API) InstallsUpdate(installID, siteID, environment string) (install, error) {
	install := install{}

	ur := struct {
		SiteID      string `json:"site_id,omitempty"`
		Environment string `json:"environment,omitempty"`
	}{
		SiteID:      siteID,
		Environment: environment,
	}

	j, err := json.Marshal(ur)
	if err != nil {
		return install, err
	}

	req, err := http.NewRequest(http.MethodPatch, fmt.Sprintf("%s/installs/%s", a.Config.BaseURL, installID), bytes.NewReader(j))
	if err != nil {
		return install, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Basic "+a.Config.AuthToken)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return install, err
	}
	defer res.Body.Close()

	err = a.checkErrorResponse(res)
	if err != nil {
		return install, err
	}

	err = json.NewDecoder(res.Body).Decode(&install)
	if err != nil {
		return install, err
	}

	return install, nil
}

func (a *API) InstallsCachePurge(installID, cacheType string) (installPurgeCacheResponse, error) {
	pr := installPurgeCacheResponse{
		CacheType: cacheType,
		IsPurged:  false,
	}

	if !isValidCacheType(cacheType) {
		return pr, fmt.Errorf("invalid cache type %s", cacheType)
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
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Basic "+a.Config.AuthToken)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return pr, err
	}

	err = a.checkErrorResponse(res)
	if err != nil {
		return pr, err
	}

	return pr, nil
}

func isValidCacheType(cacheType string) bool {
	valid := []string{"object", "page", "cdn"}

	return slices.Contains(valid, strings.ToLower(cacheType))
}

func (a *API) InstallsDelete(installID string) (objDeleted, error) {
	od := objDeleted{
		ID:        installID,
		IsDeleted: false,
	}

	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/installs/%s", a.Config.BaseURL, installID), nil)
	if err != nil {
		return od, err
	}
	req.Header.Set("Authorization", "Basic "+a.Config.BaseURL)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return od, err
	}
	defer res.Body.Close()

	err = a.checkErrorResponse(res)
	if err != nil {
		return od, err
	}

	od.IsDeleted = true

	return od, nil
}
