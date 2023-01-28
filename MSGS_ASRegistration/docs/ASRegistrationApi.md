# \ASRegistrationApi

All URIs are relative to *https://example.com/msgs-asregistration/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RegistrationsPost**](ASRegistrationApi.md#RegistrationsPost) | **Post** /registrations | Registers a new AS at a MSGin5G Server



## RegistrationsPost

> ASRegistrationAck RegistrationsPost(ctx).ASRegistration(aSRegistration).Execute()

Registers a new AS at a MSGin5G Server

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
    aSRegistration := *openapiclient.NewASRegistration("AsSvcId_example") // ASRegistration | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ASRegistrationApi.RegistrationsPost(context.Background()).ASRegistration(aSRegistration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ASRegistrationApi.RegistrationsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegistrationsPost`: ASRegistrationAck
    fmt.Fprintf(os.Stdout, "Response from `ASRegistrationApi.RegistrationsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegistrationsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aSRegistration** | [**ASRegistration**](ASRegistration.md) |  | 

### Return type

[**ASRegistrationAck**](ASRegistrationAck.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

