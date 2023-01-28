# \DefaultApi

All URIs are relative to *https://example.com/nchf-spendinglimitcontrol/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubscriptionsPost**](DefaultApi.md#SubscriptionsPost) | **Post** /subscriptions | 
[**SubscriptionsSubscriptionIdDelete**](DefaultApi.md#SubscriptionsSubscriptionIdDelete) | **Delete** /subscriptions/{subscriptionId} | 
[**SubscriptionsSubscriptionIdPut**](DefaultApi.md#SubscriptionsSubscriptionIdPut) | **Put** /subscriptions/{subscriptionId} | 



## SubscriptionsPost

> SpendingLimitStatus SubscriptionsPost(ctx).SpendingLimitContext(spendingLimitContext).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nchf_SpendingLimitControl"
)

func main() {
    spendingLimitContext := *openapiclient.NewSpendingLimitContext() // SpendingLimitContext | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.SubscriptionsPost(context.Background()).SpendingLimitContext(spendingLimitContext).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SubscriptionsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubscriptionsPost`: SpendingLimitStatus
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SubscriptionsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **spendingLimitContext** | [**SpendingLimitContext**](SpendingLimitContext.md) |  | 

### Return type

[**SpendingLimitStatus**](SpendingLimitStatus.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsSubscriptionIdDelete

> SubscriptionsSubscriptionIdDelete(ctx, subscriptionId).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nchf_SpendingLimitControl"
)

func main() {
    subscriptionId := "subscriptionId_example" // string | Identifies an individual spending limit retrieval subscription.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.SubscriptionsSubscriptionIdDelete(context.Background(), subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SubscriptionsSubscriptionIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Identifies an individual spending limit retrieval subscription. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsSubscriptionIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsSubscriptionIdPut

> SpendingLimitStatus SubscriptionsSubscriptionIdPut(ctx, subscriptionId).SpendingLimitContext(spendingLimitContext).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nchf_SpendingLimitControl"
)

func main() {
    subscriptionId := "subscriptionId_example" // string | Identifies an individual spending limit retrieval subscription.
    spendingLimitContext := *openapiclient.NewSpendingLimitContext() // SpendingLimitContext | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.SubscriptionsSubscriptionIdPut(context.Background(), subscriptionId).SpendingLimitContext(spendingLimitContext).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SubscriptionsSubscriptionIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubscriptionsSubscriptionIdPut`: SpendingLimitStatus
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SubscriptionsSubscriptionIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Identifies an individual spending limit retrieval subscription. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsSubscriptionIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **spendingLimitContext** | [**SpendingLimitContext**](SpendingLimitContext.md) |  | 

### Return type

[**SpendingLimitStatus**](SpendingLimitStatus.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

