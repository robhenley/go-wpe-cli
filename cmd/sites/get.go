package sites

import (
	"fmt"
	"os"

	"github.com/robhenley/go-wpe-cli/cmd/types"
	"github.com/robhenley/go-wpe-cli/internal/api"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get <id>",
	Short: "Get a single site",
	Long:  `Get a single site by ID`,
	Args:  cobra.ExactArgs(1),
	Run:   get,
}

func get(cmd *cobra.Command, args []string) {
	config := cmd.Root().Context().Value(types.ContextKeyCmdConfig).(types.Config)

	if len(args) != 1 {
		fmt.Fprint(os.Stderr, "Please provide a site id")
		os.Exit(1)
	}
	id := args[0]

	api := api.NewAPI(config)
	site := api.SitesGet(id)
	fmt.Fprintf(os.Stdout, "%v\n", site)

}
