/*
Npcf_PolicyAuthorization Service API

Testing PCSCFRestorationIndicationApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package Npcf_PolicyAuthorization

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_Npcf_PolicyAuthorization_PCSCFRestorationIndicationApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test PCSCFRestorationIndicationApiService PcscfRestoration", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.PCSCFRestorationIndicationApi.PcscfRestoration(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}