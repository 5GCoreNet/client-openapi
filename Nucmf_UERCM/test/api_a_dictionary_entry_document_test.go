/*
Nucmf_UECapabilityManagement

Testing ADictionaryEntryDocumentApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package Nucmf_UERCM

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_Nucmf_UERCM_ADictionaryEntryDocumentApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test ADictionaryEntryDocumentApiService CreateDictionaryEntry", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.ADictionaryEntryDocumentApi.CreateDictionaryEntry(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
