package users

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/robhenley/go-wpe-cli/cmd/types"
	"github.com/robhenley/go-wpe-cli/internal/api"
	"github.com/robhenley/go-wpe-cli/internal/helpers"
	"github.com/spf13/cobra"
)

// accountsUsersUpdateCmd represents the accounts command
var accountsUsersUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "",
	Long:  ``,
	Run:   accountsUsersUpdate,
}

func init() {
	accountsUsersUpdateCmd.Flags().StringP("account", "a", "", "The account ID where the user is located")
	accountsUsersUpdateCmd.MarkFlagRequired("account")

	accountsUsersUpdateCmd.Flags().StringP("user", "u", "", "The user ID you'd like to update")
	accountsUsersUpdateCmd.MarkFlagRequired("user")

	accountsUsersUpdateCmd.Flags().StringP("role", "r", "", "The role you'd like to assign to the user")
	accountsUsersUpdateCmd.MarkFlagRequired("role")

	accountsUsersUpdateCmd.Flags().StringSliceP(
		"install-ids",
		"i",
		[]string{},
		"Used with partial role selection. The IDs of the installs the user will have access to.")

}

func accountsUsersUpdate(cmd *cobra.Command, args []string) {
	config := cmd.Root().Context().Value(types.ContextKeyCmdConfig).(types.Config)
	api := api.NewAPI(config)

	accountID, err := cmd.Flags().GetString("account")
	cobra.CheckErr(err)

	userID, err := cmd.Flags().GetString("user")
	cobra.CheckErr(err)

	role, err := cmd.Flags().GetString("role")
	cobra.CheckErr(err)

	if !helpers.IsValidRole(role) {
		cobra.CheckErr(fmt.Errorf("invalid role specified. Valid roles are: %s", strings.Join(helpers.ValidRoles(), " | ")))
	}

	installIDs, err := cmd.Flags().GetStringSlice("install-ids")
	cobra.CheckErr(err)

	format, err := cmd.Flags().GetString("format")
	cobra.CheckErr(err)

	resp, err := api.AccountsUsersUpdate(accountID, userID, role, installIDs)
	cobra.CheckErr(err)

	if strings.ToLower(format) == "json" {
		j, err := json.Marshal(resp)
		cobra.CheckErr(err)

		fmt.Fprintf(os.Stdout, "%s\n", j)
		return
	}

	au := resp.AccountUser
	fmt.Fprintf(os.Stdout, "%s %s installs(%d)\n", au.UserID, au.Roles, len(au.Installs))
}
