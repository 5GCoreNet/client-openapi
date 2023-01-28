# \ADRFDataRetrievalSubscriptionsCollectionApi

All URIs are relative to *https://example.com/nadrf-datamanagement/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateADRFDataRetrievalSubscription**](ADRFDataRetrievalSubscriptionsCollectionApi.md#CreateADRFDataRetrievalSubscription) | **Post** /data-retrieval-subscriptions | Creates a new Individual ADRF Data Retrieval Subscription resource.



## CreateADRFDataRetrievalSubscription

> NadrfDataRetrievalSubscription CreateADRFDataRetrievalSubscription(ctx).NadrfDataRetrievalSubscription(nadrfDataRetrievalSubscription).Execute()

Creates a new Individual ADRF Data Retrieval Subscription resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nadrf_DataManagement"
)

func main() {
    nadrfDataRetrievalSubscription := openapiclient.NadrfDataRetrievalSubscription{Interface{}: new(interface{})} // NadrfDataRetrievalSubscription | Individual ADRF Data Retrieval Subscription resource to be created.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ADRFDataRetrievalSubscriptionsCollectionApi.CreateADRFDataRetrievalSubscription(context.Background()).NadrfDataRetrievalSubscription(nadrfDataRetrievalSubscription).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ADRFDataRetrievalSubscriptionsCollectionApi.CreateADRFDataRetrievalSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateADRFDataRetrievalSubscription`: NadrfDataRetrievalSubscription
    fmt.Fprintf(os.Stdout, "Response from `ADRFDataRetrievalSubscriptionsCollectionApi.CreateADRFDataRetrievalSubscription`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateADRFDataRetrievalSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nadrfDataRetrievalSubscription** | [**NadrfDataRetrievalSubscription**](NadrfDataRetrievalSubscription.md) | Individual ADRF Data Retrieval Subscription resource to be created. | 

### Return type

[**NadrfDataRetrievalSubscription**](NadrfDataRetrievalSubscription.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

