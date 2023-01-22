# \ApplicationEventSubscriptionCollectionApi

All URIs are relative to *https://example.com/naf-eventexposure/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostAfEventExposureSubsc**](ApplicationEventSubscriptionCollectionApi.md#PostAfEventExposureSubsc) | **Post** /subscriptions | Creates a new Individual Application Event Exposure Subscription resource



## PostAfEventExposureSubsc

> AfEventExposureSubsc PostAfEventExposureSubsc(ctx).AfEventExposureSubsc(afEventExposureSubsc).Execute()

Creates a new Individual Application Event Exposure Subscription resource

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
    afEventExposureSubsc := *openapiclient.NewAfEventExposureSubsc([]openapiclient.EventsSubs{*openapiclient.NewEventsSubs(*openapiclient.NewAfEvent(), *openapiclient.NewEventFilter())}, *openapiclient.NewReportingInformation(), "NotifUri_example", "NotifId_example") // AfEventExposureSubsc | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationEventSubscriptionCollectionApi.PostAfEventExposureSubsc(context.Background()).AfEventExposureSubsc(afEventExposureSubsc).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationEventSubscriptionCollectionApi.PostAfEventExposureSubsc``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostAfEventExposureSubsc`: AfEventExposureSubsc
    fmt.Fprintf(os.Stdout, "Response from `ApplicationEventSubscriptionCollectionApi.PostAfEventExposureSubsc`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAfEventExposureSubscRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **afEventExposureSubsc** | [**AfEventExposureSubsc**](AfEventExposureSubsc.md) |  | 

### Return type

[**AfEventExposureSubsc**](AfEventExposureSubsc.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

