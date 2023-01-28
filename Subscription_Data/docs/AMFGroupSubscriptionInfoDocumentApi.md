# \AMFGroupSubscriptionInfoDocumentApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAmfGroupSubscriptions**](AMFGroupSubscriptionInfoDocumentApi.md#CreateAmfGroupSubscriptions) | **Put** /subscription-data/group-data/{ueGroupId}/ee-subscriptions/{subsId}/amf-subscriptions | Create AmfSubscriptions for a group of UEs or any UE



## CreateAmfGroupSubscriptions

> []AmfSubscriptionInfo CreateAmfGroupSubscriptions(ctx, ueGroupId, subsId).AmfSubscriptionInfo(amfSubscriptionInfo).Execute()

Create AmfSubscriptions for a group of UEs or any UE

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Subscription_Data"
)

func main() {
    ueGroupId := "ueGroupId_example" // string | 
    subsId := "subsId_example" // string | 
    amfSubscriptionInfo := []openapiclient.AmfSubscriptionInfo{*openapiclient.NewAmfSubscriptionInfo("AmfInstanceId_example", "SubscriptionId_example")} // []AmfSubscriptionInfo | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AMFGroupSubscriptionInfoDocumentApi.CreateAmfGroupSubscriptions(context.Background(), ueGroupId, subsId).AmfSubscriptionInfo(amfSubscriptionInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AMFGroupSubscriptionInfoDocumentApi.CreateAmfGroupSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAmfGroupSubscriptions`: []AmfSubscriptionInfo
    fmt.Fprintf(os.Stdout, "Response from `AMFGroupSubscriptionInfoDocumentApi.CreateAmfGroupSubscriptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueGroupId** | **string** |  | 
**subsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAmfGroupSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **amfSubscriptionInfo** | [**[]AmfSubscriptionInfo**](AmfSubscriptionInfo.md) |  | 

### Return type

[**[]AmfSubscriptionInfo**](AmfSubscriptionInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

