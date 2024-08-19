package security

import (
	"encoding/json"
	"fmt"
	"github.com/goharbor/harbor-cli/pkg/api"
	"github.com/goharbor/harbor-cli/pkg/views/security/vulnerabilities"
	"github.com/spf13/cobra"
)

func listVulnerabilitiesCommand() *cobra.Command {
	var query string

	cmd := &cobra.Command{
		Use:   "vulnerabilities",
		Short: "List vulnerabilities in the system",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := api.ListVulnerabilities(query)
			if err != nil {
				return fmt.Errorf("error listing vulnerabilities: %w", err)
			}

			var vulns []vulnerabilities.Vulnerability
			if err := json.Unmarshal(response, &vulns); err != nil {
				return fmt.Errorf("error parsing vulnerabilities: %w", err)
			}

			vulnerabilities.DisplayVulnerabilities(vulns)
			return nil
		},
	}

	cmd.Flags().StringVarP(&query, "query", "q", "", "Query string to filter vulnerabilities")

	return cmd
}