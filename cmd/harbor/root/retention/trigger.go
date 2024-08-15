package retention

import (
	"fmt"
	"github.com/goharbor/harbor-cli/pkg/api"
	"github.com/spf13/cobra"
)

func newTriggerCmd() *cobra.Command {
	var id int64
	var dryRun bool
	cmd := &cobra.Command{
		Use:   "trigger",
		Short: "Trigger a retention execution",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			handler := api.NewRetentionHandler()

			execution, err := handler.TriggerRetentionExecution(ctx, id, dryRun)
			if err != nil {
				return fmt.Errorf("failed to trigger retention execution: %v", err)
			}

			fmt.Printf("Retention execution triggered: %+v\n", execution)
			return nil
		},
	}

	cmd.Flags().Int64Var(&id, "id", 0, "The ID of the retention policy")
	cmd.Flags().BoolVar(&dryRun, "dry-run", false, "If true, only simulate the execution")

	cmd.MarkFlagRequired("id")

	return cmd
}
