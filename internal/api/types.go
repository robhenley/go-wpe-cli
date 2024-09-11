package api

import "time"

type paging struct {
	Previous any `json:"previous"`
	Next     any `json:"next"`
	Count    int `json:"count"`
}

type account struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type user struct {
	UserID         string    `json:"user_id"`
	AccountID      string    `json:"account_id"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	Email          string    `json:"email"`
	Phone          string    `json:"phone"`
	InviteAccepted bool      `json:"invite_accepted"`
	MfaEnabled     bool      `json:"mfa_enabled"`
	Roles          string    `json:"roles"`
	LastOwner      bool      `json:"last_owner"`
	Installs       []install `json:"installs"`
}

type userCreateRequest struct {
	User struct {
		AccountID  string   `json:"account_id"`
		FirstName  string   `json:"first_name"`
		LastName   string   `json:"last_name"`
		Email      string   `json:"email"`
		Roles      string   `json:"roles"`
		InstallIds []string `json:"install_ids"`
	} `json:"user"`
}

type site struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Account   account   `json:"account"`
	GroupName string    `json:"group_name"`
	Tags      []string  `json:"tags"`
	Installs  []install `json:"installs"`
}

type install struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	Account    account `json:"account"`
	PhpVersion string  `json:"php_version"`
	Status     string  `json:"status"`
	Site       struct {
		ID string `json:"id"`
	} `json:"site"`
	Cname         string   `json:"cname"`
	StableIPs     []string `json:"stable_ips"`
	Environment   string   `json:"environment"`
	PrimaryDomain string   `json:"primary_domain"`
	IsMultisite   bool     `json:"is_multisite"`
}

type installCreateRequest struct {
	Name        string `json:"name"`
	AccountID   string `json:"account_id"`
	SiteID      string `json:"site_id"`
	Environment string `json:"environment"`
}

type installPurgeCacheRequest struct {
	CacheType string `json:"type"`
}

type installPurgeCacheResponse struct {
	InstallID string `json:"id"`
	CacheType string `json:"type"`
	IsPurged  bool   `json:"is_purged"`
}

type sshKey struct {
	Comment     string    `json:"comment"`
	CreatedAt   time.Time `json:"created_at"`
	Fingerprint string    `json:"fingerprint"`
	UUID        string    `json:"uuid"`
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

type accountsResponse struct {
	paging
	Results []account
}

type accountsUsersResponse struct {
	paging
	Results []user
}

type backupRequest struct {
	Description        string   `json:"description"`
	NotificationEmails []string `json:"notification_emails"`
}

type backupResponse struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

type sshKeyResponse struct {
	paging
	Results []sshKey `json:"results"`
}

type currentUser struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone_number"`
}

// type apiError struct {
// 	Resource string `json:"resource"`
// 	Field    string `json:"field"`
// 	Type     string `json:"type"`
// 	Code     string `json:"code"`
// 	Message  string `json:"message"`
// }
// type errorResponse struct {
// 	Message          string     `json:"message"`
// 	DocumentationURL string     `json:"documentation_url"`
// 	Errors           []apiError `json:"errors"`
// }

type objDeleted struct {
	ID        string `json:"id"`
	IsDeleted bool   `json:"is_deleted"`
}
