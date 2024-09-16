package api

import (
	"bytes"
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

func (a *API) InstallsDomainsGet(installID, domainID string) (domain, error) {
	d := domain{}
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/installs/%s/domains/%s", a.Config.BaseURL, installID, domainID), nil)
	if err != nil {
		return d, err
	}
	req.Header.Set("Authorization", "Basic "+a.Config.AuthToken)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return d, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return d, fmt.Errorf("%s", res.Status)
	}

	err = json.NewDecoder(res.Body).Decode(&d)
	if err != nil {
		return d, err
	}

	return d, nil
}

func (a *API) InstallsDomainsDelete(installID, domainID string) (objDeleted, error) {
	od := objDeleted{
		ID:        domainID,
		IsDeleted: false,
	}

	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/installs/%s/domains/%s", a.Config.BaseURL, installID, domainID), nil)
	if err != nil {
		return od, err
	}
	req.Header.Set("Authorization", "Basic "+a.Config.AuthToken)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return od, err
	}
	defer res.Body.Close()

	err = a.checkBadRequest(res)
	if err != nil {
		return od, err
	}

	if res.StatusCode != http.StatusNoContent {
		return od, fmt.Errorf("%s", res.Status)
	}

	od.IsDeleted = true

	return od, nil
}

// NOTE: For /installs/{install_id}/domains it returns redirectS_to not redirect_to
func (a *API) InstallsDomainsCreate(installID, name, redirect string, primary bool) (domain, error) {
	d := domain{}

	idcr := installDomainCreateRequest{
		Domain:     name,
		Primary:    primary,
		RedirectTo: redirect,
	}

	j, err := json.Marshal(idcr)
	if err != nil {
		return d, err
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/installs/%s/domains", a.Config.BaseURL, installID), bytes.NewReader(j))
	if err != nil {
		return d, err
	}
	req.Header.Set("Authorization", "Basic "+a.Config.AuthToken)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return d, err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusBadRequest {
		er := errorResponse{}
		err := json.NewDecoder(res.Body).Decode(&er)
		if err != nil {
			return d, err
		}

		return d, er
	}

	if res.StatusCode != http.StatusCreated {
		return d, fmt.Errorf("%s", res.Status)
	}

	err = json.NewDecoder(res.Body).Decode(&d)
	if err != nil {
		return d, err
	}

	return d, nil
}
