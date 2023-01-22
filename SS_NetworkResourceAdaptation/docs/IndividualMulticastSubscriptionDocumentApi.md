# \IndividualMulticastSubscriptionDocumentApi

All URIs are relative to *https://example.com/ss-nra/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMulticastSubscription**](IndividualMulticastSubscriptionDocumentApi.md#DeleteMulticastSubscription) | **Delete** /multicast-subscriptions/{multiSubId} | Delete an existing Individual Multicast Subscription
[**GetMulticastSubscription**](IndividualMulticastSubscriptionDocumentApi.md#GetMulticastSubscription) | **Get** /multicast-subscriptions/{multiSubId} | Reads an existing Individual Multicast Subscription



## DeleteMulticastSubscription

> DeleteMulticastSubscription(ctx, multiSubId).Execute()

Delete an existing Individual Multicast Subscription

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
    multiSubId := "multiSubId_example" // string | Multicast Subscription ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualMulticastSubscriptionDocumentApi.DeleteMulticastSubscription(context.Background(), multiSubId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualMulticastSubscriptionDocumentApi.DeleteMulticastSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**multiSubId** | **string** | Multicast Subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMulticastSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMulticastSubscription

> MulticastSubscription GetMulticastSubscription(ctx, multiSubId).Execute()

Reads an existing Individual Multicast Subscription

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
    multiSubId := "multiSubId_example" // string | Multicast Subscription ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualMulticastSubscriptionDocumentApi.GetMulticastSubscription(context.Background(), multiSubId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualMulticastSubscriptionDocumentApi.GetMulticastSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMulticastSubscription`: MulticastSubscription
    fmt.Fprintf(os.Stdout, "Response from `IndividualMulticastSubscriptionDocumentApi.GetMulticastSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**multiSubId** | **string** | Multicast Subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMulticastSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MulticastSubscription**](MulticastSubscription.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

