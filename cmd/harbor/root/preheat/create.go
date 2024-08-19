package preheat

import (
	"strconv"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/client/preheat"
	"fmt"
	"github.com/spf13/cobra"

	"github.com/goharbor/go-client/pkg/sdk/v2.0/models"
	"github.com/goharbor/harbor-cli/pkg/api"
)

func newCreateCommand() *cobra.Command {
	var name, description, projectName, providerID, filters, trigger string

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new preheat policy",
		Long:  `Create a new preheat policy in Harbor`,
		RunE: func(cmd *cobra.Command, args []string) error {
			providerIDInt, err := strconv.ParseInt(providerID, 10, 64)
			if err != nil {
				return fmt.Errorf("invalid provider ID: %v", err)
			}
			params := &preheat.CreatePolicyParams{
				Policy: &models.PreheatPolicy{
					Name:        name,
					Description: description,
					ProviderID:  providerIDInt,
					Filters:     filters,
					Trigger:     trigger,
				},
				ProjectName: projectName,
			}

			response, err := api.CreatePolicy(params)
			if err != nil {
				if apiErr, ok := err.(*preheat.CreatePolicyInternalServerError); ok {
					if apiErr.Payload != nil && len(apiErr.Payload.Errors) > 0 {
						return fmt.Errorf("failed to create preheat policy: code %d, message: %s", apiErr.Payload.Errors[0].Code, apiErr.Payload.Errors[0].Message)
					}
					return fmt.Errorf("failed to create preheat policy: %v", apiErr)
				}
				return fmt.Errorf("failed to create preheat policy: %v", err)
			}

			fmt.Printf("Preheat policy created successfully. ID: %d\n", response)
			return nil
		},
	}

	cmd.Flags().StringVar(&name, "name", "", "Name of the preheat policy")
	cmd.Flags().StringVar(&description, "description", "", "Description of the preheat policy")
	cmd.Flags().StringVar(&projectName, "project", "", "Name of the project")
	cmd.Flags().StringVar(&providerID, "provider", "", "ID of the provider")
	cmd.Flags().StringVar(&filters, "filters", "", "Filters for the preheat policy")
	cmd.Flags().StringVar(&trigger, "trigger", "", "Trigger for the preheat policy")

	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("project")
	cmd.MarkFlagRequired("provider")

	return cmd
}