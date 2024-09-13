package installs

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/robhenley/go-wpe-cli/cmd/types"
	"github.com/robhenley/go-wpe-cli/internal/api"
	"github.com/robhenley/go-wpe-cli/internal/ui"
	"github.com/spf13/cobra"
)

// installsListCmd represents the list command
var installsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List your WordPress installations",
	Long:  `List your WordPress installations`,
	Run:   installsList,
}

func init() {
	installsListCmd.Flags().StringP("account", "a", "", "The account ID to list installs from")
	installsListCmd.Flags().Int("page", 1, "The page to return")
	installsListCmd.Flags().Int("limit", 100, "Limit the number of results")
	installsListCmd.Flags().Bool("ui", false, "Enable the fuzzy finder (fzf) UI")
}

func installsList(cmd *cobra.Command, args []string) {
	config := cmd.Root().Context().Value(types.ContextKeyCmdConfig).(types.Config)
	api := api.NewAPI(config)

	accountID, err := cmd.Flags().GetString("account")
	cobra.CheckErr(err)

	page, err := cmd.Flags().GetInt("page")
	cobra.CheckErr(err)

	limit, err := cmd.Flags().GetInt("limit")
	cobra.CheckErr(err)
	config.Limit = limit

	enableUI, err := cmd.Flags().GetBool("ui")
	cobra.CheckErr(err)

	installs, err := api.InstallsList(page, accountID)
	cobra.CheckErr(err)

	format, err := cmd.Flags().GetString("format")
	cobra.CheckErr(err)

	if strings.ToLower(format) == "json" {

		j, err := json.Marshal(installs.Results)
		cobra.CheckErr(err)

		fmt.Fprintf(os.Stdout, "%s\n", j)
		return
	}

	if enableUI {
		var lines []string
		for _, install := range installs.Results {
			lines = append(lines, fmt.Sprintf("%s %-15s %-15s %s\n", install.ID, install.Environment, install.Name, install.PrimaryDomain))
		}

		code, err := ui.Display(lines)
		if err != nil {
			cmd.PrintErrf("Error: %s", err.Error())
			os.Exit(code)
		}

		return
	}

	for _, install := range installs.Results {
		fmt.Fprintf(os.Stdout, "%s\t%-15s\t%-15s\t%s\n", install.ID, install.Environment, install.Name, install.PrimaryDomain)
	}
}
