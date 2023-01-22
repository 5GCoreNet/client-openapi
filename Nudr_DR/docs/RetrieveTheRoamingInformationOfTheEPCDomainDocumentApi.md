# \RetrieveTheRoamingInformationOfTheEPCDomainDocumentApi

All URIs are relative to *https://example.com/nudr-dr/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QueryRoamingInformation**](RetrieveTheRoamingInformationOfTheEPCDomainDocumentApi.md#QueryRoamingInformation) | **Get** /subscription-data/{ueId}/context-data/roaming-information | Retrieves the Roaming Information of the EPC domain



## QueryRoamingInformation

> RoamingInfoUpdate QueryRoamingInformation(ctx, ueId).Execute()

Retrieves the Roaming Information of the EPC domain

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
    ueId := "ueId_example" // string | UE id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RetrieveTheRoamingInformationOfTheEPCDomainDocumentApi.QueryRoamingInformation(context.Background(), ueId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RetrieveTheRoamingInformationOfTheEPCDomainDocumentApi.QueryRoamingInformation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryRoamingInformation`: RoamingInfoUpdate
    fmt.Fprintf(os.Stdout, "Response from `RetrieveTheRoamingInformationOfTheEPCDomainDocumentApi.QueryRoamingInformation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | UE id | 

### Other Parameters

Other parameters are passed through a pointer to a apiQueryRoamingInformationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RoamingInfoUpdate**](RoamingInfoUpdate.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

