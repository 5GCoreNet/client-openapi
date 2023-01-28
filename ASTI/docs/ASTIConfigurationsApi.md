# \ASTIConfigurationsApi

All URIs are relative to *https://example.com/3gpp-asti/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNewConfiguration**](ASTIConfigurationsApi.md#CreateNewConfiguration) | **Post** /{afId}/configurations | Creates a new configuration resource
[**ReadAllConfigurations**](ASTIConfigurationsApi.md#ReadAllConfigurations) | **Get** /{afId}/configurations | read all of the active configurations of 5G access stratum time distribution for the AF
[**RetrieveStatusofConfiguration**](ASTIConfigurationsApi.md#RetrieveStatusofConfiguration) | **Post** /{afId}/configurations/retrieve | Request the status of the 5G access stratum time distribution configuration for a list of UEs.



## CreateNewConfiguration

> AccessTimeDistributionData CreateNewConfiguration(ctx, afId).AccessTimeDistributionData(accessTimeDistributionData).Execute()

Creates a new configuration resource

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/ASTI"
)

func main() {
    afId := "afId_example" // string | Identifier of the AF
    accessTimeDistributionData := openapiclient.AccessTimeDistributionData{Interface{}: new(interface{})} // AccessTimeDistributionData | new configuration creation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ASTIConfigurationsApi.CreateNewConfiguration(context.Background(), afId).AccessTimeDistributionData(accessTimeDistributionData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ASTIConfigurationsApi.CreateNewConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNewConfiguration`: AccessTimeDistributionData
    fmt.Fprintf(os.Stdout, "Response from `ASTIConfigurationsApi.CreateNewConfiguration`: %v\n", resp)
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

 **accessTimeDistributionData** | [**AccessTimeDistributionData**](AccessTimeDistributionData.md) | new configuration creation | 

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


## ReadAllConfigurations

> []AccessTimeDistributionData ReadAllConfigurations(ctx, afId).Execute()

read all of the active configurations of 5G access stratum time distribution for the AF

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/ASTI"
)

func main() {
    afId := "afId_example" // string | Identifier of the AF

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ASTIConfigurationsApi.ReadAllConfigurations(context.Background(), afId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ASTIConfigurationsApi.ReadAllConfigurations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllConfigurations`: []AccessTimeDistributionData
    fmt.Fprintf(os.Stdout, "Response from `ASTIConfigurationsApi.ReadAllConfigurations`: %v\n", resp)
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

[**[]AccessTimeDistributionData**](AccessTimeDistributionData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveStatusofConfiguration

> StatusResponseData RetrieveStatusofConfiguration(ctx, afId).StatusRequestData(statusRequestData).Execute()

Request the status of the 5G access stratum time distribution configuration for a list of UEs.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/ASTI"
)

func main() {
    afId := "afId_example" // string | Identifier of the AF
    statusRequestData := *openapiclient.NewStatusRequestData([]string{"Gpsis_example"}) // StatusRequestData | Contains the list of GPSIs.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ASTIConfigurationsApi.RetrieveStatusofConfiguration(context.Background(), afId).StatusRequestData(statusRequestData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ASTIConfigurationsApi.RetrieveStatusofConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveStatusofConfiguration`: StatusResponseData
    fmt.Fprintf(os.Stdout, "Response from `ASTIConfigurationsApi.RetrieveStatusofConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**afId** | **string** | Identifier of the AF | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveStatusofConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **statusRequestData** | [**StatusRequestData**](StatusRequestData.md) | Contains the list of GPSIs. | 

### Return type

[**StatusResponseData**](StatusResponseData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

