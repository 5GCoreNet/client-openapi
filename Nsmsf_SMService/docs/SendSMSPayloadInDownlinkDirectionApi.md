# \SendSMSPayloadInDownlinkDirectionApi

All URIs are relative to *https://example.com/nsmsf-sms/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SendMtSMS**](SendSMSPayloadInDownlinkDirectionApi.md#SendMtSMS) | **Post** /ue-contexts/{supi}/send-mt-sms | Send MT SMS payload for a given UE



## SendMtSMS

> SendMtSMS200Response SendMtSMS(ctx, supi).JsonData(jsonData).BinaryPayload(binaryPayload).Execute()

Send MT SMS payload for a given UE

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
    supi := "supi_example" // string | Subscriber Permanent Identifier (SUPI)
    jsonData := *openapiclient.NewSmsData(*openapiclient.NewRefToBinaryData("ContentId_example")) // SmsData |  (optional)
    binaryPayload := os.NewFile(1234, "some_file") // *os.File |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SendSMSPayloadInDownlinkDirectionApi.SendMtSMS(context.Background(), supi).JsonData(jsonData).BinaryPayload(binaryPayload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SendSMSPayloadInDownlinkDirectionApi.SendMtSMS``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SendMtSMS`: SendMtSMS200Response
    fmt.Fprintf(os.Stdout, "Response from `SendSMSPayloadInDownlinkDirectionApi.SendMtSMS`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supi** | **string** | Subscriber Permanent Identifier (SUPI) | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendMtSMSRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jsonData** | [**SmsData**](SmsData.md) |  | 
 **binaryPayload** | ***os.File** |  | 

### Return type

[**SendMtSMS200Response**](SendMtSMS200Response.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: multipart/related
- **Accept**: multipart/related, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

