/*
3gpp-nidd

Testing IndividualManagePortConfigurationApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package NIDD

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_NIDD_IndividualManagePortConfigurationApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test IndividualManagePortConfigurationApiService DeleteIndManagePortConfiguration", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var scsAsId string
        var configurationId string
        var portId string

        resp, httpRes, err := apiClient.IndividualManagePortConfigurationApi.DeleteIndManagePortConfiguration(context.Background(), scsAsId, configurationId, portId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test IndividualManagePortConfigurationApiService FetchIndManagePortConfiguration", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var scsAsId string
        var configurationId string
        var portId string

        resp, httpRes, err := apiClient.IndividualManagePortConfigurationApi.FetchIndManagePortConfiguration(context.Background(), scsAsId, configurationId, portId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test IndividualManagePortConfigurationApiService UpdateIndManagePortConfiguration", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var scsAsId string
        var configurationId string
        var portId string

        resp, httpRes, err := apiClient.IndividualManagePortConfigurationApi.UpdateIndManagePortConfiguration(context.Background(), scsAsId, configurationId, portId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
