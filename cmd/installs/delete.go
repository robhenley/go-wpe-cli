package installs

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "",
	Long:  ``,
	Run:   installsDelete,
}

func installsDelete(cmd *cobra.Command, args []string) {
	fmt.Fprintf(os.Stdout, "delete called\n")
}
