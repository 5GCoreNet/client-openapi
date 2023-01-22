# \ECSAddressProvisionConfigurationsCollectionApi

All URIs are relative to *https://example.com/3gpp-ecs-address-provision/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNewConfiguration**](ECSAddressProvisionConfigurationsCollectionApi.md#CreateNewConfiguration) | **Post** /{afId}/configurations | Creates a new configuration resource
[**ReadAllConfigurations**](ECSAddressProvisionConfigurationsCollectionApi.md#ReadAllConfigurations) | **Get** /{afId}/configurations | Read all active configurations for a given AF



## CreateNewConfiguration

> EcsAddressProvision CreateNewConfiguration(ctx, afId).EcsAddressProvision(ecsAddressProvision).Execute()

Creates a new configuration resource

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
    ecsAddressProvision := *openapiclient.NewEcsAddressProvision(*openapiclient.NewEcsServerAddr(), "SuppFeat_example") // EcsAddressProvision | new resource creation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ECSAddressProvisionConfigurationsCollectionApi.CreateNewConfiguration(context.Background(), afId).EcsAddressProvision(ecsAddressProvision).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ECSAddressProvisionConfigurationsCollectionApi.CreateNewConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNewConfiguration`: EcsAddressProvision
    fmt.Fprintf(os.Stdout, "Response from `ECSAddressProvisionConfigurationsCollectionApi.CreateNewConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**afId** | **string** | Identifier of the AF | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNewConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ecsAddressProvision** | [**EcsAddressProvision**](EcsAddressProvision.md) | new resource creation | 

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


## ReadAllConfigurations

> []EcsAddressProvision ReadAllConfigurations(ctx, afId).Execute()

Read all active configurations for a given AF

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ECSAddressProvisionConfigurationsCollectionApi.ReadAllConfigurations(context.Background(), afId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ECSAddressProvisionConfigurationsCollectionApi.ReadAllConfigurations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllConfigurations`: []EcsAddressProvision
    fmt.Fprintf(os.Stdout, "Response from `ECSAddressProvisionConfigurationsCollectionApi.ReadAllConfigurations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**afId** | **string** | Identifier of the AF | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAllConfigurationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]EcsAddressProvision**](EcsAddressProvision.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

