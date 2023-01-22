/*
3gpp-bdt

Testing IndividualBDTSubscriptionApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package ResourceManagementOfBdt

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_ResourceManagementOfBdt_IndividualBDTSubscriptionApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test IndividualBDTSubscriptionApiService DeleteBDTSubscription", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var scsAsId string
        var subscriptionId string

        resp, httpRes, err := apiClient.IndividualBDTSubscriptionApi.DeleteBDTSubscription(context.Background(), scsAsId, subscriptionId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test IndividualBDTSubscriptionApiService FetchIndBDTSubscription", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var scsAsId string
        var subscriptionId string

        resp, httpRes, err := apiClient.IndividualBDTSubscriptionApi.FetchIndBDTSubscription(context.Background(), scsAsId, subscriptionId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test IndividualBDTSubscriptionApiService ModifyBDTSubscription", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var scsAsId string
        var subscriptionId string

        resp, httpRes, err := apiClient.IndividualBDTSubscriptionApi.ModifyBDTSubscription(context.Background(), scsAsId, subscriptionId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test IndividualBDTSubscriptionApiService UpdateBDTSubscription", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var scsAsId string
        var subscriptionId string

        resp, httpRes, err := apiClient.IndividualBDTSubscriptionApi.UpdateBDTSubscription(context.Background(), scsAsId, subscriptionId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
