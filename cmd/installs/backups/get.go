package backups

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/robhenley/go-wpe-cli/cmd/types"
	"github.com/robhenley/go-wpe-cli/internal/api"
	"github.com/spf13/cobra"
)

// installsBackupsGetCmd represents the accounts command
var installsBackupsGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieves the status of a backup of a WordPress installation",
	Long:  `Retrieves the status of a backup of a WordPress installation`,
	Run:   installsBackupsGet,
}

func init() {
	installsBackupsGetCmd.Flags().StringP("install", "i", "", "The install ID to get a backup from")
	installsBackupsGetCmd.MarkFlagRequired("install")

	installsBackupsGetCmd.Flags().StringP("backup", "b", "", "The backup ID to get")
	installsBackupsGetCmd.MarkFlagRequired("backup")
}

func installsBackupsGet(cmd *cobra.Command, args []string) {
	config := cmd.Root().Context().Value(types.ContextKeyCmdConfig).(types.Config)
	api := api.NewAPI(config)

	installID, err := cmd.Flags().GetString("install")
	cobra.CheckErr(err)

	backupID, err := cmd.Flags().GetString("backup")
	cobra.CheckErr(err)

	format, err := cmd.Flags().GetString("format")
	cobra.CheckErr(err)

	backup, err := api.InstallsBackupsGet(installID, backupID)
	cobra.CheckErr(err)

	if strings.ToLower(format) == "json" {
		j, err := json.Marshal(backup)
		cobra.CheckErr(err)

		fmt.Fprintf(os.Stdout, "%s\n", j)
		return

	}

	fmt.Fprintf(os.Stdout, "%s %s\n", backup.ID, backup.Status)
}
