# \Class5MBSSubscriptionDataDocumentApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Query5mbsData**](Class5MBSSubscriptionDataDocumentApi.md#Query5mbsData) | **Get** /subscription-data/{ueId}/5mbs-data | Retrieves the 5mbs subscription data of a UE



## Query5mbsData

> MbsSubscriptionData Query5mbsData(ctx, ueId).SupportedFeatures(supportedFeatures).IfNoneMatch(ifNoneMatch).IfModifiedSince(ifModifiedSince).Execute()

Retrieves the 5mbs subscription data of a UE

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
    ifNoneMatch := "ifNoneMatch_example" // string | Validator for conditional requests, as described in RFC 7232, 3.2 (optional)
    ifModifiedSince := "ifModifiedSince_example" // string | Validator for conditional requests, as described in RFC 7232, 3.3 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Class5MBSSubscriptionDataDocumentApi.Query5mbsData(context.Background(), ueId).SupportedFeatures(supportedFeatures).IfNoneMatch(ifNoneMatch).IfModifiedSince(ifModifiedSince).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Class5MBSSubscriptionDataDocumentApi.Query5mbsData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Query5mbsData`: MbsSubscriptionData
    fmt.Fprintf(os.Stdout, "Response from `Class5MBSSubscriptionDataDocumentApi.Query5mbsData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | UE id | 

### Other Parameters

Other parameters are passed through a pointer to a apiQuery5mbsDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **supportedFeatures** | **string** | Supported Features | 
 **ifNoneMatch** | **string** | Validator for conditional requests, as described in RFC 7232, 3.2 | 
 **ifModifiedSince** | **string** | Validator for conditional requests, as described in RFC 7232, 3.3 | 

### Return type

[**MbsSubscriptionData**](MbsSubscriptionData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

