# \AnalyticsExposureAPIFetchAnalyticsInformationApi

All URIs are relative to *https://example.com/3gpp-analyticsexposure/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchAnalyticsInfo**](AnalyticsExposureAPIFetchAnalyticsInformationApi.md#FetchAnalyticsInfo) | **Post** /{afId}/fetch | Fetch analytics information



## FetchAnalyticsInfo

> AnalyticsData FetchAnalyticsInfo(ctx, afId).AnalyticsRequest(analyticsRequest).Execute()

Fetch analytics information

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/AnalyticsExposure"
)

func main() {
    afId := "afId_example" // string | Identifier of the AF
    analyticsRequest := *openapiclient.NewAnalyticsRequest(*openapiclient.NewAnalyticsEvent(), "SuppFeat_example") // AnalyticsRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AnalyticsExposureAPIFetchAnalyticsInformationApi.FetchAnalyticsInfo(context.Background(), afId).AnalyticsRequest(analyticsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsExposureAPIFetchAnalyticsInformationApi.FetchAnalyticsInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchAnalyticsInfo`: AnalyticsData
    fmt.Fprintf(os.Stdout, "Response from `AnalyticsExposureAPIFetchAnalyticsInformationApi.FetchAnalyticsInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**afId** | **string** | Identifier of the AF | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchAnalyticsInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **analyticsRequest** | [**AnalyticsRequest**](AnalyticsRequest.md) |  | 

### Return type

[**AnalyticsData**](AnalyticsData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

