# \ProvisionedParameterDataEntriesCollectionApi

All URIs are relative to *https://example.com/nudr-dr/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMultiplePPDataEntries**](ProvisionedParameterDataEntriesCollectionApi.md#GetMultiplePPDataEntries) | **Get** /subscription-data/{ueId}/pp-data-store | get a list of Parameter Provisioning Data Entries



## GetMultiplePPDataEntries

> PpDataEntryList GetMultiplePPDataEntries(ctx, ueId).SupportedFeatures(supportedFeatures).Execute()

get a list of Parameter Provisioning Data Entries

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
    ueId := *openapiclient.NewGetMultiplePPDataEntriesUeIdParameter() // GetMultiplePPDataEntriesUeIdParameter | Identifier of the UE
    supportedFeatures := "supportedFeatures_example" // string | Features required to be supported by the target NF (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvisionedParameterDataEntriesCollectionApi.GetMultiplePPDataEntries(context.Background(), ueId).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvisionedParameterDataEntriesCollectionApi.GetMultiplePPDataEntries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMultiplePPDataEntries`: PpDataEntryList
    fmt.Fprintf(os.Stdout, "Response from `ProvisionedParameterDataEntriesCollectionApi.GetMultiplePPDataEntries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | [**GetMultiplePPDataEntriesUeIdParameter**](.md) | Identifier of the UE | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMultiplePPDataEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **supportedFeatures** | **string** | Features required to be supported by the target NF | 

### Return type

[**PpDataEntryList**](PpDataEntryList.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

