# \IndividualNWDAFDataManagementSubscriptionDocumentApi

All URIs are relative to *https://example.com/nnwdaf-datamanagement/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteNWDAFDataSubscription**](IndividualNWDAFDataManagementSubscriptionDocumentApi.md#DeleteNWDAFDataSubscription) | **Delete** /subscriptions/{subscriptionId} | unsubscribe from notifications
[**UpdateNWDAFDataSubscription**](IndividualNWDAFDataManagementSubscriptionDocumentApi.md#UpdateNWDAFDataSubscription) | **Put** /subscriptions/{subscriptionId} | Update an existing Individual NWDAF Data Subscription.



## DeleteNWDAFDataSubscription

> DeleteNWDAFDataSubscription(ctx, subscriptionId).Execute()

unsubscribe from notifications

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
    subscriptionId := "subscriptionId_example" // string | Event Subscription ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualNWDAFDataManagementSubscriptionDocumentApi.DeleteNWDAFDataSubscription(context.Background(), subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualNWDAFDataManagementSubscriptionDocumentApi.DeleteNWDAFDataSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Event Subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNWDAFDataSubscriptionRequest struct via the builder pattern


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


## UpdateNWDAFDataSubscription

> NnwdafDataManagementSubsc UpdateNWDAFDataSubscription(ctx, subscriptionId).NnwdafDataManagementSubsc(nnwdafDataManagementSubsc).Execute()

Update an existing Individual NWDAF Data Subscription.

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
    subscriptionId := "subscriptionId_example" // string | Event Subscription ID
    nnwdafDataManagementSubsc := openapiclient.NnwdafDataManagementSubsc{Interface{}: new(interface{})} // NnwdafDataManagementSubsc | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualNWDAFDataManagementSubscriptionDocumentApi.UpdateNWDAFDataSubscription(context.Background(), subscriptionId).NnwdafDataManagementSubsc(nnwdafDataManagementSubsc).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualNWDAFDataManagementSubscriptionDocumentApi.UpdateNWDAFDataSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNWDAFDataSubscription`: NnwdafDataManagementSubsc
    fmt.Fprintf(os.Stdout, "Response from `IndividualNWDAFDataManagementSubscriptionDocumentApi.UpdateNWDAFDataSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Event Subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNWDAFDataSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nnwdafDataManagementSubsc** | [**NnwdafDataManagementSubsc**](NnwdafDataManagementSubsc.md) |  | 

### Return type

[**NnwdafDataManagementSubsc**](NnwdafDataManagementSubsc.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

