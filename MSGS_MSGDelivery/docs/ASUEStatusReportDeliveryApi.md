# \ASUEStatusReportDeliveryApi

All URIs are relative to *https://example.com/msgs-msgdelivery/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeliverReportPost**](ASUEStatusReportDeliveryApi.md#DeliverReportPost) | **Post** /deliver-report | AS or UE deliver status report to MSGin5G Server



## DeliverReportPost

> MessageDeliveryAck DeliverReportPost(ctx).DeliveryStatusReport(deliveryStatusReport).Execute()

AS or UE deliver status report to MSGin5G Server

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/MSGS_MSGDelivery"
)

func main() {
    deliveryStatusReport := *openapiclient.NewDeliveryStatusReport(*openapiclient.NewAddress(*openapiclient.NewAddressType(), "Addr_example"), *openapiclient.NewAddress(*openapiclient.NewAddressType(), "Addr_example"), "MsgId_example", *openapiclient.NewReportDeliveryStatus()) // DeliveryStatusReport | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ASUEStatusReportDeliveryApi.DeliverReportPost(context.Background()).DeliveryStatusReport(deliveryStatusReport).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ASUEStatusReportDeliveryApi.DeliverReportPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliverReportPost`: MessageDeliveryAck
    fmt.Fprintf(os.Stdout, "Response from `ASUEStatusReportDeliveryApi.DeliverReportPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeliverReportPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deliveryStatusReport** | [**DeliveryStatusReport**](DeliveryStatusReport.md) |  | 

### Return type

[**MessageDeliveryAck**](MessageDeliveryAck.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

