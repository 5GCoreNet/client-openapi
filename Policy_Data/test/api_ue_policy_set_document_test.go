/*
Unified Data Repository Service API file for policy data

Testing UEPolicySetDocumentApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package Policy_Data

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_Policy_Data_UEPolicySetDocumentApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test UEPolicySetDocumentApiService CreateOrReplaceUEPolicySet", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var ueId string

        resp, httpRes, err := apiClient.UEPolicySetDocumentApi.CreateOrReplaceUEPolicySet(context.Background(), ueId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test UEPolicySetDocumentApiService ReadUEPolicySet", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var ueId string

        resp, httpRes, err := apiClient.UEPolicySetDocumentApi.ReadUEPolicySet(context.Background(), ueId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test UEPolicySetDocumentApiService UpdateUEPolicySet", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var ueId string

        resp, httpRes, err := apiClient.UEPolicySetDocumentApi.UpdateUEPolicySet(context.Background(), ueId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
