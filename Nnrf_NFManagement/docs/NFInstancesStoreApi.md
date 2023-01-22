# \NFInstancesStoreApi

All URIs are relative to *https://example.com/nnrf-nfm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNFInstances**](NFInstancesStoreApi.md#GetNFInstances) | **Get** /nf-instances | Retrieves a collection of NF Instances
[**OptionsNFInstances**](NFInstancesStoreApi.md#OptionsNFInstances) | **Options** /nf-instances | Discover communication options supported by NRF for NF Instances



## GetNFInstances

> UriList GetNFInstances(ctx).NfType(nfType).Limit(limit).PageNumber(pageNumber).PageSize(pageSize).Execute()

Retrieves a collection of NF Instances

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
    nfType := *openapiclient.NewNFType() // NFType | Type of NF (optional)
    limit := int32(56) // int32 | How many items to return at one time (optional)
    pageNumber := int32(56) // int32 | Page number where the response shall start (optional)
    pageSize := int32(56) // int32 | Maximum number of items in each returned page (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NFInstancesStoreApi.GetNFInstances(context.Background()).NfType(nfType).Limit(limit).PageNumber(pageNumber).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NFInstancesStoreApi.GetNFInstances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNFInstances`: UriList
    fmt.Fprintf(os.Stdout, "Response from `NFInstancesStoreApi.GetNFInstances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNFInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nfType** | [**NFType**](NFType.md) | Type of NF | 
 **limit** | **int32** | How many items to return at one time | 
 **pageNumber** | **int32** | Page number where the response shall start | 
 **pageSize** | **int32** | Maximum number of items in each returned page | 

### Return type

[**UriList**](UriList.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/3gppHal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OptionsNFInstances

> OptionsResponse OptionsNFInstances(ctx).Execute()

Discover communication options supported by NRF for NF Instances

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
    resp, r, err := apiClient.NFInstancesStoreApi.OptionsNFInstances(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NFInstancesStoreApi.OptionsNFInstances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OptionsNFInstances`: OptionsResponse
    fmt.Fprintf(os.Stdout, "Response from `NFInstancesStoreApi.OptionsNFInstances`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOptionsNFInstancesRequest struct via the builder pattern


### Return type

[**OptionsResponse**](OptionsResponse.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

