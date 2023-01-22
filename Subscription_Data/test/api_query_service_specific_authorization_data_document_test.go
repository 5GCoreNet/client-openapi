/*
Unified Data Repository Service API file for subscription data

Testing QueryServiceSpecificAuthorizationDataDocumentApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package Subscription_Data

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_Subscription_Data_QueryServiceSpecificAuthorizationDataDocumentApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test QueryServiceSpecificAuthorizationDataDocumentApiService GetSSAuData", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var ueId string
        var serviceType ServiceType

        resp, httpRes, err := apiClient.QueryServiceSpecificAuthorizationDataDocumentApi.GetSSAuData(context.Background(), ueId, serviceType).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
