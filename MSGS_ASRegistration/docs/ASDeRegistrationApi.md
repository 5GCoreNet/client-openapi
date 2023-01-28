# \ASDeRegistrationApi

All URIs are relative to *https://example.com/msgs-asregistration/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RegistrationsRegistrationIdDelete**](ASDeRegistrationApi.md#RegistrationsRegistrationIdDelete) | **Delete** /registrations/{registrationId} | Delete an existing AS registration at MSGin5G Server



## RegistrationsRegistrationIdDelete

> ASRegistrationAck RegistrationsRegistrationIdDelete(ctx, registrationId).Execute()

Delete an existing AS registration at MSGin5G Server

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/MSGS_ASRegistration"
)

func main() {
    registrationId := "registrationId_example" // string | AS registration Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ASDeRegistrationApi.RegistrationsRegistrationIdDelete(context.Background(), registrationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ASDeRegistrationApi.RegistrationsRegistrationIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegistrationsRegistrationIdDelete`: ASRegistrationAck
    fmt.Fprintf(os.Stdout, "Response from `ASDeRegistrationApi.RegistrationsRegistrationIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registrationId** | **string** | AS registration Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegistrationsRegistrationIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ASRegistrationAck**](ASRegistrationAck.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

