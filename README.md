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
wpe install list --ui | wpe install backup create -
```

```bash
wpe install list | fzf | grep -oE "^([a-zA-Z0-9-]+)" | wpe install backup create -
```

```bash
wpe sites list | fzf --bind 'enter:become(echo {1})' | wpe install backup create -
```

**Purge a sites cache**
```bash
wpe installs list | fzf | grep -oE "^([a-zA-Z0-9-]+)" | wpe installs purge -
```

## Build Instructions
A placeholder for build instructions...

## TODO

- [X] GET    /status
- [ ] GET    /swagger
- [X] GET    /accounts
- [X] GET    /accounts/{account_id}
- [X] GET    /accounts/{account_id}/account_users
- [ ] POST   /accounts/{account_id}/account_users
- [X] GET    /accounts/{account_id}/account_users/{user_id}
- [ ] PATCH  /accounts/{account_id}/account_users/{user_id}
- [X] DELETE /accounts/{account_id}/account_users/{user_id}
- [X] GET    /sites
- [X] POST   /sites
- [X] GET    /sites/{site_id}
- [ ] PATCH  /sites/{site_id}
- [X] DELETE /sites/{site_id}
- [X] GET    /installs
- [ ] POST   /installs
- [X] GET    /installs/{install_id}
- [ ] DELETE /installs/{install_id}
- [ ] PATCH  /installs/{install_id}
- [ ] GET    /installs/{install_id}/domains
- [ ] POST   /installs/{install_id}/domains
- [ ] POST   /installs/{install_id}/domains/bulk
- [ ] GET    /installs/{install_id}/domains/{domain_id}
- [ ] PATCH  /installs/{install_id}/domains/{domain_id}
- [ ] DELETE /installs/{install_id}/domains/{domain_id}
- [ ] POST   /installs/{install_id}/domains/{domain_id}/check_status
- [X] POST   /installs/{install_id}/backups
- [X] GET    /installs/{install_id}/backups/{backup_id}
- [X] GET    /user
- [X] GET    /ssh_keys
- [X] POST   /ssh_keys
- [X] DELETE /ssh_keys