package retention

import (
	"strings"
	"strconv"
    "fmt"
    "github.com/goharbor/go-client/pkg/sdk/v2.0/client/retention"
    "github.com/goharbor/harbor-cli/pkg/utils"
    "github.com/spf13/cobra"
)

func DeleteRetentionCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete <retention-id>",
		Short: "Delete a retention policy",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, client, err := utils.ContextWithClient()
			if err != nil {
				return fmt.Errorf("error initializing context and client: %v", err)
			}

			retentionID, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("invalid retention ID: %v", err)
			}

			params := &retention.DeleteRetentionParams{
				ID:      retentionID,
				Context: ctx,
			}

			_, err = client.Retention.DeleteRetention(ctx, params)
			if err != nil {
				switch responseError := err.(type) {
				case *retention.CreateRetentionBadRequest:
					if strings.Contains(responseError.Payload.Errors[0].Message, "already has retention policy") {
						return fmt.Errorf("the project already has a retention policy. If you recently deleted a policy, please wait a few minutes and try again")
					}
					return fmt.Errorf("bad request: %s", formatErrorPayload(responseError.Payload))
				case *retention.CreateRetentionUnauthorized:
					return fmt.Errorf("unauthorized: %s", formatErrorPayload(responseError.Payload))
				case *retention.CreateRetentionForbidden:
					return fmt.Errorf("forbidden: %s", formatErrorPayload(responseError.Payload))
				case *retention.CreateRetentionInternalServerError:
					return fmt.Errorf("internal server error: %s", formatErrorPayload(responseError.Payload))
				default:
					return fmt.Errorf("error creating retention policy: %v", err)
				}
			}

			fmt.Printf("Retention policy %d deletion request successful.\n", retentionID)
			fmt.Println("Note: The policy may still be associated with the project. You may need to wait a few minutes before creating a new policy for this project.")
			return nil
		},
	}

	return cmd
}
// ./harbor retention delete 1
