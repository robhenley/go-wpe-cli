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
	format, err := cmd.Flags().GetString("format")
	cobra.CheckErr(err)

	config := cmd.Root().Context().Value(types.ContextKeyCmdConfig).(types.Config)
	api := api.NewAPI(config)
	user, err := api.CurrentUserGet()
	cobra.CheckErr(err)

	if strings.ToLower(format) == "json" {
		j, err := json.Marshal(user)
		cobra.CheckErr(err)

		fmt.Fprintf(os.Stdout, "%s\n", j)
		return
	}

	fmt.Fprintf(os.Stdout, "%s %s %s %s\n", user.ID, user.FirstName, user.LastName, user.Phone)
}
