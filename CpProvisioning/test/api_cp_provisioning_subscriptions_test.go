/*
3gpp-cp-parameter-provisioning

Testing CPProvisioningSubscriptionsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package CpProvisioning

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_CpProvisioning_CPProvisioningSubscriptionsApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test CPProvisioningSubscriptionsApiService CreateCPProvisioningSubscription", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var scsAsId string

        resp, httpRes, err := apiClient.CPProvisioningSubscriptionsApi.CreateCPProvisioningSubscription(context.Background(), scsAsId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test CPProvisioningSubscriptionsApiService FetchAllCPProvisioningSubscriptions", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var scsAsId string

        resp, httpRes, err := apiClient.CPProvisioningSubscriptionsApi.FetchAllCPProvisioningSubscriptions(context.Background(), scsAsId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
