package retention

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/goharbor/go-client/pkg/sdk/v2.0/client/retention"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/models"
	"github.com/goharbor/harbor-cli/pkg/utils"
	"github.com/spf13/cobra"
)

func CreateRetentionCommand() *cobra.Command {
	var scopeLevel, scopeRef, tagSelectors, action string
	var dryRun bool

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new retention policy",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, client, err := utils.ContextWithClient()
			if err != nil {
				return fmt.Errorf("error initializing context and client: %v", err)
			}

			// Parse tag selectors
			var tagSelectorsArray []*models.RetentionSelector
			err = json.Unmarshal([]byte(tagSelectors), &tagSelectorsArray)
			if err != nil {
				return fmt.Errorf("error parsing tag selectors: %v", err)
			}

			// Convert scopeRef to int64
			scopeRefInt, err := strconv.ParseInt(scopeRef, 10, 64)
			if err != nil {
				return fmt.Errorf("error parsing scope reference: %v", err)
			}

			// Construct the retention policy
			policy := &models.RetentionPolicy{
				Algorithm: "or",
				Rules: []*models.RetentionRule{
					{
						Action: action,
						ScopeSelectors: map[string][]models.RetentionSelector{
							"repository": {
								{
									Kind:       "doublestar",
									Decoration: "repoMatches",
									Pattern:    "**",
								},
							},
						},
						TagSelectors: tagSelectorsArray,
						Template:     "always",
					},
				},
				Scope: &models.RetentionPolicyScope{
					Level: scopeLevel,
					Ref:   scopeRefInt,
				},
				Trigger: &models.RetentionRuleTrigger{
					Kind: "Schedule",
					Settings: map[string]interface{}{
						"cron": "0 0 * * * *",
					},
				},
			}

			// Construct the request parameters
			params := &retention.CreateRetentionParams{
				Policy:  policy,
				Context: ctx,
			}

			// Call the API method
			_, err = client.Retention.CreateRetention(ctx, params)
			if err != nil {
				// Try to extract more detailed error information
				switch responseError := err.(type) {
				case *retention.CreateRetentionBadRequest:
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

			fmt.Println("Retention policy created successfully")
			return nil
		},
	}

	cmd.Flags().StringVar(&scopeLevel, "scope-level", "project", "Scope level (e.g., 'project')")
	cmd.Flags().StringVar(&scopeRef, "scope-ref", "", "Scope reference (e.g., project ID)")
	cmd.Flags().StringVar(&tagSelectors, "tag-selectors", "[]", "Tag selectors in JSON format")
	cmd.Flags().StringVar(&action, "action", "retain", "Action to take (retain or delete)")
	cmd.Flags().BoolVar(&dryRun, "dry-run", false, "Dry run (true or false)")

	cmd.MarkFlagRequired("scope-ref")

	return cmd
}

// Helper function to format error payload
func formatErrorPayload(payload interface{}) string {
	// Customize this function based on the actual structure of the payload
	payloadBytes, _ := json.Marshal(payload)
	return string(payloadBytes)
}


// ./harbor retention create \
//   --scope-level="project" \
//   --scope-ref="4" \
//   --tag-selectors='[{"kind":"doublestar","decoration":"repoMatches","pattern":"**"}]' \
//   --action="retain" \
//   --dry-run=false
