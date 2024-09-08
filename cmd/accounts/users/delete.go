package users

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/robhenley/go-wpe-cli/cmd/types"
	"github.com/robhenley/go-wpe-cli/internal/api"
	"github.com/spf13/cobra"
)

// accountsUsersDeleteCmd represents the accounts command
var accountsUsersDeleteCmd = &cobra.Command{
	Use:   "delete <account id> <user id>",
	Short: "Delete an account user",
	Long:  `This will remove the association this user has to this account. This delete is permanent and there is no confirmation prompt.`,
	Run:   accountsUsersDelete,
}

func accountsUsersDelete(cmd *cobra.Command, args []string) {
	if len(args) != 2 {
		cmd.Usage()
		return
	}

	accountID := args[0]
	userID := args[1]

	format, err := cmd.Flags().GetString("format")
	if err != nil {
		cmd.PrintErrf("Error: %s\n", err.Error())
		return
	}

	config := cmd.Root().Context().Value(types.ContextKeyCmdConfig).(types.Config)
	api := api.NewAPI(config)

	result, err := api.AccountsUsersDelete(accountID, userID)
	if err != nil {
		cmd.PrintErrf("Error: %s\n", err.Error())
		return
	}

	if strings.ToLower(format) == "json" {
		j, err := json.Marshal(result)
		if err != nil {
			cmd.PrintErrf("Error: %s\n", err.Error())
			return
		}

		fmt.Fprintf(os.Stdout, "%s\n", j)
		return

	}

	fmt.Fprintf(os.Stdout, "%s %t\n", result.ID, result.IsDeleted)
}
