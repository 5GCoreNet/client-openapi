/*
VAE_SessionOrientedService

Testing IndividualSessionOrientedServiceSubscriptionDocumentApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package VAE_SessionOrientedService

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_VAE_SessionOrientedService_IndividualSessionOrientedServiceSubscriptionDocumentApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test IndividualSessionOrientedServiceSubscriptionDocumentApiService DeleteSessionOrientedServiceSubscription", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var subscriptionId string

        resp, httpRes, err := apiClient.IndividualSessionOrientedServiceSubscriptionDocumentApi.DeleteSessionOrientedServiceSubscription(context.Background(), subscriptionId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test IndividualSessionOrientedServiceSubscriptionDocumentApiService ReadSessionOrientedServiceSubscription", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var subscriptionId string

        resp, httpRes, err := apiClient.IndividualSessionOrientedServiceSubscriptionDocumentApi.ReadSessionOrientedServiceSubscription(context.Background(), subscriptionId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
