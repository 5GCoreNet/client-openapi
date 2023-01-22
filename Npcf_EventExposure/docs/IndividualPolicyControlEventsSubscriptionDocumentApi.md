# \IndividualPolicyControlEventsSubscriptionDocumentApi

All URIs are relative to *https://example.com/npcf-eventexposure/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePcEventExposureSubsc**](IndividualPolicyControlEventsSubscriptionDocumentApi.md#DeletePcEventExposureSubsc) | **Delete** /subscriptions/{subscriptionId} | Cancels an existing Individual Policy Control Events Subscription 
[**GetPcEventExposureSubsc**](IndividualPolicyControlEventsSubscriptionDocumentApi.md#GetPcEventExposureSubsc) | **Get** /subscriptions/{subscriptionId} | Reads an existing Individual Policy Control Events Subscription
[**PutPcEventExposureSubsc**](IndividualPolicyControlEventsSubscriptionDocumentApi.md#PutPcEventExposureSubsc) | **Put** /subscriptions/{subscriptionId} | Modifies an existing Individual Policy Control Events Subscription 



## DeletePcEventExposureSubsc

> DeletePcEventExposureSubsc(ctx, subscriptionId).Execute()

Cancels an existing Individual Policy Control Events Subscription 

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
    subscriptionId := "subscriptionId_example" // string | Policy Control Event Subscription ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualPolicyControlEventsSubscriptionDocumentApi.DeletePcEventExposureSubsc(context.Background(), subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualPolicyControlEventsSubscriptionDocumentApi.DeletePcEventExposureSubsc``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Policy Control Event Subscription ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePcEventExposureSubscRequest struct via the builder pattern


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


## GetPcEventExposureSubsc

> PcEventExposureSubsc GetPcEventExposureSubsc(ctx, subscriptionId).Execute()

Reads an existing Individual Policy Control Events Subscription

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
    subscriptionId := "subscriptionId_example" // string | Policy Control Event Subscription ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualPolicyControlEventsSubscriptionDocumentApi.GetPcEventExposureSubsc(context.Background(), subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualPolicyControlEventsSubscriptionDocumentApi.GetPcEventExposureSubsc``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPcEventExposureSubsc`: PcEventExposureSubsc
    fmt.Fprintf(os.Stdout, "Response from `IndividualPolicyControlEventsSubscriptionDocumentApi.GetPcEventExposureSubsc`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Policy Control Event Subscription ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPcEventExposureSubscRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PcEventExposureSubsc**](PcEventExposureSubsc.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutPcEventExposureSubsc

> PcEventExposureSubsc PutPcEventExposureSubsc(ctx, subscriptionId).PcEventExposureSubsc(pcEventExposureSubsc).Execute()

Modifies an existing Individual Policy Control Events Subscription 

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
    subscriptionId := "subscriptionId_example" // string | Policy Control Event Subscription ID.
    pcEventExposureSubsc := *openapiclient.NewPcEventExposureSubsc([]openapiclient.PcEvent{*openapiclient.NewPcEvent()}, "NotifUri_example", "NotifId_example") // PcEventExposureSubsc | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualPolicyControlEventsSubscriptionDocumentApi.PutPcEventExposureSubsc(context.Background(), subscriptionId).PcEventExposureSubsc(pcEventExposureSubsc).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualPolicyControlEventsSubscriptionDocumentApi.PutPcEventExposureSubsc``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutPcEventExposureSubsc`: PcEventExposureSubsc
    fmt.Fprintf(os.Stdout, "Response from `IndividualPolicyControlEventsSubscriptionDocumentApi.PutPcEventExposureSubsc`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Policy Control Event Subscription ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutPcEventExposureSubscRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pcEventExposureSubsc** | [**PcEventExposureSubsc**](PcEventExposureSubsc.md) |  | 

### Return type

[**PcEventExposureSubsc**](PcEventExposureSubsc.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

