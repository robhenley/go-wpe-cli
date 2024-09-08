package api

import (
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

	if res.StatusCode != http.StatusOK {
		return accounts, fmt.Errorf("%s", res.Status)
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

	if res.StatusCode != http.StatusOK {
		return account, fmt.Errorf("%s", res.Status)
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

	if res.StatusCode != http.StatusOK {
		return users, fmt.Errorf("%s", res.Status)
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

	if res.StatusCode != http.StatusOK {
		return u, fmt.Errorf("%s", res.Status)
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

	if res.StatusCode != http.StatusNoContent {
		return od, fmt.Errorf("%s", res.Status)
	}

	od.IsDeleted = true

	return od, nil
}

func (a *API) AccountsUsersCreate(accountID string) error {
	// req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/accounts/%s/account_users", a.Config.BaseURL, accountID), nil)
	return nil
}
