# \DicEntryIdDocumentApi

All URIs are relative to *https://example.com/nucmf-uecm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDicEntry**](DicEntryIdDocumentApi.md#GetDicEntry) | **Get** /dic-entries/{dicEntryId} | Get an individual dictionary entry via dicEntryId



## GetDicEntry

> RetrieveDictionaryEntry200Response GetDicEntry(ctx, dicEntryId).RacFormat(racFormat).SupportedFeatures(supportedFeatures).Execute()

Get an individual dictionary entry via dicEntryId

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nucmf_UERCM"
)

func main() {
    dicEntryId := int32(56) // int32 | the ID of a dictionary entry in the UCMF
    racFormat := *openapiclient.NewRacFormat() // RacFormat | Encoding format of of RAC Info (optional)
    supportedFeatures := "supportedFeatures_example" // string | supported features of the NF consumer (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DicEntryIdDocumentApi.GetDicEntry(context.Background(), dicEntryId).RacFormat(racFormat).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DicEntryIdDocumentApi.GetDicEntry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDicEntry`: RetrieveDictionaryEntry200Response
    fmt.Fprintf(os.Stdout, "Response from `DicEntryIdDocumentApi.GetDicEntry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dicEntryId** | **int32** | the ID of a dictionary entry in the UCMF | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDicEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **racFormat** | [**RacFormat**](RacFormat.md) | Encoding format of of RAC Info | 
 **supportedFeatures** | **string** | supported features of the NF consumer | 

### Return type

[**RetrieveDictionaryEntry200Response**](RetrieveDictionaryEntry200Response.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: multipart/related, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

