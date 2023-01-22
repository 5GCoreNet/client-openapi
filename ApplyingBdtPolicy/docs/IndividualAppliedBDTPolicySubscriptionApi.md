# \IndividualAppliedBDTPolicySubscriptionApi

All URIs are relative to *https://example.com/3gpp-applying-bdt-policy/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAnSubscription**](IndividualAppliedBDTPolicySubscriptionApi.md#DeleteAnSubscription) | **Delete** /{afId}/subscriptions/{subscriptionId} | Deletes an already existing subscription
[**PartialUpdateAnSubscription**](IndividualAppliedBDTPolicySubscriptionApi.md#PartialUpdateAnSubscription) | **Patch** /{afId}/subscriptions/{subscriptionId} | Partial updates/replaces an existing subscription resource
[**ReadAnSubscription**](IndividualAppliedBDTPolicySubscriptionApi.md#ReadAnSubscription) | **Get** /{afId}/subscriptions/{subscriptionId} | read an active subscriptions for the SCS/AS and the subscription Id



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
    resp, r, err := apiClient.IndividualAppliedBDTPolicySubscriptionApi.DeleteAnSubscription(context.Background(), afId, subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualAppliedBDTPolicySubscriptionApi.DeleteAnSubscription``: %v\n", err)
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


## PartialUpdateAnSubscription

> AppliedBdtPolicy PartialUpdateAnSubscription(ctx, afId, subscriptionId).AppliedBdtPolicyPatch(appliedBdtPolicyPatch).Execute()

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
    appliedBdtPolicyPatch := *openapiclient.NewAppliedBdtPolicyPatch("BdtRefId_example") // AppliedBdtPolicyPatch | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualAppliedBDTPolicySubscriptionApi.PartialUpdateAnSubscription(context.Background(), afId, subscriptionId).AppliedBdtPolicyPatch(appliedBdtPolicyPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualAppliedBDTPolicySubscriptionApi.PartialUpdateAnSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PartialUpdateAnSubscription`: AppliedBdtPolicy
    fmt.Fprintf(os.Stdout, "Response from `IndividualAppliedBDTPolicySubscriptionApi.PartialUpdateAnSubscription`: %v\n", resp)
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


 **appliedBdtPolicyPatch** | [**AppliedBdtPolicyPatch**](AppliedBdtPolicyPatch.md) |  | 

### Return type

[**AppliedBdtPolicy**](AppliedBdtPolicy.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadAnSubscription

> AppliedBdtPolicy ReadAnSubscription(ctx, afId, subscriptionId).Execute()

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
    resp, r, err := apiClient.IndividualAppliedBDTPolicySubscriptionApi.ReadAnSubscription(context.Background(), afId, subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualAppliedBDTPolicySubscriptionApi.ReadAnSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAnSubscription`: AppliedBdtPolicy
    fmt.Fprintf(os.Stdout, "Response from `IndividualAppliedBDTPolicySubscriptionApi.ReadAnSubscription`: %v\n", resp)
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

[**AppliedBdtPolicy**](AppliedBdtPolicy.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

