# \BootstrappingInfoRetrievalCustomOperationApi

All URIs are relative to *https://example.com/nbsp-gba/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BootstrappingInfoRetrieval**](BootstrappingInfoRetrievalCustomOperationApi.md#BootstrappingInfoRetrieval) | **Post** /bootstrapping-info-retrieval | Retrieve Bootstrapping Info from GBA BSF from NAF



## BootstrappingInfoRetrieval

> BootstrappingInfoResponse BootstrappingInfoRetrieval(ctx).BootstrappingInfoRequest(bootstrappingInfoRequest).Execute()

Retrieve Bootstrapping Info from GBA BSF from NAF

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
    bootstrappingInfoRequest := *openapiclient.NewBootstrappingInfoRequest("BtId_example", *openapiclient.NewNafId("NafFqdn_example", "UaSecProtId_example")) // BootstrappingInfoRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BootstrappingInfoRetrievalCustomOperationApi.BootstrappingInfoRetrieval(context.Background()).BootstrappingInfoRequest(bootstrappingInfoRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BootstrappingInfoRetrievalCustomOperationApi.BootstrappingInfoRetrieval``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BootstrappingInfoRetrieval`: BootstrappingInfoResponse
    fmt.Fprintf(os.Stdout, "Response from `BootstrappingInfoRetrievalCustomOperationApi.BootstrappingInfoRetrieval`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBootstrappingInfoRetrievalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bootstrappingInfoRequest** | [**BootstrappingInfoRequest**](BootstrappingInfoRequest.md) |  | 

### Return type

[**BootstrappingInfoResponse**](BootstrappingInfoResponse.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

