# \IndividualDeviceTriggeringTransactionApi

All URIs are relative to *https://example.com/3gpp-device-triggering/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIndDeviceTriggeringTransaction**](IndividualDeviceTriggeringTransactionApi.md#DeleteIndDeviceTriggeringTransaction) | **Delete** /{scsAsId}/transactions/{transactionId} | Deletes an already existing device triggering transaction.
[**FetchIndDeviceTriggeringTransaction**](IndividualDeviceTriggeringTransactionApi.md#FetchIndDeviceTriggeringTransaction) | **Get** /{scsAsId}/transactions/{transactionId} | Read a device triggering transaction resource.
[**ModifyIndDeviceTriggeringTransaction**](IndividualDeviceTriggeringTransactionApi.md#ModifyIndDeviceTriggeringTransaction) | **Patch** /{scsAsId}/transactions/{transactionId} | Modify an existing Individual Device Triggering Transaction resource and the corresponding device triggering request.
[**UpdateIndDeviceTriggeringTransaction**](IndividualDeviceTriggeringTransactionApi.md#UpdateIndDeviceTriggeringTransaction) | **Put** /{scsAsId}/transactions/{transactionId} | Replace an existing device triggering transaction resource and the corresponding device trigger request.



## DeleteIndDeviceTriggeringTransaction

> DeviceTriggering DeleteIndDeviceTriggeringTransaction(ctx, scsAsId, transactionId).Execute()

Deletes an already existing device triggering transaction.

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
    transactionId := "transactionId_example" // string | Identifier of the transaction resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualDeviceTriggeringTransactionApi.DeleteIndDeviceTriggeringTransaction(context.Background(), scsAsId, transactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualDeviceTriggeringTransactionApi.DeleteIndDeviceTriggeringTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteIndDeviceTriggeringTransaction`: DeviceTriggering
    fmt.Fprintf(os.Stdout, "Response from `IndividualDeviceTriggeringTransactionApi.DeleteIndDeviceTriggeringTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of the SCS/AS | 
**transactionId** | **string** | Identifier of the transaction resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIndDeviceTriggeringTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DeviceTriggering**](DeviceTriggering.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchIndDeviceTriggeringTransaction

> DeviceTriggering FetchIndDeviceTriggeringTransaction(ctx, scsAsId, transactionId).Execute()

Read a device triggering transaction resource.

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
    transactionId := "transactionId_example" // string | Identifier of the transaction resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualDeviceTriggeringTransactionApi.FetchIndDeviceTriggeringTransaction(context.Background(), scsAsId, transactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualDeviceTriggeringTransactionApi.FetchIndDeviceTriggeringTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchIndDeviceTriggeringTransaction`: DeviceTriggering
    fmt.Fprintf(os.Stdout, "Response from `IndividualDeviceTriggeringTransactionApi.FetchIndDeviceTriggeringTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of the SCS/AS | 
**transactionId** | **string** | Identifier of the transaction resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchIndDeviceTriggeringTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DeviceTriggering**](DeviceTriggering.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyIndDeviceTriggeringTransaction

> DeviceTriggering ModifyIndDeviceTriggeringTransaction(ctx, scsAsId, transactionId).DeviceTriggeringPatch(deviceTriggeringPatch).Execute()

Modify an existing Individual Device Triggering Transaction resource and the corresponding device triggering request.

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
    transactionId := "transactionId_example" // string | Identifier of the transaction resource
    deviceTriggeringPatch := *openapiclient.NewDeviceTriggeringPatch() // DeviceTriggeringPatch | Parameters to request the modification of the existing Individual Device Triggering Transaction resource.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualDeviceTriggeringTransactionApi.ModifyIndDeviceTriggeringTransaction(context.Background(), scsAsId, transactionId).DeviceTriggeringPatch(deviceTriggeringPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualDeviceTriggeringTransactionApi.ModifyIndDeviceTriggeringTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ModifyIndDeviceTriggeringTransaction`: DeviceTriggering
    fmt.Fprintf(os.Stdout, "Response from `IndividualDeviceTriggeringTransactionApi.ModifyIndDeviceTriggeringTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of the SCS/AS | 
**transactionId** | **string** | Identifier of the transaction resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyIndDeviceTriggeringTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deviceTriggeringPatch** | [**DeviceTriggeringPatch**](DeviceTriggeringPatch.md) | Parameters to request the modification of the existing Individual Device Triggering Transaction resource. | 

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


## UpdateIndDeviceTriggeringTransaction

> DeviceTriggering UpdateIndDeviceTriggeringTransaction(ctx, scsAsId, transactionId).DeviceTriggering(deviceTriggering).Execute()

Replace an existing device triggering transaction resource and the corresponding device trigger request.

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
    transactionId := "transactionId_example" // string | Identifier of the transaction resource
    deviceTriggering := openapiclient.DeviceTriggering{Interface{}: new(interface{})} // DeviceTriggering | Parameters to update/replace the existing device triggering

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualDeviceTriggeringTransactionApi.UpdateIndDeviceTriggeringTransaction(context.Background(), scsAsId, transactionId).DeviceTriggering(deviceTriggering).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualDeviceTriggeringTransactionApi.UpdateIndDeviceTriggeringTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIndDeviceTriggeringTransaction`: DeviceTriggering
    fmt.Fprintf(os.Stdout, "Response from `IndividualDeviceTriggeringTransactionApi.UpdateIndDeviceTriggeringTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of the SCS/AS | 
**transactionId** | **string** | Identifier of the transaction resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIndDeviceTriggeringTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deviceTriggering** | [**DeviceTriggering**](DeviceTriggering.md) | Parameters to update/replace the existing device triggering | 

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

