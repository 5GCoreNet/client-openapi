# \GenerateSIPAuthDataCustomOperationApi

All URIs are relative to *https://example.com/nhss-ims-ueau/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GenerateSipAuthData**](GenerateSIPAuthDataCustomOperationApi.md#GenerateSipAuthData) | **Post** /{impi}/security-information/generate-sip-auth-data | Generate authentication data for the UE based on the Auth-Scheme provided



## GenerateSipAuthData

> SipAuthenticationInfoResult GenerateSipAuthData(ctx, impi).SipAuthenticationInfoRequest(sipAuthenticationInfoRequest).Execute()

Generate authentication data for the UE based on the Auth-Scheme provided

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nhss_imsUEAU"
)

func main() {
    impi := "impi_example" // string | IMS Private Identity for the UE (IMPI)
    sipAuthenticationInfoRequest := *openapiclient.NewSipAuthenticationInfoRequest("CscfServerName_example", *openapiclient.NewSipAuthenticationScheme()) // SipAuthenticationInfoRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GenerateSIPAuthDataCustomOperationApi.GenerateSipAuthData(context.Background(), impi).SipAuthenticationInfoRequest(sipAuthenticationInfoRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenerateSIPAuthDataCustomOperationApi.GenerateSipAuthData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateSipAuthData`: SipAuthenticationInfoResult
    fmt.Fprintf(os.Stdout, "Response from `GenerateSIPAuthDataCustomOperationApi.GenerateSipAuthData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**impi** | **string** | IMS Private Identity for the UE (IMPI) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateSipAuthDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sipAuthenticationInfoRequest** | [**SipAuthenticationInfoRequest**](SipAuthenticationInfoRequest.md) |  | 

### Return type

[**SipAuthenticationInfoResult**](SipAuthenticationInfoResult.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

