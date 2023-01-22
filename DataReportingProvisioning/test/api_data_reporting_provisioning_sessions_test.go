/*
3gpp-data-reporting-provisioning

Testing DataReportingProvisioningSessionsApiService

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

func Test_DataReportingProvisioning_DataReportingProvisioningSessionsApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test DataReportingProvisioningSessionsApiService CreateDataRepProvSession", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.DataReportingProvisioningSessionsApi.CreateDataRepProvSession(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
