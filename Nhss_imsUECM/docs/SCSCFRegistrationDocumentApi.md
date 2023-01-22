# \SCSCFRegistrationDocumentApi

All URIs are relative to *https://example.com/nhss-ims-uecm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SCSCFRegistration**](SCSCFRegistrationDocumentApi.md#SCSCFRegistration) | **Put** /{imsUeId}/scscf-registration | SCSCF registration information



## SCSCFRegistration

> ScscfRegistration SCSCFRegistration(ctx, imsUeId).ScscfRegistration(scscfRegistration).Execute()

SCSCF registration information

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
    imsUeId := "imsUeId_example" // string | IMS Identity
    scscfRegistration := *openapiclient.NewScscfRegistration(*openapiclient.NewImsRegistrationType(), "CscfServerName_example") // ScscfRegistration | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SCSCFRegistrationDocumentApi.SCSCFRegistration(context.Background(), imsUeId).ScscfRegistration(scscfRegistration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SCSCFRegistrationDocumentApi.SCSCFRegistration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SCSCFRegistration`: ScscfRegistration
    fmt.Fprintf(os.Stdout, "Response from `SCSCFRegistrationDocumentApi.SCSCFRegistration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imsUeId** | **string** | IMS Identity | 

### Other Parameters

Other parameters are passed through a pointer to a apiSCSCFRegistrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scscfRegistration** | [**ScscfRegistration**](ScscfRegistration.md) |  | 

### Return type

[**ScscfRegistration**](ScscfRegistration.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

