/*
Nhss_imsSDM

Testing CSLocationRetrievalApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package Nhss_imsSDM

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_Nhss_imsSDM_CSLocationRetrievalApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test CSLocationRetrievalApiService GetLocCsDomain", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var imsUeId string

        resp, httpRes, err := apiClient.CSLocationRetrievalApi.GetLocCsDomain(context.Background(), imsUeId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
