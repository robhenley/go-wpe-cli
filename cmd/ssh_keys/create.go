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

// SSHKeysCreateCmd represents the create command
var SSHKeysCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Add a new SSH key",
	Long:  `Use this to add a new SSH key to WP Engine.`,
	Run:   sshKeysCreate,
}

func init() {
	SSHKeysCreateCmd.Flags().String("key", "", "Path to public key file.")
	SSHKeysCreateCmd.MarkFlagRequired("key")
}

func sshKeysCreate(cmd *cobra.Command, args []string) {
	if len(args) != 0 {
		fmt.Println("Error: This command doesn't require arguments")
		cmd.Usage()
		return
	}

	config := cmd.Root().Context().Value(types.ContextKeyCmdConfig).(types.Config)

	keyFile, err := cmd.Flags().GetString("key")
	if err != nil {
		cmd.PrintErrf("Error: %s\n", err.Error())
		return
	}

	keyContents, err := os.ReadFile(keyFile)
	if err != nil {
		cmd.PrintErrf("Error: %s", err.Error())
		return
	}

	format, err := cmd.Flags().GetString("format")
	if err != nil {
		cmd.PrintErrf("Error: %s\n", err.Error())
		return
	}

	api := api.NewAPI(config)
	key, err := api.SSHKeysCreate(string(keyContents))
	if err != nil {
		cmd.PrintErrf("Error: %s", err.Error())
		return
	}

	if strings.ToLower(format) == "json" {
		j, err := json.Marshal(key)
		if err != nil {
			cmd.PrintErrf("Error: %s\n", err.Error())
			return
		}

		fmt.Fprintf(os.Stdout, "%s\n", j)
		return
	}

	fmt.Fprintf(os.Stdout, "%s %s %s %s\n", key.UUID, key.Fingerprint, key.Comment, key.CreatedAt)
}
