# \TimerSearchApi

All URIs are relative to *https://example.com/nudsf-timer/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchTimer**](TimerSearchApi.md#SearchTimer) | **Get** /{realmId}/{storageId}/timers | Timers search with get



## SearchTimer

> TimerIdList SearchTimer(ctx, realmId, storageId).Filter(filter).ExpiredFilter(expiredFilter).SupportedFeatures(supportedFeatures).Execute()

Timers search with get



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nudsf_Timer"
)

func main() {
    realmId := "Realm01" // string | Identifier of the Realm
    storageId := "Storage01" // string | Identifier of the Storage
    filter := map[string][]openapiclient.SearchExpression{ ... } // SearchExpression | Query filter using conditions on tags (optional)
    expiredFilter := openapiclient.NullValue("null") // NullValue | Used to query for expired timers. (optional)
    supportedFeatures := "supportedFeatures_example" // string | Features required to be supported by the target NF (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TimerSearchApi.SearchTimer(context.Background(), realmId, storageId).Filter(filter).ExpiredFilter(expiredFilter).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TimerSearchApi.SearchTimer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchTimer`: TimerIdList
    fmt.Fprintf(os.Stdout, "Response from `TimerSearchApi.SearchTimer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realmId** | **string** | Identifier of the Realm | 
**storageId** | **string** | Identifier of the Storage | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchTimerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | [**SearchExpression**](SearchExpression.md) | Query filter using conditions on tags | 
 **expiredFilter** | [**NullValue**](NullValue.md) | Used to query for expired timers. | 
 **supportedFeatures** | **string** | Features required to be supported by the target NF | 

### Return type

[**TimerIdList**](TimerIdList.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

