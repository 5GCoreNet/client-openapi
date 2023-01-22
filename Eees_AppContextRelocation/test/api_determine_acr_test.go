/*
Eees Application Context Relocation Service

Testing DetermineACRApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package Eees_AppContextRelocation

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_Eees_AppContextRelocation_DetermineACRApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test DetermineACRApiService Determine", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.DetermineACRApi.Determine(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
