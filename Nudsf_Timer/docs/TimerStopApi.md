# \TimerStopApi

All URIs are relative to *https://example.com/nudsf-timer/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteTimer**](TimerStopApi.md#DeleteTimer) | **Delete** /{realmId}/{storageId}/timers/{timerId} | Delete a Timer with an user provided TimerId



## DeleteTimer

> DeleteTimer(ctx, realmId, storageId, timerId).SupportedFeatures(supportedFeatures).Execute()

Delete a Timer with an user provided TimerId

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
    timerId := "timerId_example" // string | Identifier of the Timer
    supportedFeatures := "supportedFeatures_example" // string | Features required to be supported by the target NF (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TimerStopApi.DeleteTimer(context.Background(), realmId, storageId, timerId).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TimerStopApi.DeleteTimer``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteTimerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **supportedFeatures** | **string** | Features required to be supported by the target NF | 

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

