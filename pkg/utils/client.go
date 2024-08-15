package utils

import (
	"context"
	"fmt"
	"os"
	"sync"
"net/http"

	"github.com/go-openapi/runtime"
	"github.com/goharbor/go-client/pkg/harbor"
	v2client "github.com/goharbor/go-client/pkg/sdk/v2.0/client"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	clientInstance *v2client.HarborAPI
	clientOnce     sync.Once
	clientErr      error
)

func GetClient() (*v2client.HarborAPI, error) {
	clientOnce.Do(func() {
		credentialName := viper.GetString("current-credential-name")
		clientInstance = GetClientByCredentialName(credentialName)
		if clientErr != nil {
			log.Errorf("failed to initialize client: %v", clientErr)
		}
	})
	return clientInstance, clientErr
}

func ContextWithClient() (context.Context, *v2client.HarborAPI, error) {
	client, err := GetClient()
	if err != nil {
		return nil, nil, err
	}
	ctx := context.Background()
	return ctx, client, nil
}

func GetClientByConfig(clientConfig *harbor.ClientSetConfig) *v2client.HarborAPI {
	cs, err := harbor.NewClientSet(clientConfig)
	if err != nil {
		panic(err)
	}
	return cs.V2()
}

// Returns Harbor v2 client after resolving the credential name
func GetClientByCredentialName(credentialName string) *v2client.HarborAPI {
	credential, err := GetCredentials(credentialName)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	clientConfig := &harbor.ClientSetConfig{
		URL:      credential.ServerAddress,
		Username: credential.Username,
		Password: credential.Password,
	}
	return GetClientByConfig(clientConfig)
}

// IsNotFoundError checks if the error represents a 404 Not Found response
func IsNotFoundError(err error) bool {
	if err == nil {
		return false
	}

	if apiErr, ok := err.(*runtime.APIError); ok {
		return apiErr.Code == http.StatusNotFound
	}

	return false
}