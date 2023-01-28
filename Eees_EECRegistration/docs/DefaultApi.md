# \DefaultApi

All URIs are relative to *https://example.com/eees-eecregistration/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RegistrationsPost**](DefaultApi.md#RegistrationsPost) | **Post** /registrations | 
[**RegistrationsRegistrationIdDelete**](DefaultApi.md#RegistrationsRegistrationIdDelete) | **Delete** /registrations/{registrationId} | 
[**RegistrationsRegistrationIdPatch**](DefaultApi.md#RegistrationsRegistrationIdPatch) | **Patch** /registrations/{registrationId} | 
[**RegistrationsRegistrationIdPut**](DefaultApi.md#RegistrationsRegistrationIdPut) | **Put** /registrations/{registrationId} | 



## RegistrationsPost

> EECRegistration RegistrationsPost(ctx).EECRegistration(eECRegistration).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Eees_EECRegistration"
)

func main() {
    eECRegistration := *openapiclient.NewEECRegistration("EecId_example") // EECRegistration | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.RegistrationsPost(context.Background()).EECRegistration(eECRegistration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RegistrationsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegistrationsPost`: EECRegistration
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.RegistrationsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegistrationsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eECRegistration** | [**EECRegistration**](EECRegistration.md) |  | 

### Return type

[**EECRegistration**](EECRegistration.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegistrationsRegistrationIdDelete

> RegistrationsRegistrationIdDelete(ctx, registrationId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Eees_EECRegistration"
)

func main() {
    registrationId := "registrationId_example" // string | Identifies an individual EEC registration

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.RegistrationsRegistrationIdDelete(context.Background(), registrationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RegistrationsRegistrationIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registrationId** | **string** | Identifies an individual EEC registration | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegistrationsRegistrationIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegistrationsRegistrationIdPatch

> EECRegistration RegistrationsRegistrationIdPatch(ctx, registrationId).EECRegistrationPatch(eECRegistrationPatch).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Eees_EECRegistration"
)

func main() {
    registrationId := "registrationId_example" // string | Identifies an individual EEC registration
    eECRegistrationPatch := *openapiclient.NewEECRegistrationPatch() // EECRegistrationPatch | Parameters to replace the existing registration

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.RegistrationsRegistrationIdPatch(context.Background(), registrationId).EECRegistrationPatch(eECRegistrationPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RegistrationsRegistrationIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegistrationsRegistrationIdPatch`: EECRegistration
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.RegistrationsRegistrationIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registrationId** | **string** | Identifies an individual EEC registration | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegistrationsRegistrationIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eECRegistrationPatch** | [**EECRegistrationPatch**](EECRegistrationPatch.md) | Parameters to replace the existing registration | 

### Return type

[**EECRegistration**](EECRegistration.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegistrationsRegistrationIdPut

> EECRegistration RegistrationsRegistrationIdPut(ctx, registrationId).EECRegistration(eECRegistration).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Eees_EECRegistration"
)

func main() {
    registrationId := "registrationId_example" // string | Identifies an individual EEC registration
    eECRegistration := *openapiclient.NewEECRegistration("EecId_example") // EECRegistration | Parameters to replace the existing registration

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.RegistrationsRegistrationIdPut(context.Background(), registrationId).EECRegistration(eECRegistration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RegistrationsRegistrationIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegistrationsRegistrationIdPut`: EECRegistration
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.RegistrationsRegistrationIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registrationId** | **string** | Identifies an individual EEC registration | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegistrationsRegistrationIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eECRegistration** | [**EECRegistration**](EECRegistration.md) | Parameters to replace the existing registration | 

### Return type

[**EECRegistration**](EECRegistration.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

