/*
SS_NetworkResourceAdaptation

Testing IndividualMulticastSubscriptionDocumentApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package SS_NetworkResourceAdaptation

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_SS_NetworkResourceAdaptation_IndividualMulticastSubscriptionDocumentApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test IndividualMulticastSubscriptionDocumentApiService DeleteMulticastSubscription", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var multiSubId string

        resp, httpRes, err := apiClient.IndividualMulticastSubscriptionDocumentApi.DeleteMulticastSubscription(context.Background(), multiSubId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test IndividualMulticastSubscriptionDocumentApiService GetMulticastSubscription", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var multiSubId string

        resp, httpRes, err := apiClient.IndividualMulticastSubscriptionDocumentApi.GetMulticastSubscription(context.Background(), multiSubId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}