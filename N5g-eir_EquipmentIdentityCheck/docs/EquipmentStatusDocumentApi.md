# \EquipmentStatusDocumentApi

All URIs are relative to *https://example.com/n5g-eir-eic/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEquipmentStatus**](EquipmentStatusDocumentApi.md#GetEquipmentStatus) | **Get** /equipment-status | Retrieves the status of the UE



## GetEquipmentStatus

> EirResponseData GetEquipmentStatus(ctx).Pei(pei).Supi(supi).Gpsi(gpsi).SupportedFeatures(supportedFeatures).Execute()

Retrieves the status of the UE

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/N5g-eir_EquipmentIdentityCheck"
)

func main() {
    pei := "pei_example" // string | PEI of the UE
    supi := "supi_example" // string | SUPI of the UE (optional)
    gpsi := "gpsi_example" // string | GPSI of the UE (optional)
    supportedFeatures := "supportedFeatures_example" // string | supported features of the NF consumer (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EquipmentStatusDocumentApi.GetEquipmentStatus(context.Background()).Pei(pei).Supi(supi).Gpsi(gpsi).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EquipmentStatusDocumentApi.GetEquipmentStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEquipmentStatus`: EirResponseData
    fmt.Fprintf(os.Stdout, "Response from `EquipmentStatusDocumentApi.GetEquipmentStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEquipmentStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pei** | **string** | PEI of the UE | 
 **supi** | **string** | SUPI of the UE | 
 **gpsi** | **string** | GPSI of the UE | 
 **supportedFeatures** | **string** | supported features of the NF consumer | 

### Return type

[**EirResponseData**](EirResponseData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

