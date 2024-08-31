package domains

import (
	"fmt"

	"github.com/spf13/cobra"
)

// domainsGetCmd represents the installsDomainsGet command
var domainsGetCmd = &cobra.Command{
	Use:   "get <install id> <domain id>",
	Short: "Get a specific domain for a given install",
	Long:  `Returns specific domain for a given install`,
	Run:   domainsGet,
}

func domainsGet(cmd *cobra.Command, args []string) {
	if len(args) != 2 {
		fmt.Println("Error: Please provide an install ID and the domain ID")
		cmd.Usage()
		return
	}

	// installID := args[0]
	// domainID := args[1]
}
