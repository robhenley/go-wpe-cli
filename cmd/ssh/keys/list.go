package keys

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/robhenley/go-wpe-cli/cmd/types"
	"github.com/robhenley/go-wpe-cli/internal/api"
	"github.com/spf13/cobra"
)

// SSHKeysListCmd represents the list command
var SSHKeysListCmd = &cobra.Command{
	Use:   "list",
	Short: "Get your SSH keys",
	Long:  `Use this to list the SSH keys that you've added to WP Engine.`,
	Run:   sshKeysGet,
}

func init() {
	SSHKeysListCmd.Flags().Int("page", 1, "The page to return")
	SSHKeysListCmd.Flags().Int("limit", 100, "Limit the number of results")
}

func sshKeysGet(cmd *cobra.Command, args []string) {
	config := cmd.Root().Context().Value(types.ContextKeyCmdConfig).(types.Config)
	api := api.NewAPI(config)

	page, err := cmd.Flags().GetInt("page")
	cobra.CheckErr(err)

	limit, err := cmd.Flags().GetInt("page")
	cobra.CheckErr(err)
	config.Limit = limit

	format, err := cmd.Flags().GetString("format")
	cobra.CheckErr(err)

	keys, err := api.SSHKeysList(page)
	cobra.CheckErr(err)

	if strings.ToLower(format) == "json" {
		j, err := json.Marshal(keys)
		cobra.CheckErr(err)

		fmt.Fprintf(os.Stdout, "%s\n", j)
		return
	}

	if len(keys) > 0 {
		for _, key := range keys {
			fmt.Fprintf(os.Stdout, "%s %s %s %s\n", key.UUID, key.Fingerprint, key.Comment, key.CreatedAt)
		}
	} else {
		fmt.Fprint(os.Stdout, "No SSH keys where found\n")
	}
}
