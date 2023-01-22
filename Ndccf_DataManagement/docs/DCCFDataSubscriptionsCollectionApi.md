# \DCCFDataSubscriptionsCollectionApi

All URIs are relative to *https://example.com/ndccf-datamanagement/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDCCFDataSubscription**](DCCFDataSubscriptionsCollectionApi.md#CreateDCCFDataSubscription) | **Post** /data-subscriptions | Creates a new Individual DCCF Data Subscription resource.



## CreateDCCFDataSubscription

> NdccfDataSubscription CreateDCCFDataSubscription(ctx).NdccfDataSubscription(ndccfDataSubscription).Execute()

Creates a new Individual DCCF Data Subscription resource.

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
    ndccfDataSubscription := *openapiclient.NewNdccfDataSubscription(openapiclient.DataSubscription{Interface{}: new(interface{})}, "DataNotifUri_example", "DataNotifCorrId_example") // NdccfDataSubscription | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DCCFDataSubscriptionsCollectionApi.CreateDCCFDataSubscription(context.Background()).NdccfDataSubscription(ndccfDataSubscription).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DCCFDataSubscriptionsCollectionApi.CreateDCCFDataSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDCCFDataSubscription`: NdccfDataSubscription
    fmt.Fprintf(os.Stdout, "Response from `DCCFDataSubscriptionsCollectionApi.CreateDCCFDataSubscription`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDCCFDataSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ndccfDataSubscription** | [**NdccfDataSubscription**](NdccfDataSubscription.md) |  | 

### Return type

[**NdccfDataSubscription**](NdccfDataSubscription.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

