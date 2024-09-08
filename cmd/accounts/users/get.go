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

// accountsUsersGetCmd represents the accounts command
var accountsUsersGetCmd = &cobra.Command{
	Use:   "get <account id> <user id>",
	Short: "Get an account user by ID",
	Long:  `Returns a single account user`,
	Run:   accountsUsersGet,
}

func accountsUsersGet(cmd *cobra.Command, args []string) {
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

	user, err := api.AccountsUsersGet(accountID, userID)
	if err != nil {
		cmd.PrintErrf("Error: %s\n", err.Error())
		return
	}

	if strings.ToLower(format) == "json" {
		j, err := json.Marshal(user)
		if err != nil {
			cmd.PrintErrf("Error: %s\n", err.Error())
			return
		}

		fmt.Fprintf(os.Stdout, "%s\n", j)
		return

	}

	fmt.Fprintf(os.Stdout, "%s %s %s %s\n", user.UserID, user.FirstName, user.LastName, user.Roles)
}
