# \DefaultApi

All URIs are relative to *https://example.com/eees-appclientinformation/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubscriptionsPost**](DefaultApi.md#SubscriptionsPost) | **Post** /subscriptions | 
[**SubscriptionsSubscriptionIdDelete**](DefaultApi.md#SubscriptionsSubscriptionIdDelete) | **Delete** /subscriptions/{subscriptionId} | 
[**SubscriptionsSubscriptionIdGet**](DefaultApi.md#SubscriptionsSubscriptionIdGet) | **Get** /subscriptions/{subscriptionId} | 
[**SubscriptionsSubscriptionIdPatch**](DefaultApi.md#SubscriptionsSubscriptionIdPatch) | **Patch** /subscriptions/{subscriptionId} | 
[**SubscriptionsSubscriptionIdPut**](DefaultApi.md#SubscriptionsSubscriptionIdPut) | **Put** /subscriptions/{subscriptionId} | 



## SubscriptionsPost

> ACInfoSubscription SubscriptionsPost(ctx).ACInfoSubscription(aCInfoSubscription).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Eees_AppClientInformation"
)

func main() {
    aCInfoSubscription := *openapiclient.NewACInfoSubscription("EasId_example") // ACInfoSubscription | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.SubscriptionsPost(context.Background()).ACInfoSubscription(aCInfoSubscription).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SubscriptionsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubscriptionsPost`: ACInfoSubscription
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SubscriptionsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aCInfoSubscription** | [**ACInfoSubscription**](ACInfoSubscription.md) |  | 

### Return type

[**ACInfoSubscription**](ACInfoSubscription.md)

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
    openapiclient "github.com/5GCoreNet/client-openapi/Eees_AppClientInformation"
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

> ACInfoSubscription SubscriptionsSubscriptionIdGet(ctx, subscriptionId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Eees_AppClientInformation"
)

func main() {
    subscriptionId := "subscriptionId_example" // string | Subscription Id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.SubscriptionsSubscriptionIdGet(context.Background(), subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SubscriptionsSubscriptionIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubscriptionsSubscriptionIdGet`: ACInfoSubscription
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


### Return type

[**ACInfoSubscription**](ACInfoSubscription.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsSubscriptionIdPatch

> ACInfoSubscription SubscriptionsSubscriptionIdPatch(ctx, subscriptionId).ACInfoSubscriptionPatch(aCInfoSubscriptionPatch).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Eees_AppClientInformation"
)

func main() {
    subscriptionId := "subscriptionId_example" // string | Subscription Id.
    aCInfoSubscriptionPatch := *openapiclient.NewACInfoSubscriptionPatch() // ACInfoSubscriptionPatch | Partial update an existing Individual AC information Subscription.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.SubscriptionsSubscriptionIdPatch(context.Background(), subscriptionId).ACInfoSubscriptionPatch(aCInfoSubscriptionPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SubscriptionsSubscriptionIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubscriptionsSubscriptionIdPatch`: ACInfoSubscription
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

 **aCInfoSubscriptionPatch** | [**ACInfoSubscriptionPatch**](ACInfoSubscriptionPatch.md) | Partial update an existing Individual AC information Subscription. | 

### Return type

[**ACInfoSubscription**](ACInfoSubscription.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsSubscriptionIdPut

> ACInfoSubscription SubscriptionsSubscriptionIdPut(ctx, subscriptionId).ACInfoSubscription(aCInfoSubscription).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Eees_AppClientInformation"
)

func main() {
    subscriptionId := "subscriptionId_example" // string | Subscription Id.
    aCInfoSubscription := *openapiclient.NewACInfoSubscription("EasId_example") // ACInfoSubscription | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.SubscriptionsSubscriptionIdPut(context.Background(), subscriptionId).ACInfoSubscription(aCInfoSubscription).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SubscriptionsSubscriptionIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubscriptionsSubscriptionIdPut`: ACInfoSubscription
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

 **aCInfoSubscription** | [**ACInfoSubscription**](ACInfoSubscription.md) |  | 

### Return type

[**ACInfoSubscription**](ACInfoSubscription.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

