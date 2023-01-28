# \UpdateTheAuthorizationForMonitoringForAUEApi

All URIs are relative to *https://example.com/n5g-ddnmf-disc/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateMonitorAuth**](UpdateTheAuthorizationForMonitoringForAUEApi.md#UpdateMonitorAuth) | **Patch** /{ueId}/monitor-authorize/{discEntryId} | Update the authorization for monitoring for a UE



## UpdateMonitorAuth

> PatchResult UpdateMonitorAuth(ctx, ueId, discEntryId).MonitorUpdateData(monitorUpdateData).Execute()

Update the authorization for monitoring for a UE

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/N5g-ddnmf_Discovery"
)

func main() {
    ueId := "ueId_example" // string | Identifier of the UE
    discEntryId := "discEntryId_example" // string | Discovery Entry Id
    monitorUpdateData := *openapiclient.NewMonitorUpdateData(*openapiclient.NewDiscoveryType()) // MonitorUpdateData | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UpdateTheAuthorizationForMonitoringForAUEApi.UpdateMonitorAuth(context.Background(), ueId, discEntryId).MonitorUpdateData(monitorUpdateData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UpdateTheAuthorizationForMonitoringForAUEApi.UpdateMonitorAuth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMonitorAuth`: PatchResult
    fmt.Fprintf(os.Stdout, "Response from `UpdateTheAuthorizationForMonitoringForAUEApi.UpdateMonitorAuth`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | Identifier of the UE | 
**discEntryId** | **string** | Discovery Entry Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMonitorAuthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **monitorUpdateData** | [**MonitorUpdateData**](MonitorUpdateData.md) |  | 

### Return type

[**PatchResult**](PatchResult.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

