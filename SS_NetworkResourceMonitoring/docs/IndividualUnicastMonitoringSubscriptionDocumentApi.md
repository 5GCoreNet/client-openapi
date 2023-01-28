# \IndividualUnicastMonitoringSubscriptionDocumentApi

All URIs are relative to *https://example.com/ss-nrm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReadUnicastMonitoringSubscription**](IndividualUnicastMonitoringSubscriptionDocumentApi.md#ReadUnicastMonitoringSubscription) | **Get** /subscriptions/{subscriptionId} | Read an existing individual unicast monitoring subscription resource according to the subscriptionId.
[**UnsubscribeUnicastMonitoring**](IndividualUnicastMonitoringSubscriptionDocumentApi.md#UnsubscribeUnicastMonitoring) | **Delete** /subscriptions/{subscriptionId} | Remove an existing individual unicast monitoring subscription resource according to the subscriptionId.



## ReadUnicastMonitoringSubscription

> MonitoringSubscription ReadUnicastMonitoringSubscription(ctx, subscriptionId).Execute()

Read an existing individual unicast monitoring subscription resource according to the subscriptionId.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/SS_NetworkResourceMonitoring"
)

func main() {
    subscriptionId := "subscriptionId_example" // string | Represents the identifier of an individual unicast monitoring subscription resource. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualUnicastMonitoringSubscriptionDocumentApi.ReadUnicastMonitoringSubscription(context.Background(), subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualUnicastMonitoringSubscriptionDocumentApi.ReadUnicastMonitoringSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadUnicastMonitoringSubscription`: MonitoringSubscription
    fmt.Fprintf(os.Stdout, "Response from `IndividualUnicastMonitoringSubscriptionDocumentApi.ReadUnicastMonitoringSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Represents the identifier of an individual unicast monitoring subscription resource.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadUnicastMonitoringSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MonitoringSubscription**](MonitoringSubscription.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnsubscribeUnicastMonitoring

> UnsubscribeUnicastMonitoring(ctx, subscriptionId).Execute()

Remove an existing individual unicast monitoring subscription resource according to the subscriptionId.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/SS_NetworkResourceMonitoring"
)

func main() {
    subscriptionId := "subscriptionId_example" // string | Represents the identifier of an individual unicast monitoring subscription resource. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualUnicastMonitoringSubscriptionDocumentApi.UnsubscribeUnicastMonitoring(context.Background(), subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualUnicastMonitoringSubscriptionDocumentApi.UnsubscribeUnicastMonitoring``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Represents the identifier of an individual unicast monitoring subscription resource.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnsubscribeUnicastMonitoringRequest struct via the builder pattern


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

