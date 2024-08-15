package retention

import (
	"fmt"
	"encoding/json"

	"github.com/goharbor/harbor-cli/pkg/api"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/models"
	"github.com/spf13/cobra"
)
func newCreateCmd() *cobra.Command {
    var policy models.RetentionPolicy
    policy.Scope = &models.RetentionPolicyScope{}
    policy.Trigger = &models.RetentionRuleTrigger{}

    // Declare variables for rule parameters
    var ruleAction, ruleTemplate, tagKind, tagPattern, scopeKind, scopePattern string

    cmd := &cobra.Command{
        Use:   "create",
        Short: "Create a new retention policy",
        Long:  "Create a new retention policy in Harbor",
        RunE: func(cmd *cobra.Command, args []string) error {
            ctx := cmd.Context()
            handler := api.NewRetentionHandler()

            // Validate input
            if policy.Scope.Level == "" || policy.Scope.Ref == 0 {
                return fmt.Errorf("scope level and reference are required")
            }
            if policy.Trigger.Kind == "" {
                return fmt.Errorf("trigger kind is required")
            }
            if len(policy.Rules) == 0 {
                return fmt.Errorf("at least one rule is required")
            }

            err := handler.CreateRetentionPolicy(ctx, &policy)
            if err != nil {
                return err
            }
            return nil
        },
    }

    // Add flags for policy fields
    cmd.Flags().StringVar(&policy.Algorithm, "algorithm", "or", "The algorithm used for tag retention")
    cmd.Flags().StringVar(&policy.Scope.Level, "scope-level", "", "The scope level of the retention policy")
    cmd.Flags().Int64Var(&policy.Scope.Ref, "scope-ref", 0, "The scope reference of the retention policy")

    // Flags for trigger
    cmd.Flags().StringVar(&policy.Trigger.Kind, "trigger-kind", "", "The trigger kind (e.g., 'Schedule')")
    var triggerSettings string
    cmd.Flags().StringVar(&triggerSettings, "trigger-settings", "", "The trigger settings in JSON format")

    // Flags for rule
    cmd.Flags().StringVar(&ruleAction, "rule-action", "", "The action for the retention rule")
    cmd.Flags().StringVar(&ruleTemplate, "rule-template", "", "The template for the retention rule")
    cmd.Flags().StringVar(&tagKind, "rule-tag-kind", "", "The kind of tag selector")
    cmd.Flags().StringVar(&tagPattern, "rule-tag-pattern", "", "The pattern for tag selector")
    cmd.Flags().StringVar(&scopeKind, "rule-scope-kind", "", "The kind of scope selector")
    cmd.Flags().StringVar(&scopePattern, "rule-scope-pattern", "", "The pattern for scope selector")

    cmd.PreRunE = func(cmd *cobra.Command, args []string) error {
        // Parse trigger settings
        if triggerSettings != "" {
            var settings interface{}
            err := json.Unmarshal([]byte(triggerSettings), &settings)
            if err != nil {
                return fmt.Errorf("failed to parse trigger settings: %v", err)
            }
            policy.Trigger.Settings = settings
        }

        // Create and add the rule
        rule := &models.RetentionRule{
            Action:   ruleAction,
            Template: ruleTemplate,
            TagSelectors: []*models.RetentionSelector{
                {
                    Kind:    tagKind,
                    Pattern: tagPattern,
                },
            },
            ScopeSelectors: map[string][]models.RetentionSelector{
                "repository": {
                    {
                        Kind:       scopeKind,
                        Decoration: "repoMatches",
                        Pattern:    scopePattern,
                    },
                },
            },
        }
        policy.Rules = append(policy.Rules, rule)
        return nil
    }

    return cmd
}

 
// ./harbor retention create --algorithm or --scope-level project --scope-ref 1 \
//   --trigger-kind Schedule --trigger-settings '{"cron": "0 0 * * *"}' \
//   --rule-action retain --rule-template "latestPushedK" \
//   --rule-tag-kind doublestar --rule-tag-pattern "latest" \
//   --rule-scope-kind repository --rule-scope-pattern "**"
