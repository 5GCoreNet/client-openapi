# \SubscriptionsCollectionApi

All URIs are relative to *https://example.com/nnef-eventexposure/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIndividualSubcription**](SubscriptionsCollectionApi.md#CreateIndividualSubcription) | **Post** /subscriptions | subscribe to notifications



## CreateIndividualSubcription

> NefEventExposureSubsc CreateIndividualSubcription(ctx).NefEventExposureSubsc(nefEventExposureSubsc).Execute()

subscribe to notifications

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nnef_EventExposure"
)

func main() {
    nefEventExposureSubsc := *openapiclient.NewNefEventExposureSubsc([]openapiclient.NefEventSubs{*openapiclient.NewNefEventSubs(*openapiclient.NewNefEvent())}, "NotifUri_example", "NotifId_example") // NefEventExposureSubsc | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionsCollectionApi.CreateIndividualSubcription(context.Background()).NefEventExposureSubsc(nefEventExposureSubsc).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsCollectionApi.CreateIndividualSubcription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIndividualSubcription`: NefEventExposureSubsc
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionsCollectionApi.CreateIndividualSubcription`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIndividualSubcriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nefEventExposureSubsc** | [**NefEventExposureSubsc**](NefEventExposureSubsc.md) |  | 

### Return type

[**NefEventExposureSubsc**](NefEventExposureSubsc.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

