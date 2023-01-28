# \LPIParametersProvisioningsApi

All URIs are relative to *https://example.com/3gpp-lpi-pp/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNewResource**](LPIParametersProvisioningsApi.md#CreateNewResource) | **Post** /{afId}/provisionedLpis | Creates a new LPI Parameters Provisioning resource
[**ReadAllResources**](LPIParametersProvisioningsApi.md#ReadAllResources) | **Get** /{afId}/provisionedLpis | read all of the active LPI Parameters Provisioning resources for the AF



## CreateNewResource

> LpiParametersProvision CreateNewResource(ctx, afId).LpiParametersProvision(lpiParametersProvision).Execute()

Creates a new LPI Parameters Provisioning resource

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/LpiParameterProvision"
)

func main() {
    afId := "afId_example" // string | Identifier of the AF
    lpiParametersProvision := *openapiclient.NewLpiParametersProvision(*openapiclient.NewLpi(*openapiclient.NewLocationPrivacyInd()), "SuppFeat_example") // LpiParametersProvision | new resource creation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LPIParametersProvisioningsApi.CreateNewResource(context.Background(), afId).LpiParametersProvision(lpiParametersProvision).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LPIParametersProvisioningsApi.CreateNewResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNewResource`: LpiParametersProvision
    fmt.Fprintf(os.Stdout, "Response from `LPIParametersProvisioningsApi.CreateNewResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**afId** | **string** | Identifier of the AF | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNewResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lpiParametersProvision** | [**LpiParametersProvision**](LpiParametersProvision.md) | new resource creation | 

### Return type

[**LpiParametersProvision**](LpiParametersProvision.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadAllResources

> []LpiParametersProvision ReadAllResources(ctx, afId).Execute()

read all of the active LPI Parameters Provisioning resources for the AF

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/LpiParameterProvision"
)

func main() {
    afId := "afId_example" // string | Identifier of the AF

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LPIParametersProvisioningsApi.ReadAllResources(context.Background(), afId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LPIParametersProvisioningsApi.ReadAllResources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllResources`: []LpiParametersProvision
    fmt.Fprintf(os.Stdout, "Response from `LPIParametersProvisioningsApi.ReadAllResources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**afId** | **string** | Identifier of the AF | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAllResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]LpiParametersProvision**](LpiParametersProvision.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

