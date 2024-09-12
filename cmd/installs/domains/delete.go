package domains

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/robhenley/go-wpe-cli/cmd/types"
	"github.com/robhenley/go-wpe-cli/internal/api"
	"github.com/spf13/cobra"
)

// installsDomainsDeleteCmd represents the installsDomainsDelete command
var installsDomainsDeleteCmd = &cobra.Command{
	Use:   "delete <install id> <domain id>",
	Short: "Delete a specific domain for a given install",
	Long:  `Delete a specific domain for a given install`,
	Run:   installsDomainsDelete,
}

func installsDomainsDelete(cmd *cobra.Command, args []string) {
	if len(args) != 2 {
		cmd.PrintErr("Error: Please provide an install ID and the domain ID\n")
		cmd.Usage()
		return
	}

	installID := args[0]
	domainID := args[1]

	format, err := cmd.Flags().GetString("format")
	if err != nil {
		cmd.PrintErrf("Error: %s", err.Error())
		return
	}

	config := cmd.Root().Context().Value(types.ContextKeyCmdConfig).(types.Config)
	api := api.NewAPI(config)

	result, err := api.InstallsDomainsDelete(installID, domainID)
	if err != nil {
		cmd.PrintErrf("Error: %s", err.Error())
		return
	}

	if strings.ToLower(format) == "json" {
		j, err := json.Marshal(result)
		if err != nil {
			cmd.PrintErrf("Error: %s", err.Error())
			return
		}

		fmt.Fprintf(os.Stdout, "%s\n", j)
		return
	}

	fmt.Fprintf(os.Stdout, "%s deleted(%t) ", result.ID, result.IsDeleted)
}
