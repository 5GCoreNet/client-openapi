# \VALServicesApi

All URIs are relative to *https://example.com/ss-upr/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RetrieveValUserProfile**](VALServicesApi.md#RetrieveValUserProfile) | **Get** /val-services | 



## RetrieveValUserProfile

> []ProfileDoc RetrieveValUserProfile(ctx).ValTgtUe(valTgtUe).ValServiceId(valServiceId).Execute()





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
    valTgtUe := map[string][]openapiclient.ValTargetUe{ ... } // ValTargetUe | Identifying a VAL target UE.
    valServiceId := "valServiceId_example" // string | String identifying an individual VAL service (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VALServicesApi.RetrieveValUserProfile(context.Background()).ValTgtUe(valTgtUe).ValServiceId(valServiceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VALServicesApi.RetrieveValUserProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveValUserProfile`: []ProfileDoc
    fmt.Fprintf(os.Stdout, "Response from `VALServicesApi.RetrieveValUserProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveValUserProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **valTgtUe** | [**ValTargetUe**](ValTargetUe.md) | Identifying a VAL target UE. | 
 **valServiceId** | **string** | String identifying an individual VAL service | 

### Return type

[**[]ProfileDoc**](ProfileDoc.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

