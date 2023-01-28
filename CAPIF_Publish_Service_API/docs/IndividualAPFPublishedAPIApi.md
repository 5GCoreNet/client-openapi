# \IndividualAPFPublishedAPIApi

All URIs are relative to *https://example.com/published-apis/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ModifyIndAPFPubAPI**](IndividualAPFPublishedAPIApi.md#ModifyIndAPFPubAPI) | **Patch** /{apfId}/service-apis/{serviceApiId} | 



## ModifyIndAPFPubAPI

> ServiceAPIDescription ModifyIndAPFPubAPI(ctx, serviceApiId, apfId).ServiceAPIDescriptionPatch(serviceAPIDescriptionPatch).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/CAPIF_Publish_Service_API"
)

func main() {
    serviceApiId := "serviceApiId_example" // string | 
    apfId := "apfId_example" // string | 
    serviceAPIDescriptionPatch := *openapiclient.NewServiceAPIDescriptionPatch() // ServiceAPIDescriptionPatch | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualAPFPublishedAPIApi.ModifyIndAPFPubAPI(context.Background(), serviceApiId, apfId).ServiceAPIDescriptionPatch(serviceAPIDescriptionPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualAPFPublishedAPIApi.ModifyIndAPFPubAPI``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ModifyIndAPFPubAPI`: ServiceAPIDescription
    fmt.Fprintf(os.Stdout, "Response from `IndividualAPFPublishedAPIApi.ModifyIndAPFPubAPI`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceApiId** | **string** |  | 
**apfId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyIndAPFPubAPIRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **serviceAPIDescriptionPatch** | [**ServiceAPIDescriptionPatch**](ServiceAPIDescriptionPatch.md) |  | 

### Return type

[**ServiceAPIDescription**](ServiceAPIDescription.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

