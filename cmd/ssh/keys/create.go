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
	config := cmd.Root().Context().Value(types.ContextKeyCmdConfig).(types.Config)
	api := api.NewAPI(config)

	keyFile, err := cmd.Flags().GetString("key")
	cobra.CheckErr(err)

	keyContents, err := os.ReadFile(keyFile)
	cobra.CheckErr(err)

	format, err := cmd.Flags().GetString("format")
	cobra.CheckErr(err)

	key, err := api.SSHKeysCreate(string(keyContents))
	cobra.CheckErr(err)

	if strings.ToLower(format) == "json" {
		j, err := json.Marshal(key)
		cobra.CheckErr(err)

		fmt.Fprintf(os.Stdout, "%s\n", j)
		return
	}

	fmt.Fprintf(os.Stdout, "%s %s %s %s\n", key.UUID, key.Fingerprint, key.Comment, key.CreatedAt)
}
