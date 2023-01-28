# \SliceCollectionApi

All URIs are relative to *https://example.com/nnsacf-nsac/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NumOfPDUsUpdate**](SliceCollectionApi.md#NumOfPDUsUpdate) | **Post** /slices/pdus | Network Slice Admission Control on the number of PDU Sessions
[**NumOfUEsUpdate**](SliceCollectionApi.md#NumOfUEsUpdate) | **Post** /slices/ues | Network Slice Admission Control on the Number of UEs



## NumOfPDUsUpdate

> PduACResponseData NumOfPDUsUpdate(ctx).PduACRequestData(pduACRequestData).Execute()

Network Slice Admission Control on the number of PDU Sessions

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nnsacf_NSAC"
)

func main() {
    pduACRequestData := *openapiclient.NewPduACRequestData([]openapiclient.PduACRequestInfo{*openapiclient.NewPduACRequestInfo("Supi_example", openapiclient.AccessType("3GPP_ACCESS"), int32(123), []openapiclient.AcuOperationItem{*openapiclient.NewAcuOperationItem(*openapiclient.NewAcuFlag(), *openapiclient.NewSnssai(int32(123)))})}) // PduACRequestData | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SliceCollectionApi.NumOfPDUsUpdate(context.Background()).PduACRequestData(pduACRequestData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SliceCollectionApi.NumOfPDUsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NumOfPDUsUpdate`: PduACResponseData
    fmt.Fprintf(os.Stdout, "Response from `SliceCollectionApi.NumOfPDUsUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNumOfPDUsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pduACRequestData** | [**PduACRequestData**](PduACRequestData.md) |  | 

### Return type

[**PduACResponseData**](PduACResponseData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NumOfUEsUpdate

> UeACResponseData NumOfUEsUpdate(ctx).UeACRequestData(ueACRequestData).Execute()

Network Slice Admission Control on the Number of UEs

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nnsacf_NSAC"
)

func main() {
    ueACRequestData := *openapiclient.NewUeACRequestData([]openapiclient.UeACRequestInfo{*openapiclient.NewUeACRequestInfo("Supi_example", openapiclient.AccessType("3GPP_ACCESS"), []openapiclient.AcuOperationItem{*openapiclient.NewAcuOperationItem(*openapiclient.NewAcuFlag(), *openapiclient.NewSnssai(int32(123)))})}, "NfId_example") // UeACRequestData | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SliceCollectionApi.NumOfUEsUpdate(context.Background()).UeACRequestData(ueACRequestData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SliceCollectionApi.NumOfUEsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NumOfUEsUpdate`: UeACResponseData
    fmt.Fprintf(os.Stdout, "Response from `SliceCollectionApi.NumOfUEsUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNumOfUEsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ueACRequestData** | [**UeACRequestData**](UeACRequestData.md) |  | 

### Return type

[**UeACResponseData**](UeACResponseData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

