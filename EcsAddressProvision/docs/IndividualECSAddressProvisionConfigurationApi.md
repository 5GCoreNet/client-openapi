# \IndividualECSAddressProvisionConfigurationApi

All URIs are relative to *https://example.com/3gpp-ecs-address-provision/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAnConfiguration**](IndividualECSAddressProvisionConfigurationApi.md#DeleteAnConfiguration) | **Delete** /{afId}/configurations/{configurationId} | Deletes an already existing configuration resource
[**FullyUpdateAnConfiguration**](IndividualECSAddressProvisionConfigurationApi.md#FullyUpdateAnConfiguration) | **Put** /{afId}/configurations/{configurationId} | Fully updates/replaces an existing resource
[**ReadAnConfiguration**](IndividualECSAddressProvisionConfigurationApi.md#ReadAnConfiguration) | **Get** /{afId}/configurations/{configurationId} | Read an active resource for the AF and the configuration Id



## DeleteAnConfiguration

> DeleteAnConfiguration(ctx, afId, configurationId).Execute()

Deletes an already existing configuration resource

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
    configurationId := "configurationId_example" // string | Identifier of the configuration resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualECSAddressProvisionConfigurationApi.DeleteAnConfiguration(context.Background(), afId, configurationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualECSAddressProvisionConfigurationApi.DeleteAnConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**afId** | **string** | Identifier of the AF | 
**configurationId** | **string** | Identifier of the configuration resource | 

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


## FullyUpdateAnConfiguration

> EcsAddressProvision FullyUpdateAnConfiguration(ctx, afId, configurationId).EcsAddressProvision(ecsAddressProvision).Execute()

Fully updates/replaces an existing resource

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
    configurationId := "configurationId_example" // string | Identifier of the configuration resource
    ecsAddressProvision := *openapiclient.NewEcsAddressProvision(*openapiclient.NewEcsServerAddr(), "SuppFeat_example") // EcsAddressProvision | Parameters to update/replace the existing resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualECSAddressProvisionConfigurationApi.FullyUpdateAnConfiguration(context.Background(), afId, configurationId).EcsAddressProvision(ecsAddressProvision).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualECSAddressProvisionConfigurationApi.FullyUpdateAnConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FullyUpdateAnConfiguration`: EcsAddressProvision
    fmt.Fprintf(os.Stdout, "Response from `IndividualECSAddressProvisionConfigurationApi.FullyUpdateAnConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**afId** | **string** | Identifier of the AF | 
**configurationId** | **string** | Identifier of the configuration resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiFullyUpdateAnConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ecsAddressProvision** | [**EcsAddressProvision**](EcsAddressProvision.md) | Parameters to update/replace the existing resource | 

### Return type

[**EcsAddressProvision**](EcsAddressProvision.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadAnConfiguration

> EcsAddressProvision ReadAnConfiguration(ctx, afId, configurationId).Execute()

Read an active resource for the AF and the configuration Id

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
    configurationId := "configurationId_example" // string | Identifier of the configuration resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualECSAddressProvisionConfigurationApi.ReadAnConfiguration(context.Background(), afId, configurationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualECSAddressProvisionConfigurationApi.ReadAnConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAnConfiguration`: EcsAddressProvision
    fmt.Fprintf(os.Stdout, "Response from `IndividualECSAddressProvisionConfigurationApi.ReadAnConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**afId** | **string** | Identifier of the AF | 
**configurationId** | **string** | Identifier of the configuration resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAnConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EcsAddressProvision**](EcsAddressProvision.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

