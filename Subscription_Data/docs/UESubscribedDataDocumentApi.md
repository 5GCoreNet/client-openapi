# \UESubscribedDataDocumentApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QueryUeSubscribedData**](UESubscribedDataDocumentApi.md#QueryUeSubscribedData) | **Get** /subscription-data/{ueId} | Retrieve multiple subscribed data sets of a UE



## QueryUeSubscribedData

> UeSubscribedDataSets QueryUeSubscribedData(ctx, ueId).DatasetNames(datasetNames).ServingPlmn(servingPlmn).Execute()

Retrieve multiple subscribed data sets of a UE

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
    ueId := "ueId_example" // string | UE Id
    datasetNames := []openapiclient.DataSetName{*openapiclient.NewDataSetName()} // []DataSetName | List of dataset names (optional)
    servingPlmn := "servingPlmn_example" // string | Serving PLMN Id (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UESubscribedDataDocumentApi.QueryUeSubscribedData(context.Background(), ueId).DatasetNames(datasetNames).ServingPlmn(servingPlmn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UESubscribedDataDocumentApi.QueryUeSubscribedData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryUeSubscribedData`: UeSubscribedDataSets
    fmt.Fprintf(os.Stdout, "Response from `UESubscribedDataDocumentApi.QueryUeSubscribedData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | UE Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiQueryUeSubscribedDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **datasetNames** | [**[]DataSetName**](DataSetName.md) | List of dataset names | 
 **servingPlmn** | **string** | Serving PLMN Id | 

### Return type

[**UeSubscribedDataSets**](UeSubscribedDataSets.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

