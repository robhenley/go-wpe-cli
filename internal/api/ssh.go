package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func (a *API) SSHKeysList(page int) ([]sshKey, error) {
	keys := []sshKey{}

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/ssh_keys", a.Config.BaseURL), nil)
	if err != nil {
		return keys, err
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
		return keys, err
	}
	defer res.Body.Close()

	err = a.checkResponse(res)
	if err != nil {
		return keys, err
	}

	skr := sshKeyResponse{}
	err = json.NewDecoder(res.Body).Decode(&skr)
	if err != nil {
		return keys, err
	}

	return skr.Results, nil
}

func (a *API) SSHKeysCreate(publicKey string) (sshKey, error) {
	key := sshKey{}

	body := struct {
		PublicKey string `json:"public_key"`
	}{
		PublicKey: publicKey,
	}

	j, err := json.Marshal(body)
	if err != nil {
		return key, err
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/ssh_keys", a.Config.BaseURL), bytes.NewReader(j))
	if err != nil {
		return key, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Basic "+a.Config.AuthToken)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return key, err
	}
	defer res.Body.Close()

	err = a.checkResponse(res)
	if err != nil {
		return key, fmt.Errorf("%s", res.Status)
	}

	err = json.NewDecoder(res.Body).Decode(&key)
	if err != nil {
		return key, err
	}

	return key, nil
}

func (a *API) SSHKeyDelete(keyID string) (objDeleted, error) {
	od := objDeleted{
		ID:        keyID,
		IsDeleted: false,
	}

	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/ssh_keys/%s", a.Config.BaseURL, keyID), nil)
	if err != nil {
		return od, err
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
