# \SubscriptionsCollectionCollectionApi

All URIs are relative to *https://example.com/namf-comm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AMFStatusChangeSubscribe**](SubscriptionsCollectionCollectionApi.md#AMFStatusChangeSubscribe) | **Post** /subscriptions | Namf_Communication AMF Status Change Subscribe service Operation



## AMFStatusChangeSubscribe

> SubscriptionData AMFStatusChangeSubscribe(ctx).SubscriptionData(subscriptionData).Execute()

Namf_Communication AMF Status Change Subscribe service Operation

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Namf_Communication"
)

func main() {
    subscriptionData := *openapiclient.NewSubscriptionData("AmfStatusUri_example") // SubscriptionData | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionsCollectionCollectionApi.AMFStatusChangeSubscribe(context.Background()).SubscriptionData(subscriptionData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsCollectionCollectionApi.AMFStatusChangeSubscribe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AMFStatusChangeSubscribe`: SubscriptionData
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionsCollectionCollectionApi.AMFStatusChangeSubscribe`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAMFStatusChangeSubscribeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriptionData** | [**SubscriptionData**](SubscriptionData.md) |  | 

### Return type

[**SubscriptionData**](SubscriptionData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

