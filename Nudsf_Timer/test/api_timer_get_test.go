/*
Nudsf_Timer

Testing TimerGetApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package Nudsf_Timer

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_Nudsf_Timer_TimerGetApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test TimerGetApiService GetTimer", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var realmId string
        var storageId string
        var timerId string

        resp, httpRes, err := apiClient.TimerGetApi.GetTimer(context.Background(), realmId, storageId, timerId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
