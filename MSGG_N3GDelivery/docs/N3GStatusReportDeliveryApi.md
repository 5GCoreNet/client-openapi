# \N3GStatusReportDeliveryApi

All URIs are relative to *https://example.com/msgg-n3gdelivery/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeliverReportPost**](N3GStatusReportDeliveryApi.md#DeliverReportPost) | **Post** /deliver-report | deliver status report to NON-3GPP Message Gateway



## DeliverReportPost

> DeliverReportPost(ctx).DeliveryStatusReport(deliveryStatusReport).Execute()

deliver status report to NON-3GPP Message Gateway

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/MSGG_N3GDelivery"
)

func main() {
    deliveryStatusReport := *openapiclient.NewDeliveryStatusReport(*openapiclient.NewAddress(*openapiclient.NewAddressType(), "Addr_example"), *openapiclient.NewAddress(*openapiclient.NewAddressType(), "Addr_example"), "MsgId_example", *openapiclient.NewReportDeliveryStatus()) // DeliveryStatusReport | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.N3GStatusReportDeliveryApi.DeliverReportPost(context.Background()).DeliveryStatusReport(deliveryStatusReport).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `N3GStatusReportDeliveryApi.DeliverReportPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeliverReportPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deliveryStatusReport** | [**DeliveryStatusReport**](DeliveryStatusReport.md) |  | 

### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

