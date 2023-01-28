# \TMGIAllocationOrTimerExpiryRefreshApi

All URIs are relative to *https://example.com/3gpp-mbs-tmgi/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AllocateTmgi**](TMGIAllocationOrTimerExpiryRefreshApi.md#AllocateTmgi) | **Post** /allocate | Request the allocation of TMGI(s) for new MBS session(s) or the refresh of the expiry time of already allocated TMGI(s).



## AllocateTmgi

> TmgiAllocResponse AllocateTmgi(ctx).TmgiAllocRequest(tmgiAllocRequest).Execute()

Request the allocation of TMGI(s) for new MBS session(s) or the refresh of the expiry time of already allocated TMGI(s).

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/MBSTMGI"
)

func main() {
    tmgiAllocRequest := *openapiclient.NewTmgiAllocRequest("AfId_example", *openapiclient.NewTmgiAllocate()) // TmgiAllocRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TMGIAllocationOrTimerExpiryRefreshApi.AllocateTmgi(context.Background()).TmgiAllocRequest(tmgiAllocRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TMGIAllocationOrTimerExpiryRefreshApi.AllocateTmgi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AllocateTmgi`: TmgiAllocResponse
    fmt.Fprintf(os.Stdout, "Response from `TMGIAllocationOrTimerExpiryRefreshApi.AllocateTmgi`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAllocateTmgiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tmgiAllocRequest** | [**TmgiAllocRequest**](TmgiAllocRequest.md) |  | 

### Return type

[**TmgiAllocResponse**](TmgiAllocResponse.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

