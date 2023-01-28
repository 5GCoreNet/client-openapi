# \SEALEventsSubscriptionsCollectionApi

All URIs are relative to *https://example.com/ss-events/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSealEventSubsc**](SEALEventsSubscriptionsCollectionApi.md#CreateSealEventSubsc) | **Post** /subscriptions | 



## CreateSealEventSubsc

> SEALEventSubscription CreateSealEventSubsc(ctx).SEALEventSubscription(sEALEventSubscription).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/SS_Events"
)

func main() {
    sEALEventSubscription := *openapiclient.NewSEALEventSubscription("SubscriberId_example", []openapiclient.EventSubscription{*openapiclient.NewEventSubscription(*openapiclient.NewSEALEvent())}, *openapiclient.NewReportingInformation(), "NotificationDestination_example") // SEALEventSubscription | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SEALEventsSubscriptionsCollectionApi.CreateSealEventSubsc(context.Background()).SEALEventSubscription(sEALEventSubscription).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SEALEventsSubscriptionsCollectionApi.CreateSealEventSubsc``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSealEventSubsc`: SEALEventSubscription
    fmt.Fprintf(os.Stdout, "Response from `SEALEventsSubscriptionsCollectionApi.CreateSealEventSubsc`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSealEventSubscRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sEALEventSubscription** | [**SEALEventSubscription**](SEALEventSubscription.md) |  | 

### Return type

[**SEALEventSubscription**](SEALEventSubscription.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

