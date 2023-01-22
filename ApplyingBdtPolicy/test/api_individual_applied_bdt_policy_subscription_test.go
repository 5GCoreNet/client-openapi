/*
3gpp-applying-bdt-policy

Testing IndividualAppliedBDTPolicySubscriptionApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package ApplyingBdtPolicy

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_ApplyingBdtPolicy_IndividualAppliedBDTPolicySubscriptionApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test IndividualAppliedBDTPolicySubscriptionApiService DeleteAnSubscription", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var afId string
        var subscriptionId string

        resp, httpRes, err := apiClient.IndividualAppliedBDTPolicySubscriptionApi.DeleteAnSubscription(context.Background(), afId, subscriptionId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test IndividualAppliedBDTPolicySubscriptionApiService PartialUpdateAnSubscription", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var afId string
        var subscriptionId string

        resp, httpRes, err := apiClient.IndividualAppliedBDTPolicySubscriptionApi.PartialUpdateAnSubscription(context.Background(), afId, subscriptionId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test IndividualAppliedBDTPolicySubscriptionApiService ReadAnSubscription", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var afId string
        var subscriptionId string

        resp, httpRes, err := apiClient.IndividualAppliedBDTPolicySubscriptionApi.ReadAnSubscription(context.Background(), afId, subscriptionId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
