package api

import (
	"context"
	"fmt"
	"strings"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/client/retention"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/models"
	"github.com/goharbor/harbor-cli/pkg/utils"
)

type RetentionHandler struct{}

func NewRetentionHandler() *RetentionHandler {
	return &RetentionHandler{}
}

func (h *RetentionHandler) CreateRetentionPolicy(ctx context.Context, policy *models.RetentionPolicy) error {
    client, err := utils.GetClient()
    if err != nil {
        return fmt.Errorf("failed to get client: %v", err)
    }

    params := retention.NewCreateRetentionParams().WithPolicy(policy)
    _, err = client.Retention.CreateRetention(ctx, params)
    if err != nil {
        switch e := err.(type) {
        case *retention.CreateRetentionBadRequest:
            return fmt.Errorf("bad request: %s", formatErrors(e.Payload.Errors))
        case *retention.CreateRetentionUnauthorized:
            return fmt.Errorf("unauthorized: %s", formatErrors(e.Payload.Errors))
        case *retention.CreateRetentionForbidden:
            return fmt.Errorf("forbidden: %s", formatErrors(e.Payload.Errors))
        default:
            return fmt.Errorf("failed to create retention policy: %v", err)
        }
    }

    fmt.Println("Retention policy created successfully")
    return nil
}

// GetRetentionPolicy retrieves a retention policy by ID
func (h *RetentionHandler) GetRetentionPolicy(ctx context.Context, id int64) (*models.RetentionPolicy, error) {
    client, err := utils.GetClient()
    if err != nil {
        return nil, fmt.Errorf("failed to get client: %v", err)
    }

    params := retention.NewGetRetentionParams().WithID(id)
    resp, err := client.Retention.GetRetention(ctx, params)
    if err != nil {
        switch e := err.(type) {
        // case *retention.GetRetentionNotFound:
        //     return nil, fmt.Errorf("retention policy not found: %s", formatErrors(e.Payload.Errors))
        case *retention.GetRetentionUnauthorized:
            return nil, fmt.Errorf("unauthorized: %s", formatErrors(e.Payload.Errors))
        case *retention.GetRetentionForbidden:
            return nil, fmt.Errorf("forbidden: %s", formatErrors(e.Payload.Errors))
        case *retention.GetRetentionInternalServerError:
            return nil, fmt.Errorf("internal server error: %s", formatErrors(e.Payload.Errors))
        default:
            return nil, fmt.Errorf("failed to get retention policy: %v", err)
        }
    }

    return resp.Payload, nil
}

// Helper function to format error messages
func formatErrors(errors []*models.Error) string {
    var errMsgs []string
    for _, err := range errors {
        errMsgs = append(errMsgs, fmt.Sprintf("%s: %s", err.Code, err.Message))
    }
    return strings.Join(errMsgs, "; ")
}

// UpdateRetentionPolicy updates an existing retention policy
func (h *RetentionHandler) UpdateRetentionPolicy(ctx context.Context, id int64, policy *models.RetentionPolicy) error {
	client, err := utils.GetClient()
	if err != nil {
		return fmt.Errorf("failed to get client: %v", err)
	}

	params := retention.NewUpdateRetentionParams().WithID(id).WithPolicy(policy)
	_, err = client.Retention.UpdateRetention(ctx, params)
	if err != nil {
		return fmt.Errorf("failed to update retention policy: %v", err)
	}

	return nil
}

// DeleteRetentionPolicy deletes a retention policy by ID
func (h *RetentionHandler) DeleteRetentionPolicy(ctx context.Context, id int64) error {
	client, err := utils.GetClient()
	if err != nil {
		return fmt.Errorf("failed to get client: %v", err)
	}

	params := retention.NewDeleteRetentionParams().WithID(id)
	_, err = client.Retention.DeleteRetention(ctx, params)
	if err != nil {
		return fmt.Errorf("failed to delete retention policy: %v", err)
	}

	return nil
}

