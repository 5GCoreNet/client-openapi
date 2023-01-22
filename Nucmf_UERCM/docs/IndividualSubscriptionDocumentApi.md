# \IndividualSubscriptionDocumentApi

All URIs are relative to *https://example.com/nucmf-uecm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIndividualSubcription**](IndividualSubscriptionDocumentApi.md#DeleteIndividualSubcription) | **Delete** /subscriptions/{subscriptionId} | unsubscribe from notifications



## DeleteIndividualSubcription

> DeleteIndividualSubcription(ctx, subscriptionId).Execute()

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
    resp, r, err := apiClient.IndividualSubscriptionDocumentApi.DeleteIndividualSubcription(context.Background(), subscriptionId).Execute()
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
**subscriptionId** | **string** | Event Subscription ID | 

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

