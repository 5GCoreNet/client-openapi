# \IndividualACTStatusSubscriptionDocumentApi

All URIs are relative to *https://example.com/eees-eel-acr/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetACTStatusSubscription**](IndividualACTStatusSubscriptionDocumentApi.md#GetACTStatusSubscription) | **Get** /subscriptions/{subscriptionId} | Retrieve an ACT status subscription resource.



## GetACTStatusSubscription

> ACTStatusSubsc GetACTStatusSubscription(ctx, subscriptionId).Execute()

Retrieve an ACT status subscription resource.

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
    subscriptionId := "subscriptionId_example" // string | Individual ACT Status Subscription identifier.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualACTStatusSubscriptionDocumentApi.GetACTStatusSubscription(context.Background(), subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualACTStatusSubscriptionDocumentApi.GetACTStatusSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetACTStatusSubscription`: ACTStatusSubsc
    fmt.Fprintf(os.Stdout, "Response from `IndividualACTStatusSubscriptionDocumentApi.GetACTStatusSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Individual ACT Status Subscription identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetACTStatusSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ACTStatusSubsc**](ACTStatusSubsc.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

