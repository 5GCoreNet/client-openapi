/*
Nudm_SDM

Testing SubscriptionDeletionApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package Nudm_SDM

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_Nudm_SDM_SubscriptionDeletionApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test SubscriptionDeletionApiService Unsubscribe", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var ueId string
        var subscriptionId string

        resp, httpRes, err := apiClient.SubscriptionDeletionApi.Unsubscribe(context.Background(), ueId, subscriptionId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
