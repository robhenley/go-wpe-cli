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
	cobra.CheckErr(err)

	limit, err := cmd.Flags().GetInt("limit")
	cobra.CheckErr(err)

	format, err := cmd.Flags().GetString("format")
	cobra.CheckErr(err)

	config := cmd.Root().Context().Value(types.ContextKeyCmdConfig).(types.Config)
	config.Limit = limit

	api := api.NewAPI(config)
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
