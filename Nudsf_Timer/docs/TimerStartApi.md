# \TimerStartApi

All URIs are relative to *https://example.com/nudsf-timer/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrModifyTimer**](TimerStartApi.md#CreateOrModifyTimer) | **Put** /{realmId}/{storageId}/timers/{timerId} | Create/Replace Timer



## CreateOrModifyTimer

> CreateOrModifyTimer(ctx, realmId, storageId, timerId).Timer(timer).SupportedFeatures(supportedFeatures).Execute()

Create/Replace Timer



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    realmId := "Realm01" // string | Identifier(name) of the Realm
    storageId := "Storage01" // string | Identifier of the Storage
    timerId := "timerId_example" // string | Identifier of the Timer
    timer := *openapiclient.NewTimer(time.Now()) // Timer | 
    supportedFeatures := "supportedFeatures_example" // string | Features required to be supported by the target NF (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TimerStartApi.CreateOrModifyTimer(context.Background(), realmId, storageId, timerId).Timer(timer).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TimerStartApi.CreateOrModifyTimer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realmId** | **string** | Identifier(name) of the Realm | 
**storageId** | **string** | Identifier of the Storage | 
**timerId** | **string** | Identifier of the Timer | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrModifyTimerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **timer** | [**Timer**](Timer.md) |  | 
 **supportedFeatures** | **string** | Features required to be supported by the target NF | 

### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

