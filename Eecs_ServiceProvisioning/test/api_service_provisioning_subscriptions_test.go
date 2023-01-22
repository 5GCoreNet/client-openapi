/*
Eecs_ServiceProvisioning

Testing ServiceProvisioningSubscriptionsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package Eecs_ServiceProvisioning

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_Eecs_ServiceProvisioning_ServiceProvisioningSubscriptionsApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test ServiceProvisioningSubscriptionsApiService SubscriptionsPost", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.ServiceProvisioningSubscriptionsApi.SubscriptionsPost(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}