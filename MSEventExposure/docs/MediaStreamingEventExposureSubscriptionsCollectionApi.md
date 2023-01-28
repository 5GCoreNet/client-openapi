# \MediaStreamingEventExposureSubscriptionsCollectionApi

All URIs are relative to *https://example.com/3gpp-ms-event-exposure/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMSEventExposureSubsc**](MediaStreamingEventExposureSubscriptionsCollectionApi.md#CreateMSEventExposureSubsc) | **Post** /subscriptions | Request the creation of a new Individual Media Streaming Event Exposure Subscription resource.
[**RetrieveMSEventExposureSubscs**](MediaStreamingEventExposureSubscriptionsCollectionApi.md#RetrieveMSEventExposureSubscs) | **Get** /subscriptions | Retrieve all the active Media Streaming Event Exposure Subscription resources managed by the NEF.



## CreateMSEventExposureSubsc

> AfEventExposureSubsc CreateMSEventExposureSubsc(ctx).AfEventExposureSubsc(afEventExposureSubsc).Execute()

Request the creation of a new Individual Media Streaming Event Exposure Subscription resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/MSEventExposure"
)

func main() {
    afEventExposureSubsc := *openapiclient.NewAfEventExposureSubsc([]openapiclient.EventsSubs{*openapiclient.NewEventsSubs(*openapiclient.NewAfEvent(), *openapiclient.NewEventFilter())}, *openapiclient.NewReportingInformation(), "NotifUri_example", "NotifId_example") // AfEventExposureSubsc | Contains the parameters to request the creation of a new Media Streaming Event Exposure  Subscriptionat the NEF. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MediaStreamingEventExposureSubscriptionsCollectionApi.CreateMSEventExposureSubsc(context.Background()).AfEventExposureSubsc(afEventExposureSubsc).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaStreamingEventExposureSubscriptionsCollectionApi.CreateMSEventExposureSubsc``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMSEventExposureSubsc`: AfEventExposureSubsc
    fmt.Fprintf(os.Stdout, "Response from `MediaStreamingEventExposureSubscriptionsCollectionApi.CreateMSEventExposureSubsc`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMSEventExposureSubscRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **afEventExposureSubsc** | [**AfEventExposureSubsc**](AfEventExposureSubsc.md) | Contains the parameters to request the creation of a new Media Streaming Event Exposure  Subscriptionat the NEF.  | 

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


## RetrieveMSEventExposureSubscs

> []AfEventExposureSubsc RetrieveMSEventExposureSubscs(ctx).Execute()

Retrieve all the active Media Streaming Event Exposure Subscription resources managed by the NEF.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/MSEventExposure"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MediaStreamingEventExposureSubscriptionsCollectionApi.RetrieveMSEventExposureSubscs(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaStreamingEventExposureSubscriptionsCollectionApi.RetrieveMSEventExposureSubscs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveMSEventExposureSubscs`: []AfEventExposureSubsc
    fmt.Fprintf(os.Stdout, "Response from `MediaStreamingEventExposureSubscriptionsCollectionApi.RetrieveMSEventExposureSubscs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveMSEventExposureSubscsRequest struct via the builder pattern


### Return type

[**[]AfEventExposureSubsc**](AfEventExposureSubsc.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

