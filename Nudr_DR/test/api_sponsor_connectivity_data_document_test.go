/*
Nudr_DataRepository API OpenAPI file

Testing SponsorConnectivityDataDocumentApiService

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

func Test_Nudr_DR_SponsorConnectivityDataDocumentApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test SponsorConnectivityDataDocumentApiService ReadSponsorConnectivityData", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var sponsorId string

        resp, httpRes, err := apiClient.SponsorConnectivityDataDocumentApi.ReadSponsorConnectivityData(context.Background(), sponsorId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}