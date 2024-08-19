package root

import (
	"github.com/goharbor/go-client/pkg/sdk/v2.0/client/preheat"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/models"
	"github.com/goharbor/harbor-cli/pkg/utils"
	log "github.com/sirupsen/logrus"
)

// CreatePreheatPolicy creates a preheat policy under a project
func CreatePreheatPolicy(projectName string, policy *models.PreheatPolicy) (*preheat.CreatePolicyCreated, error) {
	// Create a context with the client
	ctx, client, err := utils.ContextWithClient()
	if err != nil {
		return nil, err
	}

	// Create the preheat policy
	response, err := client.Preheat.CreatePolicy(ctx, &preheat.CreatePolicyParams{
		ProjectName: projectName,
		Policy:      policy,
	})

	if err != nil {
		return nil, err
	}

	log.Infof("Preheat policy created for project %s", projectName)
	return response, nil
}

// CreateInstance creates a p2p provider instance
func CreateInstance(params *preheat.CreateInstanceParams) (*preheat.CreateInstanceCreated, error) {
	// Create a context with the client
	ctx, client, err := utils.ContextWithClient()
	if err != nil {
		return nil, err
	}

	// Create the p2p provider instance
	response, err := client.Preheat.CreateInstance(ctx, params)
	if err != nil {
		return nil, err
	}

	log.Infof("P2P provider instance created with ID %s", response)
	return response, nil
}

// DeleteInstance deletes the specified p2p provider instance
func DeleteInstance(params *preheat.DeleteInstanceParams) (*preheat.DeleteInstanceOK, error) {
	// Create a context with the client
	ctx, client, err := utils.ContextWithClient()
	if err != nil {
		return nil, err
	}

	// Delete the p2p provider instance
	response, err := client.Preheat.DeleteInstance(ctx, params)
	if err != nil {
		return nil, err
	}

	// Log the deletion status or any relevant information from the response
	log.Infof("P2P provider instance deleted with response: %+v", response)

	return response, nil
}


// DeletePolicy deletes a preheat policy
func DeletePolicy(params *preheat.DeletePolicyParams) (*preheat.DeletePolicyOK, error) {
	// Create a context with the client
	ctx, client, err := utils.ContextWithClient()
	if err != nil {
		return nil, err
	}

	// Delete the preheat policy
	response, err := client.Preheat.DeletePolicy(ctx, params)
	if err != nil {
		return nil, err
	}

	// Log the deletion status or any relevant information from the response
	log.Infof("Preheat policy deleted with response: %+v", response)

	return response, nil
}


// GetExecution gets an execution detail by id
func GetExecution(params *preheat.GetExecutionParams) (*preheat.GetExecutionOK, error) {
	// Create a context with the client
	ctx, client, err := utils.ContextWithClient()
	if err != nil {
		return nil, err
	}

	// Get the execution detail
	response, err := client.Preheat.GetExecution(ctx, params)
	if err != nil {
		return nil, err
	}

	// Log the execution details or any relevant information from the response
	log.Infof("Execution details retrieved with response: %+v", response)

	return response, nil
}



// GetInstance gets a p2p provider instance
func GetInstance(params *preheat.GetInstanceParams) (*preheat.GetInstanceOK, error) {
	// Create a context with the client
	ctx, client, err := utils.ContextWithClient()
	if err != nil {
		return nil, err
	}

	// Get the p2p provider instance
	response, err := client.Preheat.GetInstance(ctx, params)
	if err != nil {
		return nil, err
	}

	// Log the instance details or any relevant information from the response
	log.Infof("P2P provider instance details retrieved with response: %+v", response)

	return response, nil
}

// GetPolicy gets a preheat policy
func GetPolicy(params *preheat.GetPolicyParams) (*preheat.GetPolicyOK, error) {
	// Create a context with the client
	ctx, client, err := utils.ContextWithClient()
	if err != nil {
		return nil, err
	}

	// Get the preheat policy
	response, err := client.Preheat.GetPolicy(ctx, params)
	if err != nil {
		return nil, err
	}

	// Log the policy details or any relevant information from the response
	log.Infof("Preheat policy details retrieved with response: %+v", response)

	return response, nil
}

// GetPreheatLog gets the log text stream of the specified task for the given execution
func GetPreheatLog(params *preheat.GetPreheatLogParams) (*preheat.GetPreheatLogOK, error) {
	// Create a context with the client
	ctx, client, err := utils.ContextWithClient()
	if err != nil {
		return nil, err
	}

	// Get the preheat log
	response, err := client.Preheat.GetPreheatLog(ctx, params)
	if err != nil {
		return nil, err
	}

	// Log the preheat log details or any relevant information from the response
	log.Infof("Preheat log retrieved with response: %s", string(response.Payload))

	return response, nil
}

// ListExecutions lists executions for the given policy
func ListExecutions(params *preheat.ListExecutionsParams) (*preheat.ListExecutionsOK, error) {
	// Create a context with the client
	ctx, client, err := utils.ContextWithClient()
	if err != nil {
		return nil, err
	}

	// List the executions
	response, err := client.Preheat.ListExecutions(ctx, params)
	if err != nil {
		return nil, err
	}

	// Log the list of executions or any relevant information from the response
	log.Infof("List of executions retrieved with response: %+v", response)

	return response, nil
}

