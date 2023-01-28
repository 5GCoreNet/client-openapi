# \IMEISVRetrievalApi

All URIs are relative to *https://example.com/nhss-ims-sdm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetIMEISVInfo**](IMEISVRetrievalApi.md#GetIMEISVInfo) | **Get** /{imsUeId}/identities/imeisv | Retrieve the IMEISV information



## GetIMEISVInfo

> ImeiSvInformation GetIMEISVInfo(ctx, imsUeId).PrivateIdentity(privateIdentity).Execute()

Retrieve the IMEISV information

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nhss_imsSDM"
)

func main() {
    imsUeId := "imsUeId_example" // string | IMS Identity
    privateIdentity := "privateIdentity_example" // string | IMS Private Identity (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IMEISVRetrievalApi.GetIMEISVInfo(context.Background(), imsUeId).PrivateIdentity(privateIdentity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IMEISVRetrievalApi.GetIMEISVInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIMEISVInfo`: ImeiSvInformation
    fmt.Fprintf(os.Stdout, "Response from `IMEISVRetrievalApi.GetIMEISVInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imsUeId** | **string** | IMS Identity | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIMEISVInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **privateIdentity** | **string** | IMS Private Identity | 

### Return type

[**ImeiSvInformation**](ImeiSvInformation.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

