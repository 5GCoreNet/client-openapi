# \UeContextsCollectionApi

All URIs are relative to *https://example.com/namf-mt/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EnableGroupReachability**](UeContextsCollectionApi.md#EnableGroupReachability) | **Post** /ue-contexts | Namf_MT EnableGroupReachability service Operation



## EnableGroupReachability

> EnableGroupReachabilityRspData EnableGroupReachability(ctx).EnableGroupReachabilityReqData(enableGroupReachabilityReqData).Execute()

Namf_MT EnableGroupReachability service Operation

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
    enableGroupReachabilityReqData := *openapiclient.NewEnableGroupReachabilityReqData([]openapiclient.UeInfo{*openapiclient.NewUeInfo([]string{"UeList_example"})}, *openapiclient.NewTmgi("MbsServiceId_example", *openapiclient.NewPlmnId("Mcc_example", "Mnc_example"))) // EnableGroupReachabilityReqData | list of UEs requested to be made reachable for the related TMGI

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UeContextsCollectionApi.EnableGroupReachability(context.Background()).EnableGroupReachabilityReqData(enableGroupReachabilityReqData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UeContextsCollectionApi.EnableGroupReachability``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnableGroupReachability`: EnableGroupReachabilityRspData
    fmt.Fprintf(os.Stdout, "Response from `UeContextsCollectionApi.EnableGroupReachability`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnableGroupReachabilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enableGroupReachabilityReqData** | [**EnableGroupReachabilityReqData**](EnableGroupReachabilityReqData.md) | list of UEs requested to be made reachable for the related TMGI | 

### Return type

[**EnableGroupReachabilityRspData**](EnableGroupReachabilityRspData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

