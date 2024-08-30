package domains

import (
	"fmt"

	"github.com/spf13/cobra"
)

// installsDomainsGetCmd represents the installsDomainsGet command
var domainsGetCmd = &cobra.Command{
	Use:   "get <id>",
	Short: "",
	Long:  ``,
	Run:   domainsGet,
}

func domainsGet(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		fmt.Println("Error: Please provide an install ID")
		cmd.Usage()
		return
	}

	// id := args[0]
}
