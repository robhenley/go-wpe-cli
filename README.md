# The Unofficial WP Engine (WPE) CLI

The **unofficial** [WP Engine](https://wpengine.com/) CLI based off their [API](https://wpengineapi.com/).

## Setup
The config location defaults to `$HOME/.config/wpe/config.yaml` but you can also provide the `--config` flag:
```bash
wpe sites list --config path/to/config.yaml
```
The default contents are:
```yaml
base_url: https://api.wpengineapi.com/v1
auth_username: <username>
auth_password: <password>
cache_type: "object"
backup_description: "Plugin Updates"
backup_emails:
  - 1@example.com
  - 2@example.com
  - 3@example.com
```

## Examples of use

**List your sites**
```bash
wpe sites list
```
Formatted as JSON and piped into [jq](https://jqlang.github.io/jq/)
```bash
wpe sites list --format json | \
jq '.Results[] | "\(.id) \(.group_name) \(.name)"' -r
```
**Create a site**
```bash
wpe sites create <account_id> "A new site"
```

**Backup an install**
```bash
wpe install list --ui | wpe install backup create -i -
```

```bash
wpe install list | fzf | grep -oE "^([a-zA-Z0-9-]+)" | wpe install backup create -i -
```

```bash
wpe sites list | fzf --bind 'enter:become(echo {1})' | wpe install backup create -i -
```

**Purge an installs cache**
Purge an installs page cache:
```bash
wpe installs purge 912d4b68-6a7e-4b85-8a3e-95524b63ff41 page
```

Purge an installs object cache:
```bash
wpe installs list --ui | wpe installs purge -
```

Purge an installs CDN cache:
```bash
wpe installs list --ui | wpe installs purge - cdn
```

## Build Instructions
The standard `go build` should work but there is also a [GNU
Make](https://www.gnu.org/software/make/) Makefile.  It has the following usage:

```text
Usage:
  help         print this help message
  tidy         format code and tidy modfile
  audit        run quality control checks
  test         run all tests
  test/cover   run all tests and display coverage
  build        build the application
  run          run the application
```

To build just run:
```bash
make build
```

## TODO

- [X] GET    /status
- [ ] GET    /swagger
- [X] GET    /accounts
- [X] GET    /accounts/{account_id}
- [X] GET    /accounts/{account_id}/account_users
- [X] POST   /accounts/{account_id}/account_users
- [X] GET    /accounts/{account_id}/account_users/{user_id}
- [X] PATCH  /accounts/{account_id}/account_users/{user_id}
- [X] DELETE /accounts/{account_id}/account_users/{user_id}
- [X] GET    /sites
- [X] POST   /sites
- [X] GET    /sites/{site_id}
- [X] PATCH  /sites/{site_id}
- [X] DELETE /sites/{site_id}
- [X] GET    /installs
- [X] POST   /installs
- [X] GET    /installs/{install_id}
- [X] DELETE /installs/{install_id}
- [X] PATCH  /installs/{install_id}
- [X] GET    /installs/{install_id}/domains
- [X] POST   /installs/{install_id}/domains
- [ ] POST   /installs/{install_id}/domains/bulk
- [X] GET    /installs/{install_id}/domains/{domain_id}
- [X] PATCH  /installs/{install_id}/domains/{domain_id}
- [X] DELETE /installs/{install_id}/domains/{domain_id}
- [X] POST   /installs/{install_id}/domains/{domain_id}/check_status
- [X] POST   /installs/{install_id}/backups
- [X] GET    /installs/{install_id}/backups/{backup_id}
- [X] POST   /installs/{install_id}/purge_cache
- [X] GET    /user
- [X] GET    /ssh_keys
- [X] POST   /ssh_keys
- [X] DELETE /ssh_keys