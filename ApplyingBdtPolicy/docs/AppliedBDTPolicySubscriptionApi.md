# \AppliedBDTPolicySubscriptionApi

All URIs are relative to *https://example.com/3gpp-applying-bdt-policy/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNewSubscription**](AppliedBDTPolicySubscriptionApi.md#CreateNewSubscription) | **Post** /{afId}/subscriptions | Creates a new subscription resource
[**ReadAllSubscriptions**](AppliedBDTPolicySubscriptionApi.md#ReadAllSubscriptions) | **Get** /{afId}/subscriptions | read all of the active subscriptions for the AF



## CreateNewSubscription

> AppliedBdtPolicy CreateNewSubscription(ctx, afId).AppliedBdtPolicy(appliedBdtPolicy).Execute()

Creates a new subscription resource

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
    appliedBdtPolicy := openapiclient.AppliedBdtPolicy{Interface{}: new(interface{})} // AppliedBdtPolicy | Request to create a new subscription resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppliedBDTPolicySubscriptionApi.CreateNewSubscription(context.Background(), afId).AppliedBdtPolicy(appliedBdtPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppliedBDTPolicySubscriptionApi.CreateNewSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNewSubscription`: AppliedBdtPolicy
    fmt.Fprintf(os.Stdout, "Response from `AppliedBDTPolicySubscriptionApi.CreateNewSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**afId** | **string** | Identifier of the AF | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNewSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appliedBdtPolicy** | [**AppliedBdtPolicy**](AppliedBdtPolicy.md) | Request to create a new subscription resource | 

### Return type

[**AppliedBdtPolicy**](AppliedBdtPolicy.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadAllSubscriptions

> []AppliedBdtPolicy ReadAllSubscriptions(ctx, afId).Execute()

read all of the active subscriptions for the AF

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppliedBDTPolicySubscriptionApi.ReadAllSubscriptions(context.Background(), afId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppliedBDTPolicySubscriptionApi.ReadAllSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllSubscriptions`: []AppliedBdtPolicy
    fmt.Fprintf(os.Stdout, "Response from `AppliedBDTPolicySubscriptionApi.ReadAllSubscriptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**afId** | **string** | Identifier of the AF | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAllSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]AppliedBdtPolicy**](AppliedBdtPolicy.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

