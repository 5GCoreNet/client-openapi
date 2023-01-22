# \N32FForwardApi

All URIs are relative to *https://example.com/n32f-forward/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**N32fProcessOptions**](N32FForwardApi.md#N32fProcessOptions) | **Options** /n32f-process | Discover communication options supported by next hop (IPX or SEPP)
[**PostN32fProcess**](N32FForwardApi.md#PostN32fProcess) | **Post** /n32f-process | N32-f Message Forwarding



## N32fProcessOptions

> N32fProcessOptions(ctx).Execute()

Discover communication options supported by next hop (IPX or SEPP)

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.N32FForwardApi.N32fProcessOptions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `N32FForwardApi.N32fProcessOptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiN32fProcessOptionsRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostN32fProcess

> N32fReformattedRspMsg PostN32fProcess(ctx).N32fReformattedReqMsg(n32fReformattedReqMsg).ContentEncoding(contentEncoding).AcceptEncoding(acceptEncoding).Var3gppSbiMessagePriority(var3gppSbiMessagePriority).Execute()

N32-f Message Forwarding

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
    n32fReformattedReqMsg := *openapiclient.NewN32fReformattedReqMsg(*openapiclient.NewFlatJweJson("Ciphertext_example")) // N32fReformattedReqMsg | Custom operation N32-f Message Forwarding
    contentEncoding := "contentEncoding_example" // string | Content-Encoding, described in IETF RFC 7231 (optional)
    acceptEncoding := "acceptEncoding_example" // string | Accept-Encoding, described in IETF RFC 7231 (optional)
    var3gppSbiMessagePriority := "var3gppSbiMessagePriority_example" // string | 3gpp-Sbi-Message-Priority, defined in 3GPP TS 29.500 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.N32FForwardApi.PostN32fProcess(context.Background()).N32fReformattedReqMsg(n32fReformattedReqMsg).ContentEncoding(contentEncoding).AcceptEncoding(acceptEncoding).Var3gppSbiMessagePriority(var3gppSbiMessagePriority).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `N32FForwardApi.PostN32fProcess``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostN32fProcess`: N32fReformattedRspMsg
    fmt.Fprintf(os.Stdout, "Response from `N32FForwardApi.PostN32fProcess`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostN32fProcessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **n32fReformattedReqMsg** | [**N32fReformattedReqMsg**](N32fReformattedReqMsg.md) | Custom operation N32-f Message Forwarding | 
 **contentEncoding** | **string** | Content-Encoding, described in IETF RFC 7231 | 
 **acceptEncoding** | **string** | Accept-Encoding, described in IETF RFC 7231 | 
 **var3gppSbiMessagePriority** | **string** | 3gpp-Sbi-Message-Priority, defined in 3GPP TS 29.500 | 

### Return type

[**N32fReformattedRspMsg**](N32fReformattedRspMsg.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

