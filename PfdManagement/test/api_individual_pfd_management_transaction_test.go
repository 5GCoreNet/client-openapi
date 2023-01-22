/*
3gpp-pfd-management

Testing IndividualPFDManagementTransactionApiService

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

func Test_PfdManagement_IndividualPFDManagementTransactionApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test IndividualPFDManagementTransactionApiService DeleteIndPFDManagementTransaction", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var scsAsId string
        var transactionId string

        resp, httpRes, err := apiClient.IndividualPFDManagementTransactionApi.DeleteIndPFDManagementTransaction(context.Background(), scsAsId, transactionId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test IndividualPFDManagementTransactionApiService FetchIndPFDManagementTransaction", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var scsAsId string
        var transactionId string

        resp, httpRes, err := apiClient.IndividualPFDManagementTransactionApi.FetchIndPFDManagementTransaction(context.Background(), scsAsId, transactionId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test IndividualPFDManagementTransactionApiService ModifyIndPFDManagementTransaction", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var scsAsId string
        var transactionId string

        resp, httpRes, err := apiClient.IndividualPFDManagementTransactionApi.ModifyIndPFDManagementTransaction(context.Background(), scsAsId, transactionId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test IndividualPFDManagementTransactionApiService UpdateIndPFDManagementTransaction", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var scsAsId string
        var transactionId string

        resp, httpRes, err := apiClient.IndividualPFDManagementTransactionApi.UpdateIndPFDManagementTransaction(context.Background(), scsAsId, transactionId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}