package sites

type install struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Environment string `json:"environment"`
	Cname       string `json:"cname"`
	PhpVersion  string `json:"php_version"`
	IsMultisite bool   `json:"is_multisite"`
}

type listResponse struct {
	Previous any `json:"previous"`
	Next     any `json:"next"`
	Count    int `json:"count"`
	Results  []struct {
		ID      string `json:"id"`
		Name    string `json:"name"`
		Account struct {
			ID string `json:"id"`
		} `json:"account"`
		GroupName string    `json:"group_name"`
		Tags      []string  `json:"tags"`
		Installs  []install `json:"installs"`
	} `json:"results"`
}
