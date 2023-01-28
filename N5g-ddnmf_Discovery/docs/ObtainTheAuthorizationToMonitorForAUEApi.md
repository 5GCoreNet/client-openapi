# \ObtainTheAuthorizationToMonitorForAUEApi

All URIs are relative to *https://example.com/n5g-ddnmf-disc/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ObtainMonitorAuth**](ObtainTheAuthorizationToMonitorForAUEApi.md#ObtainMonitorAuth) | **Put** /{ueId}/monitor-authorize/{discEntryId} | Obtain the authorization to monitor for a UE



## ObtainMonitorAuth

> MonitorAuthRespData ObtainMonitorAuth(ctx, ueId, discEntryId).MonitorAuthReqData(monitorAuthReqData).Execute()

Obtain the authorization to monitor for a UE

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
    monitorAuthReqData := *openapiclient.NewMonitorAuthReqData(*openapiclient.NewDiscoveryType()) // MonitorAuthReqData | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObtainTheAuthorizationToMonitorForAUEApi.ObtainMonitorAuth(context.Background(), ueId, discEntryId).MonitorAuthReqData(monitorAuthReqData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObtainTheAuthorizationToMonitorForAUEApi.ObtainMonitorAuth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ObtainMonitorAuth`: MonitorAuthRespData
    fmt.Fprintf(os.Stdout, "Response from `ObtainTheAuthorizationToMonitorForAUEApi.ObtainMonitorAuth`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | Identifier of the UE | 
**discEntryId** | **string** | Discovery Entry Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiObtainMonitorAuthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **monitorAuthReqData** | [**MonitorAuthReqData**](MonitorAuthReqData.md) |  | 

### Return type

[**MonitorAuthRespData**](MonitorAuthRespData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

