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

// installsBackupsCreateCmd represents the accounts command
var installsBackupsCreateCmd = &cobra.Command{
	Use:   "create <install id>",
	Short: "Request a new backup of a WordPress installation",
	Long:  `Kicks off a backup of a WordPress installation`,
	Run:   installsBackupsCreate,
}

func init() {
	installsBackupsCreateCmd.Flags().StringSlice("emails", []string{}, "A comma separated list of emails with no spaces Ex: 1@example.com,2@example.com")
	installsBackupsCreateCmd.Flags().String("description", "", "A description of this backup")
}

func installsBackupsCreate(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		cmd.Usage()
		return
	}

	config := cmd.Root().Context().Value(types.ContextKeyCmdConfig).(types.Config)

	emails, err := cmd.Flags().GetStringSlice("emails")
	if err != nil {
		cmd.PrintErrf("Error: %s", err.Error())
		return
	}

	if len(emails) == 0 {
		if len(config.BackupEmails) == 0 {
			cmd.PrintErr("Error: Please provide the notification emails either via your config or the command flag --emails\n")
			return
		}
		emails = config.BackupEmails
	}

	description, err := cmd.Flags().GetString("description")
	if err != nil {
		cmd.PrintErrf("Error: %s", err.Error())
		return
	}

	if len(description) == 0 {
		if len(config.BackupDescription) == 0 {
			cmd.PrintErr("Error: Please provide a description either via your config or the command flag --description\n")
			return
		}

		description = config.BackupDescription
	}

	format, err := cmd.Flags().GetString("format")
	if err != nil {
		cmd.PrintErrf("Error: %s", err.Error())
		return
	}

	installID := args[0]

	api := api.NewAPI(config)
	backup, err := api.InstallsBackupsCreate(installID, description, emails)
	if err != nil {
		cmd.PrintErrf("Error: %s\n", err.Error())
		return
	}

	if strings.ToLower(format) == "json" {
		j, err := json.Marshal(backup)
		if err != nil {
			cmd.PrintErrf("Error: %s\n", err.Error())
			return
		}

		fmt.Fprintf(os.Stdout, "%s\n", j)
		return
	}

	fmt.Fprintf(os.Stdout, "%s %s\n", backup.ID, backup.Status)
}