/*
SS_LocationAreaInfoRetrieval

Testing LocationInformationCollectionApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package SS_LocationAreaInfoRetrieval

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_SS_LocationAreaInfoRetrieval_LocationInformationCollectionApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test LocationInformationCollectionApiService RetrieveUeLocInfo", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.LocationInformationCollectionApi.RetrieveUeLocInfo(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}