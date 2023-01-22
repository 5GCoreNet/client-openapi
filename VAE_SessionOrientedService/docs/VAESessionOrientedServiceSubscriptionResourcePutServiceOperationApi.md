# \VAESessionOrientedServiceSubscriptionResourcePutServiceOperationApi

All URIs are relative to *https://example.com/vae-session-Oriented-service/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubscriptionsSubscriptionIdPut**](VAESessionOrientedServiceSubscriptionResourcePutServiceOperationApi.md#SubscriptionsSubscriptionIdPut) | **Put** /subscriptions/{subscriptionId} | Updates/replaces an existing subscription resource



## SubscriptionsSubscriptionIdPut

> SessionOrientedData SubscriptionsSubscriptionIdPut(ctx, subscriptionId).SessionOrientedData(sessionOrientedData).Execute()

Updates/replaces an existing subscription resource

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
    subscriptionId := "subscriptionId_example" // string | Identifier of an Session Oriented Service Subscription resource
    sessionOrientedData := *openapiclient.NewSessionOrientedData("UeId_example", "NotifUri_example", "ServiceId_example", "AppSerId_example") // SessionOrientedData | Parameters to update/replace the existing subscription

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VAESessionOrientedServiceSubscriptionResourcePutServiceOperationApi.SubscriptionsSubscriptionIdPut(context.Background(), subscriptionId).SessionOrientedData(sessionOrientedData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VAESessionOrientedServiceSubscriptionResourcePutServiceOperationApi.SubscriptionsSubscriptionIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubscriptionsSubscriptionIdPut`: SessionOrientedData
    fmt.Fprintf(os.Stdout, "Response from `VAESessionOrientedServiceSubscriptionResourcePutServiceOperationApi.SubscriptionsSubscriptionIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Identifier of an Session Oriented Service Subscription resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsSubscriptionIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sessionOrientedData** | [**SessionOrientedData**](SessionOrientedData.md) | Parameters to update/replace the existing subscription | 

### Return type

[**SessionOrientedData**](SessionOrientedData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

