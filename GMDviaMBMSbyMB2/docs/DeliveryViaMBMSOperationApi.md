# \DeliveryViaMBMSOperationApi

All URIs are relative to *https://example.com/3gpp-group-message-delivery-mb2/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGMDViaMBMSByMB2**](DeliveryViaMBMSOperationApi.md#CreateGMDViaMBMSByMB2) | **Post** /{scsAsId}/tmgi-allocation/{tmgi}/delivery-via-mbms | Creates a new delivery via MBMS for a given SCS/AS and a TMGI.
[**FecthAllGMDViaMBMSByMB2**](DeliveryViaMBMSOperationApi.md#FecthAllGMDViaMBMSByMB2) | **Get** /{scsAsId}/tmgi-allocation/{tmgi}/delivery-via-mbms | Read all group message delivery via MBMS resource for a given SCS/AS and a TMGI.



## CreateGMDViaMBMSByMB2

> GMDViaMBMSByMb2 CreateGMDViaMBMSByMB2(ctx, scsAsId, tmgi).GMDViaMBMSByMb2(gMDViaMBMSByMb2).Execute()

Creates a new delivery via MBMS for a given SCS/AS and a TMGI.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/GMDviaMBMSbyMB2"
)

func main() {
    scsAsId := "scsAsId_example" // string | Identifier of SCS/AS
    tmgi := "tmgi_example" // string | TMGI
    gMDViaMBMSByMb2 := *openapiclient.NewGMDViaMBMSByMb2("NotificationDestination_example") // GMDViaMBMSByMb2 | representation of the GMD via MBMS by MB2 resource to be Created in the SCEF

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryViaMBMSOperationApi.CreateGMDViaMBMSByMB2(context.Background(), scsAsId, tmgi).GMDViaMBMSByMb2(gMDViaMBMSByMb2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryViaMBMSOperationApi.CreateGMDViaMBMSByMB2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateGMDViaMBMSByMB2`: GMDViaMBMSByMb2
    fmt.Fprintf(os.Stdout, "Response from `DeliveryViaMBMSOperationApi.CreateGMDViaMBMSByMB2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of SCS/AS | 
**tmgi** | **string** | TMGI | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGMDViaMBMSByMB2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **gMDViaMBMSByMb2** | [**GMDViaMBMSByMb2**](GMDViaMBMSByMb2.md) | representation of the GMD via MBMS by MB2 resource to be Created in the SCEF | 

### Return type

[**GMDViaMBMSByMb2**](GMDViaMBMSByMb2.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FecthAllGMDViaMBMSByMB2

> GMDViaMBMSByMb2 FecthAllGMDViaMBMSByMB2(ctx, scsAsId, tmgi).Execute()

Read all group message delivery via MBMS resource for a given SCS/AS and a TMGI.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/GMDviaMBMSbyMB2"
)

func main() {
    scsAsId := "scsAsId_example" // string | Identifier of SCS/AS
    tmgi := "tmgi_example" // string | TMGI

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryViaMBMSOperationApi.FecthAllGMDViaMBMSByMB2(context.Background(), scsAsId, tmgi).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryViaMBMSOperationApi.FecthAllGMDViaMBMSByMB2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FecthAllGMDViaMBMSByMB2`: GMDViaMBMSByMb2
    fmt.Fprintf(os.Stdout, "Response from `DeliveryViaMBMSOperationApi.FecthAllGMDViaMBMSByMB2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of SCS/AS | 
**tmgi** | **string** | TMGI | 

### Other Parameters

Other parameters are passed through a pointer to a apiFecthAllGMDViaMBMSByMB2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GMDViaMBMSByMb2**](GMDViaMBMSByMb2.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

