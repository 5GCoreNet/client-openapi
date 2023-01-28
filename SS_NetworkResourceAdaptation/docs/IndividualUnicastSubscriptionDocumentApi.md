# \IndividualUnicastSubscriptionDocumentApi

All URIs are relative to *https://example.com/ss-nra/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteUnicastSubscription**](IndividualUnicastSubscriptionDocumentApi.md#DeleteUnicastSubscription) | **Delete** /unicast-subscriptions/{uniSubId} | Delete an existing Individual Unicast Subscription
[**GetUnicastSubscription**](IndividualUnicastSubscriptionDocumentApi.md#GetUnicastSubscription) | **Get** /unicast-subscriptions/{uniSubId} | Reads an existing Individual Unicast Subscription



## DeleteUnicastSubscription

> DeleteUnicastSubscription(ctx, uniSubId).Execute()

Delete an existing Individual Unicast Subscription

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/SS_NetworkResourceAdaptation"
)

func main() {
    uniSubId := "uniSubId_example" // string | Unicast Subscription ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualUnicastSubscriptionDocumentApi.DeleteUnicastSubscription(context.Background(), uniSubId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualUnicastSubscriptionDocumentApi.DeleteUnicastSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uniSubId** | **string** | Unicast Subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUnicastSubscriptionRequest struct via the builder pattern


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


## GetUnicastSubscription

> UnicastSubscription GetUnicastSubscription(ctx, uniSubId).Execute()

Reads an existing Individual Unicast Subscription

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/SS_NetworkResourceAdaptation"
)

func main() {
    uniSubId := "uniSubId_example" // string | Unicast Subscription ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualUnicastSubscriptionDocumentApi.GetUnicastSubscription(context.Background(), uniSubId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualUnicastSubscriptionDocumentApi.GetUnicastSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUnicastSubscription`: UnicastSubscription
    fmt.Fprintf(os.Stdout, "Response from `IndividualUnicastSubscriptionDocumentApi.GetUnicastSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uniSubId** | **string** | Unicast Subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUnicastSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UnicastSubscription**](UnicastSubscription.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

