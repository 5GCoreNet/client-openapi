# \ServiceParameterDataStoreApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReadServiceParameterData**](ServiceParameterDataStoreApi.md#ReadServiceParameterData) | **Get** /application-data/serviceParamData | Retrieve Service Parameter Data



## ReadServiceParameterData

> []ServiceParameterData ReadServiceParameterData(ctx).ServiceParamIds(serviceParamIds).Dnns(dnns).Snssais(snssais).InternalGroupIds(internalGroupIds).Supis(supis).UeIpv4s(ueIpv4s).UeIpv6s(ueIpv6s).UeMacs(ueMacs).AnyUe(anyUe).SuppFeat(suppFeat).Execute()

Retrieve Service Parameter Data

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
    serviceParamIds := []string{"Inner_example"} // []string | Each element identifies a service. (optional)
    dnns := []string{"Inner_example"} // []string | Each element identifies a DNN. (optional)
    snssais := []openapiclient.Snssai{*openapiclient.NewSnssai(int32(123))} // []Snssai | Each element identifies a slice. (optional)
    internalGroupIds := []string{"Inner_example"} // []string | Each element identifies a group of users. (optional)
    supis := []string{"Inner_example"} // []string | Each element identifies the user. (optional)
    ueIpv4s := []string{"198.51.100.1"} // []string | Each element identifies the user. (optional)
    ueIpv6s := []openapiclient.Ipv6Addr{*openapiclient.NewIpv6Addr()} // []Ipv6Addr | Each element identifies the user. (optional)
    ueMacs := []string{"Inner_example"} // []string | Each element identifies the user. (optional)
    anyUe := true // bool | Indicates whether the request is for any UE. (optional)
    suppFeat := "suppFeat_example" // string | Supported Features (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceParameterDataStoreApi.ReadServiceParameterData(context.Background()).ServiceParamIds(serviceParamIds).Dnns(dnns).Snssais(snssais).InternalGroupIds(internalGroupIds).Supis(supis).UeIpv4s(ueIpv4s).UeIpv6s(ueIpv6s).UeMacs(ueMacs).AnyUe(anyUe).SuppFeat(suppFeat).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceParameterDataStoreApi.ReadServiceParameterData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadServiceParameterData`: []ServiceParameterData
    fmt.Fprintf(os.Stdout, "Response from `ServiceParameterDataStoreApi.ReadServiceParameterData`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadServiceParameterDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceParamIds** | **[]string** | Each element identifies a service. | 
 **dnns** | **[]string** | Each element identifies a DNN. | 
 **snssais** | [**[]Snssai**](Snssai.md) | Each element identifies a slice. | 
 **internalGroupIds** | **[]string** | Each element identifies a group of users. | 
 **supis** | **[]string** | Each element identifies the user. | 
 **ueIpv4s** | **[]string** | Each element identifies the user. | 
 **ueIpv6s** | [**[]Ipv6Addr**](Ipv6Addr.md) | Each element identifies the user. | 
 **ueMacs** | **[]string** | Each element identifies the user. | 
 **anyUe** | **bool** | Indicates whether the request is for any UE. | 
 **suppFeat** | **string** | Supported Features | 

### Return type

[**[]ServiceParameterData**](ServiceParameterData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

