# \IndividualAMInfluenceDataDocumentApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrReplaceIndividualAmInfluenceData**](IndividualAMInfluenceDataDocumentApi.md#CreateOrReplaceIndividualAmInfluenceData) | **Put** /application-data/am-influence-data/{amInfluenceId} | Create or update an individual AM Influence Data resource
[**DeleteIndividualAmInfluenceData**](IndividualAMInfluenceDataDocumentApi.md#DeleteIndividualAmInfluenceData) | **Delete** /application-data/am-influence-data/{amInfluenceId} | Delete an individual AM Influence Data resource
[**UpdateIndividualAmInfluenceData**](IndividualAMInfluenceDataDocumentApi.md#UpdateIndividualAmInfluenceData) | **Patch** /application-data/am-influence-data/{amInfluenceId} | Modify part of the properties of an individual AM Influence Data resource



## CreateOrReplaceIndividualAmInfluenceData

> AmInfluData CreateOrReplaceIndividualAmInfluenceData(ctx, amInfluenceId).AmInfluData(amInfluData).Execute()

Create or update an individual AM Influence Data resource

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
    amInfluenceId := "amInfluenceId_example" // string | The Identifier of an Individual AM Influence Data to be created or updated. It shall apply the format of Data type string. 
    amInfluData := *openapiclient.NewAmInfluData() // AmInfluData | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualAMInfluenceDataDocumentApi.CreateOrReplaceIndividualAmInfluenceData(context.Background(), amInfluenceId).AmInfluData(amInfluData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualAMInfluenceDataDocumentApi.CreateOrReplaceIndividualAmInfluenceData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrReplaceIndividualAmInfluenceData`: AmInfluData
    fmt.Fprintf(os.Stdout, "Response from `IndividualAMInfluenceDataDocumentApi.CreateOrReplaceIndividualAmInfluenceData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**amInfluenceId** | **string** | The Identifier of an Individual AM Influence Data to be created or updated. It shall apply the format of Data type string.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrReplaceIndividualAmInfluenceDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **amInfluData** | [**AmInfluData**](AmInfluData.md) |  | 

### Return type

[**AmInfluData**](AmInfluData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIndividualAmInfluenceData

> DeleteIndividualAmInfluenceData(ctx, amInfluenceId).Execute()

Delete an individual AM Influence Data resource

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
    amInfluenceId := "amInfluenceId_example" // string | The Identifier of an Individual AM Influence Data to be deleted. It shall apply the format of Data type string. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualAMInfluenceDataDocumentApi.DeleteIndividualAmInfluenceData(context.Background(), amInfluenceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualAMInfluenceDataDocumentApi.DeleteIndividualAmInfluenceData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**amInfluenceId** | **string** | The Identifier of an Individual AM Influence Data to be deleted. It shall apply the format of Data type string.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIndividualAmInfluenceDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIndividualAmInfluenceData

> AmInfluData UpdateIndividualAmInfluenceData(ctx, amInfluenceId).AmInfluDataPatch(amInfluDataPatch).Execute()

Modify part of the properties of an individual AM Influence Data resource

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
    amInfluenceId := "amInfluenceId_example" // string | The Identifier of an Individual AM Influence Data to be updated. It shall apply the format of Data type string. 
    amInfluDataPatch := *openapiclient.NewAmInfluDataPatch() // AmInfluDataPatch | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualAMInfluenceDataDocumentApi.UpdateIndividualAmInfluenceData(context.Background(), amInfluenceId).AmInfluDataPatch(amInfluDataPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualAMInfluenceDataDocumentApi.UpdateIndividualAmInfluenceData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIndividualAmInfluenceData`: AmInfluData
    fmt.Fprintf(os.Stdout, "Response from `IndividualAMInfluenceDataDocumentApi.UpdateIndividualAmInfluenceData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**amInfluenceId** | **string** | The Identifier of an Individual AM Influence Data to be updated. It shall apply the format of Data type string.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIndividualAmInfluenceDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **amInfluDataPatch** | [**AmInfluDataPatch**](AmInfluDataPatch.md) |  | 

### Return type

[**AmInfluData**](AmInfluData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

