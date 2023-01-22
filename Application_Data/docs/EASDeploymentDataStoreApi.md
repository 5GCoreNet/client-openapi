# \EASDeploymentDataStoreApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReadEasDeployData**](EASDeploymentDataStoreApi.md#ReadEasDeployData) | **Get** /application-data/eas-deploy-data | Retrieve EAS Deployment Information Data



## ReadEasDeployData

> []EasDeployInfoData ReadEasDeployData(ctx).Dnn(dnn).Snssai(snssai).InternalGroupId(internalGroupId).AppId(appId).Execute()

Retrieve EAS Deployment Information Data

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
    dnn := "dnn_example" // string | Identifies a DNN. (optional)
    snssai := map[string][]openapiclient.Snssai{ ... } // Snssai | Identifies an S-NSSAI. (optional)
    internalGroupId := "internalGroupId_example" // string | Identifies a group of users. (optional)
    appId := "appId_example" // string | Identifies an application. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EASDeploymentDataStoreApi.ReadEasDeployData(context.Background()).Dnn(dnn).Snssai(snssai).InternalGroupId(internalGroupId).AppId(appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EASDeploymentDataStoreApi.ReadEasDeployData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadEasDeployData`: []EasDeployInfoData
    fmt.Fprintf(os.Stdout, "Response from `EASDeploymentDataStoreApi.ReadEasDeployData`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadEasDeployDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dnn** | **string** | Identifies a DNN. | 
 **snssai** | [**Snssai**](Snssai.md) | Identifies an S-NSSAI. | 
 **internalGroupId** | **string** | Identifies a group of users. | 
 **appId** | **string** | Identifies an application. | 

### Return type

[**[]EasDeployInfoData**](EasDeployInfoData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

