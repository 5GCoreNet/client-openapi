# \ADictionaryEntryDocumentApi

All URIs are relative to *https://example.com/nucmf-uecm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDictionaryEntry**](ADictionaryEntryDocumentApi.md#CreateDictionaryEntry) | **Post** /dic-entries | Create a dictionary entry in the UCMF



## CreateDictionaryEntry

> DicEntryCreatedData CreateDictionaryEntry(ctx).JsonData(jsonData).BinaryDataUeRadioCapability5GS(binaryDataUeRadioCapability5GS).BinaryDataUeRadioCapabilityEPS(binaryDataUeRadioCapabilityEPS).BinaryDataUeRadioCap5GSForPaging(binaryDataUeRadioCap5GSForPaging).BinaryDataUeRadioCapEPSForPaging(binaryDataUeRadioCapEPSForPaging).Execute()

Create a dictionary entry in the UCMF

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
    jsonData := *openapiclient.NewDicEntryCreateData("TypeAllocationCode_example") // DicEntryCreateData |  (optional)
    binaryDataUeRadioCapability5GS := os.NewFile(1234, "some_file") // *os.File |  (optional)
    binaryDataUeRadioCapabilityEPS := os.NewFile(1234, "some_file") // *os.File |  (optional)
    binaryDataUeRadioCap5GSForPaging := os.NewFile(1234, "some_file") // *os.File |  (optional)
    binaryDataUeRadioCapEPSForPaging := os.NewFile(1234, "some_file") // *os.File |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ADictionaryEntryDocumentApi.CreateDictionaryEntry(context.Background()).JsonData(jsonData).BinaryDataUeRadioCapability5GS(binaryDataUeRadioCapability5GS).BinaryDataUeRadioCapabilityEPS(binaryDataUeRadioCapabilityEPS).BinaryDataUeRadioCap5GSForPaging(binaryDataUeRadioCap5GSForPaging).BinaryDataUeRadioCapEPSForPaging(binaryDataUeRadioCapEPSForPaging).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ADictionaryEntryDocumentApi.CreateDictionaryEntry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDictionaryEntry`: DicEntryCreatedData
    fmt.Fprintf(os.Stdout, "Response from `ADictionaryEntryDocumentApi.CreateDictionaryEntry`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDictionaryEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jsonData** | [**DicEntryCreateData**](DicEntryCreateData.md) |  | 
 **binaryDataUeRadioCapability5GS** | ***os.File** |  | 
 **binaryDataUeRadioCapabilityEPS** | ***os.File** |  | 
 **binaryDataUeRadioCap5GSForPaging** | ***os.File** |  | 
 **binaryDataUeRadioCapEPSForPaging** | ***os.File** |  | 

### Return type

[**DicEntryCreatedData**](DicEntryCreatedData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: multipart/related
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

