package retention

import (
	"fmt"

	"github.com/goharbor/harbor-cli/pkg/api"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/models"
	"github.com/spf13/cobra"
)

func newUpdateCmd() *cobra.Command {
	var id int64
	var policy models.RetentionPolicy

	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update an existing retention policy",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			handler := api.NewRetentionHandler()

			err := handler.UpdateRetentionPolicy(ctx, id, &policy)
			if err != nil {
				return fmt.Errorf("failed to update retention policy: %v", err)
			}

			fmt.Println("Retention policy updated successfully")
			return nil
		},
	}

	cmd.Flags().Int64Var(&id, "id", 0, "The ID of the retention policy")
	cmd.Flags().StringVar(&policy.Algorithm, "algorithm", "", "The algorithm used for tag retention")
	// Add more flags for other fields of `policy` as necessary

	cmd.MarkFlagRequired("id")

	return cmd
}
