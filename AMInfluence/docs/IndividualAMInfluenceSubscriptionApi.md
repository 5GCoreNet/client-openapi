# \IndividualAMInfluenceSubscriptionApi

All URIs are relative to *https://example.com/3gpp-am-influence/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AfIdSubscriptionsSubscriptionIdDelete**](IndividualAMInfluenceSubscriptionApi.md#AfIdSubscriptionsSubscriptionIdDelete) | **Delete** /{afId}/subscriptions/{subscriptionId} | Delete an existing subscription.
[**AfIdSubscriptionsSubscriptionIdGet**](IndividualAMInfluenceSubscriptionApi.md#AfIdSubscriptionsSubscriptionIdGet) | **Get** /{afId}/subscriptions/{subscriptionId} | Read an active subscription identified by the subscriptionId.
[**AfIdSubscriptionsSubscriptionIdPatch**](IndividualAMInfluenceSubscriptionApi.md#AfIdSubscriptionsSubscriptionIdPatch) | **Patch** /{afId}/subscriptions/{subscriptionId} | Update/Replace an existing subscription resource.
[**AfIdSubscriptionsSubscriptionIdPut**](IndividualAMInfluenceSubscriptionApi.md#AfIdSubscriptionsSubscriptionIdPut) | **Put** /{afId}/subscriptions/{subscriptionId} | Update/Replace an existing subscription resource.



## AfIdSubscriptionsSubscriptionIdDelete

> AfIdSubscriptionsSubscriptionIdDelete(ctx, afId, subscriptionId).Execute()

Delete an existing subscription.

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
    afId := "afId_example" // string | Identifier of the AF.
    subscriptionId := "subscriptionId_example" // string | Identifier of the subscription resource.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualAMInfluenceSubscriptionApi.AfIdSubscriptionsSubscriptionIdDelete(context.Background(), afId, subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualAMInfluenceSubscriptionApi.AfIdSubscriptionsSubscriptionIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**afId** | **string** | Identifier of the AF. | 
**subscriptionId** | **string** | Identifier of the subscription resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAfIdSubscriptionsSubscriptionIdDeleteRequest struct via the builder pattern


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


## AfIdSubscriptionsSubscriptionIdGet

> AmInfluSub AfIdSubscriptionsSubscriptionIdGet(ctx, afId, subscriptionId).Execute()

Read an active subscription identified by the subscriptionId.

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
    afId := "afId_example" // string | Identifier of the AF.
    subscriptionId := "subscriptionId_example" // string | Identifier of the subscription resource.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualAMInfluenceSubscriptionApi.AfIdSubscriptionsSubscriptionIdGet(context.Background(), afId, subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualAMInfluenceSubscriptionApi.AfIdSubscriptionsSubscriptionIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AfIdSubscriptionsSubscriptionIdGet`: AmInfluSub
    fmt.Fprintf(os.Stdout, "Response from `IndividualAMInfluenceSubscriptionApi.AfIdSubscriptionsSubscriptionIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**afId** | **string** | Identifier of the AF. | 
**subscriptionId** | **string** | Identifier of the subscription resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAfIdSubscriptionsSubscriptionIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AmInfluSub**](AmInfluSub.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AfIdSubscriptionsSubscriptionIdPatch

> AmInfluSub AfIdSubscriptionsSubscriptionIdPatch(ctx, afId, subscriptionId).AmInfluSubPatch(amInfluSubPatch).Execute()

Update/Replace an existing subscription resource.

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
    afId := "afId_example" // string | Identifier of the AF.
    subscriptionId := "subscriptionId_example" // string | Identifier of the subscription resource.
    amInfluSubPatch := *openapiclient.NewAmInfluSubPatch() // AmInfluSubPatch | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualAMInfluenceSubscriptionApi.AfIdSubscriptionsSubscriptionIdPatch(context.Background(), afId, subscriptionId).AmInfluSubPatch(amInfluSubPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualAMInfluenceSubscriptionApi.AfIdSubscriptionsSubscriptionIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AfIdSubscriptionsSubscriptionIdPatch`: AmInfluSub
    fmt.Fprintf(os.Stdout, "Response from `IndividualAMInfluenceSubscriptionApi.AfIdSubscriptionsSubscriptionIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**afId** | **string** | Identifier of the AF. | 
**subscriptionId** | **string** | Identifier of the subscription resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAfIdSubscriptionsSubscriptionIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **amInfluSubPatch** | [**AmInfluSubPatch**](AmInfluSubPatch.md) |  | 

### Return type

[**AmInfluSub**](AmInfluSub.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AfIdSubscriptionsSubscriptionIdPut

> AmInfluSub AfIdSubscriptionsSubscriptionIdPut(ctx, afId, subscriptionId).AmInfluSub(amInfluSub).Execute()

Update/Replace an existing subscription resource.

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
    afId := "afId_example" // string | Identifier of the AF.
    subscriptionId := "subscriptionId_example" // string | Identifier of the subscription resource.
    amInfluSub := *openapiclient.NewAmInfluSub("AfTransId_example") // AmInfluSub | Parameters to update/replace the existing subscription.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualAMInfluenceSubscriptionApi.AfIdSubscriptionsSubscriptionIdPut(context.Background(), afId, subscriptionId).AmInfluSub(amInfluSub).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualAMInfluenceSubscriptionApi.AfIdSubscriptionsSubscriptionIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AfIdSubscriptionsSubscriptionIdPut`: AmInfluSub
    fmt.Fprintf(os.Stdout, "Response from `IndividualAMInfluenceSubscriptionApi.AfIdSubscriptionsSubscriptionIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**afId** | **string** | Identifier of the AF. | 
**subscriptionId** | **string** | Identifier of the subscription resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAfIdSubscriptionsSubscriptionIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **amInfluSub** | [**AmInfluSub**](AmInfluSub.md) | Parameters to update/replace the existing subscription. | 

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

