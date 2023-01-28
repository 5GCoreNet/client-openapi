# \SCPDomainRoutingInformationSubscriptionsCollectionApi

All URIs are relative to *https://example.com/nnrf-disc/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ScpDomainRoutingInfoSubscribe**](SCPDomainRoutingInformationSubscriptionsCollectionApi.md#ScpDomainRoutingInfoSubscribe) | **Post** /scp-domain-routing-info-subs | Create a new subscription



## ScpDomainRoutingInfoSubscribe

> ScpDomainRoutingInfoSubscription ScpDomainRoutingInfoSubscribe(ctx).ScpDomainRoutingInfoSubscription(scpDomainRoutingInfoSubscription).ContentEncoding(contentEncoding).AcceptEncoding(acceptEncoding).Execute()

Create a new subscription

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nnrf_NFDiscovery"
)

func main() {
    scpDomainRoutingInfoSubscription := *openapiclient.NewScpDomainRoutingInfoSubscription("CallbackUri_example") // ScpDomainRoutingInfoSubscription | 
    contentEncoding := "contentEncoding_example" // string | Content-Encoding, described in IETF RFC 7231 (optional)
    acceptEncoding := "acceptEncoding_example" // string | Accept-Encoding, described in IETF RFC 7231 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SCPDomainRoutingInformationSubscriptionsCollectionApi.ScpDomainRoutingInfoSubscribe(context.Background()).ScpDomainRoutingInfoSubscription(scpDomainRoutingInfoSubscription).ContentEncoding(contentEncoding).AcceptEncoding(acceptEncoding).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SCPDomainRoutingInformationSubscriptionsCollectionApi.ScpDomainRoutingInfoSubscribe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ScpDomainRoutingInfoSubscribe`: ScpDomainRoutingInfoSubscription
    fmt.Fprintf(os.Stdout, "Response from `SCPDomainRoutingInformationSubscriptionsCollectionApi.ScpDomainRoutingInfoSubscribe`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiScpDomainRoutingInfoSubscribeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scpDomainRoutingInfoSubscription** | [**ScpDomainRoutingInfoSubscription**](ScpDomainRoutingInfoSubscription.md) |  | 
 **contentEncoding** | **string** | Content-Encoding, described in IETF RFC 7231 | 
 **acceptEncoding** | **string** | Accept-Encoding, described in IETF RFC 7231 | 

### Return type

[**ScpDomainRoutingInfoSubscription**](ScpDomainRoutingInfoSubscription.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

