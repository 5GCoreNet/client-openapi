# \UserConsentDataApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QueryUserConsentData**](UserConsentDataApi.md#QueryUserConsentData) | **Get** /subscription-data/{ueId}/uc-data | Retrieves the subscribed User Consent Data of a UE



## QueryUserConsentData

> UcSubscriptionData QueryUserConsentData(ctx, ueId).SupportedFeatures(supportedFeatures).UcPurpose(ucPurpose).IfNoneMatch(ifNoneMatch).IfModifiedSince(ifModifiedSince).Execute()

Retrieves the subscribed User Consent Data of a UE

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
    supportedFeatures := "supportedFeatures_example" // string | Supported Features (optional)
    ucPurpose := *openapiclient.NewUcPurpose() // UcPurpose | User consent purpose (optional)
    ifNoneMatch := "ifNoneMatch_example" // string | Validator for conditional requests, as described in RFC 7232, 3.2 (optional)
    ifModifiedSince := "ifModifiedSince_example" // string | Validator for conditional requests, as described in RFC 7232, 3.3 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserConsentDataApi.QueryUserConsentData(context.Background(), ueId).SupportedFeatures(supportedFeatures).UcPurpose(ucPurpose).IfNoneMatch(ifNoneMatch).IfModifiedSince(ifModifiedSince).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserConsentDataApi.QueryUserConsentData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryUserConsentData`: UcSubscriptionData
    fmt.Fprintf(os.Stdout, "Response from `UserConsentDataApi.QueryUserConsentData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | UE id | 

### Other Parameters

Other parameters are passed through a pointer to a apiQueryUserConsentDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **supportedFeatures** | **string** | Supported Features | 
 **ucPurpose** | [**UcPurpose**](UcPurpose.md) | User consent purpose | 
 **ifNoneMatch** | **string** | Validator for conditional requests, as described in RFC 7232, 3.2 | 
 **ifModifiedSince** | **string** | Validator for conditional requests, as described in RFC 7232, 3.3 | 

### Return type

[**UcSubscriptionData**](UcSubscriptionData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

