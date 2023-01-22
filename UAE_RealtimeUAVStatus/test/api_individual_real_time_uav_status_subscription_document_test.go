/*
UAE Server Real-time UAV Status Service

Testing IndividualRealTimeUAVStatusSubscriptionDocumentApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package UAE_RealtimeUAVStatus

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_UAE_RealtimeUAVStatus_IndividualRealTimeUAVStatusSubscriptionDocumentApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test IndividualRealTimeUAVStatusSubscriptionDocumentApiService DeleteRTUavStatusSubscription", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var subscriptionId string

        resp, httpRes, err := apiClient.IndividualRealTimeUAVStatusSubscriptionDocumentApi.DeleteRTUavStatusSubscription(context.Background(), subscriptionId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test IndividualRealTimeUAVStatusSubscriptionDocumentApiService GetRTUavStatusSubscription", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var subscriptionId string

        resp, httpRes, err := apiClient.IndividualRealTimeUAVStatusSubscriptionDocumentApi.GetRTUavStatusSubscription(context.Background(), subscriptionId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test IndividualRealTimeUAVStatusSubscriptionDocumentApiService UpdateRTUavStatusSubscription", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var subscriptionId string

        resp, httpRes, err := apiClient.IndividualRealTimeUAVStatusSubscriptionDocumentApi.UpdateRTUavStatusSubscription(context.Background(), subscriptionId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}