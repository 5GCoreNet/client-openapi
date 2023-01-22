/*
Ntsctsf_ASTI Service API

Testing ASTIConfigurationsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package Ntsctsf_ASTI

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_Ntsctsf_ASTI_ASTIConfigurationsApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test ASTIConfigurationsApiService RequestStatusof5GAccessStratumTimeDistribution", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.ASTIConfigurationsApi.RequestStatusof5GAccessStratumTimeDistribution(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
