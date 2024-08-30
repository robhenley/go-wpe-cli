package installs

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "",
	Long:  ``,
	Run:   installsList,
}

func installsList(cmd *cobra.Command, args []string) {
	fmt.Fprintf(os.Stdout, "list called\n")
}
