/*
VAE_HDMapDynamicInfo

Testing IndividualHdMapDynamicInfoSubscriptionDocumentApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package VAE_HDMapDynamicInfo

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_VAE_HDMapDynamicInfo_IndividualHdMapDynamicInfoSubscriptionDocumentApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test IndividualHdMapDynamicInfoSubscriptionDocumentApiService ReadHdMapDynamicInfoSubscription", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var subscriptionId string

        resp, httpRes, err := apiClient.IndividualHdMapDynamicInfoSubscriptionDocumentApi.ReadHdMapDynamicInfoSubscription(context.Background(), subscriptionId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}