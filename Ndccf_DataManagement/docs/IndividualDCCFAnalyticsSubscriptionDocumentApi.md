# \IndividualDCCFAnalyticsSubscriptionDocumentApi

All URIs are relative to *https://example.com/ndccf-datamanagement/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteDCCFAnalyticsSubscription**](IndividualDCCFAnalyticsSubscriptionDocumentApi.md#DeleteDCCFAnalyticsSubscription) | **Delete** /analytics-subscriptions/{subscriptionId} | Deletes an existing Individual DCCF Data Subscription.
[**UpdateDCCFAnalyticsSubscription**](IndividualDCCFAnalyticsSubscriptionDocumentApi.md#UpdateDCCFAnalyticsSubscription) | **Put** /analytics-subscriptions/{subscriptionId} | Updates an existing Individual DCCF Analytics Subscription resource.



## DeleteDCCFAnalyticsSubscription

> DeleteDCCFAnalyticsSubscription(ctx, subscriptionId).Execute()

Deletes an existing Individual DCCF Data Subscription.

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
    subscriptionId := "subscriptionId_example" // string | String identifying an analytics subscription to the Ndccf_DataManagement Service. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualDCCFAnalyticsSubscriptionDocumentApi.DeleteDCCFAnalyticsSubscription(context.Background(), subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualDCCFAnalyticsSubscriptionDocumentApi.DeleteDCCFAnalyticsSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | String identifying an analytics subscription to the Ndccf_DataManagement Service.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDCCFAnalyticsSubscriptionRequest struct via the builder pattern


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


## UpdateDCCFAnalyticsSubscription

> NdccfAnalyticsSubscription UpdateDCCFAnalyticsSubscription(ctx, subscriptionId).NdccfAnalyticsSubscription(ndccfAnalyticsSubscription).Execute()

Updates an existing Individual DCCF Analytics Subscription resource.

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
    subscriptionId := "subscriptionId_example" // string | String identifying an analytics subscription to the Ndccf_DataManagement Service. 
    ndccfAnalyticsSubscription := *openapiclient.NewNdccfAnalyticsSubscription(*openapiclient.NewNnwdafEventsSubscription([]openapiclient.EventSubscription{*openapiclient.NewEventSubscription(*openapiclient.NewNwdafEvent())}), "AnaNotifUri_example", "AnaNotifCorrId_example") // NdccfAnalyticsSubscription | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualDCCFAnalyticsSubscriptionDocumentApi.UpdateDCCFAnalyticsSubscription(context.Background(), subscriptionId).NdccfAnalyticsSubscription(ndccfAnalyticsSubscription).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualDCCFAnalyticsSubscriptionDocumentApi.UpdateDCCFAnalyticsSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDCCFAnalyticsSubscription`: NdccfAnalyticsSubscription
    fmt.Fprintf(os.Stdout, "Response from `IndividualDCCFAnalyticsSubscriptionDocumentApi.UpdateDCCFAnalyticsSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | String identifying an analytics subscription to the Ndccf_DataManagement Service.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDCCFAnalyticsSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ndccfAnalyticsSubscription** | [**NdccfAnalyticsSubscription**](NdccfAnalyticsSubscription.md) |  | 

### Return type

[**NdccfAnalyticsSubscription**](NdccfAnalyticsSubscription.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

