# \SoRInformationRetrievalApi

All URIs are relative to *https://example.com/nsoraf-sor/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSorInformation**](SoRInformationRetrievalApi.md#GetSorInformation) | **Get** /{supi}/sor-information | retrieve the steering of roaming information for a UE



## GetSorInformation

> SorInformation GetSorInformation(ctx, supi).PlmnId(plmnId).SupportedFeatures(supportedFeatures).AccessType(accessType).Execute()

retrieve the steering of roaming information for a UE

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
    supi := "supi_example" // string | Identifier of the UE
    plmnId := map[string][]openapiclient.PlmnIdNid{ ... } // PlmnIdNid | serving PLMN ID or SNPN ID
    supportedFeatures := "supportedFeatures_example" // string | Supported Features (optional)
    accessType := openapiclient.AccessType("3GPP_ACCESS") // AccessType | Access type used by the UE (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SoRInformationRetrievalApi.GetSorInformation(context.Background(), supi).PlmnId(plmnId).SupportedFeatures(supportedFeatures).AccessType(accessType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SoRInformationRetrievalApi.GetSorInformation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSorInformation`: SorInformation
    fmt.Fprintf(os.Stdout, "Response from `SoRInformationRetrievalApi.GetSorInformation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supi** | **string** | Identifier of the UE | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSorInformationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **plmnId** | [**PlmnIdNid**](PlmnIdNid.md) | serving PLMN ID or SNPN ID | 
 **supportedFeatures** | **string** | Supported Features | 
 **accessType** | [**AccessType**](AccessType.md) | Access type used by the UE | 

### Return type

[**SorInformation**](SorInformation.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

