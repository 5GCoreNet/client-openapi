# \ADRFStorageSubscriptionsApi

All URIs are relative to *https://example.com/nadrf-datamanagement/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateADRFStorageSubscription**](ADRFStorageSubscriptionsApi.md#CreateADRFStorageSubscription) | **Post** /request-storage-sub | Triggers the creation of a new ADRF Storage Subscription.
[**DeleteADRFStorageSubscription**](ADRFStorageSubscriptionsApi.md#DeleteADRFStorageSubscription) | **Post** /request-storage-sub-removal | Triggers the removal of ADRF storage subscription.



## CreateADRFStorageSubscription

> NadrfDataStoreSubscriptionRef CreateADRFStorageSubscription(ctx).NadrfDataStoreSubscription(nadrfDataStoreSubscription).Execute()

Triggers the creation of a new ADRF Storage Subscription.

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
    nadrfDataStoreSubscription := *openapiclient.NewNadrfDataStoreSubscription() // NadrfDataStoreSubscription | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ADRFStorageSubscriptionsApi.CreateADRFStorageSubscription(context.Background()).NadrfDataStoreSubscription(nadrfDataStoreSubscription).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ADRFStorageSubscriptionsApi.CreateADRFStorageSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateADRFStorageSubscription`: NadrfDataStoreSubscriptionRef
    fmt.Fprintf(os.Stdout, "Response from `ADRFStorageSubscriptionsApi.CreateADRFStorageSubscription`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateADRFStorageSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nadrfDataStoreSubscription** | [**NadrfDataStoreSubscription**](NadrfDataStoreSubscription.md) |  | 

### Return type

[**NadrfDataStoreSubscriptionRef**](NadrfDataStoreSubscriptionRef.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteADRFStorageSubscription

> DeleteADRFStorageSubscription(ctx).NadrfDataStoreSubscriptionRef(nadrfDataStoreSubscriptionRef).Execute()

Triggers the removal of ADRF storage subscription.

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
    nadrfDataStoreSubscriptionRef := *openapiclient.NewNadrfDataStoreSubscriptionRef("TransRefId_example") // NadrfDataStoreSubscriptionRef | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ADRFStorageSubscriptionsApi.DeleteADRFStorageSubscription(context.Background()).NadrfDataStoreSubscriptionRef(nadrfDataStoreSubscriptionRef).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ADRFStorageSubscriptionsApi.DeleteADRFStorageSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteADRFStorageSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nadrfDataStoreSubscriptionRef** | [**NadrfDataStoreSubscriptionRef**](NadrfDataStoreSubscriptionRef.md) |  | 

### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

