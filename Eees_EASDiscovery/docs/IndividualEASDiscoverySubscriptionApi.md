# \IndividualEASDiscoverySubscriptionApi

All URIs are relative to *https://example.com/eees-easdiscovery/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubscriptionsSubscriptionIdDelete**](IndividualEASDiscoverySubscriptionApi.md#SubscriptionsSubscriptionIdDelete) | **Delete** /subscriptions/{subscriptionId} | 
[**SubscriptionsSubscriptionIdPatch**](IndividualEASDiscoverySubscriptionApi.md#SubscriptionsSubscriptionIdPatch) | **Patch** /subscriptions/{subscriptionId} | 
[**SubscriptionsSubscriptionIdPut**](IndividualEASDiscoverySubscriptionApi.md#SubscriptionsSubscriptionIdPut) | **Put** /subscriptions/{subscriptionId} | 



## SubscriptionsSubscriptionIdDelete

> SubscriptionsSubscriptionIdDelete(ctx, subscriptionId).Execute()





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
    subscriptionId := "subscriptionId_example" // string | Identifies an individual EAS discovery subscription resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualEASDiscoverySubscriptionApi.SubscriptionsSubscriptionIdDelete(context.Background(), subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualEASDiscoverySubscriptionApi.SubscriptionsSubscriptionIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Identifies an individual EAS discovery subscription resource | 

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
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsSubscriptionIdPatch

> EasDiscoverySubscription SubscriptionsSubscriptionIdPatch(ctx, subscriptionId).EasDiscoverySubscriptionPatch(easDiscoverySubscriptionPatch).Execute()





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
    subscriptionId := "subscriptionId_example" // string | Identifies an individual EAS discovery subscription resource
    easDiscoverySubscriptionPatch := *openapiclient.NewEasDiscoverySubscriptionPatch() // EasDiscoverySubscriptionPatch | Parameters to replace the existing subscription

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualEASDiscoverySubscriptionApi.SubscriptionsSubscriptionIdPatch(context.Background(), subscriptionId).EasDiscoverySubscriptionPatch(easDiscoverySubscriptionPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualEASDiscoverySubscriptionApi.SubscriptionsSubscriptionIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubscriptionsSubscriptionIdPatch`: EasDiscoverySubscription
    fmt.Fprintf(os.Stdout, "Response from `IndividualEASDiscoverySubscriptionApi.SubscriptionsSubscriptionIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Identifies an individual EAS discovery subscription resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsSubscriptionIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **easDiscoverySubscriptionPatch** | [**EasDiscoverySubscriptionPatch**](EasDiscoverySubscriptionPatch.md) | Parameters to replace the existing subscription | 

### Return type

[**EasDiscoverySubscription**](EasDiscoverySubscription.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsSubscriptionIdPut

> EasDiscoverySubscription SubscriptionsSubscriptionIdPut(ctx, subscriptionId).EasDiscoverySubscription(easDiscoverySubscription).Execute()





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
    subscriptionId := "subscriptionId_example" // string | Identifies an individual EAS discovery subscription resource
    easDiscoverySubscription := *openapiclient.NewEasDiscoverySubscription("EecId_example", *openapiclient.NewEASDiscEventIDs()) // EasDiscoverySubscription | Parameters to replace the existing subscription

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualEASDiscoverySubscriptionApi.SubscriptionsSubscriptionIdPut(context.Background(), subscriptionId).EasDiscoverySubscription(easDiscoverySubscription).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualEASDiscoverySubscriptionApi.SubscriptionsSubscriptionIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubscriptionsSubscriptionIdPut`: EasDiscoverySubscription
    fmt.Fprintf(os.Stdout, "Response from `IndividualEASDiscoverySubscriptionApi.SubscriptionsSubscriptionIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Identifies an individual EAS discovery subscription resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsSubscriptionIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **easDiscoverySubscription** | [**EasDiscoverySubscription**](EasDiscoverySubscription.md) | Parameters to replace the existing subscription | 

### Return type

[**EasDiscoverySubscription**](EasDiscoverySubscription.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

