# \AFLevelUELocationUpdateNotificationOperationApi

All URIs are relative to *https://example.com/3gpp-mo-lcs-notify/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UELocationNotify**](AFLevelUELocationUpdateNotificationOperationApi.md#UELocationNotify) | **Post** / | UE location information update notification



## UELocationNotify

> LocUpdateDataReply UELocationNotify(ctx).LocUpdateData(locUpdateData).Execute()

UE location information update notification

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    locUpdateData := *openapiclient.NewLocUpdateData("Gpsi_example", *openapiclient.NewLocationInfo(), *openapiclient.NewLcsQosClass(), "SuppFeat_example") // LocUpdateData | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AFLevelUELocationUpdateNotificationOperationApi.UELocationNotify(context.Background()).LocUpdateData(locUpdateData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AFLevelUELocationUpdateNotificationOperationApi.UELocationNotify``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UELocationNotify`: LocUpdateDataReply
    fmt.Fprintf(os.Stdout, "Response from `AFLevelUELocationUpdateNotificationOperationApi.UELocationNotify`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUELocationNotifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **locUpdateData** | [**LocUpdateData**](LocUpdateData.md) |  | 

### Return type

[**LocUpdateDataReply**](LocUpdateDataReply.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

