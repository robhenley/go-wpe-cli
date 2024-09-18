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
	Use:   "delete",
	Short: "Delete a site",
	Long: ` This will delete the site and any installs associated with this site. 
	This delete is permanent and there is no confirmation prompt`,
	Run: sitesDelete,
}

func init() {
	deleteCmd.Flags().StringP("site", "s", "", "The ID of the site to delete")
	deleteCmd.MarkFlagRequired("site")
}

func sitesDelete(cmd *cobra.Command, args []string) {
	config := cmd.Root().Context().Value(types.ContextKeyCmdConfig).(types.Config)
	api := api.NewAPI(config)

	siteID, err := cmd.Flags().GetString("site")
	cobra.CheckErr(err)

	format, err := cmd.Flags().GetString("format")
	cobra.CheckErr(err)

	isDeleted, err := api.SitesDelete(siteID)
	cobra.CheckErr(err)

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
			cobra.CheckErr(err)

			fmt.Fprintf(os.Stdout, "%s\n", j)
			return

		}

		fmt.Fprintf(os.Stdout, "%s deleted\n", siteID)
		return
	}

}
