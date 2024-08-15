package retention

import (
	"fmt"
	"github.com/goharbor/harbor-cli/pkg/api"
	"github.com/spf13/cobra"
)

func newGetMetadataCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-metadata",
		Short: "Get retention metadata",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			handler := api.NewRetentionHandler()

			metadata, err := handler.GetRetentionMetadata(ctx)
			if err != nil {
				return fmt.Errorf("failed to get retention metadata: %v", err)
			}

			fmt.Printf("Retention Metadata: %+v\n", metadata)
			return nil
		},
	}

	return cmd
}
// ./harbor retention get-metadata
