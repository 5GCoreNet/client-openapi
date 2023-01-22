# \DeviceTriggeringTransactionsApi

All URIs are relative to *https://example.com/3gpp-device-triggering/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchAllDeviceTriggeringTransactions**](DeviceTriggeringTransactionsApi.md#FetchAllDeviceTriggeringTransactions) | **Get** /{scsAsId}/transactions | read all active device triggering transactions for a given SCS/AS.



## FetchAllDeviceTriggeringTransactions

> []DeviceTriggering FetchAllDeviceTriggeringTransactions(ctx, scsAsId).Execute()

read all active device triggering transactions for a given SCS/AS.

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
    scsAsId := "scsAsId_example" // string | Identifier of the SCS/AS

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceTriggeringTransactionsApi.FetchAllDeviceTriggeringTransactions(context.Background(), scsAsId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceTriggeringTransactionsApi.FetchAllDeviceTriggeringTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchAllDeviceTriggeringTransactions`: []DeviceTriggering
    fmt.Fprintf(os.Stdout, "Response from `DeviceTriggeringTransactionsApi.FetchAllDeviceTriggeringTransactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of the SCS/AS | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchAllDeviceTriggeringTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]DeviceTriggering**](DeviceTriggering.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

