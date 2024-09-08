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
	Use:   "create <name> <account id> <site id> <environment>",
	Short: "Create a new WordPress installation",
	Long:  `Creates a new WordPress installation`,
	Run:   installsCreate,
}

func installsCreate(cmd *cobra.Command, args []string) {
	if len(args) != 4 {
		cmd.Usage()
		return
	}

	// TODO: Validation?
	name := args[0]
	accountID := args[1]
	siteID := args[2]
	environment := args[3]

	format, err := cmd.Flags().GetString("format")
	if err != nil {
		cmd.PrintErrf("Error: %s", err.Error())
		return
	}

	config := cmd.Root().Context().Value(types.ContextKeyCmdConfig).(types.Config)
	api := api.NewAPI(config)

	install, err := api.InstallsCreate(name, accountID, siteID, environment)
	if err != nil {
		cmd.PrintErrf("Error: %s", err.Error())
		return
	}

	if strings.ToLower(format) == "json" {
		j, err := json.Marshal(install)
		if err != nil {
			cmd.PrintErrf("Error: %s", err.Error())
			return
		}

		fmt.Fprintf(os.Stdout, "%s\n", j)
		return
	}

	fmt.Fprintf(os.Stdout, "%s\t%-15s\t%-15s\t%s\n", install.ID, install.Environment, install.Name, install.PrimaryDomain)
}
