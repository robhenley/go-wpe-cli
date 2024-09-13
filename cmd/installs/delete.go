package installs

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/robhenley/go-wpe-cli/cmd/types"
	"github.com/robhenley/go-wpe-cli/internal/api"
	"github.com/spf13/cobra"
)

// installsDeleteCmd represents the delete command
var installsDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "",
	Long:  ``,
	Run:   installsDelete,
}

func init() {
	installsDeleteCmd.Flags().StringP("install", "i", "", "The install ID to create a backup from")
	installsDeleteCmd.MarkFlagRequired("install")
}

func installsDelete(cmd *cobra.Command, args []string) {
	config := cmd.Root().Context().Value(types.ContextKeyCmdConfig).(types.Config)
	api := api.NewAPI(config)

	installID, err := cmd.Flags().GetString("install")
	cobra.CheckErr(err)

	if installID == "-" {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			installID = scanner.Text()
		}

		if err := scanner.Err(); err != nil {
			cmd.PrintErrf("Error reading from stdin: %s\n", err.Error())
			return
		}

		installID = strings.Trim(installID, " ")
	}

	result, err := api.InstallsDelete(installID)
	cobra.CheckErr(err)

	format, err := cmd.Flags().GetString("format")
	cobra.CheckErr(err)

	if strings.ToLower(format) == "json" {
		j, err := json.Marshal(result)
		cobra.CheckErr(err)

		fmt.Fprintf(os.Stdout, "%s", j)
		return
	}

	fmt.Fprintf(os.Stdout, "%s deleted(%t)", result.ID, result.IsDeleted)

}
