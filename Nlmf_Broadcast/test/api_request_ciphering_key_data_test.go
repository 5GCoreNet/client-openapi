/*
LMF Broadcast

Testing RequestCipheringKeyDataApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package Nlmf_Broadcast

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_Nlmf_Broadcast_RequestCipheringKeyDataApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test RequestCipheringKeyDataApiService CipheringKeyData", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.RequestCipheringKeyDataApi.CipheringKeyData(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
