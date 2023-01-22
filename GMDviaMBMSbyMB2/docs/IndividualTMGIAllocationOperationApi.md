# \IndividualTMGIAllocationOperationApi

All URIs are relative to *https://example.com/3gpp-group-message-delivery-mb2/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteTMGIAllocation**](IndividualTMGIAllocationOperationApi.md#DeleteTMGIAllocation) | **Delete** /{scsAsId}/tmgi-allocation/{tmgi} | Deletes an existing TMGI Allocation resource for a given SCS/AS and a TMGI.
[**FetchIndTMGIAllocation**](IndividualTMGIAllocationOperationApi.md#FetchIndTMGIAllocation) | **Get** /{scsAsId}/tmgi-allocation/{tmgi} | Read a TMGI Allocation resource for a given SCS/AS and a TMGI.
[**ModifyIndTMGIAllocation**](IndividualTMGIAllocationOperationApi.md#ModifyIndTMGIAllocation) | **Patch** /{scsAsId}/tmgi-allocation/{tmgi} | Updates an existing TMGI Allocation resource for a given SCS/AS and a TMGI.
[**UpdateIndTMGIAllocation**](IndividualTMGIAllocationOperationApi.md#UpdateIndTMGIAllocation) | **Put** /{scsAsId}/tmgi-allocation/{tmgi} | Updates an existing TMGI Allocation resource for a given SCS/AS and a TMGI.



## DeleteTMGIAllocation

> DeleteTMGIAllocation(ctx, scsAsId, tmgi).Execute()

Deletes an existing TMGI Allocation resource for a given SCS/AS and a TMGI.

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
    tmgi := "tmgi_example" // string | TMGI

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualTMGIAllocationOperationApi.DeleteTMGIAllocation(context.Background(), scsAsId, tmgi).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualTMGIAllocationOperationApi.DeleteTMGIAllocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of SCS/AS | 
**tmgi** | **string** | TMGI | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTMGIAllocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchIndTMGIAllocation

> TMGIAllocation FetchIndTMGIAllocation(ctx, scsAsId, tmgi).Execute()

Read a TMGI Allocation resource for a given SCS/AS and a TMGI.

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
    tmgi := "tmgi_example" // string | TMGI

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualTMGIAllocationOperationApi.FetchIndTMGIAllocation(context.Background(), scsAsId, tmgi).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualTMGIAllocationOperationApi.FetchIndTMGIAllocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchIndTMGIAllocation`: TMGIAllocation
    fmt.Fprintf(os.Stdout, "Response from `IndividualTMGIAllocationOperationApi.FetchIndTMGIAllocation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of SCS/AS | 
**tmgi** | **string** | TMGI | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchIndTMGIAllocationRequest struct via the builder pattern


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


## ModifyIndTMGIAllocation

> TMGIAllocation ModifyIndTMGIAllocation(ctx, scsAsId, tmgi).TMGIAllocationPatch(tMGIAllocationPatch).Execute()

Updates an existing TMGI Allocation resource for a given SCS/AS and a TMGI.

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
    tmgi := "tmgi_example" // string | TMGI
    tMGIAllocationPatch := *openapiclient.NewTMGIAllocationPatch() // TMGIAllocationPatch | representation of the TMGI Allocation to be updated in the SCEF

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualTMGIAllocationOperationApi.ModifyIndTMGIAllocation(context.Background(), scsAsId, tmgi).TMGIAllocationPatch(tMGIAllocationPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualTMGIAllocationOperationApi.ModifyIndTMGIAllocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ModifyIndTMGIAllocation`: TMGIAllocation
    fmt.Fprintf(os.Stdout, "Response from `IndividualTMGIAllocationOperationApi.ModifyIndTMGIAllocation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of SCS/AS | 
**tmgi** | **string** | TMGI | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyIndTMGIAllocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **tMGIAllocationPatch** | [**TMGIAllocationPatch**](TMGIAllocationPatch.md) | representation of the TMGI Allocation to be updated in the SCEF | 

### Return type

[**TMGIAllocation**](TMGIAllocation.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIndTMGIAllocation

> TMGIAllocation UpdateIndTMGIAllocation(ctx, scsAsId, tmgi).TMGIAllocation(tMGIAllocation).Execute()

Updates an existing TMGI Allocation resource for a given SCS/AS and a TMGI.

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
    tmgi := "tmgi_example" // string | TMGI
    tMGIAllocation := *openapiclient.NewTMGIAllocation() // TMGIAllocation | representation of the TMGI Allocation to be updated in the SCEF

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualTMGIAllocationOperationApi.UpdateIndTMGIAllocation(context.Background(), scsAsId, tmgi).TMGIAllocation(tMGIAllocation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualTMGIAllocationOperationApi.UpdateIndTMGIAllocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIndTMGIAllocation`: TMGIAllocation
    fmt.Fprintf(os.Stdout, "Response from `IndividualTMGIAllocationOperationApi.UpdateIndTMGIAllocation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of SCS/AS | 
**tmgi** | **string** | TMGI | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIndTMGIAllocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **tMGIAllocation** | [**TMGIAllocation**](TMGIAllocation.md) | representation of the TMGI Allocation to be updated in the SCEF | 

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

