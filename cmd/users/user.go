package users

import (
	"github.com/spf13/cobra"
)

// UsersCmd represents the UsersCmd command
var UsersCmd = &cobra.Command{
	Use:   "users",
	Short: "Command for operations on users",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	UsersCmd.AddCommand(currentUserGetCmd)
}
