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

// installsDomainsGetCmd represents the installsDomainsGet command
var installsDomainsGetCmd = &cobra.Command{
	Use:   "get <install id> <domain id>",
	Short: "Get a specific domain for a given install",
	Long:  `Returns a specific domain for a given install`,
	Run:   installsDomainsGet,
}

func installsDomainsGet(cmd *cobra.Command, args []string) {
	if len(args) != 2 {
		fmt.Println("Error: Please provide an install ID and the domain ID")
		cmd.Usage()
		return
	}

	installID := args[0]
	domainID := args[1]

	format, err := cmd.Flags().GetString("format")
	cobra.CheckErr(err)

	config := cmd.Root().Context().Value(types.ContextKeyCmdConfig).(types.Config)
	api := api.NewAPI(config)

	domain, err := api.InstallsDomainsGet(installID, domainID)
	cobra.CheckErr(err)

	if strings.ToLower(format) == "json" {
		j, err := json.Marshal(domain)
		cobra.CheckErr(err)

		fmt.Fprintf(os.Stdout, "%s\n", j)
		return
	}

	fmt.Fprintf(os.Stdout, "%s %s %t %s", domain.ID, domain.Name, domain.Primary, domain.RedirectTo.Name)
}
