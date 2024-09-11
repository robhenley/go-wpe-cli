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
	Use:   "delete <install id>",
	Short: "",
	Long:  ``,
	Run:   installsDelete,
}

func installsDelete(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		cmd.Usage()
		return
	}

	installID := args[0]
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

	config := cmd.Root().Context().Value(types.ContextKeyCmdConfig).(types.Config)
	api := api.NewAPI(config)

	result, err := api.InstallsDelete(installID)
	if err != nil {
		cmd.PrintErrf("Error: %s", err.Error())
		return
	}

	format, err := cmd.Flags().GetString("format")
	if err != nil {
		cmd.PrintErrf("Error: %s", err.Error())
		return
	}

	if strings.ToLower(format) == "json" {
		j, err := json.Marshal(result)
		if err != nil {
			cmd.PrintErrf("Error: %s", err.Error())
			return
		}

		fmt.Fprintf(os.Stdout, "%s", j)
		return
	}

	fmt.Fprintf(os.Stdout, "%s deleted(%t)", result.ID, result.IsDeleted)

}
