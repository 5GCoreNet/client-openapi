# \TMGIAllocationOperationApi

All URIs are relative to *https://example.com/3gpp-group-message-delivery-mb2/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTMGIAllocation**](TMGIAllocationOperationApi.md#CreateTMGIAllocation) | **Post** /{scsAsId}/tmgi-allocation | Creates a new TMGI Allocation resource for a given SCS/AS.
[**FetchAllTMGIAllocations**](TMGIAllocationOperationApi.md#FetchAllTMGIAllocations) | **Get** /{scsAsId}/tmgi-allocation | read all TMGI Allocation resource for a given SCS/AS



## CreateTMGIAllocation

> TMGIAllocation CreateTMGIAllocation(ctx, scsAsId).TMGIAllocation(tMGIAllocation).Execute()

Creates a new TMGI Allocation resource for a given SCS/AS.

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
    scsAsId := "scsAsId_example" // string | Identifier of SCS/AS
    tMGIAllocation := *openapiclient.NewTMGIAllocation() // TMGIAllocation | representation of the TMGI Allocation to be created in the SCEF

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TMGIAllocationOperationApi.CreateTMGIAllocation(context.Background(), scsAsId).TMGIAllocation(tMGIAllocation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TMGIAllocationOperationApi.CreateTMGIAllocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTMGIAllocation`: TMGIAllocation
    fmt.Fprintf(os.Stdout, "Response from `TMGIAllocationOperationApi.CreateTMGIAllocation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of SCS/AS | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTMGIAllocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tMGIAllocation** | [**TMGIAllocation**](TMGIAllocation.md) | representation of the TMGI Allocation to be created in the SCEF | 

### Return type

[**TMGIAllocation**](TMGIAllocation.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchAllTMGIAllocations

> TMGIAllocation FetchAllTMGIAllocations(ctx, scsAsId).Execute()

read all TMGI Allocation resource for a given SCS/AS

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
    scsAsId := "scsAsId_example" // string | Identifier of SCS/AS

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TMGIAllocationOperationApi.FetchAllTMGIAllocations(context.Background(), scsAsId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TMGIAllocationOperationApi.FetchAllTMGIAllocations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchAllTMGIAllocations`: TMGIAllocation
    fmt.Fprintf(os.Stdout, "Response from `TMGIAllocationOperationApi.FetchAllTMGIAllocations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of SCS/AS | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchAllTMGIAllocationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TMGIAllocation**](TMGIAllocation.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

