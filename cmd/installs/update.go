package installs

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

// installsUpdateCmd represents the update command
var installsUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Move an install between sites or switch environment types",
	Long: `Allows you to move an install between different sites in your
account or switch the installs environment between the various environment
types (e.g. dev, staging, production).`,
	Run: installsUpdate,
}

func init() {
	installsUpdateCmd.Flags().StringP("install", "i", "", "The install ID to get")
	installsUpdateCmd.MarkFlagRequired("install")

	installsUpdateCmd.Flags().StringP("site", "s", "", "The site ID to move the install to")
	installsUpdateCmd.Flags().StringP("environment", "e", "", "The environment to change the install to")
}

func installsUpdate(cmd *cobra.Command, args []string) {
	config := cmd.Root().Context().Value(types.ContextKeyCmdConfig).(types.Config)
	api := api.NewAPI(config)

	installID, err := cmd.Flags().GetString("install")
	cobra.CheckErr(err)

	siteID, err := cmd.Flags().GetString("site")
	cobra.CheckErr(err)

	environment, err := cmd.Flags().GetString("environment")
	cobra.CheckErr(err)

	if len(environment) > 0 {
		if !helpers.IsValidEnvironment(environment) {
			cobra.CheckErr(fmt.Errorf("invalid environment type specified. Valid types are: %s", strings.Join(helpers.IsValidEnvironments(), " ")))
		}
	}

	format, err := cmd.Flags().GetString("format")
	cobra.CheckErr(err)

	if siteID == "" && environment == "" {
		cobra.CheckErr("You must specify either a site (--site) or environment (--environment) to update the install to")
	}

	install, err := api.InstallsUpdate(installID, siteID, environment)
	cobra.CheckErr(err)

	if strings.ToLower(format) == "json" {
		j, err := json.Marshal(install)
		cobra.CheckErr(err)

		fmt.Fprintf(os.Stdout, "%s\n", j)
		return
	}

	fmt.Fprintf(os.Stdout, "%s %s %s\n", install.ID, install.Site.ID, install.Environment)
}
