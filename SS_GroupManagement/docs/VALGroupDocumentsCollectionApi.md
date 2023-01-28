# \VALGroupDocumentsCollectionApi

All URIs are relative to *https://example.com/ss-gm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateValGroupDoc**](VALGroupDocumentsCollectionApi.md#CreateValGroupDoc) | **Post** /group-documents | 
[**RetrieveValGroupDocs**](VALGroupDocumentsCollectionApi.md#RetrieveValGroupDocs) | **Get** /group-documents | 



## CreateValGroupDoc

> VALGroupDocument CreateValGroupDoc(ctx).VALGroupDocument(vALGroupDocument).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/SS_GroupManagement"
)

func main() {
    vALGroupDocument := *openapiclient.NewVALGroupDocument("ValGroupId_example") // VALGroupDocument | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VALGroupDocumentsCollectionApi.CreateValGroupDoc(context.Background()).VALGroupDocument(vALGroupDocument).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VALGroupDocumentsCollectionApi.CreateValGroupDoc``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateValGroupDoc`: VALGroupDocument
    fmt.Fprintf(os.Stdout, "Response from `VALGroupDocumentsCollectionApi.CreateValGroupDoc`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateValGroupDocRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vALGroupDocument** | [**VALGroupDocument**](VALGroupDocument.md) |  | 

### Return type

[**VALGroupDocument**](VALGroupDocument.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveValGroupDocs

> []VALGroupDocument RetrieveValGroupDocs(ctx).ValGroupId(valGroupId).ValServiceId(valServiceId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/SS_GroupManagement"
)

func main() {
    valGroupId := "valGroupId_example" // string | String identifying the VAL group. (optional)
    valServiceId := "valServiceId_example" // string | String identifying the Val service. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VALGroupDocumentsCollectionApi.RetrieveValGroupDocs(context.Background()).ValGroupId(valGroupId).ValServiceId(valServiceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VALGroupDocumentsCollectionApi.RetrieveValGroupDocs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveValGroupDocs`: []VALGroupDocument
    fmt.Fprintf(os.Stdout, "Response from `VALGroupDocumentsCollectionApi.RetrieveValGroupDocs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveValGroupDocsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **valGroupId** | **string** | String identifying the VAL group. | 
 **valServiceId** | **string** | String identifying the Val service. | 

### Return type

[**[]VALGroupDocument**](VALGroupDocument.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

