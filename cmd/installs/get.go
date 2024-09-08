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

// installsGetCmd represents the get command
var installsGetCmd = &cobra.Command{
	Use:   "get <install id>",
	Short: "Get an install by ID",
	Long:  `Returns a single Install`,
	Run:   installsGet,
}

func installsGet(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		cmd.PrintErr("Error: Please provide an install ID")
		cmd.Usage()
		return
	}

	installID := args[0]
	config := cmd.Root().Context().Value(types.ContextKeyCmdConfig).(types.Config)
	api := api.NewAPI(config)
	install, err := api.InstallsGet(installID)
	if err != nil {
		cmd.PrintErrf("Error: %s", err.Error())
		return
	}

	format, err := cmd.Flags().GetString("format")
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

		fmt.Fprintf(os.Stdout, "%s", j)
		return
	}

	fmt.Fprintf(os.Stdout, "%s %s %s\n", install.ID, install.Name, install.Environment)
}
