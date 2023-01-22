/*
3gpp-data-reporting-provisioning

Testing DataReportingConfigurationsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package DataReportingProvisioning

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_DataReportingProvisioning_DataReportingConfigurationsApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test DataReportingConfigurationsApiService CreateDataRepConfig", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var sessionId string

        resp, httpRes, err := apiClient.DataReportingConfigurationsApi.CreateDataRepConfig(context.Background(), sessionId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}