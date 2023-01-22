/*
Nudr_DataRepository API OpenAPI file

Testing SMSSubscriptionDataDocumentApiService

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

func Test_Nudr_DR_SMSSubscriptionDataDocumentApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test SMSSubscriptionDataDocumentApiService QuerySmsData", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var ueId string
        var servingPlmnId string

        resp, httpRes, err := apiClient.SMSSubscriptionDataDocumentApi.QuerySmsData(context.Background(), ueId, servingPlmnId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
