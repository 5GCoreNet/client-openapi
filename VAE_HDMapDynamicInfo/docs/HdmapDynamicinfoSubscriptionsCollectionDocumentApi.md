# \HdmapDynamicinfoSubscriptionsCollectionDocumentApi

All URIs are relative to *https://example.com/vae-hdmap-dynamic-info/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](HdmapDynamicinfoSubscriptionsCollectionDocumentApi.md#Create) | **Post** /subscriptions | VAE_HDMapDynamicInfo resource create service Operation



## Create

> HdMapDynamicInfoData Create(ctx).HdMapDynamicInfoData(hdMapDynamicInfoData).Execute()

VAE_HDMapDynamicInfo resource create service Operation

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
    hdMapDynamicInfoData := *openapiclient.NewHdMapDynamicInfoData("UeId_example", "NotifUri_example", int32(123)) // HdMapDynamicInfoData | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HdmapDynamicinfoSubscriptionsCollectionDocumentApi.Create(context.Background()).HdMapDynamicInfoData(hdMapDynamicInfoData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HdmapDynamicinfoSubscriptionsCollectionDocumentApi.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create`: HdMapDynamicInfoData
    fmt.Fprintf(os.Stdout, "Response from `HdmapDynamicinfoSubscriptionsCollectionDocumentApi.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hdMapDynamicInfoData** | [**HdMapDynamicInfoData**](HdMapDynamicInfoData.md) |  | 

### Return type

[**HdMapDynamicInfoData**](HdMapDynamicInfoData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

