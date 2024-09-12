package cache

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

// installsCachePurgeCmd
var installsCachePurgeCmd = &cobra.Command{
	Use:   "purge <install id> [cache type]",
	Short: "Purge an installs cache",
	Long: `Purge an installs cache with supported cache types being "page",
"object", or "cdn".  Defaults to "object" cache but the default can be set
with the config key cache_type.`,
	Run: installsCachePurge,
}

func installsCachePurge(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		cmd.Usage()
		return
	}

	config := cmd.Root().Context().Value(types.ContextKeyCmdConfig).(types.Config)
	api := api.NewAPI(config)

	installID := args[0]

	// Allow for explicit stdin
	if args[0] == "-" {
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

	var cacheType string
	if len(args) == 2 {
		cacheType = args[1]
	} else {
		cacheType = setCacheType(config, cacheType)
	}

	result, err := api.InstallsCachePurge(installID, cacheType)
	cobra.CheckErr(err)

	format, err := cmd.Flags().GetString("format")
	cobra.CheckErr(err)

	if strings.ToLower(format) == "json" {
		j, err := json.Marshal(result)
		cobra.CheckErr(err)

		fmt.Fprintf(os.Stdout, "%s\n", j)
		return
	}

	fmt.Fprintf(os.Stdout, "%s %s purged(%t)\n", result.InstallID, result.CacheType, result.IsPurged)
}

func setCacheType(config types.Config, cacheType string) string {
	if cacheType == "" {
		if config.CacheType != "" {
			cacheType = config.CacheType
		} else {
			cacheType = "object"
		}
	}
	return cacheType
}
