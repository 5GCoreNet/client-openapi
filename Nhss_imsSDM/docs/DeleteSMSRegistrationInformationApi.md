# \DeleteSMSRegistrationInformationApi

All URIs are relative to *https://example.com/nhss-ims-sdm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSmsRegistrationInfo**](DeleteSMSRegistrationInformationApi.md#DeleteSmsRegistrationInfo) | **Delete** /{imsUeId}/service-data/sms-registration-info | delete the SMS registration information



## DeleteSmsRegistrationInfo

> DeleteSmsRegistrationInfo(ctx, imsUeId).PrivateIdentity(privateIdentity).Execute()

delete the SMS registration information

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
    imsUeId := "imsUeId_example" // string | Identifier of the UE
    privateIdentity := "privateIdentity_example" // string | IMS Private Identity (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeleteSMSRegistrationInformationApi.DeleteSmsRegistrationInfo(context.Background(), imsUeId).PrivateIdentity(privateIdentity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeleteSMSRegistrationInformationApi.DeleteSmsRegistrationInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imsUeId** | **string** | Identifier of the UE | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSmsRegistrationInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **privateIdentity** | **string** | IMS Private Identity | 

### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

