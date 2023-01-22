/*
M1_EventDataProcessingProvisioning

Testing DefaultApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package M1_EventDataProcessingProvisioning

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_M1_EventDataProcessingProvisioning_DefaultApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test DefaultApiService CreateEventDataProcessingConfiguration", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var provisioningSessionId string

        resp, httpRes, err := apiClient.DefaultApi.CreateEventDataProcessingConfiguration(context.Background(), provisioningSessionId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService DestroyEventDataProcessingConfiguration", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var provisioningSessionId string
        var eventDataProcessingConfigurationId string

        resp, httpRes, err := apiClient.DefaultApi.DestroyEventDataProcessingConfiguration(context.Background(), provisioningSessionId, eventDataProcessingConfigurationId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService PatchEventDataProcessingConfiguration", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var provisioningSessionId string
        var eventDataProcessingConfigurationId string

        resp, httpRes, err := apiClient.DefaultApi.PatchEventDataProcessingConfiguration(context.Background(), provisioningSessionId, eventDataProcessingConfigurationId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService RetrieveEventDataProcessingConfiguration", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var provisioningSessionId string
        var eventDataProcessingConfigurationId string

        resp, httpRes, err := apiClient.DefaultApi.RetrieveEventDataProcessingConfiguration(context.Background(), provisioningSessionId, eventDataProcessingConfigurationId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService UpdateEventDataProcessingConfiguration", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var provisioningSessionId string
        var eventDataProcessingConfigurationId string

        resp, httpRes, err := apiClient.DefaultApi.UpdateEventDataProcessingConfiguration(context.Background(), provisioningSessionId, eventDataProcessingConfigurationId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}