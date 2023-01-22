# \DefaultApi

All URIs are relative to *https://example.com/3gpp-m1/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrReserveServerCertificate**](DefaultApi.md#CreateOrReserveServerCertificate) | **Post** /provisioning-sessions/{provisioningSessionId}/certificates | Create or reserve a Service Certificate resource
[**DestroyServerCertificate**](DefaultApi.md#DestroyServerCertificate) | **Delete** /provisioning-sessions/{provisioningSessionId}/certificates/{certificateId} | Destroy an existing Server Certificate resource
[**RetrieveServerCertificate**](DefaultApi.md#RetrieveServerCertificate) | **Get** /provisioning-sessions/{provisioningSessionId}/certificates/{certificateId} | Retrieve the X.509 certificate representation of the specified Server Certificate resource
[**UploadServerCertificate**](DefaultApi.md#UploadServerCertificate) | **Put** /provisioning-sessions/{provisioningSessionId}/certificates/{certificateId} | Upload the X.509 certificate for a previously reserved Server Certificate resource



## CreateOrReserveServerCertificate

> string CreateOrReserveServerCertificate(ctx, provisioningSessionId).Csr(csr).Execute()

Create or reserve a Service Certificate resource



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
    provisioningSessionId := "provisioningSessionId_example" // string | The resource identifier of an existing Provisioning Session.'
    csr := "csr_example" // string | When present, return a Certificate Signing Request instead of generating an X.509 certificate (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateOrReserveServerCertificate(context.Background(), provisioningSessionId).Csr(csr).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateOrReserveServerCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrReserveServerCertificate`: string
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateOrReserveServerCertificate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provisioningSessionId** | **string** | The resource identifier of an existing Provisioning Session.&#39; | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrReserveServerCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **csr** | **string** | When present, return a Certificate Signing Request instead of generating an X.509 certificate | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x-pem-file

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyServerCertificate

> DestroyServerCertificate(ctx, provisioningSessionId, certificateId).Execute()

Destroy an existing Server Certificate resource

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
    provisioningSessionId := "provisioningSessionId_example" // string | The resource identifier of an existing Provisioning Session.'
    certificateId := "certificateId_example" // string | The resource identifier of an existing Server Certificate

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DestroyServerCertificate(context.Background(), provisioningSessionId, certificateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DestroyServerCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provisioningSessionId** | **string** | The resource identifier of an existing Provisioning Session.&#39; | 
**certificateId** | **string** | The resource identifier of an existing Server Certificate | 

### Other Parameters

Other parameters are passed through a pointer to a apiDestroyServerCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveServerCertificate

> string RetrieveServerCertificate(ctx, provisioningSessionId, certificateId).Execute()

Retrieve the X.509 certificate representation of the specified Server Certificate resource

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
    provisioningSessionId := "provisioningSessionId_example" // string | The resource identifier of an existing Provisioning Session.'
    certificateId := "certificateId_example" // string | The resource identifier of an existing Server Certificate

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.RetrieveServerCertificate(context.Background(), provisioningSessionId, certificateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RetrieveServerCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveServerCertificate`: string
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.RetrieveServerCertificate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provisioningSessionId** | **string** | The resource identifier of an existing Provisioning Session.&#39; | 
**certificateId** | **string** | The resource identifier of an existing Server Certificate | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveServerCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x-pem-file

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadServerCertificate

> UploadServerCertificate(ctx, provisioningSessionId, certificateId).Body(body).Execute()

Upload the X.509 certificate for a previously reserved Server Certificate resource

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
    provisioningSessionId := "provisioningSessionId_example" // string | The resource identifier of an existing Provisioning Session.'
    certificateId := "certificateId_example" // string | The resource identifier of an existing Server Certificate
    body := "body_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UploadServerCertificate(context.Background(), provisioningSessionId, certificateId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UploadServerCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provisioningSessionId** | **string** | The resource identifier of an existing Provisioning Session.&#39; | 
**certificateId** | **string** | The resource identifier of an existing Server Certificate | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadServerCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-pem-file
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

