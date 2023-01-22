/*
3gpp-analyticsexposure

Testing AnalyticsExposureAPIFetchAnalyticsInformationApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package AnalyticsExposure

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_AnalyticsExposure_AnalyticsExposureAPIFetchAnalyticsInformationApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test AnalyticsExposureAPIFetchAnalyticsInformationApiService FetchAnalyticsInfo", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var afId string

        resp, httpRes, err := apiClient.AnalyticsExposureAPIFetchAnalyticsInformationApi.FetchAnalyticsInfo(context.Background(), afId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}