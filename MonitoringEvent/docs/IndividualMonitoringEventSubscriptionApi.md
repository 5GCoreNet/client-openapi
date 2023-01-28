# \IndividualMonitoringEventSubscriptionApi

All URIs are relative to *https://example.com/3gpp-monitoring-event/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIndMonitoringEventSubscription**](IndividualMonitoringEventSubscriptionApi.md#DeleteIndMonitoringEventSubscription) | **Delete** /{scsAsId}/subscriptions/{subscriptionId} | Deletes an already existing monitoring event subscription.
[**FetchIndMonitoringEventSubscription**](IndividualMonitoringEventSubscriptionApi.md#FetchIndMonitoringEventSubscription) | **Get** /{scsAsId}/subscriptions/{subscriptionId} | Read an active subscriptions for the SCS/AS and the subscription Id.
[**ModifyIndMonitoringEventSubscription**](IndividualMonitoringEventSubscriptionApi.md#ModifyIndMonitoringEventSubscription) | **Patch** /{scsAsId}/subscriptions/{subscriptionId} | Modifies an existing subscription of monitoring event.
[**UpdateIndMonitoringEventSubscription**](IndividualMonitoringEventSubscriptionApi.md#UpdateIndMonitoringEventSubscription) | **Put** /{scsAsId}/subscriptions/{subscriptionId} | Updates/replaces an existing subscription resource.



## DeleteIndMonitoringEventSubscription

> []MonitoringEventReport DeleteIndMonitoringEventSubscription(ctx, scsAsId, subscriptionId).Execute()

Deletes an already existing monitoring event subscription.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/MonitoringEvent"
)

func main() {
    scsAsId := "scsAsId_example" // string | Identifier of the SCS/AS
    subscriptionId := "subscriptionId_example" // string | Identifier of the subscription resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualMonitoringEventSubscriptionApi.DeleteIndMonitoringEventSubscription(context.Background(), scsAsId, subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualMonitoringEventSubscriptionApi.DeleteIndMonitoringEventSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteIndMonitoringEventSubscription`: []MonitoringEventReport
    fmt.Fprintf(os.Stdout, "Response from `IndividualMonitoringEventSubscriptionApi.DeleteIndMonitoringEventSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of the SCS/AS | 
**subscriptionId** | **string** | Identifier of the subscription resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIndMonitoringEventSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]MonitoringEventReport**](MonitoringEventReport.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchIndMonitoringEventSubscription

> MonitoringEventSubscription FetchIndMonitoringEventSubscription(ctx, scsAsId, subscriptionId).Execute()

Read an active subscriptions for the SCS/AS and the subscription Id.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/MonitoringEvent"
)

func main() {
    scsAsId := "scsAsId_example" // string | Identifier of the SCS/AS
    subscriptionId := "subscriptionId_example" // string | Identifier of the subscription resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualMonitoringEventSubscriptionApi.FetchIndMonitoringEventSubscription(context.Background(), scsAsId, subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualMonitoringEventSubscriptionApi.FetchIndMonitoringEventSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchIndMonitoringEventSubscription`: MonitoringEventSubscription
    fmt.Fprintf(os.Stdout, "Response from `IndividualMonitoringEventSubscriptionApi.FetchIndMonitoringEventSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of the SCS/AS | 
**subscriptionId** | **string** | Identifier of the subscription resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchIndMonitoringEventSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**MonitoringEventSubscription**](MonitoringEventSubscription.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyIndMonitoringEventSubscription

> ModifyIndMonitoringEventSubscription(ctx, scsAsId, subscriptionId).PatchItem(patchItem).Execute()

Modifies an existing subscription of monitoring event.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/MonitoringEvent"
)

func main() {
    scsAsId := "scsAsId_example" // string | Identifier of the SCS/AS.
    subscriptionId := "subscriptionId_example" // string | Identifier of the subscription resource.
    patchItem := []openapiclient.PatchItem{*openapiclient.NewPatchItem(*openapiclient.NewPatchOperation(), "Path_example")} // []PatchItem | This is used for PATCH request for partial cancellation and/or partial addition of certain UE(s) within an active group.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualMonitoringEventSubscriptionApi.ModifyIndMonitoringEventSubscription(context.Background(), scsAsId, subscriptionId).PatchItem(patchItem).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualMonitoringEventSubscriptionApi.ModifyIndMonitoringEventSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of the SCS/AS. | 
**subscriptionId** | **string** | Identifier of the subscription resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyIndMonitoringEventSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchItem** | [**[]PatchItem**](PatchItem.md) | This is used for PATCH request for partial cancellation and/or partial addition of certain UE(s) within an active group. | 

### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIndMonitoringEventSubscription

> MonitoringEventSubscription UpdateIndMonitoringEventSubscription(ctx, scsAsId, subscriptionId).MonitoringEventSubscription(monitoringEventSubscription).Execute()

Updates/replaces an existing subscription resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/MonitoringEvent"
)

func main() {
    scsAsId := "scsAsId_example" // string | Identifier of the SCS/AS
    subscriptionId := "subscriptionId_example" // string | Identifier of the subscription resource
    monitoringEventSubscription := *openapiclient.NewMonitoringEventSubscription("NotificationDestination_example", *openapiclient.NewMonitoringType()) // MonitoringEventSubscription | Parameters to update/replace the existing subscription

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualMonitoringEventSubscriptionApi.UpdateIndMonitoringEventSubscription(context.Background(), scsAsId, subscriptionId).MonitoringEventSubscription(monitoringEventSubscription).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualMonitoringEventSubscriptionApi.UpdateIndMonitoringEventSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIndMonitoringEventSubscription`: MonitoringEventSubscription
    fmt.Fprintf(os.Stdout, "Response from `IndividualMonitoringEventSubscriptionApi.UpdateIndMonitoringEventSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of the SCS/AS | 
**subscriptionId** | **string** | Identifier of the subscription resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIndMonitoringEventSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **monitoringEventSubscription** | [**MonitoringEventSubscription**](MonitoringEventSubscription.md) | Parameters to update/replace the existing subscription | 

### Return type

[**MonitoringEventSubscription**](MonitoringEventSubscription.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

