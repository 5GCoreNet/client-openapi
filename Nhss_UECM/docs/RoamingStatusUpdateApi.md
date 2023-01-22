# \RoamingStatusUpdateApi

All URIs are relative to *https://example.com/nhss-uecm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RoamingStatusUpdate**](RoamingStatusUpdateApi.md#RoamingStatusUpdate) | **Post** /roaming-status-update | Roaming Status Update



## RoamingStatusUpdate

> RoamingStatusUpdate(ctx).RoamingStatusUpdateInfo(roamingStatusUpdateInfo).Execute()

Roaming Status Update

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
    roamingStatusUpdateInfo := *openapiclient.NewRoamingStatusUpdateInfo("Imsi_example", *openapiclient.NewPlmnId("Mcc_example", "Mnc_example")) // RoamingStatusUpdateInfo | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoamingStatusUpdateApi.RoamingStatusUpdate(context.Background()).RoamingStatusUpdateInfo(roamingStatusUpdateInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoamingStatusUpdateApi.RoamingStatusUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRoamingStatusUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **roamingStatusUpdateInfo** | [**RoamingStatusUpdateInfo**](RoamingStatusUpdateInfo.md) |  | 

### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

