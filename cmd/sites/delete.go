package sites

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/robhenley/go-wpe-cli/cmd/types"
	"github.com/robhenley/go-wpe-cli/internal/api"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete <site id>",
	Short: "Delete a site",
	Long: ` This will delete the site and any installs associated with this site. 
	This delete is permanent and there is no confirmation prompt`,
	Run: sitesDelete,
}

func sitesDelete(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		fmt.Fprint(os.Stderr, "Error: Please provide a site id\n")
		cmd.Usage()
		return
	}

	siteID := args[0]

	format := cmd.Flag("format").Value.String()
	config := cmd.Root().Context().Value(types.ContextKeyCmdConfig).(types.Config)
	api := api.NewAPI(config)

	isDeleted := api.SitesDelete(siteID)

	out := struct {
		IsDeleted bool   `json:"is_deleted"`
		SiteID    string `json:"site_id"`
	}{
		IsDeleted: isDeleted,
		SiteID:    siteID,
	}

	if isDeleted {
		if strings.ToLower(format) == "json" {

			j, err := json.Marshal(out)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
			}

			fmt.Fprintf(os.Stdout, "%s\n", j)
			return

		}

		fmt.Fprintf(os.Stdout, "%s deleted\n", siteID)
		return
	}

}
