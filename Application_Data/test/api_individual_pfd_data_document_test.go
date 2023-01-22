/*
Unified Data Repository Service API file for Application Data

Testing IndividualPFDDataDocumentApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package Application_Data

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_Application_Data_IndividualPFDDataDocumentApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test IndividualPFDDataDocumentApiService CreateOrReplaceIndividualPFDData", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var appId string

        resp, httpRes, err := apiClient.IndividualPFDDataDocumentApi.CreateOrReplaceIndividualPFDData(context.Background(), appId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test IndividualPFDDataDocumentApiService DeleteIndividualPFDData", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var appId string

        resp, httpRes, err := apiClient.IndividualPFDDataDocumentApi.DeleteIndividualPFDData(context.Background(), appId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test IndividualPFDDataDocumentApiService ReadIndividualPFDData", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var appId string

        resp, httpRes, err := apiClient.IndividualPFDDataDocumentApi.ReadIndividualPFDData(context.Background(), appId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
