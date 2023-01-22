# \IndividualASSessionWithRequiredQoSSubscriptionApi

All URIs are relative to *https://example.com/3gpp-as-session-with-qos/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIndASSessionWithQoSSubscription**](IndividualASSessionWithRequiredQoSSubscriptionApi.md#DeleteIndASSessionWithQoSSubscription) | **Delete** /{scsAsId}/subscriptions/{subscriptionId} | Deletes an already existing subscription.
[**FetchIndASSessionWithQoSSubscription**](IndividualASSessionWithRequiredQoSSubscriptionApi.md#FetchIndASSessionWithQoSSubscription) | **Get** /{scsAsId}/subscriptions/{subscriptionId} | Read an active subscriptions for the SCS/AS and the subscription Id.
[**ModifyIndASSessionWithQoSSubscription**](IndividualASSessionWithRequiredQoSSubscriptionApi.md#ModifyIndASSessionWithQoSSubscription) | **Patch** /{scsAsId}/subscriptions/{subscriptionId} | Updates/replaces an existing subscription resource.
[**UpdateIndASSessionWithQoSSubscription**](IndividualASSessionWithRequiredQoSSubscriptionApi.md#UpdateIndASSessionWithQoSSubscription) | **Put** /{scsAsId}/subscriptions/{subscriptionId} | Updates/replaces an existing subscription resource.



## DeleteIndASSessionWithQoSSubscription

> UserPlaneNotificationData DeleteIndASSessionWithQoSSubscription(ctx, scsAsId, subscriptionId).Execute()

Deletes an already existing subscription.

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
    scsAsId := "scsAsId_example" // string | Identifier of the SCS/AS
    subscriptionId := "subscriptionId_example" // string | Identifier of the subscription resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualASSessionWithRequiredQoSSubscriptionApi.DeleteIndASSessionWithQoSSubscription(context.Background(), scsAsId, subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualASSessionWithRequiredQoSSubscriptionApi.DeleteIndASSessionWithQoSSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteIndASSessionWithQoSSubscription`: UserPlaneNotificationData
    fmt.Fprintf(os.Stdout, "Response from `IndividualASSessionWithRequiredQoSSubscriptionApi.DeleteIndASSessionWithQoSSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of the SCS/AS | 
**subscriptionId** | **string** | Identifier of the subscription resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIndASSessionWithQoSSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**UserPlaneNotificationData**](UserPlaneNotificationData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchIndASSessionWithQoSSubscription

> AsSessionWithQoSSubscription FetchIndASSessionWithQoSSubscription(ctx, scsAsId, subscriptionId).Execute()

Read an active subscriptions for the SCS/AS and the subscription Id.

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
    scsAsId := "scsAsId_example" // string | Identifier of the SCS/AS
    subscriptionId := "subscriptionId_example" // string | Identifier of the subscription resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualASSessionWithRequiredQoSSubscriptionApi.FetchIndASSessionWithQoSSubscription(context.Background(), scsAsId, subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualASSessionWithRequiredQoSSubscriptionApi.FetchIndASSessionWithQoSSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchIndASSessionWithQoSSubscription`: AsSessionWithQoSSubscription
    fmt.Fprintf(os.Stdout, "Response from `IndividualASSessionWithRequiredQoSSubscriptionApi.FetchIndASSessionWithQoSSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of the SCS/AS | 
**subscriptionId** | **string** | Identifier of the subscription resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchIndASSessionWithQoSSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AsSessionWithQoSSubscription**](AsSessionWithQoSSubscription.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyIndASSessionWithQoSSubscription

> AsSessionWithQoSSubscription ModifyIndASSessionWithQoSSubscription(ctx, scsAsId, subscriptionId).AsSessionWithQoSSubscriptionPatch(asSessionWithQoSSubscriptionPatch).Execute()

Updates/replaces an existing subscription resource.

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
    scsAsId := "scsAsId_example" // string | Identifier of the SCS/AS
    subscriptionId := "subscriptionId_example" // string | Identifier of the subscription resource
    asSessionWithQoSSubscriptionPatch := *openapiclient.NewAsSessionWithQoSSubscriptionPatch() // AsSessionWithQoSSubscriptionPatch | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualASSessionWithRequiredQoSSubscriptionApi.ModifyIndASSessionWithQoSSubscription(context.Background(), scsAsId, subscriptionId).AsSessionWithQoSSubscriptionPatch(asSessionWithQoSSubscriptionPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualASSessionWithRequiredQoSSubscriptionApi.ModifyIndASSessionWithQoSSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ModifyIndASSessionWithQoSSubscription`: AsSessionWithQoSSubscription
    fmt.Fprintf(os.Stdout, "Response from `IndividualASSessionWithRequiredQoSSubscriptionApi.ModifyIndASSessionWithQoSSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of the SCS/AS | 
**subscriptionId** | **string** | Identifier of the subscription resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyIndASSessionWithQoSSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **asSessionWithQoSSubscriptionPatch** | [**AsSessionWithQoSSubscriptionPatch**](AsSessionWithQoSSubscriptionPatch.md) |  | 

### Return type

[**AsSessionWithQoSSubscription**](AsSessionWithQoSSubscription.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIndASSessionWithQoSSubscription

> AsSessionWithQoSSubscription UpdateIndASSessionWithQoSSubscription(ctx, scsAsId, subscriptionId).AsSessionWithQoSSubscription(asSessionWithQoSSubscription).Execute()

Updates/replaces an existing subscription resource.

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
    scsAsId := "scsAsId_example" // string | Identifier of the SCS/AS
    subscriptionId := "subscriptionId_example" // string | Identifier of the subscription resource
    asSessionWithQoSSubscription := *openapiclient.NewAsSessionWithQoSSubscription("NotificationDestination_example") // AsSessionWithQoSSubscription | Parameters to update/replace the existing subscription

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualASSessionWithRequiredQoSSubscriptionApi.UpdateIndASSessionWithQoSSubscription(context.Background(), scsAsId, subscriptionId).AsSessionWithQoSSubscription(asSessionWithQoSSubscription).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualASSessionWithRequiredQoSSubscriptionApi.UpdateIndASSessionWithQoSSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIndASSessionWithQoSSubscription`: AsSessionWithQoSSubscription
    fmt.Fprintf(os.Stdout, "Response from `IndividualASSessionWithRequiredQoSSubscriptionApi.UpdateIndASSessionWithQoSSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of the SCS/AS | 
**subscriptionId** | **string** | Identifier of the subscription resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIndASSessionWithQoSSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **asSessionWithQoSSubscription** | [**AsSessionWithQoSSubscription**](AsSessionWithQoSSubscription.md) | Parameters to update/replace the existing subscription | 

### Return type

[**AsSessionWithQoSSubscription**](AsSessionWithQoSSubscription.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

