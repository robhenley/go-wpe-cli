package users

import (
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
	accountsUsersUpdateCmd.Flags().StringP("account", "a", "", "The account ID to list the users from")
	accountsUsersUpdateCmd.MarkFlagRequired("account")
}

// TODO: Implement accountsUsersUpdate
func accountsUsersUpdate(cmd *cobra.Command, args []string) {
	_, err := cmd.Flags().GetString("account")
	cobra.CheckErr(err)
}
