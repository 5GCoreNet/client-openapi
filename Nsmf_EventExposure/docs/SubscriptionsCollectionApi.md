# \SubscriptionsCollectionApi

All URIs are relative to *https://example.com/nsmf-event-exposure/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIndividualSubcription**](SubscriptionsCollectionApi.md#CreateIndividualSubcription) | **Post** /subscriptions | Create an individual subscription for event notifications from the SMF



## CreateIndividualSubcription

> NsmfEventExposure CreateIndividualSubcription(ctx).NsmfEventExposure(nsmfEventExposure).Execute()

Create an individual subscription for event notifications from the SMF

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
    nsmfEventExposure := *openapiclient.NewNsmfEventExposure("NotifId_example", "NotifUri_example", []openapiclient.EventSubscription{*openapiclient.NewEventSubscription(*openapiclient.NewSmfEvent())}) // NsmfEventExposure | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionsCollectionApi.CreateIndividualSubcription(context.Background()).NsmfEventExposure(nsmfEventExposure).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsCollectionApi.CreateIndividualSubcription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIndividualSubcription`: NsmfEventExposure
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionsCollectionApi.CreateIndividualSubcription`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIndividualSubcriptionRequest struct via the builder pattern


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

