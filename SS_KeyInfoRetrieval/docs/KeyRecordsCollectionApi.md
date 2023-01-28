# \KeyRecordsCollectionApi

All URIs are relative to *https://example.com/ss-kir/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RetrieveKeyMgmtInfo**](KeyRecordsCollectionApi.md#RetrieveKeyMgmtInfo) | **Get** /key-records | 



## RetrieveKeyMgmtInfo

> ValKeyInfo RetrieveKeyMgmtInfo(ctx).ValServiceId(valServiceId).ValTgtUe(valTgtUe).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/SS_KeyInfoRetrieval"
)

func main() {
    valServiceId := "valServiceId_example" // string | String identifying an individual VAL service
    valTgtUe := map[string][]openapiclient.ValTargetUe{ ... } // ValTargetUe | Identifying a VAL target. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyRecordsCollectionApi.RetrieveKeyMgmtInfo(context.Background()).ValServiceId(valServiceId).ValTgtUe(valTgtUe).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyRecordsCollectionApi.RetrieveKeyMgmtInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveKeyMgmtInfo`: ValKeyInfo
    fmt.Fprintf(os.Stdout, "Response from `KeyRecordsCollectionApi.RetrieveKeyMgmtInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveKeyMgmtInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **valServiceId** | **string** | String identifying an individual VAL service | 
 **valTgtUe** | [**ValTargetUe**](ValTargetUe.md) | Identifying a VAL target. | 

### Return type

[**ValKeyInfo**](ValKeyInfo.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

