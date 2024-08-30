package installs

import (
	"fmt"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get <id>",
	Short: "Get an install by ID 2",
	Long:  ` `,
	Run:   installsGet,
}

func installsGet(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		fmt.Println("Error: Please provide an install ID")
		cmd.Usage()
		return
	}

	// id := args[0]

}
