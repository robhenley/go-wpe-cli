package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (a *API) CurrentUserGet() (currentUser, error) {
	cu := currentUser{}

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/user", a.Config.BaseURL), nil)
	if err != nil {
		return cu, err
	}
	req.Header.Set("Authorization", "Basic "+a.Config.AuthToken)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return cu, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return cu, fmt.Errorf("%s", res.Status)
	}

	err = json.NewDecoder(res.Body).Decode(&cu)
	if err != nil {
		return cu, err
	}

	return cu, nil

}
