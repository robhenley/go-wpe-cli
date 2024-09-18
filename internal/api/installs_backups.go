package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (a *API) InstallsBackupsCreate(installID, description string, emails []string) (backupResponse, error) {
	bres := backupResponse{}

	breq := backupRequest{
		Description:        description,
		NotificationEmails: emails,
	}

	j, err := json.Marshal(breq)
	if err != nil {
		return bres, err
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/installs/%s/backups", a.Config.BaseURL, installID), bytes.NewReader(j))
	if err != nil {
		return bres, err
	}
	req.Header.Set("Authorization", "Basic "+a.Config.AuthToken)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return bres, err
	}
	defer res.Body.Close()

	err = a.checkErrorResponse(res)
	if err != nil {
		return bres, err
	}

	err = json.NewDecoder(res.Body).Decode(&bres)
	if err != nil {
		return bres, err
	}

	return bres, nil

}

func (a *API) InstallsBackupsGet(installID, backupID string) (backupResponse, error) {
	bres := backupResponse{}
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/installs/%s/backups/%s", a.Config.BaseURL, installID, backupID), nil)
	if err != nil {
		return bres, err
	}
	req.Header.Set("Authorization", "Basic "+a.Config.AuthToken)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return bres, err
	}
	defer res.Body.Close()

	err = a.checkErrorResponse(res)
	if err != nil {
		return bres, err
	}

	err = json.NewDecoder(res.Body).Decode(&bres)
	if err != nil {
		return bres, err
	}

	return bres, nil

}
