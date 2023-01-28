# \PFDOfApplicationsByPartialUpdateApi

All URIs are relative to *https://example.com/nnef-pfdmanagement/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NnefPFDmanagementAppFetchPartialUpdate**](PFDOfApplicationsByPartialUpdateApi.md#NnefPFDmanagementAppFetchPartialUpdate) | **Post** /applications/partialpull | retrieve the PFD(s) by partial update



## NnefPFDmanagementAppFetchPartialUpdate

> []PfdDataForApp NnefPFDmanagementAppFetchPartialUpdate(ctx).ApplicationForPfdRequest(applicationForPfdRequest).Execute()

retrieve the PFD(s) by partial update

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nnef_PFDmanagement"
)

func main() {
    applicationForPfdRequest := []openapiclient.ApplicationForPfdRequest{*openapiclient.NewApplicationForPfdRequest("ApplicationId_example")} // []ApplicationForPfdRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PFDOfApplicationsByPartialUpdateApi.NnefPFDmanagementAppFetchPartialUpdate(context.Background()).ApplicationForPfdRequest(applicationForPfdRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PFDOfApplicationsByPartialUpdateApi.NnefPFDmanagementAppFetchPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NnefPFDmanagementAppFetchPartialUpdate`: []PfdDataForApp
    fmt.Fprintf(os.Stdout, "Response from `PFDOfApplicationsByPartialUpdateApi.NnefPFDmanagementAppFetchPartialUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNnefPFDmanagementAppFetchPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicationForPfdRequest** | [**[]ApplicationForPfdRequest**](ApplicationForPfdRequest.md) |  | 

### Return type

[**[]PfdDataForApp**](PfdDataForApp.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

