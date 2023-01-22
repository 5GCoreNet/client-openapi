/*
3gpp-mbs-session

Testing IndividualMBSParametersProvisioningApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package MBSSession

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_MBSSession_IndividualMBSParametersProvisioningApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test IndividualMBSParametersProvisioningApiService DeleteIndMBSParamsProvisioning", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var mbsPpId string

        resp, httpRes, err := apiClient.IndividualMBSParametersProvisioningApi.DeleteIndMBSParamsProvisioning(context.Background(), mbsPpId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test IndividualMBSParametersProvisioningApiService GetIndMBSParamsProvisioning", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var mbsPpId string

        resp, httpRes, err := apiClient.IndividualMBSParametersProvisioningApi.GetIndMBSParamsProvisioning(context.Background(), mbsPpId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test IndividualMBSParametersProvisioningApiService ModifyIndMBSParamsProvisioning", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var mbsPpId string

        resp, httpRes, err := apiClient.IndividualMBSParametersProvisioningApi.ModifyIndMBSParamsProvisioning(context.Background(), mbsPpId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test IndividualMBSParametersProvisioningApiService UpdateIndMBSParamsProvisioning", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var mbsPpId string

        resp, httpRes, err := apiClient.IndividualMBSParametersProvisioningApi.UpdateIndMBSParamsProvisioning(context.Background(), mbsPpId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
