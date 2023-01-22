# \ServiceProvisioningSubscriptionsApi

All URIs are relative to *https://example.com/eecs-serviceprovisioning/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubscriptionsPost**](ServiceProvisioningSubscriptionsApi.md#SubscriptionsPost) | **Post** /subscriptions | 



## SubscriptionsPost

> ECSServProvSubscription SubscriptionsPost(ctx).ECSServProvSubscription(eCSServProvSubscription).Execute()





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
    eCSServProvSubscription := *openapiclient.NewECSServProvSubscription("EecId_example") // ECSServProvSubscription | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceProvisioningSubscriptionsApi.SubscriptionsPost(context.Background()).ECSServProvSubscription(eCSServProvSubscription).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceProvisioningSubscriptionsApi.SubscriptionsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubscriptionsPost`: ECSServProvSubscription
    fmt.Fprintf(os.Stdout, "Response from `ServiceProvisioningSubscriptionsApi.SubscriptionsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eCSServProvSubscription** | [**ECSServProvSubscription**](ECSServProvSubscription.md) |  | 

### Return type

[**ECSServProvSubscription**](ECSServProvSubscription.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

