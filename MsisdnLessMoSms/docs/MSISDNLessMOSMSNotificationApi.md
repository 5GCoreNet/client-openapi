# \MSISDNLessMOSMSNotificationApi

All URIs are relative to *https://example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeliverMSISDNlessMOSMSNotification**](MSISDNLessMOSMSNotificationApi.md#DeliverMSISDNlessMOSMSNotification) | **Post** / | Deliver a received MSIDN-less MO SMS from the SCEF to the SCS/AS.



## DeliverMSISDNlessMOSMSNotification

> MsisdnLessMoSmsNotificationReply DeliverMSISDNlessMOSMSNotification(ctx).MsisdnLessMoSmsNotification(msisdnLessMoSmsNotification).Execute()

Deliver a received MSIDN-less MO SMS from the SCEF to the SCS/AS.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/MsisdnLessMoSms"
)

func main() {
    msisdnLessMoSmsNotification := *openapiclient.NewMsisdnLessMoSmsNotification("SupportedFeatures_example", "Sms_example", "ExternalId_example", int32(123)) // MsisdnLessMoSmsNotification | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MSISDNLessMOSMSNotificationApi.DeliverMSISDNlessMOSMSNotification(context.Background()).MsisdnLessMoSmsNotification(msisdnLessMoSmsNotification).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MSISDNLessMOSMSNotificationApi.DeliverMSISDNlessMOSMSNotification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliverMSISDNlessMOSMSNotification`: MsisdnLessMoSmsNotificationReply
    fmt.Fprintf(os.Stdout, "Response from `MSISDNLessMOSMSNotificationApi.DeliverMSISDNlessMOSMSNotification`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeliverMSISDNlessMOSMSNotificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msisdnLessMoSmsNotification** | [**MsisdnLessMoSmsNotification**](MsisdnLessMoSmsNotification.md) |  | 

### Return type

[**MsisdnLessMoSmsNotificationReply**](MsisdnLessMoSmsNotificationReply.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

