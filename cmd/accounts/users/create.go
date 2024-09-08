package users

import (
	"github.com/spf13/cobra"
)

// accountsUsersCreateCmd represents the accounts command
var accountsUsersCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "",
	Long:  ``,
	Run:   accountsUsersCreate,
}

func accountsUsersCreate(cmd *cobra.Command, args []string) {
}
