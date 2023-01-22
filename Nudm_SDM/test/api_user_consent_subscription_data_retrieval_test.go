/*
Nudm_SDM

Testing UserConsentSubscriptionDataRetrievalApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package Nudm_SDM

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_Nudm_SDM_UserConsentSubscriptionDataRetrievalApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test UserConsentSubscriptionDataRetrievalApiService GetUcData", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var supi string

        resp, httpRes, err := apiClient.UserConsentSubscriptionDataRetrievalApi.GetUcData(context.Background(), supi).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
