# \ACREventsSubscriptionCollectionApi

All URIs are relative to *https://example.com/eees-acrevents/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateACREventsSubscripton**](ACREventsSubscriptionCollectionApi.md#CreateACREventsSubscripton) | **Post** /subscriptions | 



## CreateACREventsSubscripton

> ACREventsSubscription CreateACREventsSubscripton(ctx).ACREventsSubscription(aCREventsSubscription).Execute()





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
    aCREventsSubscription := *openapiclient.NewACREventsSubscription("EecId_example", []string{"EasIds_example"}, *openapiclient.NewACREventIDs(), "NotificationDestination_example") // ACREventsSubscription | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ACREventsSubscriptionCollectionApi.CreateACREventsSubscripton(context.Background()).ACREventsSubscription(aCREventsSubscription).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ACREventsSubscriptionCollectionApi.CreateACREventsSubscripton``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateACREventsSubscripton`: ACREventsSubscription
    fmt.Fprintf(os.Stdout, "Response from `ACREventsSubscriptionCollectionApi.CreateACREventsSubscripton`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateACREventsSubscriptonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aCREventsSubscription** | [**ACREventsSubscription**](ACREventsSubscription.md) |  | 

### Return type

[**ACREventsSubscription**](ACREventsSubscription.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

