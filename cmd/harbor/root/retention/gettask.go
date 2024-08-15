package retention

import (
	"fmt"
	"github.com/goharbor/harbor-cli/pkg/api"
	"github.com/spf13/cobra"
)

func newGetTaskLogCmd() *cobra.Command {
	var id, executionID, taskID int64

	cmd := &cobra.Command{
		Use:   "get-task-log",
		Short: "Get the log for a retention task",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			handler := api.NewRetentionHandler()

			log, err := handler.GetRetentionTaskLog(ctx, id, executionID, taskID)
			if err != nil {
				return fmt.Errorf("failed to get retention task log: %v", err)
			}

			fmt.Printf("Retention Task Log:\n%s\n", log)
			return nil
		},
	}

	cmd.Flags().Int64Var(&id, "id", 0, "The ID of the retention policy")
	cmd.Flags().Int64Var(&executionID, "execution-id", 0, "The ID of the retention execution")
	cmd.Flags().Int64Var(&taskID, "task-id", 0, "The ID of the task")

	cmd.MarkFlagRequired("id")
	cmd.MarkFlagRequired("execution-id")
	cmd.MarkFlagRequired("task-id")

	return cmd
}
