# \IndividualSubscriptionDocumentApi

All URIs are relative to *https://example.com/nsmf-event-exposure/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIndividualSubcription**](IndividualSubscriptionDocumentApi.md#DeleteIndividualSubcription) | **Delete** /subscriptions/{subId} | Delete an individual subscription for event notifications from the SMF
[**GetIndividualSubcription**](IndividualSubscriptionDocumentApi.md#GetIndividualSubcription) | **Get** /subscriptions/{subId} | Read an individual subscription for event notifications from the SMF
[**ReplaceIndividualSubcription**](IndividualSubscriptionDocumentApi.md#ReplaceIndividualSubcription) | **Put** /subscriptions/{subId} | Replace an individual subscription for event notifications from the SMF



## DeleteIndividualSubcription

> DeleteIndividualSubcription(ctx, subId).Execute()

Delete an individual subscription for event notifications from the SMF

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nsmf_EventExposure"
)

func main() {
    subId := "subId_example" // string | Event Subscription ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualSubscriptionDocumentApi.DeleteIndividualSubcription(context.Background(), subId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualSubscriptionDocumentApi.DeleteIndividualSubcription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subId** | **string** | Event Subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIndividualSubcriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIndividualSubcription

> NsmfEventExposure GetIndividualSubcription(ctx, subId).Execute()

Read an individual subscription for event notifications from the SMF

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nsmf_EventExposure"
)

func main() {
    subId := "subId_example" // string | Event Subscription ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualSubscriptionDocumentApi.GetIndividualSubcription(context.Background(), subId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualSubscriptionDocumentApi.GetIndividualSubcription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIndividualSubcription`: NsmfEventExposure
    fmt.Fprintf(os.Stdout, "Response from `IndividualSubscriptionDocumentApi.GetIndividualSubcription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subId** | **string** | Event Subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIndividualSubcriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NsmfEventExposure**](NsmfEventExposure.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceIndividualSubcription

> NsmfEventExposure ReplaceIndividualSubcription(ctx, subId).NsmfEventExposure(nsmfEventExposure).Execute()

Replace an individual subscription for event notifications from the SMF

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nsmf_EventExposure"
)

func main() {
    subId := "subId_example" // string | Event Subscription ID
    nsmfEventExposure := *openapiclient.NewNsmfEventExposure("NotifId_example", "NotifUri_example", []openapiclient.EventSubscription{*openapiclient.NewEventSubscription(*openapiclient.NewSmfEvent())}) // NsmfEventExposure | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualSubscriptionDocumentApi.ReplaceIndividualSubcription(context.Background(), subId).NsmfEventExposure(nsmfEventExposure).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualSubscriptionDocumentApi.ReplaceIndividualSubcription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceIndividualSubcription`: NsmfEventExposure
    fmt.Fprintf(os.Stdout, "Response from `IndividualSubscriptionDocumentApi.ReplaceIndividualSubcription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subId** | **string** | Event Subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceIndividualSubcriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nsmfEventExposure** | [**NsmfEventExposure**](NsmfEventExposure.md) |  | 

### Return type

[**NsmfEventExposure**](NsmfEventExposure.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

