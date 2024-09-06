package users

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/robhenley/go-wpe-cli/cmd/types"
	"github.com/robhenley/go-wpe-cli/internal/api"
	"github.com/spf13/cobra"
)

// currentUserGetCmd represents the get command
var currentUserGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the current user",
	Long:  ` `,
	Run:   currentUserGet,
}

func currentUserGet(cmd *cobra.Command, args []string) {
	if len(args) != 0 {
		fmt.Println("Error: This command doesn't require arguments")
		cmd.Usage()
		return
	}

	format, err := cmd.Flags().GetString("format")
	if err != nil {
		cmd.PrintErrf("Error: %s\n", err.Error())
		return
	}

	config := cmd.Root().Context().Value(types.ContextKeyCmdConfig).(types.Config)
	api := api.NewAPI(config)
	user, err := api.CurrentUserGet()
	if err != nil {
		cmd.PrintErrf("Error: %s", err.Error())
		return
	}

	if strings.ToLower(format) == "json" {
		j, err := json.Marshal(user)
		if err != nil {
			cmd.PrintErrf("Error: %s\n", err.Error())
			return
		}

		fmt.Fprintf(os.Stdout, "%s\n", j)
		return
	}

	fmt.Fprintf(os.Stdout, "%s %s %s %s\n", user.ID, user.FirstName, user.LastName, user.Phone)
}
