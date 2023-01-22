/*
GBA BSF Nbsp_GBA Service

Testing BootstrappingInfoRetrievalCustomOperationApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package Nbsp_GBA

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_Nbsp_GBA_BootstrappingInfoRetrievalCustomOperationApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test BootstrappingInfoRetrievalCustomOperationApiService BootstrappingInfoRetrieval", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.BootstrappingInfoRetrievalCustomOperationApi.BootstrappingInfoRetrieval(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
