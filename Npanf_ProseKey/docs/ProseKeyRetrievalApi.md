# \ProseKeyRetrievalApi

All URIs are relative to *https://example.com/npanf-prosekey/&lt;apiVersion&gt;*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProseKeyRetrieval**](ProseKeyRetrievalApi.md#ProseKeyRetrieval) | **Post** /prose-keys/retrieve | retrieve the prose key



## ProseKeyRetrieval

> ProseKeyResponse ProseKeyRetrieval(ctx).ProseKeyRequest(proseKeyRequest).Execute()

retrieve the prose key

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Npanf_ProseKey"
)

func main() {
    proseKeyRequest := *openapiclient.NewProseKeyRequest("Var5gPrukId_example", int32(123)) // ProseKeyRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProseKeyRetrievalApi.ProseKeyRetrieval(context.Background()).ProseKeyRequest(proseKeyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProseKeyRetrievalApi.ProseKeyRetrieval``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProseKeyRetrieval`: ProseKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `ProseKeyRetrievalApi.ProseKeyRetrieval`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProseKeyRetrievalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **proseKeyRequest** | [**ProseKeyRequest**](ProseKeyRequest.md) |  | 

### Return type

[**ProseKeyResponse**](ProseKeyResponse.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

