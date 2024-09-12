package sites

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/robhenley/go-wpe-cli/cmd/types"
	"github.com/robhenley/go-wpe-cli/internal/api"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create <account id> <name>",
	Short: "Create a new site",
	Long:  `Create a new site`,
	Run:   sitesCreate,
}

func sitesCreate(cmd *cobra.Command, args []string) {

	if len(args) != 2 {
		fmt.Fprint(os.Stderr, "Error: Please provide an account id and a name\n")
		cmd.Usage()
		return
	}

	accountID := args[0]
	name := args[1]

	config := cmd.Root().Context().Value(types.ContextKeyCmdConfig).(types.Config)

	api := api.NewAPI(config)

	site := api.SitesCreate(accountID, name)

	format := cmd.Flag("format").Value.String()
	if strings.ToLower(format) == "json" {
		j, err := json.Marshal(site)
		cobra.CheckErr(err)

		fmt.Fprintf(os.Stdout, "%s\n", j)
		return
	}

	fmt.Fprintf(os.Stdout, "%s %s\n", site.ID, site.Name)

}
