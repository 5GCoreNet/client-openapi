# \DeleteRepositoryDataApi

All URIs are relative to *https://example.com/nhss-ims-sdm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteRepositoryDataServInd**](DeleteRepositoryDataApi.md#DeleteRepositoryDataServInd) | **Delete** /{imsUeId}/repository-data/{serviceIndication} | delete the Repository Data for a Service Indication



## DeleteRepositoryDataServInd

> DeleteRepositoryDataServInd(ctx, imsUeId, serviceIndication).Execute()

delete the Repository Data for a Service Indication

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
    imsUeId := "imsUeId_example" // string | Identifier of the UE
    serviceIndication := "serviceIndication_example" // string | Identifier of a service related data

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeleteRepositoryDataApi.DeleteRepositoryDataServInd(context.Background(), imsUeId, serviceIndication).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeleteRepositoryDataApi.DeleteRepositoryDataServInd``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imsUeId** | **string** | Identifier of the UE | 
**serviceIndication** | **string** | Identifier of a service related data | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRepositoryDataServIndRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

