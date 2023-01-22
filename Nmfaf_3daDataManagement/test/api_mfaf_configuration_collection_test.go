/*
Nmfaf_3daDataManagement

Testing MFAFConfigurationCollectionApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package Nmfaf_3daDataManagement

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_Nmfaf_3daDataManagement_MFAFConfigurationCollectionApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test MFAFConfigurationCollectionApiService CreateMFAFConfiguration", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.MFAFConfigurationCollectionApi.CreateMFAFConfiguration(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
