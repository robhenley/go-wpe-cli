package api

import "github.com/robhenley/go-wpe-cli/cmd/types"

type API struct {
	Config types.Config
}

func NewAPI(c types.Config) *API {
	return &API{
		Config: c,
	}
}
