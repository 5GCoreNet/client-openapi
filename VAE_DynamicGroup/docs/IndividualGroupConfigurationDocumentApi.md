# \IndividualGroupConfigurationDocumentApi

All URIs are relative to *https://example.com/vae-dynamic-group/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteGroupConfiguration**](IndividualGroupConfigurationDocumentApi.md#DeleteGroupConfiguration) | **Delete** /group-configurations/{configId} | VAE Group Configuration resource delete service Operation
[**ReadDynamicGroupConfiguration**](IndividualGroupConfigurationDocumentApi.md#ReadDynamicGroupConfiguration) | **Get** /group-configurations/{configId} | VAE Group Configuration resource read service Operation



## DeleteGroupConfiguration

> DeleteGroupConfiguration(ctx, configId).Execute()

VAE Group Configuration resource delete service Operation

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/VAE_DynamicGroup"
)

func main() {
    configId := "configId_example" // string | Unique ID of the group configuration to be deleted.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualGroupConfigurationDocumentApi.DeleteGroupConfiguration(context.Background(), configId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualGroupConfigurationDocumentApi.DeleteGroupConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configId** | **string** | Unique ID of the group configuration to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGroupConfigurationRequest struct via the builder pattern


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


## ReadDynamicGroupConfiguration

> GroupConfigurationData ReadDynamicGroupConfiguration(ctx, configId).Execute()

VAE Group Configuration resource read service Operation

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/VAE_DynamicGroup"
)

func main() {
    configId := "configId_example" // string | Identifier of an group configuration resource.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualGroupConfigurationDocumentApi.ReadDynamicGroupConfiguration(context.Background(), configId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualGroupConfigurationDocumentApi.ReadDynamicGroupConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadDynamicGroupConfiguration`: GroupConfigurationData
    fmt.Fprintf(os.Stdout, "Response from `IndividualGroupConfigurationDocumentApi.ReadDynamicGroupConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configId** | **string** | Identifier of an group configuration resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadDynamicGroupConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GroupConfigurationData**](GroupConfigurationData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

