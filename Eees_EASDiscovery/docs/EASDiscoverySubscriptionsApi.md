# \EASDiscoverySubscriptionsApi

All URIs are relative to *https://example.com/eees-easdiscovery/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubscriptionsPost**](EASDiscoverySubscriptionsApi.md#SubscriptionsPost) | **Post** /subscriptions | 



## SubscriptionsPost

> EasDiscoverySubscription SubscriptionsPost(ctx).EasDiscoverySubscription(easDiscoverySubscription).Execute()





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
    easDiscoverySubscription := *openapiclient.NewEasDiscoverySubscription("EecId_example", *openapiclient.NewEASDiscEventIDs()) // EasDiscoverySubscription | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EASDiscoverySubscriptionsApi.SubscriptionsPost(context.Background()).EasDiscoverySubscription(easDiscoverySubscription).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EASDiscoverySubscriptionsApi.SubscriptionsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubscriptionsPost`: EasDiscoverySubscription
    fmt.Fprintf(os.Stdout, "Response from `EASDiscoverySubscriptionsApi.SubscriptionsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **easDiscoverySubscription** | [**EasDiscoverySubscription**](EasDiscoverySubscription.md) |  | 

### Return type

[**EasDiscoverySubscription**](EasDiscoverySubscription.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

