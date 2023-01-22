# \ReferenceLocationInfoRetrievalApi

All URIs are relative to *https://example.com/nhss-ims-sdm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetReferenceLocationInfo**](ReferenceLocationInfoRetrievalApi.md#GetReferenceLocationInfo) | **Get** /{imsUeId}/access-data/wireline-domain/reference-location | Retrieve the reference location information



## GetReferenceLocationInfo

> ReferenceLocationInformation GetReferenceLocationInfo(ctx, imsUeId).SupportedFeatures(supportedFeatures).PrivateIdentity(privateIdentity).Execute()

Retrieve the reference location information

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
    supportedFeatures := "supportedFeatures_example" // string | Supported Features (optional)
    privateIdentity := "privateIdentity_example" // string | IMS Private Identity (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReferenceLocationInfoRetrievalApi.GetReferenceLocationInfo(context.Background(), imsUeId).SupportedFeatures(supportedFeatures).PrivateIdentity(privateIdentity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReferenceLocationInfoRetrievalApi.GetReferenceLocationInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReferenceLocationInfo`: ReferenceLocationInformation
    fmt.Fprintf(os.Stdout, "Response from `ReferenceLocationInfoRetrievalApi.GetReferenceLocationInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imsUeId** | **string** | IMS Identity | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReferenceLocationInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **supportedFeatures** | **string** | Supported Features | 
 **privateIdentity** | **string** | IMS Private Identity | 

### Return type

[**ReferenceLocationInformation**](ReferenceLocationInformation.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

