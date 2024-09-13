package installs

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/robhenley/go-wpe-cli/cmd/types"
	"github.com/robhenley/go-wpe-cli/internal/api"
	"github.com/spf13/cobra"
)

// installsCreateCmd represents the create command
var installsCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new WordPress installation",
	Long:  `Creates a new WordPress installation`,
	Run:   installsCreate,
}

func init() {
	installsCreateCmd.Flags().StringP("name", "n", "", "The name of the install to create")
	installsCreateCmd.MarkFlagRequired("name")

	installsCreateCmd.Flags().StringP("account", "a", "", "The account ID to create the install in")
	installsCreateCmd.MarkFlagRequired("account")

	installsCreateCmd.Flags().StringP("site", "s", "", "The site ID to create the install in")
	installsCreateCmd.MarkFlagRequired("site")

	installsCreateCmd.Flags().StringP("environment", "e", "", "The environment to create install in")
	installsCreateCmd.MarkFlagRequired("environment")
}

func installsCreate(cmd *cobra.Command, args []string) {
	config := cmd.Root().Context().Value(types.ContextKeyCmdConfig).(types.Config)
	api := api.NewAPI(config)

	name, err := cmd.Flags().GetString("name")
	cobra.CheckErr(err)

	accountID, err := cmd.Flags().GetString("account")
	cobra.CheckErr(err)

	siteID, err := cmd.Flags().GetString("site")
	cobra.CheckErr(err)

	environment, err := cmd.Flags().GetString("environment")
	cobra.CheckErr(err)

	format, err := cmd.Flags().GetString("format")
	cobra.CheckErr(err)

	install, err := api.InstallsCreate(name, accountID, siteID, environment)
	cobra.CheckErr(err)

	if strings.ToLower(format) == "json" {
		j, err := json.Marshal(install)
		cobra.CheckErr(err)

		fmt.Fprintf(os.Stdout, "%s\n", j)
		return
	}

	fmt.Fprintf(os.Stdout, "%s\t%-15s\t%-15s\t%s\n", install.ID, install.Environment, install.Name, install.PrimaryDomain)
}
