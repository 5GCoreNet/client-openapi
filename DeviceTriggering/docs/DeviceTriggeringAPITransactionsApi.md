# \DeviceTriggeringAPITransactionsApi

All URIs are relative to *https://example.com/3gpp-device-triggering/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDeviceTriggeringTransaction**](DeviceTriggeringAPITransactionsApi.md#CreateDeviceTriggeringTransaction) | **Post** /{scsAsId}/transactions | Create a long-term transaction for a device triggering.



## CreateDeviceTriggeringTransaction

> DeviceTriggering CreateDeviceTriggeringTransaction(ctx, scsAsId).DeviceTriggering(deviceTriggering).Execute()

Create a long-term transaction for a device triggering.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/DeviceTriggering"
)

func main() {
    scsAsId := "scsAsId_example" // string | Identifier of the SCS/AS
    deviceTriggering := openapiclient.DeviceTriggering{Interface{}: new(interface{})} // DeviceTriggering | Parameters to request a device triggering delivery.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceTriggeringAPITransactionsApi.CreateDeviceTriggeringTransaction(context.Background(), scsAsId).DeviceTriggering(deviceTriggering).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceTriggeringAPITransactionsApi.CreateDeviceTriggeringTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDeviceTriggeringTransaction`: DeviceTriggering
    fmt.Fprintf(os.Stdout, "Response from `DeviceTriggeringAPITransactionsApi.CreateDeviceTriggeringTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of the SCS/AS | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceTriggeringTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deviceTriggering** | [**DeviceTriggering**](DeviceTriggering.md) | Parameters to request a device triggering delivery. | 

### Return type

[**DeviceTriggering**](DeviceTriggering.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

