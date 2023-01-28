# \IndividualAnalyticsExposureSubscriptionApi

All URIs are relative to *https://example.com/3gpp-analyticsexposure/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAnSubscription**](IndividualAnalyticsExposureSubscriptionApi.md#DeleteAnSubscription) | **Delete** /{afId}/subscriptions/{subscriptionId} | Deletes an already existing subscription
[**FullyUpdateAnSubscription**](IndividualAnalyticsExposureSubscriptionApi.md#FullyUpdateAnSubscription) | **Put** /{afId}/subscriptions/{subscriptionId} | Fully updates/replaces an existing subscription resource
[**ReadAnSubscription**](IndividualAnalyticsExposureSubscriptionApi.md#ReadAnSubscription) | **Get** /{afId}/subscriptions/{subscriptionId} | read an active subscription for the AF and the subscription Id



## DeleteAnSubscription

> DeleteAnSubscription(ctx, afId, subscriptionId).Execute()

Deletes an already existing subscription

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/AnalyticsExposure"
)

func main() {
    afId := "afId_example" // string | Identifier of the AF
    subscriptionId := "subscriptionId_example" // string | Identifier of the subscription resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualAnalyticsExposureSubscriptionApi.DeleteAnSubscription(context.Background(), afId, subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualAnalyticsExposureSubscriptionApi.DeleteAnSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**afId** | **string** | Identifier of the AF | 
**subscriptionId** | **string** | Identifier of the subscription resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAnSubscriptionRequest struct via the builder pattern


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


## FullyUpdateAnSubscription

> AnalyticsExposureSubsc FullyUpdateAnSubscription(ctx, afId, subscriptionId).AnalyticsExposureSubsc(analyticsExposureSubsc).Execute()

Fully updates/replaces an existing subscription resource

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/AnalyticsExposure"
)

func main() {
    afId := "afId_example" // string | Identifier of the AF
    subscriptionId := "subscriptionId_example" // string | Identifier of the subscription resource
    analyticsExposureSubsc := *openapiclient.NewAnalyticsExposureSubsc([]openapiclient.AnalyticsEventSubsc{*openapiclient.NewAnalyticsEventSubsc(*openapiclient.NewAnalyticsEvent())}, "NotifUri_example", "NotifId_example") // AnalyticsExposureSubsc | Parameters to update/replace the existing subscription

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualAnalyticsExposureSubscriptionApi.FullyUpdateAnSubscription(context.Background(), afId, subscriptionId).AnalyticsExposureSubsc(analyticsExposureSubsc).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualAnalyticsExposureSubscriptionApi.FullyUpdateAnSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FullyUpdateAnSubscription`: AnalyticsExposureSubsc
    fmt.Fprintf(os.Stdout, "Response from `IndividualAnalyticsExposureSubscriptionApi.FullyUpdateAnSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**afId** | **string** | Identifier of the AF | 
**subscriptionId** | **string** | Identifier of the subscription resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiFullyUpdateAnSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **analyticsExposureSubsc** | [**AnalyticsExposureSubsc**](AnalyticsExposureSubsc.md) | Parameters to update/replace the existing subscription | 

### Return type

[**AnalyticsExposureSubsc**](AnalyticsExposureSubsc.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadAnSubscription

> AnalyticsExposureSubsc ReadAnSubscription(ctx, afId, subscriptionId).SuppFeat(suppFeat).Execute()

read an active subscription for the AF and the subscription Id

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/AnalyticsExposure"
)

func main() {
    afId := "afId_example" // string | Identifier of the AF
    subscriptionId := "subscriptionId_example" // string | Identifier of the subscription resource
    suppFeat := "suppFeat_example" // string | Features supported by the NF service consumer (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualAnalyticsExposureSubscriptionApi.ReadAnSubscription(context.Background(), afId, subscriptionId).SuppFeat(suppFeat).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualAnalyticsExposureSubscriptionApi.ReadAnSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAnSubscription`: AnalyticsExposureSubsc
    fmt.Fprintf(os.Stdout, "Response from `IndividualAnalyticsExposureSubscriptionApi.ReadAnSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**afId** | **string** | Identifier of the AF | 
**subscriptionId** | **string** | Identifier of the subscription resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAnSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **suppFeat** | **string** | Features supported by the NF service consumer | 

### Return type

[**AnalyticsExposureSubsc**](AnalyticsExposureSubsc.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

