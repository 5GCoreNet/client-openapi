# \RepositoryDataListApi

All URIs are relative to *https://example.com/nhss-ims-sdm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRepositoryDataServIndList**](RepositoryDataListApi.md#GetRepositoryDataServIndList) | **Get** /{imsUeId}/repository-data | Retrieve the repository data associated to an IMPU and service indication list



## GetRepositoryDataServIndList

> RepositoryDataList GetRepositoryDataServIndList(ctx, imsUeId).ServiceIndications(serviceIndications).SupportedFeatures(supportedFeatures).Execute()

Retrieve the repository data associated to an IMPU and service indication list

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
    serviceIndications := []string{"Inner_example"} // []string | Identifiers of a services related data
    supportedFeatures := "supportedFeatures_example" // string | Supported Features (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoryDataListApi.GetRepositoryDataServIndList(context.Background(), imsUeId).ServiceIndications(serviceIndications).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoryDataListApi.GetRepositoryDataServIndList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRepositoryDataServIndList`: RepositoryDataList
    fmt.Fprintf(os.Stdout, "Response from `RepositoryDataListApi.GetRepositoryDataServIndList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imsUeId** | **string** | IMS Identity | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepositoryDataServIndListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serviceIndications** | **[]string** | Identifiers of a services related data | 
 **supportedFeatures** | **string** | Supported Features | 

### Return type

[**RepositoryDataList**](RepositoryDataList.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

