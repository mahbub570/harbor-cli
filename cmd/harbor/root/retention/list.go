package retention

import (

	"fmt"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/client/retention"
	"github.com/goharbor/harbor-cli/pkg/utils"
	"github.com/spf13/cobra"
	"log"
)

func ListRetentionCommand() *cobra.Command {
	var retentionID int64
	var page, pageSize int64

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List retention executions",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Create context and client
			ctx, client, err := utils.ContextWithClient()
			if err != nil {
				return fmt.Errorf("error initializing context and client: %v", err)
			}

			// Set up the parameters
			params := &retention.ListRetentionExecutionsParams{
				ID:        retentionID,
				Page:      &page,
				PageSize:  &pageSize,
			}

			// Debug logging
			log.Printf("Sending request with RetentionID: %d, Page: %d, PageSize: %d\n", retentionID, page, pageSize)

			// Call the API method with the context and parameters
			result, err := client.Retention.ListRetentionExecutions(ctx, params)
			if err != nil {
				// Capture and log the full error message
				log.Printf("API Error: %v\n", err)
				return fmt.Errorf("error fetching retention executions: %v", err)
			}

			// Display the results
			fmt.Println("Retention Executions:")
			for _, execution := range result.Payload {
				fmt.Printf("ID: %d\n", execution.ID)
				fmt.Printf("Trigger: %s\n", execution.Trigger)
				fmt.Printf("Status: %s\n", execution.Status)
				fmt.Println("---")
			}

			return nil
		},
	}

	// Add flags
	cmd.Flags().Int64Var(&retentionID, "retention-id", 0, "Retention policy ID to filter executions")
	cmd.Flags().Int64Var(&page, "page", 1, "Page number")
	cmd.Flags().Int64Var(&pageSize, "page-size", 10, "Number of executions per page")

	return cmd
}
