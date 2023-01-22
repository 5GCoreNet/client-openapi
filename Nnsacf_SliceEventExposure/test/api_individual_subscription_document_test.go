/*
Nnsacf_SliceEventExposure

Testing IndividualSubscriptionDocumentApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package Nnsacf_SliceEventExposure

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_Nnsacf_SliceEventExposure_IndividualSubscriptionDocumentApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test IndividualSubscriptionDocumentApiService CompleteModifySubscription", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var subscriptionId string

        resp, httpRes, err := apiClient.IndividualSubscriptionDocumentApi.CompleteModifySubscription(context.Background(), subscriptionId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test IndividualSubscriptionDocumentApiService DeleteSubscription", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var subscriptionId string

        resp, httpRes, err := apiClient.IndividualSubscriptionDocumentApi.DeleteSubscription(context.Background(), subscriptionId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test IndividualSubscriptionDocumentApiService PartialModifySubscription", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var subscriptionId string

        resp, httpRes, err := apiClient.IndividualSubscriptionDocumentApi.PartialModifySubscription(context.Background(), subscriptionId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}