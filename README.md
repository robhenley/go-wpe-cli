# The Unofficial WPEngine (WPE) CLI

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

**Backup a site**
```bash
wpe sites list | fzf | grep -oE "^([a-zA-Z0-9-]+)" | wp install backup
```

**Purge a sites cache**
```bash
wpe installs list | fzf | grep -oE "^([a-zA-Z0-9-]+)" | wpe installs purge
```
