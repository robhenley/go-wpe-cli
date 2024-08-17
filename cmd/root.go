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

	"github.com/robhenley/go-wpe-cli/cmd/sites"
	"github.com/robhenley/go-wpe-cli/cmd/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "wpe",
	Short: "Unofficial WP Engine CLI",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Without this the config doesn't get set
	// Run: func(cmd *cobra.Command, args []string) {},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
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

	rootCmd.AddCommand(sites.SitesCmd)
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

	data := fmt.Sprintf("%s:%s", username, password)
	token := base64.StdEncoding.EncodeToString([]byte(data))
	config := types.Config{
		BaseURL:   baseURL,
		AuthToken: token,
	}

	ctx := context.WithValue(context.Background(), types.ContextKeyCmdConfig, config)
	rootCmd.SetContext(ctx)

}
