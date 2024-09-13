package backups

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

// installsBackupsCreateCmd represents the accounts command
var installsBackupsCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Request a new backup of a WordPress installation",
	Long:  `Kicks off a backup of a WordPress installation`,
	Run:   installsBackupsCreate,
}

func init() {
	installsBackupsCreateCmd.Flags().StringP("install", "i", "", "The install ID to create a backup from")
	installsBackupsCreateCmd.MarkFlagRequired("install")

	installsBackupsCreateCmd.Flags().StringSlice("emails", []string{}, "A comma separated list of emails with no spaces Ex: 1@example.com,2@example.com")
	installsBackupsCreateCmd.Flags().String("description", "", "A description of this backup")
}

func installsBackupsCreate(cmd *cobra.Command, args []string) {
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

	emails, err := cmd.Flags().GetStringSlice("emails")
	cobra.CheckErr(err)

	if len(emails) == 0 {
		if len(config.BackupEmails) == 0 {
			cmd.PrintErr("Error: Please provide the notification emails either via your config or the command flag --emails\n")
			return
		}
		emails = config.BackupEmails
	}

	description, err := cmd.Flags().GetString("description")
	cobra.CheckErr(err)

	if len(description) == 0 {
		if len(config.BackupDescription) == 0 {
			cmd.PrintErr("Error: Please provide a description either via your config or the command flag --description\n")
			return
		}

		description = config.BackupDescription
	}

	format, err := cmd.Flags().GetString("format")
	cobra.CheckErr(err)

	backup, err := api.InstallsBackupsCreate(installID, description, emails)
	cobra.CheckErr(err)

	if strings.ToLower(format) == "json" {
		j, err := json.Marshal(backup)
		cobra.CheckErr(err)

		fmt.Fprintf(os.Stdout, "%s\n", j)
		return
	}

	fmt.Fprintf(os.Stdout, "%s %s\n", backup.ID, backup.Status)
}