func (h *RetentionHandler) TriggerRetentionExecution(ctx context.Context, id int64, dryRun bool) (*models.RetentionExecution, error) {
    client, err := utils.GetClient()
    if err != nil {
        return nil, fmt.Errorf("failed to get client: %v", err)
    }

    body := retention.TriggerRetentionExecutionBody{
        DryRun: dryRun,
    }
    params := retention.NewTriggerRetentionExecutionParams().WithID(id).WithBody(body)
    
    respOK, respCreated, err := client.Retention.TriggerRetentionExecution(ctx, params)
    if err != nil {
        return nil, fmt.Errorf("failed to trigger retention execution: %v", err)
    }
	fmt.Printf("%+v\n", respOK)
	fmt.Printf("%+v\n", respCreated)
    
    return nil, fmt.Errorf("received unexpected response: respOK=%v, respCreated=%v", respOK, respCreated)
}


// ListRetentionExecutions lists retention executions for a policy
func (h *RetentionHandler) ListRetentionExecutions(ctx context.Context, id int64, page, pageSize int64) ([]*models.RetentionExecution, error) {
	client, err := utils.GetClient()
	if err != nil {
		return nil, fmt.Errorf("failed to get client: %v", err)
	}

	params := retention.NewListRetentionExecutionsParams().WithID(id).WithPage(&page).WithPageSize(&pageSize)
	resp, err := client.Retention.ListRetentionExecutions(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("failed to list retention executions: %v", err)
	}

	return resp.Payload, nil
}

// GetRetentionTaskLog retrieves the log for a retention task
func (h *RetentionHandler) GetRetentionTaskLog(ctx context.Context, id, executionID, taskID int64) (string, error) {
	client, err := utils.GetClient()
	if err != nil {
		return "", fmt.Errorf("failed to get client: %v", err)
	}

	params := retention.NewGetRetentionTaskLogParams().WithID(id).WithEid(executionID).WithTid(taskID)
	resp, err := client.Retention.GetRetentionTaskLog(ctx, params)
	if err != nil {
		return "", fmt.Errorf("failed to get retention task log: %v", err)
	}

	return string(resp.Payload), nil
}

func (h *RetentionHandler) GetRetentionMetadata(ctx context.Context) (*models.RetentionMetadata, error) {
    client, err := utils.GetClient()
    if err != nil {
        return nil, fmt.Errorf("failed to get client: %v", err)
    }

    // Use the generated params and reader types
    params := retention.NewGetRentenitionMetadataParams()
    resp, err := client.Retention.GetRentenitionMetadata(ctx, params)
    if err != nil {
        return nil, fmt.Errorf("failed to get retention metadata: %v", err)
    }

    // The response should already be of type *GetRentenitionMetadataOK
    metadata := resp.GetPayload()

    // Print Scope Selectors
    fmt.Println("Scope Selectors:")
    for i, selector := range metadata.ScopeSelectors {
        fmt.Printf("  Selector %d:\n", i+1)
        fmt.Printf("    Kind: %s\n", selector.Kind)
        // Print fields based on actual struct definition
        fmt.Printf("    Decorations: %v\n", selector.Decorations)
        // Add more fields as necessary
    }

    // Print Tag Selectors
    fmt.Println("\nTag Selectors:")
    for i, selector := range metadata.TagSelectors {
        fmt.Printf("  Selector %d:\n", i+1)
        fmt.Printf("    Kind: %s\n", selector.Kind)
        // Print fields based on actual struct definition
        fmt.Printf("    Decorations: %v\n", selector.Decorations)
        // Add more fields as necessary
    }

    // Print Templates
    fmt.Println("\nTemplates:")
    for i, tmpl := range metadata.Templates {
        fmt.Printf("  Template %d:\n", i+1)
        fmt.Printf("    Action: %s\n", tmpl.Action)
        fmt.Printf("    Display Text: %s\n", tmpl.DisplayText)
        fmt.Printf("    Rule Template: %s\n", tmpl.RuleTemplate)
        if len(tmpl.Params) > 0 {
            fmt.Printf("    Parameters:\n")
            for j, param := range tmpl.Params {
                fmt.Printf("      Param %d: %s\n", j+1, param)
            }
        } else {
            fmt.Println("    Parameters: None")
        }
    }

    return metadata, nil
}
