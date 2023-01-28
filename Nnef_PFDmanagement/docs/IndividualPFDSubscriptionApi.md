# \IndividualPFDSubscriptionApi

All URIs are relative to *https://example.com/nnef-pfdmanagement/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NnefPFDmanagementModifySubscr**](IndividualPFDSubscriptionApi.md#NnefPFDmanagementModifySubscr) | **Put** /subscriptions/{subscriptionId} | Updates/replaces an existing subscription resource
[**NnefPFDmanagementUnsubscribe**](IndividualPFDSubscriptionApi.md#NnefPFDmanagementUnsubscribe) | **Delete** /subscriptions/{subscriptionId} | Delete a subscription of PFD change notification.



## NnefPFDmanagementModifySubscr

> PfdSubscription NnefPFDmanagementModifySubscr(ctx, subscriptionId).PfdSubscription(pfdSubscription).Execute()

Updates/replaces an existing subscription resource

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nnef_PFDmanagement"
)

func main() {
    subscriptionId := "subscriptionId_example" // string | Identify the subscription.
    pfdSubscription := *openapiclient.NewPfdSubscription("NotifyUri_example", "SupportedFeatures_example") // PfdSubscription | Parameters to update/replace the existing subscription

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualPFDSubscriptionApi.NnefPFDmanagementModifySubscr(context.Background(), subscriptionId).PfdSubscription(pfdSubscription).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualPFDSubscriptionApi.NnefPFDmanagementModifySubscr``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NnefPFDmanagementModifySubscr`: PfdSubscription
    fmt.Fprintf(os.Stdout, "Response from `IndividualPFDSubscriptionApi.NnefPFDmanagementModifySubscr`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Identify the subscription. | 

### Other Parameters

Other parameters are passed through a pointer to a apiNnefPFDmanagementModifySubscrRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pfdSubscription** | [**PfdSubscription**](PfdSubscription.md) | Parameters to update/replace the existing subscription | 

### Return type

[**PfdSubscription**](PfdSubscription.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NnefPFDmanagementUnsubscribe

> NnefPFDmanagementUnsubscribe(ctx, subscriptionId).Execute()

Delete a subscription of PFD change notification.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nnef_PFDmanagement"
)

func main() {
    subscriptionId := "subscriptionId_example" // string | Identify the subscription.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualPFDSubscriptionApi.NnefPFDmanagementUnsubscribe(context.Background(), subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualPFDSubscriptionApi.NnefPFDmanagementUnsubscribe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Identify the subscription. | 

### Other Parameters

Other parameters are passed through a pointer to a apiNnefPFDmanagementUnsubscribeRequest struct via the builder pattern


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

