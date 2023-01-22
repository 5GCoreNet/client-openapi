# \DefaultApi

All URIs are relative to *https://example.com/3gpp-m5/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDynamicPolicy**](DefaultApi.md#CreateDynamicPolicy) | **Post** /dynamic-policies | Create (and optionally upload) a new Dynamic Policy resource
[**DestroyDynamicPolicy**](DefaultApi.md#DestroyDynamicPolicy) | **Delete** /dynamic-policies/{dynamicPolicyId} | Destroy an existing Dynamic Policy resource
[**PatchDynamicPolicy**](DefaultApi.md#PatchDynamicPolicy) | **Patch** /dynamic-policies/{dynamicPolicyId} | Patch an existing Dynamic Policy resource
[**RetrieveDynamicPolicy**](DefaultApi.md#RetrieveDynamicPolicy) | **Get** /dynamic-policies/{dynamicPolicyId} | Retrieve an existing Dynamic Policy resource
[**UpdateDynamicPolicy**](DefaultApi.md#UpdateDynamicPolicy) | **Put** /dynamic-policies/{dynamicPolicyId} | Update an existing Dynamic Policy resource



## CreateDynamicPolicy

> DynamicPolicy CreateDynamicPolicy(ctx).DynamicPolicy(dynamicPolicy).Execute()

Create (and optionally upload) a new Dynamic Policy resource

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
    dynamicPolicy := *openapiclient.NewDynamicPolicy("DynamicPolicyId_example", "PolicyTemplateId_example", []openapiclient.ServiceDataFlowDescription{*openapiclient.NewServiceDataFlowDescription()}, "ProvisioningSessionId_example") // DynamicPolicy | An optional JSON representation of a Dynamic Policy resource (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateDynamicPolicy(context.Background()).DynamicPolicy(dynamicPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateDynamicPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDynamicPolicy`: DynamicPolicy
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateDynamicPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDynamicPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dynamicPolicy** | [**DynamicPolicy**](DynamicPolicy.md) | An optional JSON representation of a Dynamic Policy resource | 

### Return type

[**DynamicPolicy**](DynamicPolicy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyDynamicPolicy

> DestroyDynamicPolicy(ctx, dynamicPolicyId).Execute()

Destroy an existing Dynamic Policy resource

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
    dynamicPolicyId := "dynamicPolicyId_example" // string | The resource identifier of a Dynamic Policy resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DestroyDynamicPolicy(context.Background(), dynamicPolicyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DestroyDynamicPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dynamicPolicyId** | **string** | The resource identifier of a Dynamic Policy resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiDestroyDynamicPolicyRequest struct via the builder pattern


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


## PatchDynamicPolicy

> DynamicPolicy PatchDynamicPolicy(ctx, dynamicPolicyId).DynamicPolicy(dynamicPolicy).Execute()

Patch an existing Dynamic Policy resource

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
    dynamicPolicyId := "dynamicPolicyId_example" // string | The resource identifier of a Dynamic Policy resource
    dynamicPolicy := *openapiclient.NewDynamicPolicy("DynamicPolicyId_example", "PolicyTemplateId_example", []openapiclient.ServiceDataFlowDescription{*openapiclient.NewServiceDataFlowDescription()}, "ProvisioningSessionId_example") // DynamicPolicy | A JSON patch to a Dynamic Policy resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.PatchDynamicPolicy(context.Background(), dynamicPolicyId).DynamicPolicy(dynamicPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.PatchDynamicPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchDynamicPolicy`: DynamicPolicy
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.PatchDynamicPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dynamicPolicyId** | **string** | The resource identifier of a Dynamic Policy resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchDynamicPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dynamicPolicy** | [**DynamicPolicy**](DynamicPolicy.md) | A JSON patch to a Dynamic Policy resource | 

### Return type

[**DynamicPolicy**](DynamicPolicy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/merge-patch+json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveDynamicPolicy

> DynamicPolicy RetrieveDynamicPolicy(ctx, dynamicPolicyId).Execute()

Retrieve an existing Dynamic Policy resource

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
    dynamicPolicyId := "dynamicPolicyId_example" // string | The resource identifier of a Dynamic Policy resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.RetrieveDynamicPolicy(context.Background(), dynamicPolicyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RetrieveDynamicPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveDynamicPolicy`: DynamicPolicy
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.RetrieveDynamicPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dynamicPolicyId** | **string** | The resource identifier of a Dynamic Policy resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveDynamicPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DynamicPolicy**](DynamicPolicy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDynamicPolicy

> UpdateDynamicPolicy(ctx, dynamicPolicyId).DynamicPolicy(dynamicPolicy).Execute()

Update an existing Dynamic Policy resource

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
    dynamicPolicyId := "dynamicPolicyId_example" // string | The resource identifier of a Dynamic Policy resource
    dynamicPolicy := *openapiclient.NewDynamicPolicy("DynamicPolicyId_example", "PolicyTemplateId_example", []openapiclient.ServiceDataFlowDescription{*openapiclient.NewServiceDataFlowDescription()}, "ProvisioningSessionId_example") // DynamicPolicy | A replacement JSON representation of a Dynamic Policy resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateDynamicPolicy(context.Background(), dynamicPolicyId).DynamicPolicy(dynamicPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateDynamicPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dynamicPolicyId** | **string** | The resource identifier of a Dynamic Policy resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDynamicPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dynamicPolicy** | [**DynamicPolicy**](DynamicPolicy.md) | A replacement JSON representation of a Dynamic Policy resource | 

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

