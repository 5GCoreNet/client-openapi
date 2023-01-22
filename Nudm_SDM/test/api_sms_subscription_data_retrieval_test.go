/*
Nudm_SDM

Testing SMSSubscriptionDataRetrievalApiService

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

func Test_Nudm_SDM_SMSSubscriptionDataRetrievalApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test SMSSubscriptionDataRetrievalApiService GetSmsData", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var supi string

        resp, httpRes, err := apiClient.SMSSubscriptionDataRetrievalApi.GetSmsData(context.Background(), supi).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
