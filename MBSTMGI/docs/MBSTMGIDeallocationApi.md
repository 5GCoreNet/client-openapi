# \MBSTMGIDeallocationApi

All URIs are relative to *https://example.com/3gpp-mbs-tmgi/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeallocateTmgi**](MBSTMGIDeallocationApi.md#DeallocateTmgi) | **Post** /deallocate | Request the deallocation of MBS TMGI(s).



## DeallocateTmgi

> DeallocateTmgi(ctx).TmgiDeallocRequest(tmgiDeallocRequest).Execute()

Request the deallocation of MBS TMGI(s).

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
    tmgiDeallocRequest := *openapiclient.NewTmgiDeallocRequest("AfId_example", []openapiclient.Tmgi{*openapiclient.NewTmgi("MbsServiceId_example", *openapiclient.NewPlmnId("Mcc_example", "Mnc_example"))}) // TmgiDeallocRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MBSTMGIDeallocationApi.DeallocateTmgi(context.Background()).TmgiDeallocRequest(tmgiDeallocRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MBSTMGIDeallocationApi.DeallocateTmgi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeallocateTmgiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tmgiDeallocRequest** | [**TmgiDeallocRequest**](TmgiDeallocRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

