# \IndividualACREventsSubscriptionApi

All URIs are relative to *https://example.com/eees-acrevents/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteACREventsSubscription**](IndividualACREventsSubscriptionApi.md#DeleteACREventsSubscription) | **Delete** /subscriptions/{subscriptionId} | 
[**ModifyACREventsSubscription**](IndividualACREventsSubscriptionApi.md#ModifyACREventsSubscription) | **Patch** /subscriptions/{subscriptionId} | 
[**UpdateACREventsSubscription**](IndividualACREventsSubscriptionApi.md#UpdateACREventsSubscription) | **Put** /subscriptions/{subscriptionId} | 



## DeleteACREventsSubscription

> DeleteACREventsSubscription(ctx, subscriptionId).Execute()





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
    subscriptionId := "subscriptionId_example" // string | Identifies an individual ACR Events subscription resource.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualACREventsSubscriptionApi.DeleteACREventsSubscription(context.Background(), subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualACREventsSubscriptionApi.DeleteACREventsSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Identifies an individual ACR Events subscription resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteACREventsSubscriptionRequest struct via the builder pattern


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


## ModifyACREventsSubscription

> ACREventsSubscription ModifyACREventsSubscription(ctx, subscriptionId).ACREventsSubscriptionPatch(aCREventsSubscriptionPatch).Execute()





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
    subscriptionId := "subscriptionId_example" // string | Identifies an individual ACR Events subscription resource.
    aCREventsSubscriptionPatch := *openapiclient.NewACREventsSubscriptionPatch() // ACREventsSubscriptionPatch | Parameters to replace the existing subscription

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualACREventsSubscriptionApi.ModifyACREventsSubscription(context.Background(), subscriptionId).ACREventsSubscriptionPatch(aCREventsSubscriptionPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualACREventsSubscriptionApi.ModifyACREventsSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ModifyACREventsSubscription`: ACREventsSubscription
    fmt.Fprintf(os.Stdout, "Response from `IndividualACREventsSubscriptionApi.ModifyACREventsSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Identifies an individual ACR Events subscription resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyACREventsSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aCREventsSubscriptionPatch** | [**ACREventsSubscriptionPatch**](ACREventsSubscriptionPatch.md) | Parameters to replace the existing subscription | 

### Return type

[**ACREventsSubscription**](ACREventsSubscription.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateACREventsSubscription

> ACREventsSubscription UpdateACREventsSubscription(ctx, subscriptionId).ACREventsSubscription(aCREventsSubscription).Execute()





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
    subscriptionId := "subscriptionId_example" // string | Identifies an individual ACR Events subscription resource
    aCREventsSubscription := *openapiclient.NewACREventsSubscription("EecId_example", []string{"EasIds_example"}, *openapiclient.NewACREventIDs(), "NotificationDestination_example") // ACREventsSubscription | Parameters to replace the existing subscription

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualACREventsSubscriptionApi.UpdateACREventsSubscription(context.Background(), subscriptionId).ACREventsSubscription(aCREventsSubscription).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualACREventsSubscriptionApi.UpdateACREventsSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateACREventsSubscription`: ACREventsSubscription
    fmt.Fprintf(os.Stdout, "Response from `IndividualACREventsSubscriptionApi.UpdateACREventsSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Identifies an individual ACR Events subscription resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateACREventsSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aCREventsSubscription** | [**ACREventsSubscription**](ACREventsSubscription.md) | Parameters to replace the existing subscription | 

### Return type

[**ACREventsSubscription**](ACREventsSubscription.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

