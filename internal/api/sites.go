package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/robhenley/go-wpe-cli/internal/helpers"
)

func (a *API) SitesList(page int, filters []string) ([]site, error) {
	sites := []site{}

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/sites", a.Config.BaseURL), nil)
	if err != nil {
		return sites, err
	}
	req.Header.Set("Authorization", "Basic "+a.Config.AuthToken)

	offset := (page - 1) * a.Config.Limit
	q := req.URL.Query()
	if offset > 0 {
		q.Add("offset", fmt.Sprintf("%d", offset))
	}

	if a.Config.Limit > 0 {
		q.Add("limit", fmt.Sprintf("%d", a.Config.Limit))
	}
	req.URL.RawQuery = q.Encode()

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return sites, err
	}
	defer response.Body.Close()

	err = a.checkResponse(response)
	if err != nil {
		return sites, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return sites, err
	}

	slr := sitesListResponse{}
	err = json.Unmarshal(body, &slr)
	if err != nil {
		return sites, err
	}

	if len(filters) > 0 {
		fm := helpers.PrepareFilters(filters)
		for _, site := range slr.Results {
			if helpers.HasTags(fm["tag"], site.Tags) || helpers.HasGroup(fm["group"], site.GroupName) {
				sites = append(sites, site)
			}
		}
		return sites, nil
	}

	return slr.Results, nil

}

func (a *API) SitesGet(id string) (site, error) {
	s := site{}

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/sites/%s", a.Config.BaseURL, id), nil)
	if err != nil {
		return s, err
	}
	req.Header.Set("Authorization", "Basic "+a.Config.AuthToken)

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return s, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return s, err
	}

	err = a.checkResponse(response)
	if err != nil {
		return s, err
	}

	err = json.Unmarshal(body, &s)
	if err != nil {
		return s, err
	}

	return s, nil

}

func (a *API) SitesCreate(accountID, name string) (site, error) {
	s := site{}

	pr := sitesCreateRequest{
		Name:      name,
		AccountID: accountID,
	}

	j, err := json.Marshal(pr)
	if err != nil {
		return s, err
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/sites", a.Config.BaseURL), bytes.NewReader(j))
	if err != nil {
		return s, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+a.Config.AuthToken)

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return s, err
	}
	defer response.Body.Close()

	err = a.checkResponse(response)
	if err != nil {
		return s, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return s, err
	}

	err = json.Unmarshal(body, &s)
	if err != nil {
		return s, err
	}

	return s, nil
}

func (a *API) SitesDelete(id string) (bool, error) {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/sites/%s", a.Config.BaseURL, id), nil)
	if err != nil {
		return false, err
	}
	req.Header.Set("Authorization", "Basic "+a.Config.AuthToken)

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, err
	}
	defer response.Body.Close()

	err = a.checkResponse(response)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (a *API) SitesUpdate(siteID, siteName string) (site, error) {
	s := site{}

	su := site{
		Name: siteName,
	}

	j, err := json.Marshal(su)
	if err != nil {
		return s, err
	}

	req, err := http.NewRequest(http.MethodPatch, fmt.Sprintf("%s/sites/%s", a.Config.BaseURL, siteID), bytes.NewReader(j))
	if err != nil {
		return s, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+a.Config.AuthToken)

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return s, err
	}
	defer response.Body.Close()

	err = a.checkResponse(response)
	if err != nil {
		return s, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return s, err
	}

	err = json.Unmarshal(body, &s)
	if err != nil {
		return s, err
	}

	return s, nil

}
