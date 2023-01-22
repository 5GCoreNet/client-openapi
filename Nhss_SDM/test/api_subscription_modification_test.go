/*
Nhss_SDM

Testing SubscriptionModificationApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package Nhss_SDM

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_Nhss_SDM_SubscriptionModificationApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test SubscriptionModificationApiService Modify", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var ueId string
        var subscriptionId string

        resp, httpRes, err := apiClient.SubscriptionModificationApi.Modify(context.Background(), ueId, subscriptionId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
