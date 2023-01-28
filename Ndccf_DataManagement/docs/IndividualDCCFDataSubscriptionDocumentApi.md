# \IndividualDCCFDataSubscriptionDocumentApi

All URIs are relative to *https://example.com/ndccf-datamanagement/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteDCCFDataSubscription**](IndividualDCCFDataSubscriptionDocumentApi.md#DeleteDCCFDataSubscription) | **Delete** /data-subscriptions/{subscriptionId} | Deletes an existing Individual DCCF Data Subscription resource.
[**UpdateDCCFDataSubscription**](IndividualDCCFDataSubscriptionDocumentApi.md#UpdateDCCFDataSubscription) | **Put** /data-subscriptions/{subscriptionId} | Updates an existing Individual DCCF Data Subscription resource.



## DeleteDCCFDataSubscription

> DeleteDCCFDataSubscription(ctx, subscriptionId).Execute()

Deletes an existing Individual DCCF Data Subscription resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Ndccf_DataManagement"
)

func main() {
    subscriptionId := "subscriptionId_example" // string | String identifying a data subscription to the Ndccf_DataManagement Service.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualDCCFDataSubscriptionDocumentApi.DeleteDCCFDataSubscription(context.Background(), subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualDCCFDataSubscriptionDocumentApi.DeleteDCCFDataSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | String identifying a data subscription to the Ndccf_DataManagement Service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDCCFDataSubscriptionRequest struct via the builder pattern


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


## UpdateDCCFDataSubscription

> NdccfDataSubscription UpdateDCCFDataSubscription(ctx, subscriptionId).NdccfDataSubscription(ndccfDataSubscription).Execute()

Updates an existing Individual DCCF Data Subscription resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Ndccf_DataManagement"
)

func main() {
    subscriptionId := "subscriptionId_example" // string | String identifying a data subscription to the Ndccf_DataManagement Service. 
    ndccfDataSubscription := *openapiclient.NewNdccfDataSubscription(openapiclient.DataSubscription{Interface{}: new(interface{})}, "DataNotifUri_example", "DataNotifCorrId_example") // NdccfDataSubscription | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualDCCFDataSubscriptionDocumentApi.UpdateDCCFDataSubscription(context.Background(), subscriptionId).NdccfDataSubscription(ndccfDataSubscription).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualDCCFDataSubscriptionDocumentApi.UpdateDCCFDataSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDCCFDataSubscription`: NdccfDataSubscription
    fmt.Fprintf(os.Stdout, "Response from `IndividualDCCFDataSubscriptionDocumentApi.UpdateDCCFDataSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | String identifying a data subscription to the Ndccf_DataManagement Service.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDCCFDataSubscriptionRequest struct via the builder pattern


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

