/*
3gpp-chargeable-party

Testing ChargeablePartyTransactionOperationApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package ChargeableParty

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_ChargeableParty_ChargeablePartyTransactionOperationApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test ChargeablePartyTransactionOperationApiService CreateChargeablePartyTransaction", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var scsAsId string

        resp, httpRes, err := apiClient.ChargeablePartyTransactionOperationApi.CreateChargeablePartyTransaction(context.Background(), scsAsId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ChargeablePartyTransactionOperationApiService FetchAllChargeablePartyTransactions", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var scsAsId string

        resp, httpRes, err := apiClient.ChargeablePartyTransactionOperationApi.FetchAllChargeablePartyTransactions(context.Background(), scsAsId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}