/*
Nnef_SMContext

Testing SMContextsCollectionCollectionApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package Nnef_SMContext

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_Nnef_SMContext_SMContextsCollectionCollectionApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test SMContextsCollectionCollectionApiService Create", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.SMContextsCollectionCollectionApi.Create(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
