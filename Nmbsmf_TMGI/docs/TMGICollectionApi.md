# \TMGICollectionApi

All URIs are relative to *https://example.com/nmbsmf-tmgi/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AllocateTmgi**](TMGICollectionApi.md#AllocateTmgi) | **Post** /tmgi | Allocate TMGIs
[**TMGIDeallocate**](TMGICollectionApi.md#TMGIDeallocate) | **Delete** /tmgi | Deallocate one or more TMGIs



## AllocateTmgi

> TmgiAllocated AllocateTmgi(ctx).TmgiAllocate(tmgiAllocate).Execute()

Allocate TMGIs

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
    tmgiAllocate := *openapiclient.NewTmgiAllocate() // TmgiAllocate | representation of the TMGIs to be created in the MB-SMF

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TMGICollectionApi.AllocateTmgi(context.Background()).TmgiAllocate(tmgiAllocate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TMGICollectionApi.AllocateTmgi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AllocateTmgi`: TmgiAllocated
    fmt.Fprintf(os.Stdout, "Response from `TMGICollectionApi.AllocateTmgi`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAllocateTmgiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tmgiAllocate** | [**TmgiAllocate**](TmgiAllocate.md) | representation of the TMGIs to be created in the MB-SMF | 

### Return type

[**TmgiAllocated**](TmgiAllocated.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TMGIDeallocate

> TMGIDeallocate(ctx).TmgiList(tmgiList).Execute()

Deallocate one or more TMGIs

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
    tmgiList := []openapiclient.Tmgi{*openapiclient.NewTmgi("MbsServiceId_example", *openapiclient.NewPlmnId("Mcc_example", "Mnc_example"))} // []Tmgi | One of more TMGIs to be deallocated (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TMGICollectionApi.TMGIDeallocate(context.Background()).TmgiList(tmgiList).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TMGICollectionApi.TMGIDeallocate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTMGIDeallocateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tmgiList** | [**[]Tmgi**](Tmgi.md) | One of more TMGIs to be deallocated | 

### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

