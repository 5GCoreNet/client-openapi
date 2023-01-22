# \MBSSessionPolicyControlDataDocumentApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMBSSessPolCtrlData**](MBSSessionPolicyControlDataDocumentApi.md#GetMBSSessPolCtrlData) | **Get** /policy-data/mbs-session-pol-data/{polSessionId} | Retrieve MBS Session Policy Control Data for an MBS Session.



## GetMBSSessPolCtrlData

> MbsSessPolCtrlData GetMBSSessPolCtrlData(ctx, polSessionId).Execute()

Retrieve MBS Session Policy Control Data for an MBS Session.

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
    polSessionId := map[string][]openapiclient.MbsSessPolDataId{ ... } // MbsSessPolDataId | Represents the identifier of the MBS Session Policy Control Data. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MBSSessionPolicyControlDataDocumentApi.GetMBSSessPolCtrlData(context.Background(), polSessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MBSSessionPolicyControlDataDocumentApi.GetMBSSessPolCtrlData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMBSSessPolCtrlData`: MbsSessPolCtrlData
    fmt.Fprintf(os.Stdout, "Response from `MBSSessionPolicyControlDataDocumentApi.GetMBSSessPolCtrlData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**polSessionId** | [**MbsSessPolDataId**](.md) | Represents the identifier of the MBS Session Policy Control Data.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMBSSessPolCtrlDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MbsSessPolCtrlData**](MbsSessPolCtrlData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

