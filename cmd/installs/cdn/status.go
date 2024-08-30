package cdn

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// cdnStatusCmd represents the CDN status command
var cdnStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "",
	Long:  ``,
	Run:   cdnStatus,
}

func cdnStatus(cmd *cobra.Command, args []string) {
	fmt.Fprintf(os.Stdout, "status called\n")
}
