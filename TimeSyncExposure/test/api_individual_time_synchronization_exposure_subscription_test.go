/*
3gpp-time-sync-exposure

Testing IndividualTimeSynchronizationExposureSubscriptionApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package TimeSyncExposure

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_TimeSyncExposure_IndividualTimeSynchronizationExposureSubscriptionApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test IndividualTimeSynchronizationExposureSubscriptionApiService DeleteAnSubscription", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var afId string
        var subscriptionId string

        resp, httpRes, err := apiClient.IndividualTimeSynchronizationExposureSubscriptionApi.DeleteAnSubscription(context.Background(), afId, subscriptionId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test IndividualTimeSynchronizationExposureSubscriptionApiService FullyUpdateAnSubscription", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var afId string
        var subscriptionId string

        resp, httpRes, err := apiClient.IndividualTimeSynchronizationExposureSubscriptionApi.FullyUpdateAnSubscription(context.Background(), afId, subscriptionId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test IndividualTimeSynchronizationExposureSubscriptionApiService ReadAnSubscription", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var afId string
        var subscriptionId string

        resp, httpRes, err := apiClient.IndividualTimeSynchronizationExposureSubscriptionApi.ReadAnSubscription(context.Background(), afId, subscriptionId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test IndividualTimeSynchronizationExposureSubscriptionApiService ReadTimeSynSubscription", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var afId string
        var subscriptionId string
        var instanceReference string

        resp, httpRes, err := apiClient.IndividualTimeSynchronizationExposureSubscriptionApi.ReadTimeSynSubscription(context.Background(), afId, subscriptionId, instanceReference).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}