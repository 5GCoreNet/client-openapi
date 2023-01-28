# \RetrievalOfAssociatedIMSPrivateIdentitiesApi

All URIs are relative to *https://example.com/nhss-ims-sdm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetImsPrivateIds**](RetrievalOfAssociatedIMSPrivateIdentitiesApi.md#GetImsPrivateIds) | **Get** /{imsUeId}/identities/private-identities | Retrieve the associated private identities to the IMS public identity included in the service request 



## GetImsPrivateIds

> PrivateIdentities GetImsPrivateIds(ctx, imsUeId).SupportedFeatures(supportedFeatures).Impi(impi).Execute()

Retrieve the associated private identities to the IMS public identity included in the service request 

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
    imsUeId := "imsUeId_example" // string | IMS Public Identity
    supportedFeatures := "supportedFeatures_example" // string | Supported Features (optional)
    impi := "impi_example" // string | Private Identity of type IMPI (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RetrievalOfAssociatedIMSPrivateIdentitiesApi.GetImsPrivateIds(context.Background(), imsUeId).SupportedFeatures(supportedFeatures).Impi(impi).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RetrievalOfAssociatedIMSPrivateIdentitiesApi.GetImsPrivateIds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetImsPrivateIds`: PrivateIdentities
    fmt.Fprintf(os.Stdout, "Response from `RetrievalOfAssociatedIMSPrivateIdentitiesApi.GetImsPrivateIds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imsUeId** | **string** | IMS Public Identity | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetImsPrivateIdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **supportedFeatures** | **string** | Supported Features | 
 **impi** | **string** | Private Identity of type IMPI | 

### Return type

[**PrivateIdentities**](PrivateIdentities.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

