/*
Nudm_PP

Testing SubscriptionDataUpdateApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package Nudm_PP

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_Nudm_PP_SubscriptionDataUpdateApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test SubscriptionDataUpdateApiService Update", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var ueId UpdateUeIdParameter

        resp, httpRes, err := apiClient.SubscriptionDataUpdateApi.Update(context.Background(), ueId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
