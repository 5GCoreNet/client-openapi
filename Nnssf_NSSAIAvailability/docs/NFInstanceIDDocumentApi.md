# \NFInstanceIDDocumentApi

All URIs are relative to *https://example.com/nnssf-nssaiavailability/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NSSAIAvailabilityDelete**](NFInstanceIDDocumentApi.md#NSSAIAvailabilityDelete) | **Delete** /nssai-availability/{nfId} | Deletes an already existing S-NSSAIs per TA provided by the NF service consumer (e.g AMF)
[**NSSAIAvailabilityPatch**](NFInstanceIDDocumentApi.md#NSSAIAvailabilityPatch) | **Patch** /nssai-availability/{nfId} | Updates an already existing S-NSSAIs per TA provided by the NF service consumer (e.g AMF)
[**NSSAIAvailabilityPut**](NFInstanceIDDocumentApi.md#NSSAIAvailabilityPut) | **Put** /nssai-availability/{nfId} | Updates/replaces the NSSF with the S-NSSAIs the NF service consumer (e.g AMF)supports per TA



## NSSAIAvailabilityDelete

> NSSAIAvailabilityDelete(ctx, nfId).Execute()

Deletes an already existing S-NSSAIs per TA provided by the NF service consumer (e.g AMF)

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
    nfId := "nfId_example" // string | Identifier of the NF service consumer instance

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NFInstanceIDDocumentApi.NSSAIAvailabilityDelete(context.Background(), nfId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NFInstanceIDDocumentApi.NSSAIAvailabilityDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nfId** | **string** | Identifier of the NF service consumer instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiNSSAIAvailabilityDeleteRequest struct via the builder pattern


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


## NSSAIAvailabilityPatch

> AuthorizedNssaiAvailabilityInfo NSSAIAvailabilityPatch(ctx, nfId).PatchItem(patchItem).ContentEncoding(contentEncoding).AcceptEncoding(acceptEncoding).Execute()

Updates an already existing S-NSSAIs per TA provided by the NF service consumer (e.g AMF)

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
    nfId := "nfId_example" // string | Identifier of the NF service consumer instance
    patchItem := []openapiclient.PatchItem{*openapiclient.NewPatchItem(*openapiclient.NewPatchOperation(), "Path_example")} // []PatchItem | JSON Patch instructions to update at the NSSF, the S-NSSAIs supported per TA
    contentEncoding := "contentEncoding_example" // string | Content-Encoding, described in IETF RFC 7231 (optional)
    acceptEncoding := "acceptEncoding_example" // string | Accept-Encoding, described in IETF RFC 7231 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NFInstanceIDDocumentApi.NSSAIAvailabilityPatch(context.Background(), nfId).PatchItem(patchItem).ContentEncoding(contentEncoding).AcceptEncoding(acceptEncoding).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NFInstanceIDDocumentApi.NSSAIAvailabilityPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NSSAIAvailabilityPatch`: AuthorizedNssaiAvailabilityInfo
    fmt.Fprintf(os.Stdout, "Response from `NFInstanceIDDocumentApi.NSSAIAvailabilityPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nfId** | **string** | Identifier of the NF service consumer instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiNSSAIAvailabilityPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchItem** | [**[]PatchItem**](PatchItem.md) | JSON Patch instructions to update at the NSSF, the S-NSSAIs supported per TA | 
 **contentEncoding** | **string** | Content-Encoding, described in IETF RFC 7231 | 
 **acceptEncoding** | **string** | Accept-Encoding, described in IETF RFC 7231 | 

### Return type

[**AuthorizedNssaiAvailabilityInfo**](AuthorizedNssaiAvailabilityInfo.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json-patch+json:
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NSSAIAvailabilityPut

> AuthorizedNssaiAvailabilityInfo NSSAIAvailabilityPut(ctx, nfId).NssaiAvailabilityInfo(nssaiAvailabilityInfo).ContentEncoding(contentEncoding).AcceptEncoding(acceptEncoding).Execute()

Updates/replaces the NSSF with the S-NSSAIs the NF service consumer (e.g AMF)supports per TA

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
    nfId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Identifier of the NF service consumer instance
    nssaiAvailabilityInfo := *openapiclient.NewNssaiAvailabilityInfo([]openapiclient.SupportedNssaiAvailabilityData{*openapiclient.NewSupportedNssaiAvailabilityData(*openapiclient.NewTai(*openapiclient.NewPlmnId("Mcc_example", "Mnc_example"), "Tac_example"), []openapiclient.ExtSnssai{*openapiclient.NewExtSnssai(int32(123))})}) // NssaiAvailabilityInfo | Parameters to update/replace at the NSSF, the S-NSSAIs supported per TA
    contentEncoding := "contentEncoding_example" // string | Content-Encoding, described in IETF RFC 7231 (optional)
    acceptEncoding := "acceptEncoding_example" // string | Accept-Encoding, described in IETF RFC 7231 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NFInstanceIDDocumentApi.NSSAIAvailabilityPut(context.Background(), nfId).NssaiAvailabilityInfo(nssaiAvailabilityInfo).ContentEncoding(contentEncoding).AcceptEncoding(acceptEncoding).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NFInstanceIDDocumentApi.NSSAIAvailabilityPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NSSAIAvailabilityPut`: AuthorizedNssaiAvailabilityInfo
    fmt.Fprintf(os.Stdout, "Response from `NFInstanceIDDocumentApi.NSSAIAvailabilityPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nfId** | **string** | Identifier of the NF service consumer instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiNSSAIAvailabilityPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nssaiAvailabilityInfo** | [**NssaiAvailabilityInfo**](NssaiAvailabilityInfo.md) | Parameters to update/replace at the NSSF, the S-NSSAIs supported per TA | 
 **contentEncoding** | **string** | Content-Encoding, described in IETF RFC 7231 | 
 **acceptEncoding** | **string** | Accept-Encoding, described in IETF RFC 7231 | 

### Return type

[**AuthorizedNssaiAvailabilityInfo**](AuthorizedNssaiAvailabilityInfo.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

