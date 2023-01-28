# \Individual5GLANParametersProvisionSubscriptionApi

All URIs are relative to *https://example.com/3gpp-5glan-pp/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAnSubscription**](Individual5GLANParametersProvisionSubscriptionApi.md#DeleteAnSubscription) | **Delete** /{afId}/subscriptions/{subscriptionId} | Deletes an already existing subscription
[**FullyUpdateAnSubscription**](Individual5GLANParametersProvisionSubscriptionApi.md#FullyUpdateAnSubscription) | **Put** /{afId}/subscriptions/{subscriptionId} | Fully updates/replaces an existing subscription resource
[**PartialUpdateAnSubscription**](Individual5GLANParametersProvisionSubscriptionApi.md#PartialUpdateAnSubscription) | **Patch** /{afId}/subscriptions/{subscriptionId} | Partial updates an existing subscription resource
[**ReadAnSubscription**](Individual5GLANParametersProvisionSubscriptionApi.md#ReadAnSubscription) | **Get** /{afId}/subscriptions/{subscriptionId} | read an active subscription for the AF and the subscription Id



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
    openapiclient "github.com/5GCoreNet/client-openapi/5GLANParameterProvision"
)

func main() {
    afId := "afId_example" // string | Identifier of the AF
    subscriptionId := "subscriptionId_example" // string | Identifier of the subscription resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Individual5GLANParametersProvisionSubscriptionApi.DeleteAnSubscription(context.Background(), afId, subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Individual5GLANParametersProvisionSubscriptionApi.DeleteAnSubscription``: %v\n", err)
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

> Model5GLanParametersProvision FullyUpdateAnSubscription(ctx, afId, subscriptionId).Model5GLanParametersProvision(model5GLanParametersProvision).Execute()

Fully updates/replaces an existing subscription resource

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/5GLANParameterProvision"
)

func main() {
    afId := "afId_example" // string | Identifier of the AF
    subscriptionId := "subscriptionId_example" // string | Identifier of the subscription resource
    model5GLanParametersProvision := *openapiclient.NewModel5GLanParametersProvision(*openapiclient.NewModel5GLanParameters("ExterGroupId_example", map[string]string{"key": "Inner_example"}, "Dnn_example", *openapiclient.NewSnssai(int32(123)), *openapiclient.NewPduSessionType(), map[string]AppDescriptor{"key": *openapiclient.NewAppDescriptor("OsId_example", map[string]string{"key": "Inner_example"})}), "SuppFeat_example") // Model5GLanParametersProvision | Parameters to update/replace the existing subscription

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Individual5GLANParametersProvisionSubscriptionApi.FullyUpdateAnSubscription(context.Background(), afId, subscriptionId).Model5GLanParametersProvision(model5GLanParametersProvision).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Individual5GLANParametersProvisionSubscriptionApi.FullyUpdateAnSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FullyUpdateAnSubscription`: Model5GLanParametersProvision
    fmt.Fprintf(os.Stdout, "Response from `Individual5GLANParametersProvisionSubscriptionApi.FullyUpdateAnSubscription`: %v\n", resp)
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


 **model5GLanParametersProvision** | [**Model5GLanParametersProvision**](Model5GLanParametersProvision.md) | Parameters to update/replace the existing subscription | 

### Return type

[**Model5GLanParametersProvision**](Model5GLanParametersProvision.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartialUpdateAnSubscription

> Model5GLanParametersProvision PartialUpdateAnSubscription(ctx, afId, subscriptionId).Model5GLanParametersProvisionPatch(model5GLanParametersProvisionPatch).Execute()

Partial updates an existing subscription resource

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/5GLANParameterProvision"
)

func main() {
    afId := "afId_example" // string | Identifier of the AF
    subscriptionId := "subscriptionId_example" // string | Identifier of the subscription resource
    model5GLanParametersProvisionPatch := *openapiclient.NewModel5GLanParametersProvisionPatch() // Model5GLanParametersProvisionPatch | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Individual5GLANParametersProvisionSubscriptionApi.PartialUpdateAnSubscription(context.Background(), afId, subscriptionId).Model5GLanParametersProvisionPatch(model5GLanParametersProvisionPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Individual5GLANParametersProvisionSubscriptionApi.PartialUpdateAnSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PartialUpdateAnSubscription`: Model5GLanParametersProvision
    fmt.Fprintf(os.Stdout, "Response from `Individual5GLANParametersProvisionSubscriptionApi.PartialUpdateAnSubscription`: %v\n", resp)
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


 **model5GLanParametersProvisionPatch** | [**Model5GLanParametersProvisionPatch**](Model5GLanParametersProvisionPatch.md) |  | 

### Return type

[**Model5GLanParametersProvision**](Model5GLanParametersProvision.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadAnSubscription

> Model5GLanParametersProvision ReadAnSubscription(ctx, afId, subscriptionId).Execute()

read an active subscription for the AF and the subscription Id

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/5GLANParameterProvision"
)

func main() {
    afId := "afId_example" // string | Identifier of the AF
    subscriptionId := "subscriptionId_example" // string | Identifier of the subscription resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Individual5GLANParametersProvisionSubscriptionApi.ReadAnSubscription(context.Background(), afId, subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Individual5GLANParametersProvisionSubscriptionApi.ReadAnSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAnSubscription`: Model5GLanParametersProvision
    fmt.Fprintf(os.Stdout, "Response from `Individual5GLANParametersProvisionSubscriptionApi.ReadAnSubscription`: %v\n", resp)
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

[**Model5GLanParametersProvision**](Model5GLanParametersProvision.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

