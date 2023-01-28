# \TimersDeleteApi

All URIs are relative to *https://example.com/nudsf-timer/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteTimers**](TimersDeleteApi.md#DeleteTimers) | **Delete** /{realmId}/{storageId}/timers | Delete one or multiple timers based on filter



## DeleteTimers

> TimerIdList DeleteTimers(ctx, realmId, storageId).SupportedFeatures(supportedFeatures).Filter(filter).ExpiredFilter(expiredFilter).Execute()

Delete one or multiple timers based on filter

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
    realmId := "Realm01" // string | Identifier(name) of the Realm
    storageId := "Storage01" // string | Identifier of the Storage
    supportedFeatures := "supportedFeatures_example" // string | Features required to be supported by the target NF (optional)
    filter := map[string][]openapiclient.SearchExpression{ ... } // SearchExpression | A filter that determines the set of timers to be deleted (optional)
    expiredFilter := openapiclient.NullValue("null") // NullValue | Presence indicates that only expired timers are to be deleted. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TimersDeleteApi.DeleteTimers(context.Background(), realmId, storageId).SupportedFeatures(supportedFeatures).Filter(filter).ExpiredFilter(expiredFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TimersDeleteApi.DeleteTimers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteTimers`: TimerIdList
    fmt.Fprintf(os.Stdout, "Response from `TimersDeleteApi.DeleteTimers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realmId** | **string** | Identifier(name) of the Realm | 
**storageId** | **string** | Identifier of the Storage | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTimersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **supportedFeatures** | **string** | Features required to be supported by the target NF | 
 **filter** | [**SearchExpression**](SearchExpression.md) | A filter that determines the set of timers to be deleted | 
 **expiredFilter** | [**NullValue**](NullValue.md) | Presence indicates that only expired timers are to be deleted. | 

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

