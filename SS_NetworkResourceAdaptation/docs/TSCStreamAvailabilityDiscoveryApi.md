# \TSCStreamAvailabilityDiscoveryApi

All URIs are relative to *https://example.com/ss-nra/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTscStreamAvailability**](TSCStreamAvailabilityDiscoveryApi.md#GetTscStreamAvailability) | **Get** /tsc-stream-availability | Discover the TSC stream availability information.



## GetTscStreamAvailability

> []TscStreamAvailability GetTscStreamAvailability(ctx).StreamSpecs(streamSpecs).Execute()

Discover the TSC stream availability information.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/SS_NetworkResourceAdaptation"
)

func main() {
    streamSpecs := []openapiclient.StreamSpecification{*openapiclient.NewStreamSpecification("SrcMacAddr_example", "DstMacAddr_example")} // []StreamSpecification | The MAC address(es) of the source DS-TT port(s) and the destination DS-TT port(s). 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TSCStreamAvailabilityDiscoveryApi.GetTscStreamAvailability(context.Background()).StreamSpecs(streamSpecs).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TSCStreamAvailabilityDiscoveryApi.GetTscStreamAvailability``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTscStreamAvailability`: []TscStreamAvailability
    fmt.Fprintf(os.Stdout, "Response from `TSCStreamAvailabilityDiscoveryApi.GetTscStreamAvailability`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTscStreamAvailabilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **streamSpecs** | [**[]StreamSpecification**](StreamSpecification.md) | The MAC address(es) of the source DS-TT port(s) and the destination DS-TT port(s).  | 

### Return type

[**[]TscStreamAvailability**](TscStreamAvailability.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

