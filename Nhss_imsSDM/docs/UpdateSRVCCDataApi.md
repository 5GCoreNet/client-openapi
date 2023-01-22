# \UpdateSRVCCDataApi

All URIs are relative to *https://example.com/nhss-ims-sdm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateSrvccData**](UpdateSRVCCDataApi.md#UpdateSrvccData) | **Patch** /{imsUeId}/srvcc-data | Patch



## UpdateSrvccData

> PatchResult UpdateSrvccData(ctx, imsUeId).PatchItem(patchItem).SupportedFeatures(supportedFeatures).PrivateIdentity(privateIdentity).Execute()

Patch

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
    imsUeId := "imsUeId_example" // string | IMS Public Identity or IMS Private Identity
    patchItem := []openapiclient.PatchItem{*openapiclient.NewPatchItem(*openapiclient.NewPatchOperation(), "Path_example")} // []PatchItem | 
    supportedFeatures := "supportedFeatures_example" // string | Supported Features (optional)
    privateIdentity := "privateIdentity_example" // string | IMS Private Identity (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UpdateSRVCCDataApi.UpdateSrvccData(context.Background(), imsUeId).PatchItem(patchItem).SupportedFeatures(supportedFeatures).PrivateIdentity(privateIdentity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UpdateSRVCCDataApi.UpdateSrvccData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSrvccData`: PatchResult
    fmt.Fprintf(os.Stdout, "Response from `UpdateSRVCCDataApi.UpdateSrvccData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imsUeId** | **string** | IMS Public Identity or IMS Private Identity | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSrvccDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchItem** | [**[]PatchItem**](PatchItem.md) |  | 
 **supportedFeatures** | **string** | Supported Features | 
 **privateIdentity** | **string** | IMS Private Identity | 

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

