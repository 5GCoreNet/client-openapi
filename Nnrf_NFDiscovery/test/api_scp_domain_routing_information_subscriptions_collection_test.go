/*
NRF NFDiscovery Service

Testing SCPDomainRoutingInformationSubscriptionsCollectionApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package Nnrf_NFDiscovery

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_Nnrf_NFDiscovery_SCPDomainRoutingInformationSubscriptionsCollectionApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test SCPDomainRoutingInformationSubscriptionsCollectionApiService ScpDomainRoutingInfoSubscribe", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.SCPDomainRoutingInformationSubscriptionsCollectionApi.ScpDomainRoutingInfoSubscribe(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
