/*
Namf_Communication

Testing N1N2IndividualSubscriptionDocumentApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package Namf_Communication

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_Namf_Communication_N1N2IndividualSubscriptionDocumentApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test N1N2IndividualSubscriptionDocumentApiService N1N2MessageUnSubscribe", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var ueContextId string
        var subscriptionId string

        resp, httpRes, err := apiClient.N1N2IndividualSubscriptionDocumentApi.N1N2MessageUnSubscribe(context.Background(), ueContextId, subscriptionId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
