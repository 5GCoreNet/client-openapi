# \QueryAMFSubscriptionInfoDocumentApi

All URIs are relative to *https://example.com/nudr-dr/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAmfGroupSubscriptions**](QueryAMFSubscriptionInfoDocumentApi.md#GetAmfGroupSubscriptions) | **Get** /subscription-data/group-data/{ueGroupId}/ee-subscriptions/{subsId}/amf-subscriptions | Retrieve AMF subscription Info for a group of UEs or any UE
[**GetAmfSubscriptionInfo**](QueryAMFSubscriptionInfoDocumentApi.md#GetAmfSubscriptionInfo) | **Get** /subscription-data/{ueId}/context-data/ee-subscriptions/{subsId}/amf-subscriptions | Retrieve AMF subscription Info



## GetAmfGroupSubscriptions

> []AmfSubscriptionInfo GetAmfGroupSubscriptions(ctx, ueGroupId, subsId).Execute()

Retrieve AMF subscription Info for a group of UEs or any UE

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QueryAMFSubscriptionInfoDocumentApi.GetAmfGroupSubscriptions(context.Background(), ueGroupId, subsId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueryAMFSubscriptionInfoDocumentApi.GetAmfGroupSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAmfGroupSubscriptions`: []AmfSubscriptionInfo
    fmt.Fprintf(os.Stdout, "Response from `QueryAMFSubscriptionInfoDocumentApi.GetAmfGroupSubscriptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueGroupId** | **string** |  | 
**subsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAmfGroupSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]AmfSubscriptionInfo**](AmfSubscriptionInfo.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAmfSubscriptionInfo

> []AmfSubscriptionInfo GetAmfSubscriptionInfo(ctx, ueId, subsId).Execute()

Retrieve AMF subscription Info

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
    ueId := "ueId_example" // string | 
    subsId := "subsId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QueryAMFSubscriptionInfoDocumentApi.GetAmfSubscriptionInfo(context.Background(), ueId, subsId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueryAMFSubscriptionInfoDocumentApi.GetAmfSubscriptionInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAmfSubscriptionInfo`: []AmfSubscriptionInfo
    fmt.Fprintf(os.Stdout, "Response from `QueryAMFSubscriptionInfoDocumentApi.GetAmfSubscriptionInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** |  | 
**subsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAmfSubscriptionInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]AmfSubscriptionInfo**](AmfSubscriptionInfo.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

