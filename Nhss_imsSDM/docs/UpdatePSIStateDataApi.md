# \UpdatePSIStateDataApi

All URIs are relative to *https://example.com/nhss-ims-sdm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdatePsiState**](UpdatePSIStateDataApi.md#UpdatePsiState) | **Patch** /{imsUeId}/service-data/psi-status | Patch



## UpdatePsiState

> PatchResult UpdatePsiState(ctx, imsUeId).PatchItem(patchItem).SupportedFeatures(supportedFeatures).Execute()

Patch

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nhss_imsSDM"
)

func main() {
    imsUeId := "imsUeId_example" // string | IMS Public Service Identity
    patchItem := []openapiclient.PatchItem{*openapiclient.NewPatchItem(*openapiclient.NewPatchOperation(), "Path_example")} // []PatchItem | 
    supportedFeatures := "supportedFeatures_example" // string | Supported Features (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UpdatePSIStateDataApi.UpdatePsiState(context.Background(), imsUeId).PatchItem(patchItem).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UpdatePSIStateDataApi.UpdatePsiState``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePsiState`: PatchResult
    fmt.Fprintf(os.Stdout, "Response from `UpdatePSIStateDataApi.UpdatePsiState`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imsUeId** | **string** | IMS Public Service Identity | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePsiStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchItem** | [**[]PatchItem**](PatchItem.md) |  | 
 **supportedFeatures** | **string** | Supported Features | 

### Return type

[**PatchResult**](PatchResult.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

