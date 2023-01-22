/*
Nudm_UEAU

Testing DeleteAuthApiService

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

func Test_Nudm_UEAU_DeleteAuthApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test DeleteAuthApiService DeleteAuth", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var supi string
        var authEventId string

        resp, httpRes, err := apiClient.DeleteAuthApi.DeleteAuth(context.Background(), supi, authEventId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
