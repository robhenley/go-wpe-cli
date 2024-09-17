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

// accountsUsersCreateCmd represents the accounts command
var accountsUsersCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "",
	Long:  ``,
	Run:   accountsUsersCreate,
}

func init() {
	accountsUsersCreateCmd.Flags().StringP("account", "a", "", "The account ID to create the user in")
	accountsUsersCreateCmd.MarkFlagRequired("account")

	accountsUsersCreateCmd.Flags().String("firstname", "", "The first name of the user")
	accountsUsersCreateCmd.MarkFlagRequired("firstname")

	accountsUsersCreateCmd.Flags().String("lastname", "", "The last name of the user")
	accountsUsersCreateCmd.MarkFlagRequired("lastname")

	accountsUsersCreateCmd.Flags().StringP("email", "e", "", "The email of the user")
	accountsUsersCreateCmd.MarkFlagRequired("email")

	accountsUsersCreateCmd.Flags().StringP("role", "r", "", "The role to assign the user")
	accountsUsersCreateCmd.MarkFlagRequired("role")

	accountsUsersCreateCmd.Flags().StringSlice("installs", []string{}, "The install IDs to assign the user")
}

func accountsUsersCreate(cmd *cobra.Command, args []string) {
	config := cmd.Root().Context().Value(types.ContextKeyCmdConfig).(types.Config)
	api := api.NewAPI(config)

	account, err := cmd.Flags().GetString("account")
	cobra.CheckErr(err)

	firstname, err := cmd.Flags().GetString("firstname")
	cobra.CheckErr(err)

	lastname, err := cmd.Flags().GetString("lastname")
	cobra.CheckErr(err)

	email, err := cmd.Flags().GetString("email")
	cobra.CheckErr(err)

	role, err := cmd.Flags().GetString("role")
	cobra.CheckErr(err)

	installs, err := cmd.Flags().GetStringSlice("installs")
	cobra.CheckErr(err)

	format, err := cmd.Flags().GetString("format")
	cobra.CheckErr(err)

	resp, err := api.AccountsUsersCreate(account, firstname, lastname, email, role, installs)
	cobra.CheckErr(err)

	if strings.ToLower(format) == "json" {
		j, err := json.Marshal(resp)
		cobra.CheckErr(err)

		fmt.Fprintf(os.Stdout, "%s\n", j)
	}

	u := resp.AccountUser
	fmt.Fprintf(os.Stdout, "%s %s %s %s\n", u.UserID, u.FirstName, u.LastName, u.Email)

}
