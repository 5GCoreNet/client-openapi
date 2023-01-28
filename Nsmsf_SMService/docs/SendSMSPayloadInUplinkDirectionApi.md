# \SendSMSPayloadInUplinkDirectionApi

All URIs are relative to *https://example.com/nsmsf-sms/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SendSMS**](SendSMSPayloadInUplinkDirectionApi.md#SendSMS) | **Post** /ue-contexts/{supi}/sendsms | Send SMS payload for a given UE



## SendSMS

> SmsRecordDeliveryData SendSMS(ctx, supi).JsonData(jsonData).BinaryPayload(binaryPayload).Execute()

Send SMS payload for a given UE

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nsmsf_SMService"
)

func main() {
    supi := "supi_example" // string | Subscriber Permanent Identifier (SUPI)
    jsonData := *openapiclient.NewSmsRecordData("SmsRecordId_example", *openapiclient.NewRefToBinaryData("ContentId_example")) // SmsRecordData |  (optional)
    binaryPayload := os.NewFile(1234, "some_file") // *os.File |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SendSMSPayloadInUplinkDirectionApi.SendSMS(context.Background(), supi).JsonData(jsonData).BinaryPayload(binaryPayload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SendSMSPayloadInUplinkDirectionApi.SendSMS``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SendSMS`: SmsRecordDeliveryData
    fmt.Fprintf(os.Stdout, "Response from `SendSMSPayloadInUplinkDirectionApi.SendSMS`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supi** | **string** | Subscriber Permanent Identifier (SUPI) | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendSMSRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jsonData** | [**SmsRecordData**](SmsRecordData.md) |  | 
 **binaryPayload** | ***os.File** |  | 

### Return type

[**SmsRecordDeliveryData**](SmsRecordDeliveryData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: multipart/related
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

