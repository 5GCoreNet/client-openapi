/*
Unified Data Repository Service API file for subscription data

Testing SMFRegistrationDocumentApiService

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

func Test_Subscription_Data_SMFRegistrationDocumentApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test SMFRegistrationDocumentApiService CreateOrUpdateSmfRegistration", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var ueId string
        var pduSessionId int32

        resp, httpRes, err := apiClient.SMFRegistrationDocumentApi.CreateOrUpdateSmfRegistration(context.Background(), ueId, pduSessionId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test SMFRegistrationDocumentApiService DeleteSmfRegistration", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var ueId string
        var pduSessionId int32

        resp, httpRes, err := apiClient.SMFRegistrationDocumentApi.DeleteSmfRegistration(context.Background(), ueId, pduSessionId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test SMFRegistrationDocumentApiService QuerySmfRegistration", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var ueId string
        var pduSessionId int32

        resp, httpRes, err := apiClient.SMFRegistrationDocumentApi.QuerySmfRegistration(context.Background(), ueId, pduSessionId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test SMFRegistrationDocumentApiService UpdateSmfContext", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var ueId string
        var pduSessionId int32

        resp, httpRes, err := apiClient.SMFRegistrationDocumentApi.UpdateSmfContext(context.Background(), ueId, pduSessionId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
