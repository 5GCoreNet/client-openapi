# \DefaultApi

All URIs are relative to *https://example.com/3gpp-m1/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProvisioningSession**](DefaultApi.md#CreateProvisioningSession) | **Post** /provisioning-sessions | Create a new Provisioning Session
[**DestroyProvisioningSession**](DefaultApi.md#DestroyProvisioningSession) | **Delete** /provisioning-sessions/{provisioningSessionId} | Destroy an existing Provisioning Session
[**GetProvisioningSessionById**](DefaultApi.md#GetProvisioningSessionById) | **Get** /provisioning-sessions/{provisioningSessionId} | Retrieve an existing Provisioning Session



## CreateProvisioningSession

> ProvisioningSession CreateProvisioningSession(ctx).Execute()

Create a new Provisioning Session

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateProvisioningSession(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateProvisioningSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProvisioningSession`: ProvisioningSession
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateProvisioningSession`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateProvisioningSessionRequest struct via the builder pattern


### Return type

[**ProvisioningSession**](ProvisioningSession.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyProvisioningSession

> DestroyProvisioningSession(ctx, provisioningSessionId).Execute()

Destroy an existing Provisioning Session

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
    provisioningSessionId := "provisioningSessionId_example" // string | The resource identifier of an existing Provisioning Session.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DestroyProvisioningSession(context.Background(), provisioningSessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DestroyProvisioningSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provisioningSessionId** | **string** | The resource identifier of an existing Provisioning Session. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDestroyProvisioningSessionRequest struct via the builder pattern


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


## GetProvisioningSessionById

> ProvisioningSession GetProvisioningSessionById(ctx, provisioningSessionId).Execute()

Retrieve an existing Provisioning Session

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
    provisioningSessionId := "provisioningSessionId_example" // string | The resource identifier of an existing Provisioning Session.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetProvisioningSessionById(context.Background(), provisioningSessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetProvisioningSessionById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProvisioningSessionById`: ProvisioningSession
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetProvisioningSessionById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provisioningSessionId** | **string** | The resource identifier of an existing Provisioning Session. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProvisioningSessionByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProvisioningSession**](ProvisioningSession.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

