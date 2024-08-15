package retention

import (
	"fmt"
	"github.com/goharbor/harbor-cli/pkg/api"
	"github.com/spf13/cobra"
)

func newListExecutionsCmd() *cobra.Command {
	var id int64
	var page, pageSize int64

	cmd := &cobra.Command{
		Use:   "list-executions",
		Short: "List retention executions for a policy",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			handler := api.NewRetentionHandler()

			executions, err := handler.ListRetentionExecutions(ctx, id, page, pageSize)
			if err != nil {
				return fmt.Errorf("failed to list retention executions: %v", err)
			}

			fmt.Printf("Retention Executions: %+v\n", executions)
			return nil
		},
	}

	cmd.Flags().Int64Var(&id, "id", 0, "The ID of the retention policy")
	cmd.Flags().Int64Var(&page, "page", 1, "The page number")
	cmd.Flags().Int64Var(&pageSize, "page-size", 10, "The number of executions per page")

	cmd.MarkFlagRequired("id")

	return cmd
}
