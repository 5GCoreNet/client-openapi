/*
Nudr_DataRepository API OpenAPI file

Testing AuthenticationSubscriptionDocumentApiService

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

func Test_Nudr_DR_AuthenticationSubscriptionDocumentApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test AuthenticationSubscriptionDocumentApiService ModifyAuthenticationSubscription", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var ueId string

        resp, httpRes, err := apiClient.AuthenticationSubscriptionDocumentApi.ModifyAuthenticationSubscription(context.Background(), ueId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
