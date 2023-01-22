# \IndividualSCPDomainRoutingInformationSubscriptionDocumentApi

All URIs are relative to *https://example.com/nnrf-disc/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ScpDomainRoutingInfoUnsubscribe**](IndividualSCPDomainRoutingInformationSubscriptionDocumentApi.md#ScpDomainRoutingInfoUnsubscribe) | **Delete** /scp-domain-routing-info-subs/{subscriptionID} | Deletes a subscription



## ScpDomainRoutingInfoUnsubscribe

> ScpDomainRoutingInfoUnsubscribe(ctx, subscriptionID).Execute()

Deletes a subscription

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
    subscriptionID := "subscriptionID_example" // string | Unique ID of the subscription to remove

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualSCPDomainRoutingInformationSubscriptionDocumentApi.ScpDomainRoutingInfoUnsubscribe(context.Background(), subscriptionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualSCPDomainRoutingInformationSubscriptionDocumentApi.ScpDomainRoutingInfoUnsubscribe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionID** | **string** | Unique ID of the subscription to remove | 

### Other Parameters

Other parameters are passed through a pointer to a apiScpDomainRoutingInfoUnsubscribeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

