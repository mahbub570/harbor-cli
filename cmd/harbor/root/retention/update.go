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

func UpdateRetentionCommand() *cobra.Command {
	var id int64
	var scopeLevel, scopeRef, tagSelectors, action string

	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update an existing retention policy",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, client, err := utils.ContextWithClient()
			if err != nil {
				return fmt.Errorf("error initializing context and client: %v", err)
			}

			// First, get the existing policy
			getParams := &retention.GetRetentionParams{
				ID:      id,
				Context: ctx,
			}
			existingPolicy, err := client.Retention.GetRetention(ctx, getParams)
			if err != nil {
				return fmt.Errorf("error fetching existing policy: %v", err)
			}

			// Update only the fields that were provided
			policy := existingPolicy.Payload
			if action != "" {
				policy.Rules[0].Action = action
			}
			if scopeLevel != "" {
				policy.Scope.Level = scopeLevel
			}
			if scopeRef != "" {
				scopeRefInt, err := strconv.ParseInt(scopeRef, 10, 64)
				if err != nil {
					return fmt.Errorf("error parsing scope reference: %v", err)
				}
				policy.Scope.Ref = scopeRefInt
			}
			if tagSelectors != "" {
				var tagSelectorsArray []*models.RetentionSelector
				err = json.Unmarshal([]byte(tagSelectors), &tagSelectorsArray)
				if err != nil {
					return fmt.Errorf("error parsing tag selectors: %v", err)
				}
				policy.Rules[0].TagSelectors = tagSelectorsArray
			}

			// Construct the request parameters
			params := &retention.UpdateRetentionParams{
				ID:      id,
				Policy:  policy,
				Context: ctx,
			}

			// Call the API method
			_, err = client.Retention.UpdateRetention(ctx, params)
			if err != nil {
				// Handle specific error types as before
				// ...
			}

			fmt.Println("Retention policy updated successfully")
			return nil
		},
	}

	cmd.Flags().Int64Var(&id, "id", 0, "ID of the retention policy to update")
	cmd.Flags().StringVar(&scopeLevel, "scope-level", "", "Scope level (e.g., 'project')")
	cmd.Flags().StringVar(&scopeRef, "scope-ref", "", "Scope reference (e.g., project ID)")
	cmd.Flags().StringVar(&tagSelectors, "tag-selectors", "", "Tag selectors in JSON format")
	cmd.Flags().StringVar(&action, "action", "", "Action to take (retain or delete)")

	cmd.MarkFlagRequired("id")

	return cmd
}