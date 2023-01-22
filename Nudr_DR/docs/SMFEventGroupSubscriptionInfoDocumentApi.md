# \SMFEventGroupSubscriptionInfoDocumentApi

All URIs are relative to *https://example.com/nudr-dr/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSmfGroupSubscriptions**](SMFEventGroupSubscriptionInfoDocumentApi.md#CreateSmfGroupSubscriptions) | **Put** /subscription-data/group-data/{ueGroupId}/ee-subscriptions/{subsId}/smf-subscriptions | Create SMF Subscription Info for a group of UEs or any YE



## CreateSmfGroupSubscriptions

> SmfSubscriptionInfo CreateSmfGroupSubscriptions(ctx, ueGroupId, subsId).SmfSubscriptionInfo(smfSubscriptionInfo).Execute()

Create SMF Subscription Info for a group of UEs or any YE

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
    ueGroupId := "ueGroupId_example" // string | 
    subsId := "subsId_example" // string | 
    smfSubscriptionInfo := *openapiclient.NewSmfSubscriptionInfo([]openapiclient.SmfSubscriptionItem{*openapiclient.NewSmfSubscriptionItem("SmfInstanceId_example", "SubscriptionId_example")}) // SmfSubscriptionInfo | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SMFEventGroupSubscriptionInfoDocumentApi.CreateSmfGroupSubscriptions(context.Background(), ueGroupId, subsId).SmfSubscriptionInfo(smfSubscriptionInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SMFEventGroupSubscriptionInfoDocumentApi.CreateSmfGroupSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSmfGroupSubscriptions`: SmfSubscriptionInfo
    fmt.Fprintf(os.Stdout, "Response from `SMFEventGroupSubscriptionInfoDocumentApi.CreateSmfGroupSubscriptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueGroupId** | **string** |  | 
**subsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSmfGroupSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **smfSubscriptionInfo** | [**SmfSubscriptionInfo**](SmfSubscriptionInfo.md) |  | 

### Return type

[**SmfSubscriptionInfo**](SmfSubscriptionInfo.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

