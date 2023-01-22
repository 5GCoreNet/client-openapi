# \EASProfilesApi

All URIs are relative to *https://example.com/eees-easdiscovery/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EasProfilesRequestDiscoveryPost**](EASProfilesApi.md#EasProfilesRequestDiscoveryPost) | **Post** /eas-profiles/request-discovery | 



## EasProfilesRequestDiscoveryPost

> EasDiscoveryResp EasProfilesRequestDiscoveryPost(ctx).EasDiscoveryReq(easDiscoveryReq).Execute()





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
    easDiscoveryReq := *openapiclient.NewEasDiscoveryReq(openapiclient.RequestorId{Interface{}: new(interface{})}) // EasDiscoveryReq | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EASProfilesApi.EasProfilesRequestDiscoveryPost(context.Background()).EasDiscoveryReq(easDiscoveryReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EASProfilesApi.EasProfilesRequestDiscoveryPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EasProfilesRequestDiscoveryPost`: EasDiscoveryResp
    fmt.Fprintf(os.Stdout, "Response from `EASProfilesApi.EasProfilesRequestDiscoveryPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEasProfilesRequestDiscoveryPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **easDiscoveryReq** | [**EasDiscoveryReq**](EasDiscoveryReq.md) |  | 

### Return type

[**EasDiscoveryResp**](EasDiscoveryResp.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

