/*
Copyright © 2024 Rob Henley <rob.henley@gmail.com>
*/
package cmd

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"path"

	"github.com/robhenley/go-wpe-cli/cmd/accounts"
	"github.com/robhenley/go-wpe-cli/cmd/installs"
	"github.com/robhenley/go-wpe-cli/cmd/sites"
	"github.com/robhenley/go-wpe-cli/cmd/ssh/keys"
	"github.com/robhenley/go-wpe-cli/cmd/types"
	"github.com/robhenley/go-wpe-cli/cmd/users"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "wpe",
	Short: "The Unofficial WP Engine CLI",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/wpe/config.yaml)")
	rootCmd.PersistentFlags().String("format", "human-readable", "Use this output format. Can be 'human-readable' or 'json'")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(statusCmd)
	rootCmd.AddCommand(sites.SitesCmd)
	rootCmd.AddCommand(installs.InstallsCmd)
	rootCmd.AddCommand(accounts.AccountsCmd)
	rootCmd.AddCommand(users.UsersCmd)
	rootCmd.AddCommand(keys.SSHKeysCmd)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(path.Join(home, ".config/wpe/"))
		viper.SetConfigType("yaml")
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())

	}

	username := viper.GetString("auth_username")
	password := viper.GetString("auth_password")
	baseURL := viper.GetString("base_url")
	backupDescription := viper.GetString("backup_description")
	backupEmails := viper.GetStringSlice("backup_emails")
	cacheType := viper.GetString("cache_type")

	data := fmt.Sprintf("%s:%s", username, password)
	token := base64.StdEncoding.EncodeToString([]byte(data))
	config := types.Config{
		BaseURL:           baseURL,
		AuthToken:         token,
		BackupDescription: backupDescription,
		BackupEmails:      backupEmails,
		CacheType:         cacheType,
	}

	ctx := context.WithValue(context.Background(), types.ContextKeyCmdConfig, config)
	rootCmd.SetContext(ctx)

}
