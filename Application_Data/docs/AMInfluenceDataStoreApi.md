# \AMInfluenceDataStoreApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReadAmInfluenceData**](AMInfluenceDataStoreApi.md#ReadAmInfluenceData) | **Get** /application-data/am-influence-data | Retrieve AM Influence Data



## ReadAmInfluenceData

> []AmInfluData ReadAmInfluenceData(ctx).AmInfluenceIds(amInfluenceIds).Dnns(dnns).Snssais(snssais).DnnSnssaiInfos(dnnSnssaiInfos).InternalGroupIds(internalGroupIds).Supis(supis).AnyUe(anyUe).SuppFeat(suppFeat).Execute()

Retrieve AM Influence Data

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
    amInfluenceIds := []string{"Inner_example"} // []string | Each element identifies a service. (optional)
    dnns := []string{"Inner_example"} // []string | Each element identifies a DNN. (optional)
    snssais := []openapiclient.Snssai{*openapiclient.NewSnssai(int32(123))} // []Snssai | Each element identifies a slice. (optional)
    dnnSnssaiInfos := []openapiclient.DnnSnssaiInformation{*openapiclient.NewDnnSnssaiInformation()} // []DnnSnssaiInformation | Each element identifies a combination of (DNN, S-NSSAI). (optional)
    internalGroupIds := []string{"Inner_example"} // []string | Each element identifies a group of users. (optional)
    supis := []string{"Inner_example"} // []string | Each element identifies the user. (optional)
    anyUe := true // bool | Indicates whether the request is for any UE. (optional)
    suppFeat := "suppFeat_example" // string | Supported Features (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AMInfluenceDataStoreApi.ReadAmInfluenceData(context.Background()).AmInfluenceIds(amInfluenceIds).Dnns(dnns).Snssais(snssais).DnnSnssaiInfos(dnnSnssaiInfos).InternalGroupIds(internalGroupIds).Supis(supis).AnyUe(anyUe).SuppFeat(suppFeat).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AMInfluenceDataStoreApi.ReadAmInfluenceData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAmInfluenceData`: []AmInfluData
    fmt.Fprintf(os.Stdout, "Response from `AMInfluenceDataStoreApi.ReadAmInfluenceData`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadAmInfluenceDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **amInfluenceIds** | **[]string** | Each element identifies a service. | 
 **dnns** | **[]string** | Each element identifies a DNN. | 
 **snssais** | [**[]Snssai**](Snssai.md) | Each element identifies a slice. | 
 **dnnSnssaiInfos** | [**[]DnnSnssaiInformation**](DnnSnssaiInformation.md) | Each element identifies a combination of (DNN, S-NSSAI). | 
 **internalGroupIds** | **[]string** | Each element identifies a group of users. | 
 **supis** | **[]string** | Each element identifies the user. | 
 **anyUe** | **bool** | Indicates whether the request is for any UE. | 
 **suppFeat** | **string** | Supported Features | 

### Return type

[**[]AmInfluData**](AmInfluData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

