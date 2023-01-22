# \IndividualApplicationRequirementDocumentApi

All URIs are relative to *https://example.com/vae-app-req/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApplicationRequirement**](IndividualApplicationRequirementDocumentApi.md#DeleteApplicationRequirement) | **Delete** /application-requirements/{requirementId} | VAE Application Requirement resource delete service Operation
[**ReadApplicationRequirement**](IndividualApplicationRequirementDocumentApi.md#ReadApplicationRequirement) | **Get** /application-requirements/{requirementId} | VAE Application Requirement resource read service Operation



## DeleteApplicationRequirement

> DeleteApplicationRequirement(ctx, requirementId).Execute()

VAE Application Requirement resource delete service Operation

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
    requirementId := "requirementId_example" // string | Unique ID of the application requirement to be deleted

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualApplicationRequirementDocumentApi.DeleteApplicationRequirement(context.Background(), requirementId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualApplicationRequirementDocumentApi.DeleteApplicationRequirement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requirementId** | **string** | Unique ID of the application requirement to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApplicationRequirementRequest struct via the builder pattern


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


## ReadApplicationRequirement

> ApplicationRequirementData ReadApplicationRequirement(ctx, requirementId).Execute()

VAE Application Requirement resource read service Operation

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
    requirementId := "requirementId_example" // string | Identifier of an application requirement resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualApplicationRequirementDocumentApi.ReadApplicationRequirement(context.Background(), requirementId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualApplicationRequirementDocumentApi.ReadApplicationRequirement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadApplicationRequirement`: ApplicationRequirementData
    fmt.Fprintf(os.Stdout, "Response from `IndividualApplicationRequirementDocumentApi.ReadApplicationRequirement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requirementId** | **string** | Identifier of an application requirement resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadApplicationRequirementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApplicationRequirementData**](ApplicationRequirementData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

