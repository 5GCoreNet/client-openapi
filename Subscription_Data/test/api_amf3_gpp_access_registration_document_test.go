/*
Unified Data Repository Service API file for subscription data

Testing AMF3GPPAccessRegistrationDocumentApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package Subscription_Data

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_Subscription_Data_AMF3GPPAccessRegistrationDocumentApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test AMF3GPPAccessRegistrationDocumentApiService AmfContext3gpp", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var ueId string

        resp, httpRes, err := apiClient.AMF3GPPAccessRegistrationDocumentApi.AmfContext3gpp(context.Background(), ueId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test AMF3GPPAccessRegistrationDocumentApiService CreateAmfContext3gpp", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var ueId string

        resp, httpRes, err := apiClient.AMF3GPPAccessRegistrationDocumentApi.CreateAmfContext3gpp(context.Background(), ueId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test AMF3GPPAccessRegistrationDocumentApiService QueryAmfContext3gpp", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var ueId string

        resp, httpRes, err := apiClient.AMF3GPPAccessRegistrationDocumentApi.QueryAmfContext3gpp(context.Background(), ueId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
