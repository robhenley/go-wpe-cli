package installs

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// installsUpdateCmd represents the update command
var installsUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "",
	Long:  ``,
	Run:   installsUpdate,
}

func installsUpdate(cmd *cobra.Command, args []string) {
	fmt.Fprintf(os.Stdout, "update called\n")
}
