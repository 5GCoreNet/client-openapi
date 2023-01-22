/*
Eees_ACREvents

Testing ACREventsSubscriptionCollectionApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package Eees_ACREvents

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_Eees_ACREvents_ACREventsSubscriptionCollectionApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test ACREventsSubscriptionCollectionApiService CreateACREventsSubscripton", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.ACREventsSubscriptionCollectionApi.CreateACREventsSubscripton(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
