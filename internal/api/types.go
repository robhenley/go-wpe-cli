package api

type paging struct {
	Previous any `json:"previous"`
	Next     any `json:"next"`
	Count    int `json:"count"`
}

type account struct {
	Account struct {
		ID string `json:"id"`
	} `json:"account"`
}

type site struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	account
	GroupName string    `json:"group_name"`
	Tags      []string  `json:"tags"`
	Installs  []install `json:"installs"`
}

type install struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	account
	PhpVersion string `json:"php_version"`
	Status     string `json:"status"`
	Site       struct {
		ID string `json:"id"`
	} `json:"site"`
	Cname         string   `json:"cname"`
	StableIPs     []string `json:"stable_ips"`
	Environment   string   `json:"environment"`
	PrimaryDomain string   `json:"primary_domain"`
	IsMultisite   bool     `json:"is_multisite"`
}

type sitesListResponse struct {
	paging
	Results []site
}

type sitesCreateRequest struct {
	Name      string `json:"name"`
	AccountID string `json:"account_id"`
}

type redirects struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type domain struct {
	Name        string      `json:"name"`
	Duplicate   bool        `json:"duplicate"`
	Primary     bool        `json:"primary"`
	ID          string      `json:"id"`
	RedirectsTo []redirects `json:"redirects_to"`
}

type installDomainsListResponse struct {
	paging
	Results []domain
}

type InstallDomainCDNStatusResponse struct {
	Name        string      `json:"name"`
	Duplicate   bool        `json:"duplicate"`
	Primary     bool        `json:"primary"`
	ID          string      `json:"id"`
	RedirectsTo []redirects `json:"redirects_to"`
}

type installResponse struct {
	paging
	Results []install `json:"results"`
}
