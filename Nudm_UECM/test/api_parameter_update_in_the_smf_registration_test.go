/*
Nudm_UECM

Testing ParameterUpdateInTheSMFRegistrationApiService

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

func Test_Nudm_UECM_ParameterUpdateInTheSMFRegistrationApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test ParameterUpdateInTheSMFRegistrationApiService UpdateSmfRegistration", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var ueId string
        var pduSessionId int32

        resp, httpRes, err := apiClient.ParameterUpdateInTheSMFRegistrationApi.UpdateSmfRegistration(context.Background(), ueId, pduSessionId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
