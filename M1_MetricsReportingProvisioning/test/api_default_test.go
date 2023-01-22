/*
M1_MetricsReportingProvisioning

Testing DefaultApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package M1_MetricsReportingProvisioning

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_M1_MetricsReportingProvisioning_DefaultApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test DefaultApiService ActivateMetricsReporting", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var provisioningSessionId string

        resp, httpRes, err := apiClient.DefaultApi.ActivateMetricsReporting(context.Background(), provisioningSessionId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService DestroyMetricsReportingConfiguration", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var provisioningSessionId string
        var metricsReportingConfigurationId string

        resp, httpRes, err := apiClient.DefaultApi.DestroyMetricsReportingConfiguration(context.Background(), provisioningSessionId, metricsReportingConfigurationId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService PatchMetricsReportingConfiguration", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var provisioningSessionId string
        var metricsReportingConfigurationId string

        resp, httpRes, err := apiClient.DefaultApi.PatchMetricsReportingConfiguration(context.Background(), provisioningSessionId, metricsReportingConfigurationId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService RetrieveMetricsReportingConfiguration", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var provisioningSessionId string
        var metricsReportingConfigurationId string

        resp, httpRes, err := apiClient.DefaultApi.RetrieveMetricsReportingConfiguration(context.Background(), provisioningSessionId, metricsReportingConfigurationId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test DefaultApiService UpdateMetricsReportingConfiguration", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var provisioningSessionId string
        var metricsReportingConfigurationId string

        resp, httpRes, err := apiClient.DefaultApi.UpdateMetricsReportingConfiguration(context.Background(), provisioningSessionId, metricsReportingConfigurationId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
