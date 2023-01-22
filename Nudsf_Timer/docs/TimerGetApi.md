# \TimerGetApi

All URIs are relative to *https://example.com/nudsf-timer/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTimer**](TimerGetApi.md#GetTimer) | **Get** /{realmId}/{storageId}/timers/{timerId} | Timer access



## GetTimer

> Timer GetTimer(ctx, realmId, storageId, timerId).SupportedFeatures(supportedFeatures).Execute()

Timer access



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
    realmId := "Realm01" // string | Identifier of the Realm
    storageId := "Storage01" // string | Identifier of the Storage
    timerId := "timerId_example" // string | Identifier of the Timer
    supportedFeatures := "supportedFeatures_example" // string | Features required to be supported by the target NF (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TimerGetApi.GetTimer(context.Background(), realmId, storageId, timerId).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TimerGetApi.GetTimer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTimer`: Timer
    fmt.Fprintf(os.Stdout, "Response from `TimerGetApi.GetTimer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realmId** | **string** | Identifier of the Realm | 
**storageId** | **string** | Identifier of the Storage | 
**timerId** | **string** | Identifier of the Timer | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTimerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **supportedFeatures** | **string** | Features required to be supported by the target NF | 

### Return type

[**Timer**](Timer.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

