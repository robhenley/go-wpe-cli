package types

type contextKey string

var (
	ContextKeyCmdConfig = contextKey("app-config")
)

type Config struct {
	BaseURL           string
	AuthToken         string
	BackupDescription string
	BackupEmails      []string
	Limit             int
	CacheType         string
}
