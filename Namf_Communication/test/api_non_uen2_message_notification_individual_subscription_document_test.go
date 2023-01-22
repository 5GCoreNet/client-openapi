/*
Namf_Communication

Testing NonUEN2MessageNotificationIndividualSubscriptionDocumentApiService

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

func Test_Namf_Communication_NonUEN2MessageNotificationIndividualSubscriptionDocumentApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test NonUEN2MessageNotificationIndividualSubscriptionDocumentApiService NonUeN2InfoUnSubscribe", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var n2NotifySubscriptionId string

        resp, httpRes, err := apiClient.NonUEN2MessageNotificationIndividualSubscriptionDocumentApi.NonUeN2InfoUnSubscribe(context.Background(), n2NotifySubscriptionId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}