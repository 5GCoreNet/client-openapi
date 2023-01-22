# \HSSEventGroupSubscriptionInfoDocumentApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateHssGroupSubscriptions**](HSSEventGroupSubscriptionInfoDocumentApi.md#CreateHssGroupSubscriptions) | **Put** /subscription-data/group-data/{ueGroupId}/ee-subscriptions/{subsId}/hss-subscriptions | Create HSS Subscription Info for a group of UEs



## CreateHssGroupSubscriptions

> HssSubscriptionInfo CreateHssGroupSubscriptions(ctx, externalGroupId, subsId).HssSubscriptionInfo(hssSubscriptionInfo).Execute()

Create HSS Subscription Info for a group of UEs

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
    externalGroupId := "externalGroupId_example" // string | 
    subsId := "subsId_example" // string | 
    hssSubscriptionInfo := *openapiclient.NewHssSubscriptionInfo([]openapiclient.HssSubscriptionItem{*openapiclient.NewHssSubscriptionItem("HssInstanceId_example", "SubscriptionId_example")}) // HssSubscriptionInfo | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HSSEventGroupSubscriptionInfoDocumentApi.CreateHssGroupSubscriptions(context.Background(), externalGroupId, subsId).HssSubscriptionInfo(hssSubscriptionInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HSSEventGroupSubscriptionInfoDocumentApi.CreateHssGroupSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateHssGroupSubscriptions`: HssSubscriptionInfo
    fmt.Fprintf(os.Stdout, "Response from `HSSEventGroupSubscriptionInfoDocumentApi.CreateHssGroupSubscriptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalGroupId** | **string** |  | 
**subsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateHssGroupSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **hssSubscriptionInfo** | [**HssSubscriptionInfo**](HssSubscriptionInfo.md) |  | 

### Return type

[**HssSubscriptionInfo**](HssSubscriptionInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

