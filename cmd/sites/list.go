package sites

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/robhenley/go-wpe-cli/cmd/types"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List your sites",
	Long:  `List the sites you have access to.`,
	Run:   list,
}

func init() {
	listCmd.Flags().Int("limit", 5, "Limit the number of results")
}

func list(cmd *cobra.Command, args []string) {
	config := cmd.Root().Context().Value(types.ContextKeyCmdConfig).(types.Config)

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/sites", config.BaseURL), nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	req.Header.Set("Authorization", "Basic "+config.AuthToken)

	limit, err := cmd.Flags().GetInt("limit")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	q := req.URL.Query()
	q.Add("limit", fmt.Sprintf("%d", limit))
	req.URL.RawQuery = q.Encode()

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "Error: %v\n", response.Status)
		os.Exit(1)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	lr := listResponse{}
	err = json.Unmarshal(body, &lr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	for _, result := range lr.Results {
		for _, install := range result.Installs {
			fmt.Printf("%s\t%s\t%s\t%s\n", install.ID, install.Environment, install.Name, result.Name)
		}
	}

}
