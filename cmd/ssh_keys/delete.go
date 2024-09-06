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

// SSHKeysDeleteCmd represents the delete command
var SSHKeysDeleteCmd = &cobra.Command{
	Use:   "delete <key id>",
	Short: "Delete an existing SSH key",
	Long:  `This will delete the SSH key.`,
	Run:   sshKeysDelete,
}

func sshKeysDelete(cmd *cobra.Command, args []string) {

	if len(args) != 1 {
		fmt.Println("Error: Please provide the key ID")
		cmd.Usage()
		return
	}

	keyID := args[0]
	config := cmd.Root().Context().Value(types.ContextKeyCmdConfig).(types.Config)

	format, err := cmd.Flags().GetString("format")
	if err != nil {
		cmd.PrintErrf("Error: %s\n", err.Error())
		return
	}

	api := api.NewAPI(config)
	objDeleted, err := api.SSHKeyDelete(keyID)
	if err != nil {
		cmd.PrintErrf("Error: %s", err.Error())
		return
	}

	if strings.ToLower(format) == "json" {
		j, err := json.Marshal(objDeleted)
		if err != nil {
			cmd.PrintErrf("Error: %s\n", err.Error())
			return
		}

		fmt.Fprintf(os.Stdout, "%s\n", j)
		return
	}

	fmt.Fprintf(os.Stdout, "%s %t\n", objDeleted.ID, objDeleted.IsDeleted)
}
