/*
Unified Data Repository Service API file for Application Data

Testing IPTVConfigurationDataStoreApiService

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

func Test_Application_Data_IPTVConfigurationDataStoreApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test IPTVConfigurationDataStoreApiService ReadIPTVCongifurationData", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.IPTVConfigurationDataStoreApi.ReadIPTVCongifurationData(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
