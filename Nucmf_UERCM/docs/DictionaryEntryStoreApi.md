# \DictionaryEntryStoreApi

All URIs are relative to *https://example.com/nucmf-uecm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RetrieveDictionaryEntry**](DictionaryEntryStoreApi.md#RetrieveDictionaryEntry) | **Get** /dic-entries | retrieve a dictionary entry matching query parameters



## RetrieveDictionaryEntry

> RetrieveDictionaryEntry200Response RetrieveDictionaryEntry(ctx).UeRadioCapaId(ueRadioCapaId).RacFormat(racFormat).SupportedFeatures(supportedFeatures).Execute()

retrieve a dictionary entry matching query parameters

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
    ueRadioCapaId := map[string][]openapiclient.UeRadioCapaId{ ... } // UeRadioCapaId | UE Radio Capability ID, either PLMN Assigned or Manufacturer Assigned
    racFormat := *openapiclient.NewRacFormat() // RacFormat | Encoding format of RAC Info (optional)
    supportedFeatures := "supportedFeatures_example" // string | supported features of the NF consumer (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DictionaryEntryStoreApi.RetrieveDictionaryEntry(context.Background()).UeRadioCapaId(ueRadioCapaId).RacFormat(racFormat).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DictionaryEntryStoreApi.RetrieveDictionaryEntry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveDictionaryEntry`: RetrieveDictionaryEntry200Response
    fmt.Fprintf(os.Stdout, "Response from `DictionaryEntryStoreApi.RetrieveDictionaryEntry`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveDictionaryEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ueRadioCapaId** | [**UeRadioCapaId**](UeRadioCapaId.md) | UE Radio Capability ID, either PLMN Assigned or Manufacturer Assigned | 
 **racFormat** | [**RacFormat**](RacFormat.md) | Encoding format of RAC Info | 
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

