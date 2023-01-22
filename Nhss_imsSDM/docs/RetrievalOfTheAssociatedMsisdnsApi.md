# \RetrievalOfTheAssociatedMsisdnsApi

All URIs are relative to *https://example.com/nhss-ims-sdm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMsisdns**](RetrievalOfTheAssociatedMsisdnsApi.md#GetMsisdns) | **Get** /{imsUeId}/identities/msisdns | retrieve the Msisdns associated to requested identity



## GetMsisdns

> MsisdnList GetMsisdns(ctx, imsUeId).PrivateId(privateId).Execute()

retrieve the Msisdns associated to requested identity

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
    imsUeId := "imsUeId_example" // string | IMS Identity
    privateId := "privateId_example" // string | Private identity (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RetrievalOfTheAssociatedMsisdnsApi.GetMsisdns(context.Background(), imsUeId).PrivateId(privateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RetrievalOfTheAssociatedMsisdnsApi.GetMsisdns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsisdns`: MsisdnList
    fmt.Fprintf(os.Stdout, "Response from `RetrievalOfTheAssociatedMsisdnsApi.GetMsisdns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imsUeId** | **string** | IMS Identity | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMsisdnsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **privateId** | **string** | Private identity | 

### Return type

[**MsisdnList**](MsisdnList.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

