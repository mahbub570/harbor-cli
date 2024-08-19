package security

import (
	"encoding/json"
	"fmt"
	"github.com/goharbor/harbor-cli/pkg/api"
	summaryview "github.com/goharbor/harbor-cli/pkg/views/security/summary"
	"github.com/spf13/cobra"
)

func getSecuritySummaryCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "summary",
		Short: "Get the security summary of the system",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := api.GetSecuritySummary()
			if err != nil {
				return fmt.Errorf("error getting security summary: %w", err)
			}

			var summaryData summaryview.SecuritySummary
			if err := json.Unmarshal(response, &summaryData); err != nil {
				return fmt.Errorf("error parsing security summary: %w", err)
			}

			summaryview.DisplaySecuritySummary(&summaryData)
			return nil
		},
	}

	return cmd
}