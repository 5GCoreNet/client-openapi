# \IndividualNWDAFEventSubscriptionTransferDocumentApi

All URIs are relative to *https://example.com/nnwdaf-eventssubscription/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteNWDAFEventSubscriptionTransfer**](IndividualNWDAFEventSubscriptionTransferDocumentApi.md#DeleteNWDAFEventSubscriptionTransfer) | **Delete** /transfers/{transferId} | Delete an existing Individual NWDAF Event Subscription Transfer
[**UpdateNWDAFEventSubscriptionTransfer**](IndividualNWDAFEventSubscriptionTransferDocumentApi.md#UpdateNWDAFEventSubscriptionTransfer) | **Put** /transfers/{transferId} | Update an existing Individual NWDAF Event Subscription Transfer



## DeleteNWDAFEventSubscriptionTransfer

> DeleteNWDAFEventSubscriptionTransfer(ctx, transferId).Execute()

Delete an existing Individual NWDAF Event Subscription Transfer

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nnwdaf_EventsSubscription"
)

func main() {
    transferId := "transferId_example" // string | String identifying a request for an analytics subscription transfer to the  Nnwdaf_EventsSubscription Service 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualNWDAFEventSubscriptionTransferDocumentApi.DeleteNWDAFEventSubscriptionTransfer(context.Background(), transferId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualNWDAFEventSubscriptionTransferDocumentApi.DeleteNWDAFEventSubscriptionTransfer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transferId** | **string** | String identifying a request for an analytics subscription transfer to the  Nnwdaf_EventsSubscription Service  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNWDAFEventSubscriptionTransferRequest struct via the builder pattern


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


## UpdateNWDAFEventSubscriptionTransfer

> UpdateNWDAFEventSubscriptionTransfer(ctx, transferId).AnalyticsSubscriptionsTransfer(analyticsSubscriptionsTransfer).Execute()

Update an existing Individual NWDAF Event Subscription Transfer

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nnwdaf_EventsSubscription"
)

func main() {
    transferId := "transferId_example" // string | String identifying a request for an analytics subscription transfer to the  Nnwdaf_EventsSubscription Service 
    analyticsSubscriptionsTransfer := *openapiclient.NewAnalyticsSubscriptionsTransfer([]openapiclient.SubscriptionTransferInfo{*openapiclient.NewSubscriptionTransferInfo(*openapiclient.NewTransferRequestType(), *openapiclient.NewNnwdafEventsSubscription([]openapiclient.EventSubscription{*openapiclient.NewEventSubscription(*openapiclient.NewNwdafEvent())}), "ConsumerId_example")}) // AnalyticsSubscriptionsTransfer | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualNWDAFEventSubscriptionTransferDocumentApi.UpdateNWDAFEventSubscriptionTransfer(context.Background(), transferId).AnalyticsSubscriptionsTransfer(analyticsSubscriptionsTransfer).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualNWDAFEventSubscriptionTransferDocumentApi.UpdateNWDAFEventSubscriptionTransfer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transferId** | **string** | String identifying a request for an analytics subscription transfer to the  Nnwdaf_EventsSubscription Service  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNWDAFEventSubscriptionTransferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **analyticsSubscriptionsTransfer** | [**AnalyticsSubscriptionsTransfer**](AnalyticsSubscriptionsTransfer.md) |  | 

### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

