# \IndividualSubscriptionDocumentApi

All URIs are relative to *https://example.com/nbsf-management/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIndividualSubcription**](IndividualSubscriptionDocumentApi.md#DeleteIndividualSubcription) | **Delete** /subscriptions/{subId} | Delete an individual subscription for event notifications from the BSF
[**ReplaceIndividualSubcription**](IndividualSubscriptionDocumentApi.md#ReplaceIndividualSubcription) | **Put** /subscriptions/{subId} | Replace an individual subscription for event notifications from the BSF



## DeleteIndividualSubcription

> DeleteIndividualSubcription(ctx, subId).Execute()

Delete an individual subscription for event notifications from the BSF

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
    subId := "subId_example" // string | Subscription correlation ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualSubscriptionDocumentApi.DeleteIndividualSubcription(context.Background(), subId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualSubscriptionDocumentApi.DeleteIndividualSubcription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subId** | **string** | Subscription correlation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIndividualSubcriptionRequest struct via the builder pattern


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


## ReplaceIndividualSubcription

> BsfSubscriptionResp ReplaceIndividualSubcription(ctx, subId).BsfSubscription(bsfSubscription).Execute()

Replace an individual subscription for event notifications from the BSF

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
    subId := "subId_example" // string | Subscription correlation ID
    bsfSubscription := *openapiclient.NewBsfSubscription([]openapiclient.BsfEvent{*openapiclient.NewBsfEvent()}, "NotifUri_example", "NotifCorreId_example", "Supi_example") // BsfSubscription | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualSubscriptionDocumentApi.ReplaceIndividualSubcription(context.Background(), subId).BsfSubscription(bsfSubscription).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualSubscriptionDocumentApi.ReplaceIndividualSubcription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceIndividualSubcription`: BsfSubscriptionResp
    fmt.Fprintf(os.Stdout, "Response from `IndividualSubscriptionDocumentApi.ReplaceIndividualSubcription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subId** | **string** | Subscription correlation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceIndividualSubcriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bsfSubscription** | [**BsfSubscription**](BsfSubscription.md) |  | 

### Return type

[**BsfSubscriptionResp**](BsfSubscriptionResp.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

