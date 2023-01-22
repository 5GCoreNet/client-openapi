/*
Nudr_DataRepository API OpenAPI file

Testing AMFGroupSubscriptionInfoDocumentApiService

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

func Test_Nudr_DR_AMFGroupSubscriptionInfoDocumentApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test AMFGroupSubscriptionInfoDocumentApiService CreateAmfGroupSubscriptions", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var ueGroupId string
        var subsId string

        resp, httpRes, err := apiClient.AMFGroupSubscriptionInfoDocumentApi.CreateAmfGroupSubscriptions(context.Background(), ueGroupId, subsId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
