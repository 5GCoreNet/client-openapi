# \RetrievalOfTheSCSCFSelectionAssistanceInformationForTheIMSSubscriptionApi

All URIs are relative to *https://example.com/nhss-ims-sdm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetScscfSelectionAssistanceInfo**](RetrievalOfTheSCSCFSelectionAssistanceInformationForTheIMSSubscriptionApi.md#GetScscfSelectionAssistanceInfo) | **Get** /{imsUeId}/ims-data/location-data/scscf-selection-assistance-info | Retrieve the S-CSCF selection assistance info



## GetScscfSelectionAssistanceInfo

> ScscfSelectionAssistanceInformation GetScscfSelectionAssistanceInfo(ctx, imsUeId).Execute()

Retrieve the S-CSCF selection assistance info

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
    imsUeId := "imsUeId_example" // string | IMS Identity

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RetrievalOfTheSCSCFSelectionAssistanceInformationForTheIMSSubscriptionApi.GetScscfSelectionAssistanceInfo(context.Background(), imsUeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RetrievalOfTheSCSCFSelectionAssistanceInformationForTheIMSSubscriptionApi.GetScscfSelectionAssistanceInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetScscfSelectionAssistanceInfo`: ScscfSelectionAssistanceInformation
    fmt.Fprintf(os.Stdout, "Response from `RetrievalOfTheSCSCFSelectionAssistanceInformationForTheIMSSubscriptionApi.GetScscfSelectionAssistanceInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imsUeId** | **string** | IMS Identity | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetScscfSelectionAssistanceInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ScscfSelectionAssistanceInformation**](ScscfSelectionAssistanceInformation.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

