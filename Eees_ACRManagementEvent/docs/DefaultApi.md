# \DefaultApi

All URIs are relative to *https://example.com/eees-acrmgntevent/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubscriptionsGet**](DefaultApi.md#SubscriptionsGet) | **Get** /subscriptions | 
[**SubscriptionsPost**](DefaultApi.md#SubscriptionsPost) | **Post** /subscriptions | 
[**SubscriptionsSubscriptionIdDelete**](DefaultApi.md#SubscriptionsSubscriptionIdDelete) | **Delete** /subscriptions/{subscriptionId} | 
[**SubscriptionsSubscriptionIdGet**](DefaultApi.md#SubscriptionsSubscriptionIdGet) | **Get** /subscriptions/{subscriptionId} | 
[**SubscriptionsSubscriptionIdPatch**](DefaultApi.md#SubscriptionsSubscriptionIdPatch) | **Patch** /subscriptions/{subscriptionId} | 
[**SubscriptionsSubscriptionIdPut**](DefaultApi.md#SubscriptionsSubscriptionIdPut) | **Put** /subscriptions/{subscriptionId} | 



## SubscriptionsGet

> []AcrMgntEventsSubscription SubscriptionsGet(ctx).SuppFeat(suppFeat).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Eees_ACRManagementEvent"
)

func main() {
    suppFeat := "suppFeat_example" // string | Features supported by the EAS. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.SubscriptionsGet(context.Background()).SuppFeat(suppFeat).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SubscriptionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubscriptionsGet`: []AcrMgntEventsSubscription
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SubscriptionsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **suppFeat** | **string** | Features supported by the EAS. | 

### Return type

[**[]AcrMgntEventsSubscription**](AcrMgntEventsSubscription.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsPost

> AcrMgntEventsSubscription SubscriptionsPost(ctx).AcrMgntEventsSubscription(acrMgntEventsSubscription).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Eees_ACRManagementEvent"
)

func main() {
    acrMgntEventsSubscription := *openapiclient.NewAcrMgntEventsSubscription("EasId_example", []openapiclient.AcrMgntEventSubsc{*openapiclient.NewAcrMgntEventSubsc(*openapiclient.NewAcrMgntEvent())}, "NotificationDestination_example") // AcrMgntEventsSubscription | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.SubscriptionsPost(context.Background()).AcrMgntEventsSubscription(acrMgntEventsSubscription).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SubscriptionsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubscriptionsPost`: AcrMgntEventsSubscription
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SubscriptionsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acrMgntEventsSubscription** | [**AcrMgntEventsSubscription**](AcrMgntEventsSubscription.md) |  | 

### Return type

[**AcrMgntEventsSubscription**](AcrMgntEventsSubscription.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsSubscriptionIdDelete

> SubscriptionsSubscriptionIdDelete(ctx, subscriptionId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Eees_ACRManagementEvent"
)

func main() {
    subscriptionId := "subscriptionId_example" // string | Subscription Id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.SubscriptionsSubscriptionIdDelete(context.Background(), subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SubscriptionsSubscriptionIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Subscription Id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsSubscriptionIdDeleteRequest struct via the builder pattern


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


## SubscriptionsSubscriptionIdGet

> AcrMgntEventsSubscription SubscriptionsSubscriptionIdGet(ctx, subscriptionId).SuppFeat(suppFeat).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Eees_ACRManagementEvent"
)

func main() {
    subscriptionId := "subscriptionId_example" // string | Subscription Id.
    suppFeat := "suppFeat_example" // string | Features supported by the EAS. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.SubscriptionsSubscriptionIdGet(context.Background(), subscriptionId).SuppFeat(suppFeat).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SubscriptionsSubscriptionIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubscriptionsSubscriptionIdGet`: AcrMgntEventsSubscription
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SubscriptionsSubscriptionIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Subscription Id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsSubscriptionIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **suppFeat** | **string** | Features supported by the EAS. | 

### Return type

[**AcrMgntEventsSubscription**](AcrMgntEventsSubscription.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsSubscriptionIdPatch

> AcrMgntEventsSubscription SubscriptionsSubscriptionIdPatch(ctx, subscriptionId).AcrMgntEventsSubscriptionPatch(acrMgntEventsSubscriptionPatch).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Eees_ACRManagementEvent"
)

func main() {
    subscriptionId := "subscriptionId_example" // string | Subscription Id.
    acrMgntEventsSubscriptionPatch := *openapiclient.NewAcrMgntEventsSubscriptionPatch() // AcrMgntEventsSubscriptionPatch | Partial update an existing Individual ACR Management Events Subscription.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.SubscriptionsSubscriptionIdPatch(context.Background(), subscriptionId).AcrMgntEventsSubscriptionPatch(acrMgntEventsSubscriptionPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SubscriptionsSubscriptionIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubscriptionsSubscriptionIdPatch`: AcrMgntEventsSubscription
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SubscriptionsSubscriptionIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Subscription Id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsSubscriptionIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acrMgntEventsSubscriptionPatch** | [**AcrMgntEventsSubscriptionPatch**](AcrMgntEventsSubscriptionPatch.md) | Partial update an existing Individual ACR Management Events Subscription. | 

### Return type

[**AcrMgntEventsSubscription**](AcrMgntEventsSubscription.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsSubscriptionIdPut

> AcrMgntEventsSubscription SubscriptionsSubscriptionIdPut(ctx, subscriptionId).AcrMgntEventsSubscription(acrMgntEventsSubscription).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Eees_ACRManagementEvent"
)

func main() {
    subscriptionId := "subscriptionId_example" // string | Subscription Id.
    acrMgntEventsSubscription := *openapiclient.NewAcrMgntEventsSubscription("EasId_example", []openapiclient.AcrMgntEventSubsc{*openapiclient.NewAcrMgntEventSubsc(*openapiclient.NewAcrMgntEvent())}, "NotificationDestination_example") // AcrMgntEventsSubscription | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.SubscriptionsSubscriptionIdPut(context.Background(), subscriptionId).AcrMgntEventsSubscription(acrMgntEventsSubscription).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SubscriptionsSubscriptionIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubscriptionsSubscriptionIdPut`: AcrMgntEventsSubscription
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SubscriptionsSubscriptionIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Subscription Id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsSubscriptionIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acrMgntEventsSubscription** | [**AcrMgntEventsSubscription**](AcrMgntEventsSubscription.md) |  | 

### Return type

[**AcrMgntEventsSubscription**](AcrMgntEventsSubscription.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

