/*
Nudr_DataRepository API OpenAPI file

Testing SubsToNotifyCollectionApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package Nudr_DR

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_Nudr_DR_SubsToNotifyCollectionApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test SubsToNotifyCollectionApiService QuerySubsToNotify", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.SubsToNotifyCollectionApi.QuerySubsToNotify(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test SubsToNotifyCollectionApiService RemoveMultipleSubscriptionDataSubscriptions", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.SubsToNotifyCollectionApi.RemoveMultipleSubscriptionDataSubscriptions(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test SubsToNotifyCollectionApiService SubscriptionDataSubscriptions", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.SubsToNotifyCollectionApi.SubscriptionDataSubscriptions(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}