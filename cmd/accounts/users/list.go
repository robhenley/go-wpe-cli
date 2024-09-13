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

// accountsUsersListCmd represents the accounts command
var accountsUsersListCmd = &cobra.Command{
	Use:   "list <account id>",
	Short: "List your WP Engine accounts",
	Long:  `Use this to list your WP Engine accounts.`,
	Run:   accountsUsersList,
}

func init() {
	accountsUsersListCmd.Flags().StringP("account", "a", "", "The account ID to list the users from")
	accountsUsersListCmd.MarkFlagRequired("account")

	accountsUsersListCmd.Flags().Int("page", 1, "The page to return")
	accountsUsersListCmd.Flags().Int("limit", 100, "Limit the number of results")
}

func accountsUsersList(cmd *cobra.Command, args []string) {
	config := cmd.Root().Context().Value(types.ContextKeyCmdConfig).(types.Config)
	api := api.NewAPI(config)

	accountID, err := cmd.Flags().GetString("account")
	cobra.CheckErr(err)

	page, err := cmd.Flags().GetInt("page")
	cobra.CheckErr(err)

	limit, err := cmd.Flags().GetInt("limit")
	cobra.CheckErr(err)
	config.Limit = limit

	format, err := cmd.Flags().GetString("format")
	cobra.CheckErr(err)

	users, err := api.AccountsUsersList(accountID, page)
	cobra.CheckErr(err)

	if len(users) > 0 {
		if strings.ToLower(format) == "json" {
			j, err := json.Marshal(users)
			cobra.CheckErr(err)

			fmt.Fprintf(os.Stdout, "%s\n", j)
			return

		}

		for _, user := range users {
			fmt.Fprintf(os.Stdout, "%s %s %s %s\n", user.UserID, user.FirstName, user.LastName, user.Roles)
		}
	}
}