// ListInstances lists P2P provider instances
func ListInstances(params *preheat.ListInstancesParams) (*preheat.ListInstancesOK, error) {
	// Create a context with the client
	ctx, client, err := utils.ContextWithClient()
	if err != nil {
		return nil, err
	}

	// List the P2P provider instances
	response, err := client.Preheat.ListInstances(ctx, params)
	if err != nil {
		return nil, err
	}

	// Log the list of instances or any relevant information from the response
	log.Infof("List of P2P provider instances retrieved with response: %+v", response)

	return response, nil
}

// ListPolicies lists preheat policies
func ListPolicies(params *preheat.ListPoliciesParams) (*preheat.ListPoliciesOK, error) {
	// Create a context with the client
	ctx, client, err := utils.ContextWithClient()
	if err != nil {
		return nil, err
	}

	// List the preheat policies
	response, err := client.Preheat.ListPolicies(ctx, params)
	if err != nil {
		return nil, err
	}

	// Log the list of policies or any relevant information from the response
	log.Infof("List of preheat policies retrieved with response: %+v", response)

	return response, nil
}


// ListProviders lists P2P providers
func ListProviders(params *preheat.ListProvidersParams) (*preheat.ListProvidersOK, error) {
	// Create a context with the client
	ctx, client, err := utils.ContextWithClient()
	if err != nil {
		return nil, err
	}

	// List the P2P providers
	response, err := client.Preheat.ListProviders(ctx, params)
	if err != nil {
		return nil, err
	}

	// Log the list of providers or any relevant information from the response
	log.Infof("List of P2P providers retrieved with response: %+v", response)

	return response, nil
}


// ListProvidersUnderProject gets all providers at project level
func ListProvidersUnderProject(params *preheat.ListProvidersUnderProjectParams) (*preheat.ListProvidersUnderProjectOK, error) {
	// Create a context with the client
	ctx, client, err := utils.ContextWithClient()
	if err != nil {
		return nil, err
	}

	// List providers under the project
	response, err := client.Preheat.ListProvidersUnderProject(ctx, params)
	if err != nil {
		return nil, err
	}

	// Log the list of providers under the project or any relevant information from the response
	log.Infof("List of providers under project retrieved with response: %+v", response)

	return response, nil
}

// ListTasks lists all the related tasks for the given execution
func ListTasks(params *preheat.ListTasksParams) (*preheat.ListTasksOK, error) {
	// Create a context with the client
	ctx, client, err := utils.ContextWithClient()
	if err != nil {
		return nil, err
	}

	// List the tasks related to the execution
	response, err := client.Preheat.ListTasks(ctx, params)
	if err != nil {
		return nil, err
	}

	// Log the list of tasks or any relevant information from the response
	log.Infof("List of tasks for execution retrieved with response: %+v", response)

	return response, nil
}



// ManualPreheat performs a manual preheat operation
func ManualPreheat(params *preheat.ManualPreheatParams) (*preheat.ManualPreheatCreated, error) {
	// Create a context with the client
	ctx, client, err := utils.ContextWithClient()
	if err != nil {
		return nil, err
	}

	// Perform the manual preheat
	response, err := client.Preheat.ManualPreheat(ctx, params)
	if err != nil {
		return nil, err
	}

	// Log the response or any relevant information
	log.Infof("Manual preheat performed with response: %+v", response)

	return response, nil
}



// PingInstances checks the status of a P2P provider instance
func PingInstances(params *preheat.PingInstancesParams) (*preheat.PingInstancesOK, error) {
	// Create a context with the client
	ctx, client, err := utils.ContextWithClient()
	if err != nil {
		return nil, err
	}

	// Ping the instance
	response, err := client.Preheat.PingInstances(ctx, params)
	if err != nil {
		return nil, err
	}

	// Log the response or any relevant information
	log.Infof("Pinging instances with response: %+v", response)

	return response, nil
}


// StopExecution stops a preheat execution
func StopExecution(params *preheat.StopExecutionParams) (*preheat.StopExecutionOK, error) {
	// Create a context with the client
	ctx, client, err := utils.ContextWithClient()
	if err != nil {
		return nil, err
	}

	// Stop the execution
	response, err := client.Preheat.StopExecution(ctx, params)
	if err != nil {
		return nil, err
	}

	// Log the response or any relevant information
	log.Infof("Stopping execution with response: %+v", response)

	return response, nil
}
// UpdateInstance updates the specified P2P provider instance
func UpdateInstance(params *preheat.UpdateInstanceParams) (*preheat.UpdateInstanceOK, error) {
	// Create a context with the client
	ctx, client, err := utils.ContextWithClient()
	if err != nil {
		return nil, err
	}

	// Update the instance
	response, err := client.Preheat.UpdateInstance(ctx, params)
	if err != nil {
		return nil, err
	}

	// Log the response or any relevant information
	log.Infof("Updating instance with response: %+v", response)

	return response, nil
}



// UpdatePolicy updates a preheat policy
func UpdatePolicy(params *preheat.UpdatePolicyParams) (*preheat.UpdatePolicyOK, error) {
	// Create a context with the client
	ctx, client, err := utils.ContextWithClient()
	if err != nil {
		return nil, err
	}

	// Update the policy
	response, err := client.Preheat.UpdatePolicy(ctx, params)
	if err != nil {
		return nil, err
	}

	// Log the response or any relevant information
	log.Infof("Updating policy with response: %+v", response)

	return response, nil
}

