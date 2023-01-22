/*
nmbsf-mbs-ud-ingest

Testing MBSUserDataIngestSessionStatusSubscriptionsCollectionApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package Nmbsf_MBSUserDataIngestSession

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_Nmbsf_MBSUserDataIngestSession_MBSUserDataIngestSessionStatusSubscriptionsCollectionApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test MBSUserDataIngestSessionStatusSubscriptionsCollectionApiService CreateMBSUserDataIngStatSubsc", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.MBSUserDataIngestSessionStatusSubscriptionsCollectionApi.CreateMBSUserDataIngStatSubsc(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test MBSUserDataIngestSessionStatusSubscriptionsCollectionApiService RetrieveMBSUserDataIngStatSubscs", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.MBSUserDataIngestSessionStatusSubscriptionsCollectionApi.RetrieveMBSUserDataIngStatSubscs(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
