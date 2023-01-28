# \BootstrappingRequestApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BootstrappingInfoRequest**](BootstrappingRequestApi.md#BootstrappingInfoRequest) | **Get** /bootstrapping | Bootstrapping Info Request



## BootstrappingInfoRequest

> BootstrappingInfo BootstrappingInfoRequest(ctx).IfNoneMatch(ifNoneMatch).Execute()

Bootstrapping Info Request

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nnrf_Bootstrapping"
)

func main() {
    ifNoneMatch := "ifNoneMatch_example" // string | Validator for conditional requests, as described in IETF RFC 7232, 3.2 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BootstrappingRequestApi.BootstrappingInfoRequest(context.Background()).IfNoneMatch(ifNoneMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BootstrappingRequestApi.BootstrappingInfoRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BootstrappingInfoRequest`: BootstrappingInfo
    fmt.Fprintf(os.Stdout, "Response from `BootstrappingRequestApi.BootstrappingInfoRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBootstrappingInfoRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ifNoneMatch** | **string** | Validator for conditional requests, as described in IETF RFC 7232, 3.2 | 

### Return type

[**BootstrappingInfo**](BootstrappingInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/3gppHal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

