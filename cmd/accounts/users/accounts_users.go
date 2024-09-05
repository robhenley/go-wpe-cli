package users

import (
	"github.com/spf13/cobra"
)

// AccountsUsersCmd represents the accounts users command
var AccountsUsersCmd = &cobra.Command{
	Use:   "users",
	Short: "Commands related to the users of an account",
	Long:  `Use this to operate on the users of your account.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Usage()
			return
		}
	},
}

func init() {
	AccountsUsersCmd.AddCommand(accountsUsersListCmd)
}
