package cdn

import (
	"github.com/spf13/cobra"
)

// cdnCmd represents the cdn command
var CdnCmd = &cobra.Command{
	Use:   "cdn",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

func init() {
	CdnCmd.AddCommand(cdnStatusCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cdnCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cdnCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
