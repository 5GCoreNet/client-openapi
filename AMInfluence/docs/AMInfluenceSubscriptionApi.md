# \AMInfluenceSubscriptionApi

All URIs are relative to *https://example.com/3gpp-am-influence/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AfIdSubscriptionsGet**](AMInfluenceSubscriptionApi.md#AfIdSubscriptionsGet) | **Get** /{afId}/subscriptions | Read all of the active subscriptions for the AF.
[**CreateAMInfluenceSubcription**](AMInfluenceSubscriptionApi.md#CreateAMInfluenceSubcription) | **Post** /{afId}/subscriptions | Create a new subscription to AM influence.



## AfIdSubscriptionsGet

> []AmInfluSub AfIdSubscriptionsGet(ctx, afId).Execute()

Read all of the active subscriptions for the AF.

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
    afId := "afId_example" // string | Identifier of the AF

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AMInfluenceSubscriptionApi.AfIdSubscriptionsGet(context.Background(), afId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AMInfluenceSubscriptionApi.AfIdSubscriptionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AfIdSubscriptionsGet`: []AmInfluSub
    fmt.Fprintf(os.Stdout, "Response from `AMInfluenceSubscriptionApi.AfIdSubscriptionsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**afId** | **string** | Identifier of the AF | 

### Other Parameters

Other parameters are passed through a pointer to a apiAfIdSubscriptionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]AmInfluSub**](AmInfluSub.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAMInfluenceSubcription

> AmInfluSub CreateAMInfluenceSubcription(ctx, afId).AmInfluSub(amInfluSub).Execute()

Create a new subscription to AM influence.

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
    afId := "afId_example" // string | Identifier of the AF
    amInfluSub := *openapiclient.NewAmInfluSub("AfTransId_example") // AmInfluSub | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AMInfluenceSubscriptionApi.CreateAMInfluenceSubcription(context.Background(), afId).AmInfluSub(amInfluSub).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AMInfluenceSubscriptionApi.CreateAMInfluenceSubcription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAMInfluenceSubcription`: AmInfluSub
    fmt.Fprintf(os.Stdout, "Response from `AMInfluenceSubscriptionApi.CreateAMInfluenceSubcription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**afId** | **string** | Identifier of the AF | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAMInfluenceSubcriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **amInfluSub** | [**AmInfluSub**](AmInfluSub.md) |  | 

### Return type

[**AmInfluSub**](AmInfluSub.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

