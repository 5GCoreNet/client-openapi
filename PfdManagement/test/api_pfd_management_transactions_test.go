/*
3gpp-pfd-management

Testing PFDManagementTransactionsApiService

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

func Test_PfdManagement_PFDManagementTransactionsApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test PFDManagementTransactionsApiService CreatePFDManagementTransaction", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var scsAsId string

        resp, httpRes, err := apiClient.PFDManagementTransactionsApi.CreatePFDManagementTransaction(context.Background(), scsAsId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test PFDManagementTransactionsApiService FetchAllPFDManagementTransactions", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var scsAsId string

        resp, httpRes, err := apiClient.PFDManagementTransactionsApi.FetchAllPFDManagementTransactions(context.Background(), scsAsId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
