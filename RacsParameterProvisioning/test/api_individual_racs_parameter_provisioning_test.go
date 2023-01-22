/*
3gpp-racs-parameter-provisioning

Testing IndividualRACSParameterProvisioningApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package RacsParameterProvisioning

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_RacsParameterProvisioning_IndividualRACSParameterProvisioningApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test IndividualRACSParameterProvisioningApiService DeleteIndRACSParameterProvisioning", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var scsAsId string
        var provisioningId string

        resp, httpRes, err := apiClient.IndividualRACSParameterProvisioningApi.DeleteIndRACSParameterProvisioning(context.Background(), scsAsId, provisioningId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test IndividualRACSParameterProvisioningApiService FetchIndRACSParameterProvisioning", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var scsAsId string
        var provisioningId string

        resp, httpRes, err := apiClient.IndividualRACSParameterProvisioningApi.FetchIndRACSParameterProvisioning(context.Background(), scsAsId, provisioningId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test IndividualRACSParameterProvisioningApiService ModifyIndRACSParameterProvisioning", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var scsAsId string
        var provisioningId string

        resp, httpRes, err := apiClient.IndividualRACSParameterProvisioningApi.ModifyIndRACSParameterProvisioning(context.Background(), scsAsId, provisioningId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test IndividualRACSParameterProvisioningApiService UpdateIndRACSParameterProvisioning", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var scsAsId string
        var provisioningId string

        resp, httpRes, err := apiClient.IndividualRACSParameterProvisioningApi.UpdateIndRACSParameterProvisioning(context.Background(), scsAsId, provisioningId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
