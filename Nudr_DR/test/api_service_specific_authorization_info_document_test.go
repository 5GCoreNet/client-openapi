/*
Nudr_DataRepository API OpenAPI file

Testing ServiceSpecificAuthorizationInfoDocumentApiService

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

func Test_Nudr_DR_ServiceSpecificAuthorizationInfoDocumentApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test ServiceSpecificAuthorizationInfoDocumentApiService CreateServiceSpecificAuthorizationInfo", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var ueId string
        var serviceType ServiceType

        resp, httpRes, err := apiClient.ServiceSpecificAuthorizationInfoDocumentApi.CreateServiceSpecificAuthorizationInfo(context.Background(), ueId, serviceType).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ServiceSpecificAuthorizationInfoDocumentApiService GetServiceSpecificAuthorizationInfo", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var ueId string
        var serviceType ServiceType

        resp, httpRes, err := apiClient.ServiceSpecificAuthorizationInfoDocumentApi.GetServiceSpecificAuthorizationInfo(context.Background(), ueId, serviceType).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ServiceSpecificAuthorizationInfoDocumentApiService ModifyServiceSpecificAuthorizationInfo", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var ueId string
        var serviceType ServiceType

        resp, httpRes, err := apiClient.ServiceSpecificAuthorizationInfoDocumentApi.ModifyServiceSpecificAuthorizationInfo(context.Background(), ueId, serviceType).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ServiceSpecificAuthorizationInfoDocumentApiService RemoveServiceSpecificAuthorizationInfo", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var ueId string
        var serviceType ServiceType

        resp, httpRes, err := apiClient.ServiceSpecificAuthorizationInfoDocumentApi.RemoveServiceSpecificAuthorizationInfo(context.Background(), ueId, serviceType).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
