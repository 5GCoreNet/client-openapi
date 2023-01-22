# \TimerUpdateApi

All URIs are relative to *https://example.com/nudsf-timer/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateTimer**](TimerUpdateApi.md#UpdateTimer) | **Patch** /{realmId}/{storageId}/timers/{timerId} | Timer modification



## UpdateTimer

> PatchResult UpdateTimer(ctx, realmId, storageId, timerId).PatchItem(patchItem).SupportedFeatures(supportedFeatures).Execute()

Timer modification



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
    patchItem := []openapiclient.PatchItem{*openapiclient.NewPatchItem(*openapiclient.NewPatchOperation(), "Path_example")} // []PatchItem | Timer data to patch
    supportedFeatures := "supportedFeatures_example" // string | Features required to be supported by the target NF (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TimerUpdateApi.UpdateTimer(context.Background(), realmId, storageId, timerId).PatchItem(patchItem).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TimerUpdateApi.UpdateTimer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTimer`: PatchResult
    fmt.Fprintf(os.Stdout, "Response from `TimerUpdateApi.UpdateTimer`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiUpdateTimerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **patchItem** | [**[]PatchItem**](PatchItem.md) | Timer data to patch | 
 **supportedFeatures** | **string** | Features required to be supported by the target NF | 

### Return type

[**PatchResult**](PatchResult.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

