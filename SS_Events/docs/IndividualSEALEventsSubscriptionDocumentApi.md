# \IndividualSEALEventsSubscriptionDocumentApi

All URIs are relative to *https://example.com/ss-events/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIndSealEventSubsc**](IndividualSEALEventsSubscriptionDocumentApi.md#DeleteIndSealEventSubsc) | **Delete** /subscriptions/{subscriptionId} | 
[**ModifyIndSealEventSubsc**](IndividualSEALEventsSubscriptionDocumentApi.md#ModifyIndSealEventSubsc) | **Patch** /subscriptions/{subscriptionId} | 
[**UpdateIndSealEventSubsc**](IndividualSEALEventsSubscriptionDocumentApi.md#UpdateIndSealEventSubsc) | **Put** /subscriptions/{subscriptionId} | 



## DeleteIndSealEventSubsc

> DeleteIndSealEventSubsc(ctx, subscriptionId).Execute()





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
    subscriptionId := "subscriptionId_example" // string | Identifier of an individual Events Subscription

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualSEALEventsSubscriptionDocumentApi.DeleteIndSealEventSubsc(context.Background(), subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualSEALEventsSubscriptionDocumentApi.DeleteIndSealEventSubsc``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Identifier of an individual Events Subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIndSealEventSubscRequest struct via the builder pattern


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


## ModifyIndSealEventSubsc

> SEALEventSubscription ModifyIndSealEventSubsc(ctx, subscriptionId).SEALEventSubscriptionPatch(sEALEventSubscriptionPatch).Execute()





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
    subscriptionId := "subscriptionId_example" // string | Identifier of an individual Events Subscription
    sEALEventSubscriptionPatch := *openapiclient.NewSEALEventSubscriptionPatch() // SEALEventSubscriptionPatch | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualSEALEventsSubscriptionDocumentApi.ModifyIndSealEventSubsc(context.Background(), subscriptionId).SEALEventSubscriptionPatch(sEALEventSubscriptionPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualSEALEventsSubscriptionDocumentApi.ModifyIndSealEventSubsc``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ModifyIndSealEventSubsc`: SEALEventSubscription
    fmt.Fprintf(os.Stdout, "Response from `IndividualSEALEventsSubscriptionDocumentApi.ModifyIndSealEventSubsc`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Identifier of an individual Events Subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyIndSealEventSubscRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sEALEventSubscriptionPatch** | [**SEALEventSubscriptionPatch**](SEALEventSubscriptionPatch.md) |  | 

### Return type

[**SEALEventSubscription**](SEALEventSubscription.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIndSealEventSubsc

> SEALEventSubscription UpdateIndSealEventSubsc(ctx, subscriptionId).SEALEventSubscription(sEALEventSubscription).Execute()





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
    subscriptionId := "subscriptionId_example" // string | Identifier of an individual Events Subscription
    sEALEventSubscription := *openapiclient.NewSEALEventSubscription("SubscriberId_example", []openapiclient.EventSubscription{*openapiclient.NewEventSubscription(*openapiclient.NewSEALEvent())}, *openapiclient.NewReportingInformation(), "NotificationDestination_example") // SEALEventSubscription | Individual SEAL events subscription to be replaced.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualSEALEventsSubscriptionDocumentApi.UpdateIndSealEventSubsc(context.Background(), subscriptionId).SEALEventSubscription(sEALEventSubscription).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualSEALEventsSubscriptionDocumentApi.UpdateIndSealEventSubsc``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIndSealEventSubsc`: SEALEventSubscription
    fmt.Fprintf(os.Stdout, "Response from `IndividualSEALEventsSubscriptionDocumentApi.UpdateIndSealEventSubsc`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Identifier of an individual Events Subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIndSealEventSubscRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sEALEventSubscription** | [**SEALEventSubscription**](SEALEventSubscription.md) | Individual SEAL events subscription to be replaced. | 

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

