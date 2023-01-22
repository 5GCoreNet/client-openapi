/*
NRF NFManagement Service

Testing NFInstancesStoreApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package Nnrf_NFManagement

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_Nnrf_NFManagement_NFInstancesStoreApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test NFInstancesStoreApiService GetNFInstances", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.NFInstancesStoreApi.GetNFInstances(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test NFInstancesStoreApiService OptionsNFInstances", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.NFInstancesStoreApi.OptionsNFInstances(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
