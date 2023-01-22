# \RepositoryDataApi

All URIs are relative to *https://example.com/nhss-ims-sdm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRepositoryDataServInd**](RepositoryDataApi.md#GetRepositoryDataServInd) | **Get** /{imsUeId}/repository-data/{serviceIndication} | Retrieve the repository data associated to an IMPU and service indication



## GetRepositoryDataServInd

> RepositoryData GetRepositoryDataServInd(ctx, imsUeId, serviceIndication).SupportedFeatures(supportedFeatures).Execute()

Retrieve the repository data associated to an IMPU and service indication

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
    imsUeId := "imsUeId_example" // string | IMS Identity
    serviceIndication := "serviceIndication_example" // string | Identifier of a service related data
    supportedFeatures := "supportedFeatures_example" // string | Supported Features (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RepositoryDataApi.GetRepositoryDataServInd(context.Background(), imsUeId, serviceIndication).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RepositoryDataApi.GetRepositoryDataServInd``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRepositoryDataServInd`: RepositoryData
    fmt.Fprintf(os.Stdout, "Response from `RepositoryDataApi.GetRepositoryDataServInd`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imsUeId** | **string** | IMS Identity | 
**serviceIndication** | **string** | Identifier of a service related data | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepositoryDataServIndRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **supportedFeatures** | **string** | Supported Features | 

### Return type

[**RepositoryData**](RepositoryData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

