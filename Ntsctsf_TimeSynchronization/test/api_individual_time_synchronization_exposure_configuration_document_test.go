/*
Ntsctsf_TimeSynchronization Service API

Testing IndividualTimeSynchronizationExposureConfigurationDocumentApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package Ntsctsf_TimeSynchronization

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_Ntsctsf_TimeSynchronization_IndividualTimeSynchronizationExposureConfigurationDocumentApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test IndividualTimeSynchronizationExposureConfigurationDocumentApiService CreateIndividualTimeSynchronizationExposureConfiguration", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var subscriptionId string

        resp, httpRes, err := apiClient.IndividualTimeSynchronizationExposureConfigurationDocumentApi.CreateIndividualTimeSynchronizationExposureConfiguration(context.Background(), subscriptionId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test IndividualTimeSynchronizationExposureConfigurationDocumentApiService DeleteIndividualTimeSynchronizationExposureConfiguration", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var subscriptionId string
        var configurationId string

        resp, httpRes, err := apiClient.IndividualTimeSynchronizationExposureConfigurationDocumentApi.DeleteIndividualTimeSynchronizationExposureConfiguration(context.Background(), subscriptionId, configurationId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test IndividualTimeSynchronizationExposureConfigurationDocumentApiService GetIndividualTimeSynchronizationExposureConfiguration", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var subscriptionId string
        var configurationId string

        resp, httpRes, err := apiClient.IndividualTimeSynchronizationExposureConfigurationDocumentApi.GetIndividualTimeSynchronizationExposureConfiguration(context.Background(), subscriptionId, configurationId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test IndividualTimeSynchronizationExposureConfigurationDocumentApiService ReplaceIndividualTimeSynchronizationExposureConfiguration", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var subscriptionId string
        var configurationId string

        resp, httpRes, err := apiClient.IndividualTimeSynchronizationExposureConfigurationDocumentApi.ReplaceIndividualTimeSynchronizationExposureConfiguration(context.Background(), subscriptionId, configurationId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
