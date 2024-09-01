/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/robhenley/go-wpe-cli/cmd/types"
	"github.com/robhenley/go-wpe-cli/internal/api"
	"github.com/spf13/cobra"
)

// statusCmd represents the status of the WP Engine public API
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "The status of the WP Engine public API",
	Long:  `This endpoint will report the system status and any outages that might be occurring.`,
	Run:   status,
}

func status(cmd *cobra.Command, args []string) {
	config := cmd.Root().Context().Value(types.ContextKeyCmdConfig).(types.Config)
	api := api.NewAPI(config)
	status, err := api.Status()
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
		j, err := json.Marshal(status)
		if err != nil {
			cmd.PrintErrf("Error: %s", err.Error())
			return
		}

		fmt.Fprintf(os.Stdout, "%s\n", j)
		return
	}

	friendlyStatus := "DOWN"
	if status.Success {
		friendlyStatus = "UP"
	}

	fmt.Fprintf(os.Stdout, "%s (%s)\n", friendlyStatus, status.CreatedOn)

}
