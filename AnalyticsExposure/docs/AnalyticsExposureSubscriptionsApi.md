# \AnalyticsExposureSubscriptionsApi

All URIs are relative to *https://example.com/3gpp-analyticsexposure/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNewSubscription**](AnalyticsExposureSubscriptionsApi.md#CreateNewSubscription) | **Post** /{afId}/subscriptions | Creates a new subscription resource
[**ReadAllSubscriptions**](AnalyticsExposureSubscriptionsApi.md#ReadAllSubscriptions) | **Get** /{afId}/subscriptions | read all of the active subscriptions for the AF



## CreateNewSubscription

> AnalyticsExposureSubsc CreateNewSubscription(ctx, afId).AnalyticsExposureSubsc(analyticsExposureSubsc).Execute()

Creates a new subscription resource

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
    analyticsExposureSubsc := *openapiclient.NewAnalyticsExposureSubsc([]openapiclient.AnalyticsEventSubsc{*openapiclient.NewAnalyticsEventSubsc(*openapiclient.NewAnalyticsEvent())}, "NotifUri_example", "NotifId_example") // AnalyticsExposureSubsc | new subscription creation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AnalyticsExposureSubscriptionsApi.CreateNewSubscription(context.Background(), afId).AnalyticsExposureSubsc(analyticsExposureSubsc).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsExposureSubscriptionsApi.CreateNewSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNewSubscription`: AnalyticsExposureSubsc
    fmt.Fprintf(os.Stdout, "Response from `AnalyticsExposureSubscriptionsApi.CreateNewSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**afId** | **string** | Identifier of the AF | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNewSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **analyticsExposureSubsc** | [**AnalyticsExposureSubsc**](AnalyticsExposureSubsc.md) | new subscription creation | 

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


## ReadAllSubscriptions

> []AnalyticsExposureSubsc ReadAllSubscriptions(ctx, afId).SuppFeat(suppFeat).Execute()

read all of the active subscriptions for the AF

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
    suppFeat := "suppFeat_example" // string | Features supported by the NF service consumer (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AnalyticsExposureSubscriptionsApi.ReadAllSubscriptions(context.Background(), afId).SuppFeat(suppFeat).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsExposureSubscriptionsApi.ReadAllSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllSubscriptions`: []AnalyticsExposureSubsc
    fmt.Fprintf(os.Stdout, "Response from `AnalyticsExposureSubscriptionsApi.ReadAllSubscriptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**afId** | **string** | Identifier of the AF | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAllSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **suppFeat** | **string** | Features supported by the NF service consumer | 

### Return type

[**[]AnalyticsExposureSubsc**](AnalyticsExposureSubsc.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

