/*
Nudr_DataRepository API OpenAPI file

Testing IndividualExposureDataSubscriptionDocumentApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package Nudr_DR

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_Nudr_DR_IndividualExposureDataSubscriptionDocumentApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test IndividualExposureDataSubscriptionDocumentApiService DeleteIndividualExposureDataSubscription", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var subId string

        resp, httpRes, err := apiClient.IndividualExposureDataSubscriptionDocumentApi.DeleteIndividualExposureDataSubscription(context.Background(), subId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test IndividualExposureDataSubscriptionDocumentApiService ReplaceIndividualExposureDataSubscription", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var subId string

        resp, httpRes, err := apiClient.IndividualExposureDataSubscriptionDocumentApi.ReplaceIndividualExposureDataSubscription(context.Background(), subId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}