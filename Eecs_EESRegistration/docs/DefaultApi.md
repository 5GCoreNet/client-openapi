# \DefaultApi

All URIs are relative to *https://example.com/eecs-eesregistration/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RegistrationsPost**](DefaultApi.md#RegistrationsPost) | **Post** /registrations | 
[**RegistrationsRegistrationIdDelete**](DefaultApi.md#RegistrationsRegistrationIdDelete) | **Delete** /registrations/{registrationId} | 
[**RegistrationsRegistrationIdGet**](DefaultApi.md#RegistrationsRegistrationIdGet) | **Get** /registrations/{registrationId} | 
[**RegistrationsRegistrationIdPatch**](DefaultApi.md#RegistrationsRegistrationIdPatch) | **Patch** /registrations/{registrationId} | 
[**RegistrationsRegistrationIdPut**](DefaultApi.md#RegistrationsRegistrationIdPut) | **Put** /registrations/{registrationId} | 



## RegistrationsPost

> EESRegistration RegistrationsPost(ctx).EESRegistration(eESRegistration).Execute()





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
    eESRegistration := *openapiclient.NewEESRegistration(*openapiclient.NewEESProfile("EesId_example", openapiclient.EndPoint{Interface{}: new(interface{})}, false)) // EESRegistration | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.RegistrationsPost(context.Background()).EESRegistration(eESRegistration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RegistrationsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegistrationsPost`: EESRegistration
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.RegistrationsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegistrationsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eESRegistration** | [**EESRegistration**](EESRegistration.md) |  | 

### Return type

[**EESRegistration**](EESRegistration.md)

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
    openapiclient "./openapi"
)

func main() {
    registrationId := "registrationId_example" // string | Registration Id.

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
**registrationId** | **string** | Registration Id. | 

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

> EESRegistration RegistrationsRegistrationIdGet(ctx, registrationId).Execute()





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
    registrationId := "registrationId_example" // string | Registration Id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.RegistrationsRegistrationIdGet(context.Background(), registrationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RegistrationsRegistrationIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegistrationsRegistrationIdGet`: EESRegistration
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

[**EESRegistration**](EESRegistration.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegistrationsRegistrationIdPatch

> EESRegistration RegistrationsRegistrationIdPatch(ctx, registrationId).EESRegistrationPatch(eESRegistrationPatch).Execute()





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
    registrationId := "registrationId_example" // string | EES registration Id.
    eESRegistrationPatch := *openapiclient.NewEESRegistrationPatch() // EESRegistrationPatch | Partial update an existing EES registration resource.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.RegistrationsRegistrationIdPatch(context.Background(), registrationId).EESRegistrationPatch(eESRegistrationPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RegistrationsRegistrationIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegistrationsRegistrationIdPatch`: EESRegistration
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.RegistrationsRegistrationIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registrationId** | **string** | EES registration Id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegistrationsRegistrationIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eESRegistrationPatch** | [**EESRegistrationPatch**](EESRegistrationPatch.md) | Partial update an existing EES registration resource. | 

### Return type

[**EESRegistration**](EESRegistration.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegistrationsRegistrationIdPut

> EESRegistration RegistrationsRegistrationIdPut(ctx, registrationId).EESRegistration(eESRegistration).Execute()





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
    registrationId := "registrationId_example" // string | EES Registration Id.
    eESRegistration := *openapiclient.NewEESRegistration(*openapiclient.NewEESProfile("EesId_example", openapiclient.EndPoint{Interface{}: new(interface{})}, false)) // EESRegistration | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.RegistrationsRegistrationIdPut(context.Background(), registrationId).EESRegistration(eESRegistration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RegistrationsRegistrationIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegistrationsRegistrationIdPut`: EESRegistration
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.RegistrationsRegistrationIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registrationId** | **string** | EES Registration Id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegistrationsRegistrationIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eESRegistration** | [**EESRegistration**](EESRegistration.md) |  | 

### Return type

[**EESRegistration**](EESRegistration.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

