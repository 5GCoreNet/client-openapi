# \UnicastMonitoringSubscriptionsCollectionApi

All URIs are relative to *https://example.com/ss-nrm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubscribeUnicastMonitoring**](UnicastMonitoringSubscriptionsCollectionApi.md#SubscribeUnicastMonitoring) | **Post** /subscriptions | Create individual unicast monitoring subscription resource or obtain unicast QoS monitoring data for VAL UEs, VAL Group, or VAL Streams.



## SubscribeUnicastMonitoring

> MonitoringReport SubscribeUnicastMonitoring(ctx).MonitoringSubscription(monitoringSubscription).Execute()

Create individual unicast monitoring subscription resource or obtain unicast QoS monitoring data for VAL UEs, VAL Group, or VAL Streams.

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
    monitoringSubscription := openapiclient.MonitoringSubscription{Interface{}: new(interface{})} // MonitoringSubscription | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UnicastMonitoringSubscriptionsCollectionApi.SubscribeUnicastMonitoring(context.Background()).MonitoringSubscription(monitoringSubscription).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UnicastMonitoringSubscriptionsCollectionApi.SubscribeUnicastMonitoring``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubscribeUnicastMonitoring`: MonitoringReport
    fmt.Fprintf(os.Stdout, "Response from `UnicastMonitoringSubscriptionsCollectionApi.SubscribeUnicastMonitoring`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscribeUnicastMonitoringRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **monitoringSubscription** | [**MonitoringSubscription**](MonitoringSubscription.md) |  | 

### Return type

[**MonitoringReport**](MonitoringReport.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

