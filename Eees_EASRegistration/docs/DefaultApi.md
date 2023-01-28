# \DefaultApi

All URIs are relative to *https://example.com/eees-easregistration/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RegistrationsPost**](DefaultApi.md#RegistrationsPost) | **Post** /registrations | 
[**RegistrationsRegistrationIdDelete**](DefaultApi.md#RegistrationsRegistrationIdDelete) | **Delete** /registrations/{registrationId} | 
[**RegistrationsRegistrationIdGet**](DefaultApi.md#RegistrationsRegistrationIdGet) | **Get** /registrations/{registrationId} | 
[**RegistrationsRegistrationIdPatch**](DefaultApi.md#RegistrationsRegistrationIdPatch) | **Patch** /registrations/{registrationId} | 
[**RegistrationsRegistrationIdPut**](DefaultApi.md#RegistrationsRegistrationIdPut) | **Put** /registrations/{registrationId} | 



## RegistrationsPost

> EASRegistration RegistrationsPost(ctx).EASRegistration(eASRegistration).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Eees_EASRegistration"
)

func main() {
    eASRegistration := *openapiclient.NewEASRegistration(*openapiclient.NewEASProfile("EasId_example", openapiclient.EndPoint{Interface{}: new(interface{})})) // EASRegistration | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.RegistrationsPost(context.Background()).EASRegistration(eASRegistration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RegistrationsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegistrationsPost`: EASRegistration
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.RegistrationsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegistrationsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eASRegistration** | [**EASRegistration**](EASRegistration.md) |  | 

### Return type

[**EASRegistration**](EASRegistration.md)

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
    openapiclient "github.com/5GCoreNet/client-openapi/Eees_EASRegistration"
)

func main() {
    registrationId := "registrationId_example" // string | EAS registration Id.

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
**registrationId** | **string** | EAS registration Id. | 

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


## RegistrationsRegistrationIdGet

> EASRegistration RegistrationsRegistrationIdGet(ctx, registrationId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Eees_EASRegistration"
)

func main() {
    registrationId := "registrationId_example" // string | Registration Id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.RegistrationsRegistrationIdGet(context.Background(), registrationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RegistrationsRegistrationIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegistrationsRegistrationIdGet`: EASRegistration
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.RegistrationsRegistrationIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registrationId** | **string** | Registration Id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegistrationsRegistrationIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EASRegistration**](EASRegistration.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegistrationsRegistrationIdPatch

> EASRegistration RegistrationsRegistrationIdPatch(ctx, registrationId).EASRegistrationPatch(eASRegistrationPatch).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Eees_EASRegistration"
)

func main() {
    registrationId := "registrationId_example" // string | EAS registration Id.
    eASRegistrationPatch := *openapiclient.NewEASRegistrationPatch() // EASRegistrationPatch | Partial update of an existing EAS registration resource.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.RegistrationsRegistrationIdPatch(context.Background(), registrationId).EASRegistrationPatch(eASRegistrationPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RegistrationsRegistrationIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegistrationsRegistrationIdPatch`: EASRegistration
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.RegistrationsRegistrationIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registrationId** | **string** | EAS registration Id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegistrationsRegistrationIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eASRegistrationPatch** | [**EASRegistrationPatch**](EASRegistrationPatch.md) | Partial update of an existing EAS registration resource. | 

### Return type

[**EASRegistration**](EASRegistration.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegistrationsRegistrationIdPut

> EASRegistration RegistrationsRegistrationIdPut(ctx, registrationId).EASRegistration(eASRegistration).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Eees_EASRegistration"
)

func main() {
    registrationId := "registrationId_example" // string | EAS registration Id.
    eASRegistration := *openapiclient.NewEASRegistration(*openapiclient.NewEASProfile("EasId_example", openapiclient.EndPoint{Interface{}: new(interface{})})) // EASRegistration | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.RegistrationsRegistrationIdPut(context.Background(), registrationId).EASRegistration(eASRegistration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RegistrationsRegistrationIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegistrationsRegistrationIdPut`: EASRegistration
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.RegistrationsRegistrationIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registrationId** | **string** | EAS registration Id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegistrationsRegistrationIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eASRegistration** | [**EASRegistration**](EASRegistration.md) |  | 

### Return type

[**EASRegistration**](EASRegistration.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

