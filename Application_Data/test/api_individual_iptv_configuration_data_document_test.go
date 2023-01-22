/*
Unified Data Repository Service API file for Application Data

Testing IndividualIPTVConfigurationDataDocumentApiService

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

func Test_Application_Data_IndividualIPTVConfigurationDataDocumentApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test IndividualIPTVConfigurationDataDocumentApiService CreateOrReplaceIndividualIPTVConfigurationData", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var configurationId string

        resp, httpRes, err := apiClient.IndividualIPTVConfigurationDataDocumentApi.CreateOrReplaceIndividualIPTVConfigurationData(context.Background(), configurationId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test IndividualIPTVConfigurationDataDocumentApiService DeleteIndividualIPTVConfigurationData", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var configurationId string

        resp, httpRes, err := apiClient.IndividualIPTVConfigurationDataDocumentApi.DeleteIndividualIPTVConfigurationData(context.Background(), configurationId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test IndividualIPTVConfigurationDataDocumentApiService PartialReplaceIndividualIPTVConfigurationData", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var configurationId string

        resp, httpRes, err := apiClient.IndividualIPTVConfigurationDataDocumentApi.PartialReplaceIndividualIPTVConfigurationData(context.Background(), configurationId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
