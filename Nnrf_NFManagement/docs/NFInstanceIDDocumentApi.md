# \NFInstanceIDDocumentApi

All URIs are relative to *https://example.com/nnrf-nfm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeregisterNFInstance**](NFInstanceIDDocumentApi.md#DeregisterNFInstance) | **Delete** /nf-instances/{nfInstanceID} | Deregisters a given NF Instance
[**GetNFInstance**](NFInstanceIDDocumentApi.md#GetNFInstance) | **Get** /nf-instances/{nfInstanceID} | Read the profile of a given NF Instance
[**RegisterNFInstance**](NFInstanceIDDocumentApi.md#RegisterNFInstance) | **Put** /nf-instances/{nfInstanceID} | Register a new NF Instance
[**UpdateNFInstance**](NFInstanceIDDocumentApi.md#UpdateNFInstance) | **Patch** /nf-instances/{nfInstanceID} | Update NF Instance profile



## DeregisterNFInstance

> DeregisterNFInstance(ctx, nfInstanceID).Execute()

Deregisters a given NF Instance

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nnrf_NFManagement"
)

func main() {
    nfInstanceID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique ID of the NF Instance to deregister

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NFInstanceIDDocumentApi.DeregisterNFInstance(context.Background(), nfInstanceID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NFInstanceIDDocumentApi.DeregisterNFInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nfInstanceID** | **string** | Unique ID of the NF Instance to deregister | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeregisterNFInstanceRequest struct via the builder pattern


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


## GetNFInstance

> NFProfile GetNFInstance(ctx, nfInstanceID).RequesterFeatures(requesterFeatures).Execute()

Read the profile of a given NF Instance

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nnrf_NFManagement"
)

func main() {
    nfInstanceID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique ID of the NF Instance
    requesterFeatures := "requesterFeatures_example" // string | Features supported by the NF Service Consumer (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NFInstanceIDDocumentApi.GetNFInstance(context.Background(), nfInstanceID).RequesterFeatures(requesterFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NFInstanceIDDocumentApi.GetNFInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNFInstance`: NFProfile
    fmt.Fprintf(os.Stdout, "Response from `NFInstanceIDDocumentApi.GetNFInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nfInstanceID** | **string** | Unique ID of the NF Instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNFInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requesterFeatures** | **string** | Features supported by the NF Service Consumer | 

### Return type

[**NFProfile**](NFProfile.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterNFInstance

> NFProfile RegisterNFInstance(ctx, nfInstanceID).NFProfile(nFProfile).ContentEncoding(contentEncoding).AcceptEncoding(acceptEncoding).Execute()

Register a new NF Instance

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nnrf_NFManagement"
)

func main() {
    nfInstanceID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique ID of the NF Instance to register
    nFProfile := *openapiclient.NewNFProfile("NfInstanceId_example", *openapiclient.NewNFType(), *openapiclient.NewNFStatus()) // NFProfile | 
    contentEncoding := "contentEncoding_example" // string | Content-Encoding, described in IETF RFC 7231 (optional)
    acceptEncoding := "acceptEncoding_example" // string | Accept-Encoding, described in IETF RFC 7231 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NFInstanceIDDocumentApi.RegisterNFInstance(context.Background(), nfInstanceID).NFProfile(nFProfile).ContentEncoding(contentEncoding).AcceptEncoding(acceptEncoding).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NFInstanceIDDocumentApi.RegisterNFInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegisterNFInstance`: NFProfile
    fmt.Fprintf(os.Stdout, "Response from `NFInstanceIDDocumentApi.RegisterNFInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nfInstanceID** | **string** | Unique ID of the NF Instance to register | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegisterNFInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nFProfile** | [**NFProfile**](NFProfile.md) |  | 
 **contentEncoding** | **string** | Content-Encoding, described in IETF RFC 7231 | 
 **acceptEncoding** | **string** | Accept-Encoding, described in IETF RFC 7231 | 

### Return type

[**NFProfile**](NFProfile.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNFInstance

> NFProfile UpdateNFInstance(ctx, nfInstanceID).PatchItem(patchItem).ContentEncoding(contentEncoding).AcceptEncoding(acceptEncoding).IfMatch(ifMatch).Execute()

Update NF Instance profile

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nnrf_NFManagement"
)

func main() {
    nfInstanceID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique ID of the NF Instance to update
    patchItem := []openapiclient.PatchItem{*openapiclient.NewPatchItem(*openapiclient.NewPatchOperation(), "Path_example")} // []PatchItem | 
    contentEncoding := "contentEncoding_example" // string | Content-Encoding, described in IETF RFC 7231 (optional)
    acceptEncoding := "acceptEncoding_example" // string | Accept-Encoding, described in IETF RFC 7231 (optional)
    ifMatch := "ifMatch_example" // string | Validator for conditional requests, as described in IETF RFC 7232, 3.2 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NFInstanceIDDocumentApi.UpdateNFInstance(context.Background(), nfInstanceID).PatchItem(patchItem).ContentEncoding(contentEncoding).AcceptEncoding(acceptEncoding).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NFInstanceIDDocumentApi.UpdateNFInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNFInstance`: NFProfile
    fmt.Fprintf(os.Stdout, "Response from `NFInstanceIDDocumentApi.UpdateNFInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nfInstanceID** | **string** | Unique ID of the NF Instance to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNFInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchItem** | [**[]PatchItem**](PatchItem.md) |  | 
 **contentEncoding** | **string** | Content-Encoding, described in IETF RFC 7231 | 
 **acceptEncoding** | **string** | Accept-Encoding, described in IETF RFC 7231 | 
 **ifMatch** | **string** | Validator for conditional requests, as described in IETF RFC 7232, 3.2 | 

### Return type

[**NFProfile**](NFProfile.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

