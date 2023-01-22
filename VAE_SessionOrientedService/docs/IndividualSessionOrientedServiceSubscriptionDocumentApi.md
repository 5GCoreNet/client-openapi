# \IndividualSessionOrientedServiceSubscriptionDocumentApi

All URIs are relative to *https://example.com/vae-session-Oriented-service/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSessionOrientedServiceSubscription**](IndividualSessionOrientedServiceSubscriptionDocumentApi.md#DeleteSessionOrientedServiceSubscription) | **Delete** /subscriptions/{subscriptionId} | VAE Session Oriented Service Subscription resource delete service Operation
[**ReadSessionOrientedServiceSubscription**](IndividualSessionOrientedServiceSubscriptionDocumentApi.md#ReadSessionOrientedServiceSubscription) | **Get** /subscriptions/{subscriptionId} | VAE Session Oriented Service Subscription resource read service Operation



## DeleteSessionOrientedServiceSubscription

> DeleteSessionOrientedServiceSubscription(ctx, subscriptionId).Execute()

VAE Session Oriented Service Subscription resource delete service Operation

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
    subscriptionId := "subscriptionId_example" // string | Unique ID of the Session Oriented Service Subscription n to be deleted

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualSessionOrientedServiceSubscriptionDocumentApi.DeleteSessionOrientedServiceSubscription(context.Background(), subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualSessionOrientedServiceSubscriptionDocumentApi.DeleteSessionOrientedServiceSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Unique ID of the Session Oriented Service Subscription n to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSessionOrientedServiceSubscriptionRequest struct via the builder pattern


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


## ReadSessionOrientedServiceSubscription

> SessionOrientedData ReadSessionOrientedServiceSubscription(ctx, subscriptionId).Execute()

VAE Session Oriented Service Subscription resource read service Operation

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
    subscriptionId := "subscriptionId_example" // string | Identifier of an Session Oriented Service Subscription resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualSessionOrientedServiceSubscriptionDocumentApi.ReadSessionOrientedServiceSubscription(context.Background(), subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualSessionOrientedServiceSubscriptionDocumentApi.ReadSessionOrientedServiceSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadSessionOrientedServiceSubscription`: SessionOrientedData
    fmt.Fprintf(os.Stdout, "Response from `IndividualSessionOrientedServiceSubscriptionDocumentApi.ReadSessionOrientedServiceSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Identifier of an Session Oriented Service Subscription resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadSessionOrientedServiceSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SessionOrientedData**](SessionOrientedData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

