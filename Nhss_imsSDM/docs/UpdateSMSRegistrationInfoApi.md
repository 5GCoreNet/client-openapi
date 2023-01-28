# \UpdateSMSRegistrationInfoApi

All URIs are relative to *https://example.com/nhss-ims-sdm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateSmsRegistrationInfo**](UpdateSMSRegistrationInfoApi.md#UpdateSmsRegistrationInfo) | **Put** /{imsUeId}/service-data/sms-registration-info | Update the SMS registration information associated to a user



## UpdateSmsRegistrationInfo

> SmsRegistrationInfo UpdateSmsRegistrationInfo(ctx, imsUeId).IpSmGwAddress(ipSmGwAddress).PrivateIdentity(privateIdentity).Execute()

Update the SMS registration information associated to a user

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nhss_imsSDM"
)

func main() {
    imsUeId := "imsUeId_example" // string | IMS Identity
    ipSmGwAddress := *openapiclient.NewIpSmGwAddress("IpSmGwNumber_example") // IpSmGwAddress | 
    privateIdentity := "privateIdentity_example" // string | IMS Private Identity (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UpdateSMSRegistrationInfoApi.UpdateSmsRegistrationInfo(context.Background(), imsUeId).IpSmGwAddress(ipSmGwAddress).PrivateIdentity(privateIdentity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UpdateSMSRegistrationInfoApi.UpdateSmsRegistrationInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSmsRegistrationInfo`: SmsRegistrationInfo
    fmt.Fprintf(os.Stdout, "Response from `UpdateSMSRegistrationInfoApi.UpdateSmsRegistrationInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imsUeId** | **string** | IMS Identity | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSmsRegistrationInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ipSmGwAddress** | [**IpSmGwAddress**](IpSmGwAddress.md) |  | 
 **privateIdentity** | **string** | IMS Private Identity | 

### Return type

[**SmsRegistrationInfo**](SmsRegistrationInfo.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

