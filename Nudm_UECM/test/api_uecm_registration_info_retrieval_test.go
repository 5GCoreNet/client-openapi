/*
Nudm_UECM

Testing UECMRegistrationInfoRetrievalApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package Nudm_UECM

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_Nudm_UECM_UECMRegistrationInfoRetrievalApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test UECMRegistrationInfoRetrievalApiService GetRegistrations", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var ueId string

        resp, httpRes, err := apiClient.UECMRegistrationInfoRetrievalApi.GetRegistrations(context.Background(), ueId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}