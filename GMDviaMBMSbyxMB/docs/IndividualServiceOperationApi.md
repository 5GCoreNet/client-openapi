# \IndividualServiceOperationApi

All URIs are relative to *https://example.com/3gpp-group-message-delivery-xmb/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletexMBService**](IndividualServiceOperationApi.md#DeletexMBService) | **Delete** /{scsAsId}/services/{serviceId} | Deletes an existing service resource for a given SCS/AS and a service id.
[**FetchIndxMBService**](IndividualServiceOperationApi.md#FetchIndxMBService) | **Get** /{scsAsId}/services/{serviceId} | Read a service resource for a given SCS/AS and a Service Id.



## DeletexMBService

> DeletexMBService(ctx, scsAsId, serviceId).Execute()

Deletes an existing service resource for a given SCS/AS and a service id.

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
    scsAsId := "scsAsId_example" // string | Identifier of SCS/AS
    serviceId := "serviceId_example" // string | Service Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualServiceOperationApi.DeletexMBService(context.Background(), scsAsId, serviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualServiceOperationApi.DeletexMBService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of SCS/AS | 
**serviceId** | **string** | Service Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletexMBServiceRequest struct via the builder pattern


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


## FetchIndxMBService

> ServiceCreation FetchIndxMBService(ctx, scsAsId, serviceId).Execute()

Read a service resource for a given SCS/AS and a Service Id.

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
    scsAsId := "scsAsId_example" // string | Identifier of SCS/AS
    serviceId := "serviceId_example" // string | Service Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualServiceOperationApi.FetchIndxMBService(context.Background(), scsAsId, serviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualServiceOperationApi.FetchIndxMBService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchIndxMBService`: ServiceCreation
    fmt.Fprintf(os.Stdout, "Response from `IndividualServiceOperationApi.FetchIndxMBService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of SCS/AS | 
**serviceId** | **string** | Service Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchIndxMBServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ServiceCreation**](ServiceCreation.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

