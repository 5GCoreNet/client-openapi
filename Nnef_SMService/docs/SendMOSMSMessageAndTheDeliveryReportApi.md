# \SendMOSMSMessageAndTheDeliveryReportApi

All URIs are relative to *https://example.com/nnef-smservice/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SendSMS**](SendMOSMSMessageAndTheDeliveryReportApi.md#SendSMS) | **Post** /sm-contexts/{supi}/sendsms | Send SMS payload for a given UE



## SendSMS

> SendSMS200Response SendSMS(ctx, supi).JsonData(jsonData).BinaryPayload(binaryPayload).Execute()

Send SMS payload for a given UE

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nnef_SMService"
)

func main() {
    supi := "supi_example" // string | Subscription Permanent Identifier (SUPI)
    jsonData := *openapiclient.NewSmsData(*openapiclient.NewRefToBinaryData("ContentId_example")) // SmsData |  (optional)
    binaryPayload := os.NewFile(1234, "some_file") // *os.File |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SendMOSMSMessageAndTheDeliveryReportApi.SendSMS(context.Background(), supi).JsonData(jsonData).BinaryPayload(binaryPayload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SendMOSMSMessageAndTheDeliveryReportApi.SendSMS``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SendSMS`: SendSMS200Response
    fmt.Fprintf(os.Stdout, "Response from `SendMOSMSMessageAndTheDeliveryReportApi.SendSMS`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supi** | **string** | Subscription Permanent Identifier (SUPI) | 

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
- **Accept**: multipart/related, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

