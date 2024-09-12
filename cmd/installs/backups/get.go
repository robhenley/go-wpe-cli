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
	Use:   "get <install id> <backup id>",
	Short: "Retrieves the status of a backup of a WordPress installation",
	Long:  `Retrieves the status of a backup of a WordPress installation`,
	Run:   installsBackupsGet,
}

func installsBackupsGet(cmd *cobra.Command, args []string) {
	if len(args) != 2 {
		cmd.Usage()
		return
	}

	installID := args[0]
	backupID := args[1]

	format, err := cmd.Flags().GetString("format")
	cobra.CheckErr(err)

	config := cmd.Root().Context().Value(types.ContextKeyCmdConfig).(types.Config)

	api := api.NewAPI(config)
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
