# \DefaultApi

All URIs are relative to *https://example.com/3gpp-m5/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkAssistanceSession**](DefaultApi.md#CreateNetworkAssistanceSession) | **Post** /network-assistance/ | Create a new Network Assistance Session.
[**DestroyNetworkAssistanceSession**](DefaultApi.md#DestroyNetworkAssistanceSession) | **Delete** /network-assistance/{naSessionId} | Destroy an existing Network Assistance Session resource
[**PatchNetworkAssistanceSession**](DefaultApi.md#PatchNetworkAssistanceSession) | **Patch** /network-assistance/{naSessionId} | Patch an existing Network Assistance Session resource
[**RequestBitRateRecommendation**](DefaultApi.md#RequestBitRateRecommendation) | **Get** /network-assistance/{naSessionId}/recommendation | Obtain a bit rate recommendation for the next recommendation window
[**RequestDeliveryBoost**](DefaultApi.md#RequestDeliveryBoost) | **Post** /network-assistance/{naSessionId}/boost-request | Request a delivery boost
[**RetrieveNetworkAssistanceSession**](DefaultApi.md#RetrieveNetworkAssistanceSession) | **Get** /network-assistance/{naSessionId} | Retrieve an existing Network Assistance Session resource
[**UpdateNetworkAssistanceSession**](DefaultApi.md#UpdateNetworkAssistanceSession) | **Put** /network-assistance/{naSessionId} | Update an existing Network Assistance Session resource



## CreateNetworkAssistanceSession

> NetworkAssistanceSession CreateNetworkAssistanceSession(ctx).Execute()

Create a new Network Assistance Session.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/M5_NetworkAssistance"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateNetworkAssistanceSession(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateNetworkAssistanceSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkAssistanceSession`: NetworkAssistanceSession
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateNetworkAssistanceSession`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkAssistanceSessionRequest struct via the builder pattern


### Return type

[**NetworkAssistanceSession**](NetworkAssistanceSession.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyNetworkAssistanceSession

> DestroyNetworkAssistanceSession(ctx, naSessionId).Execute()

Destroy an existing Network Assistance Session resource

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/M5_NetworkAssistance"
)

func main() {
    naSessionId := "naSessionId_example" // string | The resource identifier of an existing Network Assistance Session resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DestroyNetworkAssistanceSession(context.Background(), naSessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DestroyNetworkAssistanceSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**naSessionId** | **string** | The resource identifier of an existing Network Assistance Session resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiDestroyNetworkAssistanceSessionRequest struct via the builder pattern


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


## PatchNetworkAssistanceSession

> NetworkAssistanceSession PatchNetworkAssistanceSession(ctx, naSessionId).NetworkAssistanceSession(networkAssistanceSession).Execute()

Patch an existing Network Assistance Session resource

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/M5_NetworkAssistance"
)

func main() {
    naSessionId := "naSessionId_example" // string | The resource identifier of an existing Network Assistance Session resource
    networkAssistanceSession := *openapiclient.NewNetworkAssistanceSession("NaSessionId_example") // NetworkAssistanceSession | A JSON patch to a Network Assistance Session resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.PatchNetworkAssistanceSession(context.Background(), naSessionId).NetworkAssistanceSession(networkAssistanceSession).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.PatchNetworkAssistanceSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchNetworkAssistanceSession`: NetworkAssistanceSession
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.PatchNetworkAssistanceSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**naSessionId** | **string** | The resource identifier of an existing Network Assistance Session resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchNetworkAssistanceSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkAssistanceSession** | [**NetworkAssistanceSession**](NetworkAssistanceSession.md) | A JSON patch to a Network Assistance Session resource | 

### Return type

[**NetworkAssistanceSession**](NetworkAssistanceSession.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/merge-patch+json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestBitRateRecommendation

> M5QoSSpecification RequestBitRateRecommendation(ctx, naSessionId).Execute()

Obtain a bit rate recommendation for the next recommendation window

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/M5_NetworkAssistance"
)

func main() {
    naSessionId := "naSessionId_example" // string | The resource identifier of an existing Network Assistance Session resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.RequestBitRateRecommendation(context.Background(), naSessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RequestBitRateRecommendation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RequestBitRateRecommendation`: M5QoSSpecification
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.RequestBitRateRecommendation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**naSessionId** | **string** | The resource identifier of an existing Network Assistance Session resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequestBitRateRecommendationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**M5QoSSpecification**](M5QoSSpecification.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestDeliveryBoost

> OperationSuccessResponse RequestDeliveryBoost(ctx, naSessionId).Execute()

Request a delivery boost

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/M5_NetworkAssistance"
)

func main() {
    naSessionId := "naSessionId_example" // string | The resource identifier of an existing Network Assistance Session resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.RequestDeliveryBoost(context.Background(), naSessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RequestDeliveryBoost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RequestDeliveryBoost`: OperationSuccessResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.RequestDeliveryBoost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**naSessionId** | **string** | The resource identifier of an existing Network Assistance Session resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequestDeliveryBoostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperationSuccessResponse**](OperationSuccessResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveNetworkAssistanceSession

> NetworkAssistanceSession RetrieveNetworkAssistanceSession(ctx, naSessionId).Execute()

Retrieve an existing Network Assistance Session resource

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/M5_NetworkAssistance"
)

func main() {
    naSessionId := "naSessionId_example" // string | The resource identifier of an existing Network Assistance Session resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.RetrieveNetworkAssistanceSession(context.Background(), naSessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RetrieveNetworkAssistanceSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveNetworkAssistanceSession`: NetworkAssistanceSession
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.RetrieveNetworkAssistanceSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**naSessionId** | **string** | The resource identifier of an existing Network Assistance Session resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveNetworkAssistanceSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NetworkAssistanceSession**](NetworkAssistanceSession.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkAssistanceSession

> UpdateNetworkAssistanceSession(ctx, naSessionId).NetworkAssistanceSession(networkAssistanceSession).Execute()

Update an existing Network Assistance Session resource

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/M5_NetworkAssistance"
)

func main() {
    naSessionId := "naSessionId_example" // string | The resource identifier of an existing Network Assistance Session resource
    networkAssistanceSession := *openapiclient.NewNetworkAssistanceSession("NaSessionId_example") // NetworkAssistanceSession | A replacement JSON representation of a Network Assistance Session resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateNetworkAssistanceSession(context.Background(), naSessionId).NetworkAssistanceSession(networkAssistanceSession).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateNetworkAssistanceSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**naSessionId** | **string** | The resource identifier of an existing Network Assistance Session resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkAssistanceSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkAssistanceSession** | [**NetworkAssistanceSession**](NetworkAssistanceSession.md) | A replacement JSON representation of a Network Assistance Session resource | 

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

