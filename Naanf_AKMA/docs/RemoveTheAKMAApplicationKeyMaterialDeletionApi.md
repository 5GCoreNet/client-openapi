# \RemoveTheAKMAApplicationKeyMaterialDeletionApi

All URIs are relative to *https://example.com/naanf-akma/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RemoveContext**](RemoveTheAKMAApplicationKeyMaterialDeletionApi.md#RemoveContext) | **Post** /remove-context | Request to remove the AKMA related key material.



## RemoveContext

> RemoveContext(ctx).CtxRemove(ctxRemove).Execute()

Request to remove the AKMA related key material.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Naanf_AKMA"
)

func main() {
    ctxRemove := *openapiclient.NewCtxRemove() // CtxRemove | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RemoveTheAKMAApplicationKeyMaterialDeletionApi.RemoveContext(context.Background()).CtxRemove(ctxRemove).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemoveTheAKMAApplicationKeyMaterialDeletionApi.RemoveContext``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveContextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctxRemove** | [**CtxRemove**](CtxRemove.md) |  | 

### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

