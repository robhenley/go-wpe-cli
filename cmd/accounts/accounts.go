package accounts

import (
	"github.com/robhenley/go-wpe-cli/cmd/accounts/users"
	"github.com/spf13/cobra"
)

// AccountsCmd represents the accounts command
var AccountsCmd = &cobra.Command{
	Use:   "accounts",
	Short: "List your WP Engine accounts",
	Long:  `Use this to list your WP Engine accounts.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Usage()
			return
		}
	},
}

func init() {
	AccountsCmd.AddCommand(accountsListCmd)
	AccountsCmd.AddCommand(accountsGetCmd)
	AccountsCmd.AddCommand(users.AccountsUsersCmd)
}
