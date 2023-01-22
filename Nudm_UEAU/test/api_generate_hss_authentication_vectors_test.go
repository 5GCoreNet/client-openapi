/*
Nudm_UEAU

Testing GenerateHSSAuthenticationVectorsApiService

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

func Test_Nudm_UEAU_GenerateHSSAuthenticationVectorsApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test GenerateHSSAuthenticationVectorsApiService GenerateAv", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var supi string
        var hssAuthType HssAuthTypeInUri

        resp, httpRes, err := apiClient.GenerateHSSAuthenticationVectorsApi.GenerateAv(context.Background(), supi, hssAuthType).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}