/*
3gpp-device-triggering

Testing DeviceTriggeringTransactionsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package DeviceTriggering

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_DeviceTriggering_DeviceTriggeringTransactionsApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test DeviceTriggeringTransactionsApiService FetchAllDeviceTriggeringTransactions", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var scsAsId string

        resp, httpRes, err := apiClient.DeviceTriggeringTransactionsApi.FetchAllDeviceTriggeringTransactions(context.Background(), scsAsId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}