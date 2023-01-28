# \ServiceOperationApi

All URIs are relative to *https://example.com/3gpp-group-message-delivery-xmb/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatexMBService**](ServiceOperationApi.md#CreatexMBService) | **Post** /{scsAsId}/services | Creates a new service creation resource for a given SCS/AS.
[**FetchAllxMBServices**](ServiceOperationApi.md#FetchAllxMBServices) | **Get** /{scsAsId}/services | Read all service resources for a given SCS/AS.



## CreatexMBService

> ServiceCreation CreatexMBService(ctx, scsAsId).ServiceCreation(serviceCreation).Execute()

Creates a new service creation resource for a given SCS/AS.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/GMDviaMBMSbyxMB"
)

func main() {
    scsAsId := "scsAsId_example" // string | Identifier of SCS/AS
    serviceCreation := *openapiclient.NewServiceCreation() // ServiceCreation | representation of the service to be created in the SCEF

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceOperationApi.CreatexMBService(context.Background(), scsAsId).ServiceCreation(serviceCreation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceOperationApi.CreatexMBService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatexMBService`: ServiceCreation
    fmt.Fprintf(os.Stdout, "Response from `ServiceOperationApi.CreatexMBService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of SCS/AS | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatexMBServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serviceCreation** | [**ServiceCreation**](ServiceCreation.md) | representation of the service to be created in the SCEF | 

### Return type

[**ServiceCreation**](ServiceCreation.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchAllxMBServices

> []ServiceCreation FetchAllxMBServices(ctx, scsAsId).Execute()

Read all service resources for a given SCS/AS.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/GMDviaMBMSbyxMB"
)

func main() {
    scsAsId := "scsAsId_example" // string | Identifier of SCS/AS

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceOperationApi.FetchAllxMBServices(context.Background(), scsAsId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceOperationApi.FetchAllxMBServices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchAllxMBServices`: []ServiceCreation
    fmt.Fprintf(os.Stdout, "Response from `ServiceOperationApi.FetchAllxMBServices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of SCS/AS | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchAllxMBServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ServiceCreation**](ServiceCreation.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

