package installs

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "",
	Long:  ``,
	Run:   installsUpdate,
}

func installsUpdate(cmd *cobra.Command, args []string) {
	fmt.Fprintf(os.Stdout, "update called\n")
}
