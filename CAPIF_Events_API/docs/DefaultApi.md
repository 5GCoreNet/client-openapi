# \DefaultApi

All URIs are relative to *https://example.com/capif-events/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubscriberIdSubscriptionsPost**](DefaultApi.md#SubscriberIdSubscriptionsPost) | **Post** /{subscriberId}/subscriptions | 
[**SubscriberIdSubscriptionsSubscriptionIdDelete**](DefaultApi.md#SubscriberIdSubscriptionsSubscriptionIdDelete) | **Delete** /{subscriberId}/subscriptions/{subscriptionId} | 



## SubscriberIdSubscriptionsPost

> EventSubscription SubscriberIdSubscriptionsPost(ctx, subscriberId).EventSubscription(eventSubscription).Execute()





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
    subscriberId := "subscriberId_example" // string | Identifier of the Subscriber
    eventSubscription := *openapiclient.NewEventSubscription([]openapiclient.CAPIFEvent{*openapiclient.NewCAPIFEvent()}, "NotificationDestination_example") // EventSubscription | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.SubscriberIdSubscriptionsPost(context.Background(), subscriberId).EventSubscription(eventSubscription).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SubscriberIdSubscriptionsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubscriberIdSubscriptionsPost`: EventSubscription
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SubscriberIdSubscriptionsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriberId** | **string** | Identifier of the Subscriber | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriberIdSubscriptionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eventSubscription** | [**EventSubscription**](EventSubscription.md) |  | 

### Return type

[**EventSubscription**](EventSubscription.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriberIdSubscriptionsSubscriptionIdDelete

> SubscriberIdSubscriptionsSubscriptionIdDelete(ctx, subscriberId, subscriptionId).Execute()





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
    subscriberId := "subscriberId_example" // string | Identifier of the Subscriber
    subscriptionId := "subscriptionId_example" // string | Identifier of an individual Events Subscription

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.SubscriberIdSubscriptionsSubscriptionIdDelete(context.Background(), subscriberId, subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SubscriberIdSubscriptionsSubscriptionIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriberId** | **string** | Identifier of the Subscriber | 
**subscriptionId** | **string** | Identifier of an individual Events Subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriberIdSubscriptionsSubscriptionIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

