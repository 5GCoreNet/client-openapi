# \IndividualEASDeploymentDataDocumentApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrReplaceIndividualEasDeployData**](IndividualEASDeploymentDataDocumentApi.md#CreateOrReplaceIndividualEasDeployData) | **Put** /application-data/eas-deploy-data/{easDeployInfoId} | Create or update an individual EAS Deployment Data resource
[**ReadIndividualEasDeployData**](IndividualEASDeploymentDataDocumentApi.md#ReadIndividualEasDeployData) | **Get** /application-data/eas-deploy-data/{easDeployInfoId} | Retrieve an individual EAS Deployment Data resource



## CreateOrReplaceIndividualEasDeployData

> EasDeployInfoData CreateOrReplaceIndividualEasDeployData(ctx, easDeployInfoId).EasDeployInfoData(easDeployInfoData).Execute()

Create or update an individual EAS Deployment Data resource

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Application_Data"
)

func main() {
    easDeployInfoId := "easDeployInfoId_example" // string | The Identifier of an Individual EAS Deployment Data to be created or updated. It shall apply the format of Data type string. 
    easDeployInfoData := *openapiclient.NewEasDeployInfoData([]openapiclient.FqdnPatternMatchingRule{openapiclient.FqdnPatternMatchingRule{Interface{}: new(interface{})}}) // EasDeployInfoData | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualEASDeploymentDataDocumentApi.CreateOrReplaceIndividualEasDeployData(context.Background(), easDeployInfoId).EasDeployInfoData(easDeployInfoData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualEASDeploymentDataDocumentApi.CreateOrReplaceIndividualEasDeployData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrReplaceIndividualEasDeployData`: EasDeployInfoData
    fmt.Fprintf(os.Stdout, "Response from `IndividualEASDeploymentDataDocumentApi.CreateOrReplaceIndividualEasDeployData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**easDeployInfoId** | **string** | The Identifier of an Individual EAS Deployment Data to be created or updated. It shall apply the format of Data type string.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrReplaceIndividualEasDeployDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **easDeployInfoData** | [**EasDeployInfoData**](EasDeployInfoData.md) |  | 

### Return type

[**EasDeployInfoData**](EasDeployInfoData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadIndividualEasDeployData

> EasDeployInfoData ReadIndividualEasDeployData(ctx, easDeployInfoId).Execute()

Retrieve an individual EAS Deployment Data resource

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Application_Data"
)

func main() {
    easDeployInfoId := "easDeployInfoId_example" // string | String identifying an Individual EAS Deployment Information Data resource. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualEASDeploymentDataDocumentApi.ReadIndividualEasDeployData(context.Background(), easDeployInfoId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualEASDeploymentDataDocumentApi.ReadIndividualEasDeployData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadIndividualEasDeployData`: EasDeployInfoData
    fmt.Fprintf(os.Stdout, "Response from `IndividualEASDeploymentDataDocumentApi.ReadIndividualEasDeployData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**easDeployInfoId** | **string** | String identifying an Individual EAS Deployment Information Data resource.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadIndividualEasDeployDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EasDeployInfoData**](EasDeployInfoData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

