/*
Nhss_imsUEAU

Testing GenerateSIPAuthDataCustomOperationApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package Nhss_imsUEAU

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_Nhss_imsUEAU_GenerateSIPAuthDataCustomOperationApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test GenerateSIPAuthDataCustomOperationApiService GenerateSipAuthData", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var impi string

        resp, httpRes, err := apiClient.GenerateSIPAuthDataCustomOperationApi.GenerateSipAuthData(context.Background(), impi).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
