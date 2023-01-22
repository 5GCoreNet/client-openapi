/*
Nnssaaf_NSSAA

Testing SliceAuthenticationContextCreationApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package Nnssaaf_NSSAA

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_Nnssaaf_NSSAA_SliceAuthenticationContextCreationApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test SliceAuthenticationContextCreationApiService CreateSliceAuthenticationContext", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.SliceAuthenticationContextCreationApi.CreateSliceAuthenticationContext(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
