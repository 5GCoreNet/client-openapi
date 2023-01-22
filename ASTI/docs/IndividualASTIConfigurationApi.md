# \IndividualASTIConfigurationApi

All URIs are relative to *https://example.com/3gpp-asti/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAnConfiguration**](IndividualASTIConfigurationApi.md#DeleteAnConfiguration) | **Delete** /{afId}/configurations/{configId} | Deletes an already existing configuration
[**FullyModifyAnConfiguration**](IndividualASTIConfigurationApi.md#FullyModifyAnConfiguration) | **Put** /{afId}/configurations/{configId} | Modifies an active configuration for the AF and the configuration Id
[**ReadAnConfiguration**](IndividualASTIConfigurationApi.md#ReadAnConfiguration) | **Get** /{afId}/configurations/{configId} | Reads an active configuration for the AF and the configuration Id



## DeleteAnConfiguration

> DeleteAnConfiguration(ctx, afId, configId).Execute()

Deletes an already existing configuration

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
    afId := "afId_example" // string | Identifier of the AF
    configId := "configId_example" // string | Identifier of the configuration resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualASTIConfigurationApi.DeleteAnConfiguration(context.Background(), afId, configId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualASTIConfigurationApi.DeleteAnConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**afId** | **string** | Identifier of the AF | 
**configId** | **string** | Identifier of the configuration resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAnConfigurationRequest struct via the builder pattern


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


## FullyModifyAnConfiguration

> AccessTimeDistributionData FullyModifyAnConfiguration(ctx, afId, configId).AccessTimeDistributionData(accessTimeDistributionData).Execute()

Modifies an active configuration for the AF and the configuration Id

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
    afId := "afId_example" // string | Identifier of the AF
    configId := "configId_example" // string | Identifier of the configuration resource
    accessTimeDistributionData := openapiclient.AccessTimeDistributionData{Interface{}: new(interface{})} // AccessTimeDistributionData | Parameters to update/replace the existing configuration

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualASTIConfigurationApi.FullyModifyAnConfiguration(context.Background(), afId, configId).AccessTimeDistributionData(accessTimeDistributionData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualASTIConfigurationApi.FullyModifyAnConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FullyModifyAnConfiguration`: AccessTimeDistributionData
    fmt.Fprintf(os.Stdout, "Response from `IndividualASTIConfigurationApi.FullyModifyAnConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**afId** | **string** | Identifier of the AF | 
**configId** | **string** | Identifier of the configuration resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiFullyModifyAnConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accessTimeDistributionData** | [**AccessTimeDistributionData**](AccessTimeDistributionData.md) | Parameters to update/replace the existing configuration | 

### Return type

[**AccessTimeDistributionData**](AccessTimeDistributionData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadAnConfiguration

> AccessTimeDistributionData ReadAnConfiguration(ctx, afId, configId).Execute()

Reads an active configuration for the AF and the configuration Id

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
    afId := "afId_example" // string | Identifier of the AF
    configId := "configId_example" // string | Identifier of the configuration resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualASTIConfigurationApi.ReadAnConfiguration(context.Background(), afId, configId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualASTIConfigurationApi.ReadAnConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAnConfiguration`: AccessTimeDistributionData
    fmt.Fprintf(os.Stdout, "Response from `IndividualASTIConfigurationApi.ReadAnConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**afId** | **string** | Identifier of the AF | 
**configId** | **string** | Identifier of the configuration resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAnConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AccessTimeDistributionData**](AccessTimeDistributionData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

