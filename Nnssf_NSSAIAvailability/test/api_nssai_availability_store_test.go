/*
NSSF NSSAI Availability

Testing NSSAIAvailabilityStoreApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package Nnssf_NSSAIAvailability

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_Nnssf_NSSAIAvailability_NSSAIAvailabilityStoreApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test NSSAIAvailabilityStoreApiService NSSAIAvailabilityOptions", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.NSSAIAvailabilityStoreApi.NSSAIAvailabilityOptions(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}