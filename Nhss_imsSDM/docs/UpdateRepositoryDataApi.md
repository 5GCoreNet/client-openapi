# \UpdateRepositoryDataApi

All URIs are relative to *https://example.com/nhss-ims-sdm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateRepositoryDataServInd**](UpdateRepositoryDataApi.md#UpdateRepositoryDataServInd) | **Put** /{imsUeId}/repository-data/{serviceIndication} | Update the repository data associated to an IMPU and service indication



## UpdateRepositoryDataServInd

> RepositoryData UpdateRepositoryDataServInd(ctx, imsUeId, serviceIndication).RepositoryData(repositoryData).Execute()

Update the repository data associated to an IMPU and service indication

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
    serviceIndication := "serviceIndication_example" // string | Identifier of a service related data
    repositoryData := *openapiclient.NewRepositoryData(int32(123), string(123)) // RepositoryData | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UpdateRepositoryDataApi.UpdateRepositoryDataServInd(context.Background(), imsUeId, serviceIndication).RepositoryData(repositoryData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UpdateRepositoryDataApi.UpdateRepositoryDataServInd``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRepositoryDataServInd`: RepositoryData
    fmt.Fprintf(os.Stdout, "Response from `UpdateRepositoryDataApi.UpdateRepositoryDataServInd`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imsUeId** | **string** | IMS Identity | 
**serviceIndication** | **string** | Identifier of a service related data | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRepositoryDataServIndRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **repositoryData** | [**RepositoryData**](RepositoryData.md) |  | 

### Return type

[**RepositoryData**](RepositoryData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

