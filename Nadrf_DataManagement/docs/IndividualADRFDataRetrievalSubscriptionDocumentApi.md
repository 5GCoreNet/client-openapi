# \IndividualADRFDataRetrievalSubscriptionDocumentApi

All URIs are relative to *https://example.com/nadrf-datamanagement/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteADRFDataRetrievalSubscription**](IndividualADRFDataRetrievalSubscriptionDocumentApi.md#DeleteADRFDataRetrievalSubscription) | **Delete** /data-retrieval-subscriptions/{subscriptionId} | Delete an existing Individual ADRF Data Retrieval Subscription resource.



## DeleteADRFDataRetrievalSubscription

> DeleteADRFDataRetrievalSubscription(ctx, subscriptionId).Execute()

Delete an existing Individual ADRF Data Retrieval Subscription resource.

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
    subscriptionId := "subscriptionId_example" // string | String identifying a data retrieval subscription to the Nadrf_DataManagement  Service. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualADRFDataRetrievalSubscriptionDocumentApi.DeleteADRFDataRetrievalSubscription(context.Background(), subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualADRFDataRetrievalSubscriptionDocumentApi.DeleteADRFDataRetrievalSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | String identifying a data retrieval subscription to the Nadrf_DataManagement  Service.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteADRFDataRetrievalSubscriptionRequest struct via the builder pattern


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

