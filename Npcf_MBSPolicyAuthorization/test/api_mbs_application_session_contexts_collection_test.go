/*
Npcf_MBSPolicyAuthorization API

Testing MBSApplicationSessionContextsCollectionApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package Npcf_MBSPolicyAuthorization

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_Npcf_MBSPolicyAuthorization_MBSApplicationSessionContextsCollectionApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test MBSApplicationSessionContextsCollectionApiService CreateMBSAppSessionCtxt", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.MBSApplicationSessionContextsCollectionApi.CreateMBSAppSessionCtxt(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}