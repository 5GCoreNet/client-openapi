/*
Nudm_UEAU

Testing GenerateGBAAuthenticationVectorsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package Nudm_UEAU

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_Nudm_UEAU_GenerateGBAAuthenticationVectorsApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test GenerateGBAAuthenticationVectorsApiService GenerateGbaAv", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var supi string

        resp, httpRes, err := apiClient.GenerateGBAAuthenticationVectorsApi.GenerateGbaAv(context.Background(), supi).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}