# \SendMTSMSMessageAndTheDeliveryReportApi

All URIs are relative to *https://example.com/nipsmgw-smservice/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SendSMS**](SendMTSMSMessageAndTheDeliveryReportApi.md#SendSMS) | **Post** /mt-sm-infos/{gpsi}/sendsms | Send SMS payload for a given UE



## SendSMS

> SendSMS200Response SendSMS(ctx, gpsi).JsonData(jsonData).BinaryPayload(binaryPayload).Execute()

Send SMS payload for a given UE

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
    gpsi := "gpsi_example" // string | Generic Public Subscription Identifier (GPSI)
    jsonData := *openapiclient.NewSmsData(*openapiclient.NewRefToBinaryData("ContentId_example")) // SmsData |  (optional)
    binaryPayload := os.NewFile(1234, "some_file") // *os.File |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SendMTSMSMessageAndTheDeliveryReportApi.SendSMS(context.Background(), gpsi).JsonData(jsonData).BinaryPayload(binaryPayload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SendMTSMSMessageAndTheDeliveryReportApi.SendSMS``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SendSMS`: SendSMS200Response
    fmt.Fprintf(os.Stdout, "Response from `SendMTSMSMessageAndTheDeliveryReportApi.SendSMS`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gpsi** | **string** | Generic Public Subscription Identifier (GPSI) | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendSMSRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jsonData** | [**SmsData**](SmsData.md) |  | 
 **binaryPayload** | ***os.File** |  | 

### Return type

[**SendSMS200Response**](SendSMS200Response.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: multipart/related
- **Accept**: multipart/related, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

