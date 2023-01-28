# \DefaultApi

All URIs are relative to *https://example.com/api-provider-management/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RegistrationsPost**](DefaultApi.md#RegistrationsPost) | **Post** /registrations | 
[**RegistrationsRegistrationIdDelete**](DefaultApi.md#RegistrationsRegistrationIdDelete) | **Delete** /registrations/{registrationId} | 
[**RegistrationsRegistrationIdPut**](DefaultApi.md#RegistrationsRegistrationIdPut) | **Put** /registrations/{registrationId} | 



## RegistrationsPost

> APIProviderEnrolmentDetails RegistrationsPost(ctx).APIProviderEnrolmentDetails(aPIProviderEnrolmentDetails).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/CAPIF_API_Provider_Management_API"
)

func main() {
    aPIProviderEnrolmentDetails := *openapiclient.NewAPIProviderEnrolmentDetails("RegSec_example") // APIProviderEnrolmentDetails | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.RegistrationsPost(context.Background()).APIProviderEnrolmentDetails(aPIProviderEnrolmentDetails).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RegistrationsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegistrationsPost`: APIProviderEnrolmentDetails
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.RegistrationsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegistrationsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aPIProviderEnrolmentDetails** | [**APIProviderEnrolmentDetails**](APIProviderEnrolmentDetails.md) |  | 

### Return type

[**APIProviderEnrolmentDetails**](APIProviderEnrolmentDetails.md)

### Authorization

No authorization required

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
    openapiclient "github.com/5GCoreNet/client-openapi/CAPIF_API_Provider_Management_API"
)

func main() {
    registrationId := "registrationId_example" // string | String identifying an registered API provider domain resource.

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
**registrationId** | **string** | String identifying an registered API provider domain resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegistrationsRegistrationIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegistrationsRegistrationIdPut

> APIProviderEnrolmentDetails RegistrationsRegistrationIdPut(ctx, registrationId).APIProviderEnrolmentDetails(aPIProviderEnrolmentDetails).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/CAPIF_API_Provider_Management_API"
)

func main() {
    registrationId := "registrationId_example" // string | String identifying an registered API provider domain resource.
    aPIProviderEnrolmentDetails := *openapiclient.NewAPIProviderEnrolmentDetails("RegSec_example") // APIProviderEnrolmentDetails | Representation of the API provider domain registration details to be updated in CAPIF core function. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.RegistrationsRegistrationIdPut(context.Background(), registrationId).APIProviderEnrolmentDetails(aPIProviderEnrolmentDetails).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RegistrationsRegistrationIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegistrationsRegistrationIdPut`: APIProviderEnrolmentDetails
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.RegistrationsRegistrationIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registrationId** | **string** | String identifying an registered API provider domain resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegistrationsRegistrationIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aPIProviderEnrolmentDetails** | [**APIProviderEnrolmentDetails**](APIProviderEnrolmentDetails.md) | Representation of the API provider domain registration details to be updated in CAPIF core function.  | 

### Return type

[**APIProviderEnrolmentDetails**](APIProviderEnrolmentDetails.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

