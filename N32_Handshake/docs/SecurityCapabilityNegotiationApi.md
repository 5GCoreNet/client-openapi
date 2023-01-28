# \SecurityCapabilityNegotiationApi

All URIs are relative to *https://example.com/n32c-handshake/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostExchangeCapability**](SecurityCapabilityNegotiationApi.md#PostExchangeCapability) | **Post** /exchange-capability | Security Capability Negotiation



## PostExchangeCapability

> SecNegotiateRspData PostExchangeCapability(ctx).SecNegotiateReqData(secNegotiateReqData).Execute()

Security Capability Negotiation

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/N32_Handshake"
)

func main() {
    secNegotiateReqData := *openapiclient.NewSecNegotiateReqData("Sender_example", []openapiclient.SecurityCapability{*openapiclient.NewSecurityCapability()}) // SecNegotiateReqData | Custom operation for security capability negotiation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityCapabilityNegotiationApi.PostExchangeCapability(context.Background()).SecNegotiateReqData(secNegotiateReqData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityCapabilityNegotiationApi.PostExchangeCapability``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostExchangeCapability`: SecNegotiateRspData
    fmt.Fprintf(os.Stdout, "Response from `SecurityCapabilityNegotiationApi.PostExchangeCapability`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostExchangeCapabilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **secNegotiateReqData** | [**SecNegotiateReqData**](SecNegotiateReqData.md) | Custom operation for security capability negotiation | 

### Return type

[**SecNegotiateRspData**](SecNegotiateRspData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

