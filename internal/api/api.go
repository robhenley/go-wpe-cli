package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/robhenley/go-wpe-cli/cmd/types"
)

type API struct {
	Config types.Config
}

type status struct {
	Success   bool   `json:"success"`
	CreatedOn string `json:"created_on"`
}

func NewAPI(c types.Config) *API {
	return &API{
		Config: c,
	}
}

func (a *API) Status() (status, error) {
	s := status{}
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/status", a.Config.BaseURL), nil)
	if err != nil {
		return s, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return s, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return s, fmt.Errorf("%s", res.Status)
	}

	err = json.NewDecoder(res.Body).Decode(&s)
	if err != nil {
		return s, err
	}

	return s, nil
}

func (e errorResponse) Error() string {
	text := fmt.Sprint(e.Message)

	if len(e.Errors) > 0 {
		text += "\n\nErrors:"
		for _, v := range e.Errors {
			text += fmt.Sprintf("\nResource: %s\nField:%s\nType:%s\nCode:%s\nMessage: %s\n", v.Resource, v.Field, v.Type, v.Code, v.Message)
		}
	}

	return text
}

func (a *API) checkErrorResponse(res *http.Response) error {
	switch res.StatusCode {
	case http.StatusBadRequest, http.StatusInternalServerError, http.StatusUnauthorized:
		er := errorResponse{}
		err := json.NewDecoder(res.Body).Decode(&er)
		if err != nil {
			return err
		}
		return er
	case http.StatusTooManyRequests:
		return errorResponse{
			Message: "API rate limit exceeded",
		}
	case http.StatusServiceUnavailable:
		return errorResponse{
			Message: "API is currently unavailable",
		}

	}

	return nil
}
