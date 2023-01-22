/*
Ndccf_DataManagement

Testing DCCFDataSubscriptionsCollectionApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package Ndccf_DataManagement

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_Ndccf_DataManagement_DCCFDataSubscriptionsCollectionApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test DCCFDataSubscriptionsCollectionApiService CreateDCCFDataSubscription", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.DCCFDataSubscriptionsCollectionApi.CreateDCCFDataSubscription(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
