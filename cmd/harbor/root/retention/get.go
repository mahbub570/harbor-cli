package retention

import (
	"fmt"
	"github.com/goharbor/harbor-cli/pkg/api"
	"github.com/spf13/cobra"
)

func newGetCmd() *cobra.Command {
    var id int64
    cmd := &cobra.Command{
        Use:   "get",
        Short: "Get a retention policy by ID",
        RunE: func(cmd *cobra.Command, args []string) error {
            ctx := cmd.Context()
            handler := api.NewRetentionHandler()

            policy, err := handler.GetRetentionPolicy(ctx, id)
            if err != nil {
                return err // Now we can return the error directly, as it's more informative
            }

            fmt.Printf("Retention Policy: %+v\n", policy)
            return nil
        },
    }

    cmd.Flags().Int64Var(&id, "id", 0, "The ID of the retention policy")
    cmd.MarkFlagRequired("id")

    return cmd
}

// ./harbor retention get --id 6
