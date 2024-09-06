package ssh_keys

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
	if len(args) != 0 {
		fmt.Println("Error: This command doesn't require arguments")
		cmd.Usage()
		return
	}

	config := cmd.Root().Context().Value(types.ContextKeyCmdConfig).(types.Config)

	page, err := cmd.Flags().GetInt("page")
	if err != nil {
		cmd.PrintErrf("Error: %s\n", err.Error())
		return
	}

	limit, err := cmd.Flags().GetInt("page")
	if err != nil {
		cmd.PrintErrf("Error: %s\n", err.Error())
		return
	}
	config.Limit = limit

	format, err := cmd.Flags().GetString("format")
	if err != nil {
		cmd.PrintErrf("Error: %s\n", err.Error())
		return
	}

	api := api.NewAPI(config)
	keys, err := api.SSHKeysList(page)
	if err != nil {
		cmd.PrintErrf("Error: %s", err.Error())
		return
	}

	if strings.ToLower(format) == "json" {
		j, err := json.Marshal(keys)
		if err != nil {
			cmd.PrintErrf("Error: %s\n", err.Error())
			return
		}

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
