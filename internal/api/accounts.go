package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func (a *API) AccountsList(page int) ([]account, error) {
	accounts := []account{}

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/accounts", a.Config.BaseURL), nil)
	if err != nil {
		return accounts, err
	}
	req.Header.Set("Authorization", "Basic "+a.Config.AuthToken)

	q := req.URL.Query()
	offset := (page - 1) * a.Config.Limit
	if offset > 0 {
		q.Add("offset", strconv.Itoa(offset))
	}

	if a.Config.Limit > 0 {
		q.Add("limit", strconv.Itoa(a.Config.Limit))
	}
	req.URL.RawQuery = q.Encode()

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return accounts, err
	}
	defer res.Body.Close()

	err = a.checkResponse(res)
	if err != nil {
		return accounts, err
	}

	ar := accountsResponse{}
	err = json.NewDecoder(res.Body).Decode(&ar)
	if err != nil {
		return accounts, err
	}

	return ar.Results, nil
}

func (a *API) AccountsGet(accountID string) (account, error) {
	account := account{}
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/accounts/%s", a.Config.BaseURL, accountID), nil)
	if err != nil {
		return account, err
	}
	req.Header.Set("Authorization", "Basic "+a.Config.AuthToken)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return account, err
	}
	defer res.Body.Close()

	err = a.checkResponse(res)
	if err != nil {
		return account, err
	}

	err = json.NewDecoder(res.Body).Decode(&account)
	if err != nil {
		return account, err
	}

	return account, nil
}

func (a *API) AccountsUsersList(accountID string, page int) ([]user, error) {
	users := []user{}

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/accounts/%s/account_users", a.Config.BaseURL, accountID), nil)
	if err != nil {
		return users, err
	}
	req.Header.Set("Authorization", "Basic "+a.Config.AuthToken)

	q := req.URL.Query()
	offset := (page - 1) * a.Config.Limit
	if offset > 0 {
		q.Add("offset", strconv.Itoa(offset))
	}

	if a.Config.Limit > 0 {
		q.Add("limit", strconv.Itoa(a.Config.Limit))
	}
	req.URL.RawQuery = q.Encode()

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return users, err
	}
	defer res.Body.Close()

	err = a.checkResponse(res)
	if err != nil {
		return users, err
	}

	ur := accountsUsersResponse{}
	err = json.NewDecoder(res.Body).Decode(&ur)
	if err != nil {
		return users, err
	}

	return ur.Results, nil
}

func (a *API) AccountsUsersGet(accountID, userID string) (user, error) {
	u := user{}

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/accounts/%s/account_users/%s", a.Config.BaseURL, accountID, userID), nil)
	if err != nil {
		return u, err
	}
	req.Header.Set("Authorization", "Basic "+a.Config.AuthToken)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return u, err
	}
	defer res.Body.Close()

	err = a.checkResponse(res)
	if err != nil {
		return u, err
	}

	err = json.NewDecoder(res.Body).Decode(&u)
	if err != nil {
		return u, err
	}

	return u, nil
}

func (a *API) AccountsUsersDelete(accountID, userID string) (objDeleted, error) {
	od := objDeleted{
		ID:        userID,
		IsDeleted: false,
	}

	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/accounts/%s/account_users/%s", a.Config.BaseURL, accountID, userID), nil)
	if err != nil {
		return od, nil
	}
	req.Header.Set("Authorization", "Basic "+a.Config.AuthToken)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return od, err
	}
	defer res.Body.Close()

	err = a.checkResponse(res)
	if err != nil {
		return od, err
	}

	od.IsDeleted = true

	return od, nil
}

func (a *API) AccountsUsersCreate(accountID, firstname, lastname, email, role string, installs []string) (accountUserResponse, error) {
	ucr := accountUserResponse{}

	uc := userCreateRequest{
		userCreateUser{
			FirstName:  firstname,
			LastName:   lastname,
			Email:      email,
			Role:       role,
			InstallIDs: installs,
		},
	}

	j, err := json.Marshal(uc)
	if err != nil {
		return ucr, err
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/accounts/%s/account_users", a.Config.BaseURL, accountID), bytes.NewBuffer(j))
	if err != nil {
		return ucr, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Basic "+a.Config.AuthToken)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return ucr, err
	}
	defer res.Body.Close()

	err = a.checkResponse(res)
	if err != nil {
		return ucr, err
	}

	err = json.NewDecoder(res.Body).Decode(&ucr)
	if err != nil {
		return ucr, err
	}

	return ucr, nil
}

func (a *API) AccountsUsersUpdate(accountID, userID, role string, installs []string) (accountUserResponse, error) {
	aur := accountUserResponse{}

	ur := struct {
		Role       string   `json:"roles"`
		InstallIDs []string `json:"install_ids,omitempty"`
	}{
		Role:       role,
		InstallIDs: installs,
	}

	j, err := json.Marshal(ur)
	if err != nil {
		return aur, err
	}

	req, err := http.NewRequest(http.MethodPatch, fmt.Sprintf("%s/accounts/%s/account_users/%s", a.Config.BaseURL, accountID, userID), bytes.NewBuffer(j))
	if err != nil {
		return aur, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Basic "+a.Config.AuthToken)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return aur, err
	}
	defer res.Body.Close()

	err = a.checkResponse(res)
	if err != nil {
		return aur, err
	}

	err = json.NewDecoder(res.Body).Decode(&aur)
	if err != nil {
		return aur, err
	}

	return aur, nil
}
