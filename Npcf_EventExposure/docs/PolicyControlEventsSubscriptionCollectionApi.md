# \PolicyControlEventsSubscriptionCollectionApi

All URIs are relative to *https://example.com/npcf-eventexposure/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostPcEventExposureSubsc**](PolicyControlEventsSubscriptionCollectionApi.md#PostPcEventExposureSubsc) | **Post** /subscriptions | Creates a new Individual Policy Control Events Subscription resource



## PostPcEventExposureSubsc

> PcEventExposureSubsc PostPcEventExposureSubsc(ctx).PcEventExposureSubsc(pcEventExposureSubsc).Execute()

Creates a new Individual Policy Control Events Subscription resource

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Npcf_EventExposure"
)

func main() {
    pcEventExposureSubsc := *openapiclient.NewPcEventExposureSubsc([]openapiclient.PcEvent{*openapiclient.NewPcEvent()}, "NotifUri_example", "NotifId_example") // PcEventExposureSubsc | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyControlEventsSubscriptionCollectionApi.PostPcEventExposureSubsc(context.Background()).PcEventExposureSubsc(pcEventExposureSubsc).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyControlEventsSubscriptionCollectionApi.PostPcEventExposureSubsc``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostPcEventExposureSubsc`: PcEventExposureSubsc
    fmt.Fprintf(os.Stdout, "Response from `PolicyControlEventsSubscriptionCollectionApi.PostPcEventExposureSubsc`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostPcEventExposureSubscRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pcEventExposureSubsc** | [**PcEventExposureSubsc**](PcEventExposureSubsc.md) |  | 

### Return type

[**PcEventExposureSubsc**](PcEventExposureSubsc.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

