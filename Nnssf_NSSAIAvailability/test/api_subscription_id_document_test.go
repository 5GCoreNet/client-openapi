/*
NSSF NSSAI Availability

Testing SubscriptionIDDocumentApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package Nnssf_NSSAIAvailability

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_Nnssf_NSSAIAvailability_SubscriptionIDDocumentApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test SubscriptionIDDocumentApiService NSSAIAvailabilitySubModifyPatch", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var subscriptionId string

        resp, httpRes, err := apiClient.SubscriptionIDDocumentApi.NSSAIAvailabilitySubModifyPatch(context.Background(), subscriptionId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test SubscriptionIDDocumentApiService NSSAIAvailabilityUnsubscribe", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var subscriptionId string

        resp, httpRes, err := apiClient.SubscriptionIDDocumentApi.NSSAIAvailabilityUnsubscribe(context.Background(), subscriptionId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
