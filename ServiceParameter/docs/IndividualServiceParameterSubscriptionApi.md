# \IndividualServiceParameterSubscriptionApi

All URIs are relative to *https://example.com/3gpp-service-parameter/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAnSubscription**](IndividualServiceParameterSubscriptionApi.md#DeleteAnSubscription) | **Delete** /{afId}/subscriptions/{subscriptionId} | Deletes an already existing subscription
[**FullyUpdateAnSubscription**](IndividualServiceParameterSubscriptionApi.md#FullyUpdateAnSubscription) | **Put** /{afId}/subscriptions/{subscriptionId} | Fully updates/replaces an existing subscription resource
[**PartialUpdateAnSubscription**](IndividualServiceParameterSubscriptionApi.md#PartialUpdateAnSubscription) | **Patch** /{afId}/subscriptions/{subscriptionId} | Partial updates/replaces an existing subscription resource
[**ReadAnSubscription**](IndividualServiceParameterSubscriptionApi.md#ReadAnSubscription) | **Get** /{afId}/subscriptions/{subscriptionId} | read an active subscriptions for the SCS/AS and the subscription Id



## DeleteAnSubscription

> DeleteAnSubscription(ctx, afId, subscriptionId).Execute()

Deletes an already existing subscription

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
    afId := "afId_example" // string | Identifier of the AF
    subscriptionId := "subscriptionId_example" // string | Identifier of the subscription resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualServiceParameterSubscriptionApi.DeleteAnSubscription(context.Background(), afId, subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualServiceParameterSubscriptionApi.DeleteAnSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**afId** | **string** | Identifier of the AF | 
**subscriptionId** | **string** | Identifier of the subscription resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAnSubscriptionRequest struct via the builder pattern


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


## FullyUpdateAnSubscription

> ServiceParameterData FullyUpdateAnSubscription(ctx, afId, subscriptionId).ServiceParameterData(serviceParameterData).Execute()

Fully updates/replaces an existing subscription resource

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
    afId := "afId_example" // string | Identifier of the AF
    subscriptionId := "subscriptionId_example" // string | Identifier of the subscription resource
    serviceParameterData := *openapiclient.NewServiceParameterData() // ServiceParameterData | Parameters to update/replace the existing subscription

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualServiceParameterSubscriptionApi.FullyUpdateAnSubscription(context.Background(), afId, subscriptionId).ServiceParameterData(serviceParameterData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualServiceParameterSubscriptionApi.FullyUpdateAnSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FullyUpdateAnSubscription`: ServiceParameterData
    fmt.Fprintf(os.Stdout, "Response from `IndividualServiceParameterSubscriptionApi.FullyUpdateAnSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**afId** | **string** | Identifier of the AF | 
**subscriptionId** | **string** | Identifier of the subscription resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiFullyUpdateAnSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **serviceParameterData** | [**ServiceParameterData**](ServiceParameterData.md) | Parameters to update/replace the existing subscription | 

### Return type

[**ServiceParameterData**](ServiceParameterData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartialUpdateAnSubscription

> ServiceParameterData PartialUpdateAnSubscription(ctx, afId, subscriptionId).ServiceParameterDataPatch(serviceParameterDataPatch).Execute()

Partial updates/replaces an existing subscription resource

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
    afId := "afId_example" // string | Identifier of the AF
    subscriptionId := "subscriptionId_example" // string | Identifier of the subscription resource
    serviceParameterDataPatch := *openapiclient.NewServiceParameterDataPatch() // ServiceParameterDataPatch | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualServiceParameterSubscriptionApi.PartialUpdateAnSubscription(context.Background(), afId, subscriptionId).ServiceParameterDataPatch(serviceParameterDataPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualServiceParameterSubscriptionApi.PartialUpdateAnSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PartialUpdateAnSubscription`: ServiceParameterData
    fmt.Fprintf(os.Stdout, "Response from `IndividualServiceParameterSubscriptionApi.PartialUpdateAnSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**afId** | **string** | Identifier of the AF | 
**subscriptionId** | **string** | Identifier of the subscription resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartialUpdateAnSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **serviceParameterDataPatch** | [**ServiceParameterDataPatch**](ServiceParameterDataPatch.md) |  | 

### Return type

[**ServiceParameterData**](ServiceParameterData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadAnSubscription

> ServiceParameterData ReadAnSubscription(ctx, afId, subscriptionId).Execute()

read an active subscriptions for the SCS/AS and the subscription Id

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
    afId := "afId_example" // string | Identifier of the AF
    subscriptionId := "subscriptionId_example" // string | Identifier of the subscription resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualServiceParameterSubscriptionApi.ReadAnSubscription(context.Background(), afId, subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualServiceParameterSubscriptionApi.ReadAnSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAnSubscription`: ServiceParameterData
    fmt.Fprintf(os.Stdout, "Response from `IndividualServiceParameterSubscriptionApi.ReadAnSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**afId** | **string** | Identifier of the AF | 
**subscriptionId** | **string** | Identifier of the subscription resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAnSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ServiceParameterData**](ServiceParameterData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

