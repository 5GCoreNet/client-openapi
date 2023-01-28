# \AuthenticationUPUDocumentApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAuthenticationUPU**](AuthenticationUPUDocumentApi.md#CreateAuthenticationUPU) | **Put** /subscription-data/{ueId}/ue-update-confirmation-data/upu-data | To store the UPU acknowledgement information of a UE
[**QueryAuthUPU**](AuthenticationUPUDocumentApi.md#QueryAuthUPU) | **Get** /subscription-data/{ueId}/ue-update-confirmation-data/upu-data | Retrieves the UPU acknowledgement information of a UE



## CreateAuthenticationUPU

> CreateAuthenticationUPU(ctx, ueId).UpuData(upuData).SupportedFeatures(supportedFeatures).Execute()

To store the UPU acknowledgement information of a UE

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/5GCoreNet/client-openapi/Subscription_Data"
)

func main() {
    ueId := "ueId_example" // string | UE id
    upuData := *openapiclient.NewUpuData(time.Now(), openapiclient.UeUpdateStatus("NOT_SENT")) // UpuData | 
    supportedFeatures := "supportedFeatures_example" // string | Supported Features (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticationUPUDocumentApi.CreateAuthenticationUPU(context.Background(), ueId).UpuData(upuData).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationUPUDocumentApi.CreateAuthenticationUPU``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | UE id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAuthenticationUPURequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **upuData** | [**UpuData**](UpuData.md) |  | 
 **supportedFeatures** | **string** | Supported Features | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryAuthUPU

> UpuData QueryAuthUPU(ctx, ueId).SupportedFeatures(supportedFeatures).Execute()

Retrieves the UPU acknowledgement information of a UE

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Subscription_Data"
)

func main() {
    ueId := "ueId_example" // string | UE id
    supportedFeatures := "supportedFeatures_example" // string | Supported Features (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticationUPUDocumentApi.QueryAuthUPU(context.Background(), ueId).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationUPUDocumentApi.QueryAuthUPU``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryAuthUPU`: UpuData
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationUPUDocumentApi.QueryAuthUPU`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | UE id | 

### Other Parameters

Other parameters are passed through a pointer to a apiQueryAuthUPURequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **supportedFeatures** | **string** | Supported Features | 

### Return type

[**UpuData**](UpuData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

