/*
3gpp-lpi-pp

Testing LPIParametersProvisioningsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package LpiParameterProvision

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_LpiParameterProvision_LPIParametersProvisioningsApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test LPIParametersProvisioningsApiService CreateNewResource", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var afId string

        resp, httpRes, err := apiClient.LPIParametersProvisioningsApi.CreateNewResource(context.Background(), afId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test LPIParametersProvisioningsApiService ReadAllResources", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var afId string

        resp, httpRes, err := apiClient.LPIParametersProvisioningsApi.ReadAllResources(context.Background(), afId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
