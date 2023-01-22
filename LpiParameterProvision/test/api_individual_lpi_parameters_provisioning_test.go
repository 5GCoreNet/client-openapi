/*
3gpp-lpi-pp

Testing IndividualLPIParametersProvisioningApiService

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

func Test_LpiParameterProvision_IndividualLPIParametersProvisioningApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test IndividualLPIParametersProvisioningApiService DeleteAnResource", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var afId string
        var provisionedLpiId string

        resp, httpRes, err := apiClient.IndividualLPIParametersProvisioningApi.DeleteAnResource(context.Background(), afId, provisionedLpiId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test IndividualLPIParametersProvisioningApiService FullyUpdateAnResource", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var afId string
        var provisionedLpiId string

        resp, httpRes, err := apiClient.IndividualLPIParametersProvisioningApi.FullyUpdateAnResource(context.Background(), afId, provisionedLpiId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test IndividualLPIParametersProvisioningApiService PartialUpdateAnResource", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var afId string
        var provisionedLpiId string

        resp, httpRes, err := apiClient.IndividualLPIParametersProvisioningApi.PartialUpdateAnResource(context.Background(), afId, provisionedLpiId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test IndividualLPIParametersProvisioningApiService ReadAnResource", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var afId string
        var provisionedLpiId string

        resp, httpRes, err := apiClient.IndividualLPIParametersProvisioningApi.ReadAnResource(context.Background(), afId, provisionedLpiId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}