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

// accountsUsersUpdateCmd represents the accounts command
var accountsUsersUpdateCmd = &cobra.Command{
	Use:   "update <account id>",
	Short: "",
	Long:  ``,
	Run:   accountsUsersUpdate,
}

func init() {
	accountsUsersUpdateCmd.Flags().Int("page", 1, "The page to return")
	accountsUsersUpdateCmd.Flags().Int("limit", 100, "Limit the number of results")
}

func accountsUsersUpdate(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		cmd.Usage()
		return
	}

	accountID := args[0]

	page, err := cmd.Flags().GetInt("page")
	if err != nil {
		cmd.PrintErrf("Error: %s\n", err.Error())
		return
	}

	limit, err := cmd.Flags().GetInt("limit")
	if err != nil {
		cmd.PrintErrf("Error: %s\n", err.Error())
		return
	}

	format, err := cmd.Flags().GetString("format")
	if err != nil {
		cmd.PrintErrf("Error: %s\n", err.Error())
		return
	}

	config := cmd.Root().Context().Value(types.ContextKeyCmdConfig).(types.Config)
	config.Limit = limit

	api := api.NewAPI(config)
	users, err := api.AccountsUsersList(accountID, page)
	if err != nil {
		cmd.PrintErrf("Error: %s\n", err.Error())
		return
	}

	if len(users) > 0 {
		if strings.ToLower(format) == "json" {
			j, err := json.Marshal(users)
			if err != nil {
				cmd.PrintErrf("Error: %s\n", err.Error())
				return
			}

			fmt.Fprintf(os.Stdout, "%s\n", j)
			return

		}

		for _, user := range users {
			fmt.Fprintf(os.Stdout, "%s %s %s %s\n", user.UserID, user.FirstName, user.LastName, user.Roles)
		}
	}
}
