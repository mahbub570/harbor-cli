package retention

import (
	"fmt"
	"github.com/goharbor/harbor-cli/pkg/api"
	"github.com/spf13/cobra"
)

func newDeleteCmd() *cobra.Command {
	var id int64
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a retention policy by ID",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			handler := api.NewRetentionHandler()

			err := handler.DeleteRetentionPolicy(ctx, id)
			if err != nil {
				return fmt.Errorf("failed to delete retention policy: %v", err)
			}

			fmt.Println("Retention policy deleted successfully")
			return nil
		},
	}

	cmd.Flags().Int64Var(&id, "id", 0, "The ID of the retention policy")
	cmd.MarkFlagRequired("id")

	return cmd
}
// ./harbor retention delete --id 1