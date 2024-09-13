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

func init() {
	installsGetCmd.Flags().StringP("install", "i", "", "The install ID to get")
	installsGetCmd.MarkFlagRequired("install")
}

func installsGet(cmd *cobra.Command, args []string) {
	config := cmd.Root().Context().Value(types.ContextKeyCmdConfig).(types.Config)
	api := api.NewAPI(config)

	installID, err := cmd.Flags().GetString("install")
	cobra.CheckErr(err)

	install, err := api.InstallsGet(installID)
	cobra.CheckErr(err)

	format, err := cmd.Flags().GetString("format")
	cobra.CheckErr(err)

	if strings.ToLower(format) == "json" {
		j, err := json.Marshal(install)
		cobra.CheckErr(err)

		fmt.Fprintf(os.Stdout, "%s", j)
		return
	}

	fmt.Fprintf(os.Stdout, "%s %s %s\n", install.ID, install.Name, install.Environment)
}
