package api

import (
	"fmt"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/client/preheat"
 	"github.com/goharbor/harbor-cli/pkg/utils"
 
)

func CreatePolicy(params *preheat.CreatePolicyParams) (*preheat.CreatePolicyCreated, error) {
	ctx, client, err := utils.ContextWithClient()
	if err != nil {
		return nil, err
	}

	response, err := client.Preheat.CreatePolicy(ctx, params)
	if err != nil {
		// Print more details about the error
		fmt.Printf("Error details: %+v\n", err)
		return nil, err
	}

	return response, nil
}