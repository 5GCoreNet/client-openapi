# \IndividualBaselineDNSPatternApi

All URIs are relative to *https://example.com/neasdf-baselinednspattern/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrReplaceBaseDnsPattern**](IndividualBaselineDNSPatternApi.md#CreateOrReplaceBaseDnsPattern) | **Put** /base-dns-patterns/{smfId}/{smfImplementationSegmentPaths} | Creates or Updates the Baseline DNS Pattern (complete replacement)
[**DeleteBaseDnsPattern**](IndividualBaselineDNSPatternApi.md#DeleteBaseDnsPattern) | **Delete** /base-dns-patterns/{smfId}/{smfImplementationSegmentPaths} | Deletes a Baseline DNS Pattern
[**UpdateBaseDNSPattern**](IndividualBaselineDNSPatternApi.md#UpdateBaseDNSPattern) | **Patch** /base-dns-patterns/{smfId}/{smfImplementationSegmentPaths} | Updates the Baseline DNS Pattern



## CreateOrReplaceBaseDnsPattern

> BaseDnsPatternCreatedData CreateOrReplaceBaseDnsPattern(ctx, smfId, smfImplementationSegmentPaths).BaseDnsPatternCreateData(baseDnsPatternCreateData).ContentEncoding(contentEncoding).Execute()

Creates or Updates the Baseline DNS Pattern (complete replacement)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Neasdf_BaselineDNSPattern"
)

func main() {
    smfId := map[string][]openapiclient.VarNfId{ ... } // VarNfId | SMF or SMF set identifier or Set Id part in SMF set identifier
    smfImplementationSegmentPaths := "smfImplementationSegmentPaths_example" // string | SMF Implementation Dependent Segment Paths
    baseDnsPatternCreateData := *openapiclient.NewBaseDnsPatternCreateData() // BaseDnsPatternCreateData | 
    contentEncoding := "contentEncoding_example" // string | Content-Encoding, described in IETF RFC 7231 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualBaselineDNSPatternApi.CreateOrReplaceBaseDnsPattern(context.Background(), smfId, smfImplementationSegmentPaths).BaseDnsPatternCreateData(baseDnsPatternCreateData).ContentEncoding(contentEncoding).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualBaselineDNSPatternApi.CreateOrReplaceBaseDnsPattern``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrReplaceBaseDnsPattern`: BaseDnsPatternCreatedData
    fmt.Fprintf(os.Stdout, "Response from `IndividualBaselineDNSPatternApi.CreateOrReplaceBaseDnsPattern`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**smfId** | [**VarNfId**](.md) | SMF or SMF set identifier or Set Id part in SMF set identifier | 
**smfImplementationSegmentPaths** | **string** | SMF Implementation Dependent Segment Paths | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrReplaceBaseDnsPatternRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **baseDnsPatternCreateData** | [**BaseDnsPatternCreateData**](BaseDnsPatternCreateData.md) |  | 
 **contentEncoding** | **string** | Content-Encoding, described in IETF RFC 7231 | 

### Return type

[**BaseDnsPatternCreatedData**](BaseDnsPatternCreatedData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBaseDnsPattern

> DeleteBaseDnsPattern(ctx, smfId, smfImplementationSegmentPaths).Execute()

Deletes a Baseline DNS Pattern

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Neasdf_BaselineDNSPattern"
)

func main() {
    smfId := map[string][]openapiclient.VarNfId{ ... } // VarNfId | SMF or SMF set identifier or Set Id part in SMF set identifier
    smfImplementationSegmentPaths := "smfImplementationSegmentPaths_example" // string | SMF Implementation Dependent Segment Paths

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualBaselineDNSPatternApi.DeleteBaseDnsPattern(context.Background(), smfId, smfImplementationSegmentPaths).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualBaselineDNSPatternApi.DeleteBaseDnsPattern``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**smfId** | [**VarNfId**](.md) | SMF or SMF set identifier or Set Id part in SMF set identifier | 
**smfImplementationSegmentPaths** | **string** | SMF Implementation Dependent Segment Paths | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBaseDnsPatternRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBaseDNSPattern

> PatchResult UpdateBaseDNSPattern(ctx, smfId, smfImplementationSegmentPaths).PatchItem(patchItem).ContentEncoding(contentEncoding).Execute()

Updates the Baseline DNS Pattern

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Neasdf_BaselineDNSPattern"
)

func main() {
    smfId := map[string][]openapiclient.VarNfId{ ... } // VarNfId | SMF or SMF set identifier or Set Id part in SMF set identifier
    smfImplementationSegmentPaths := "smfImplementationSegmentPaths_example" // string | SMF Implementation Dependent Segment Paths
    patchItem := []openapiclient.PatchItem{*openapiclient.NewPatchItem(*openapiclient.NewPatchOperation(), "Path_example")} // []PatchItem | 
    contentEncoding := "contentEncoding_example" // string | Content-Encoding, described in IETF RFC 7231 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualBaselineDNSPatternApi.UpdateBaseDNSPattern(context.Background(), smfId, smfImplementationSegmentPaths).PatchItem(patchItem).ContentEncoding(contentEncoding).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualBaselineDNSPatternApi.UpdateBaseDNSPattern``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBaseDNSPattern`: PatchResult
    fmt.Fprintf(os.Stdout, "Response from `IndividualBaselineDNSPatternApi.UpdateBaseDNSPattern`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**smfId** | [**VarNfId**](.md) | SMF or SMF set identifier or Set Id part in SMF set identifier | 
**smfImplementationSegmentPaths** | **string** | SMF Implementation Dependent Segment Paths | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBaseDNSPatternRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchItem** | [**[]PatchItem**](PatchItem.md) |  | 
 **contentEncoding** | **string** | Content-Encoding, described in IETF RFC 7231 | 

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

