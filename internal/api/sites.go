package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func (a *API) SitesList(page int) sitesListResponse {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/sites", a.Config.BaseURL), nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
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
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "Error: %v\n", response.Status)
		os.Exit(1)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	lr := sitesListResponse{}
	err = json.Unmarshal(body, &lr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	return lr
}

func (a *API) SitesGet(id string) site {

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/sites/%s", a.Config.BaseURL, id), nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	req.Header.Set("Authorization", "Basic "+a.Config.AuthToken)

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stdout, "Error: %v\n", err)
		os.Exit(1)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	if response.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "Error: %v\n", response.Status)
		os.Exit(1)
	}

	s := site{}
	err = json.Unmarshal(body, &s)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	return s

}
