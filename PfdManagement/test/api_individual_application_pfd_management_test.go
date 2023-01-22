/*
3gpp-pfd-management

Testing IndividualApplicationPFDManagementApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package PfdManagement

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_PfdManagement_IndividualApplicationPFDManagementApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test IndividualApplicationPFDManagementApiService DeleteIndApplicationPFDManagement", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var scsAsId string
        var transactionId string
        var appId string

        resp, httpRes, err := apiClient.IndividualApplicationPFDManagementApi.DeleteIndApplicationPFDManagement(context.Background(), scsAsId, transactionId, appId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test IndividualApplicationPFDManagementApiService FetchIndApplicationPFDManagement", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var scsAsId string
        var transactionId string
        var appId string

        resp, httpRes, err := apiClient.IndividualApplicationPFDManagementApi.FetchIndApplicationPFDManagement(context.Background(), scsAsId, transactionId, appId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test IndividualApplicationPFDManagementApiService ModifyIndApplicationPFDManagement", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var scsAsId string
        var transactionId string
        var appId string

        resp, httpRes, err := apiClient.IndividualApplicationPFDManagementApi.ModifyIndApplicationPFDManagement(context.Background(), scsAsId, transactionId, appId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test IndividualApplicationPFDManagementApiService UpdateIndApplicationPFDManagement", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var scsAsId string
        var transactionId string
        var appId string

        resp, httpRes, err := apiClient.IndividualApplicationPFDManagementApi.UpdateIndApplicationPFDManagement(context.Background(), scsAsId, transactionId, appId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}