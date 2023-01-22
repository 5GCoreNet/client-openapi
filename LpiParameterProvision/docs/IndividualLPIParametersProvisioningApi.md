# \IndividualLPIParametersProvisioningApi

All URIs are relative to *https://example.com/3gpp-lpi-pp/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAnResource**](IndividualLPIParametersProvisioningApi.md#DeleteAnResource) | **Delete** /{afId}/provisionedLpis/{provisionedLpiId} | Deletes an already existing LPI Parameters Provisioning resource
[**FullyUpdateAnResource**](IndividualLPIParametersProvisioningApi.md#FullyUpdateAnResource) | **Put** /{afId}/provisionedLpis/{provisionedLpiId} | Fully updates/replaces an existing LPI Parameters Provisioning resource
[**PartialUpdateAnResource**](IndividualLPIParametersProvisioningApi.md#PartialUpdateAnResource) | **Patch** /{afId}/provisionedLpis/{provisionedLpiId} | Partially modifies an existing LPI Parameters Provisioning resource.
[**ReadAnResource**](IndividualLPIParametersProvisioningApi.md#ReadAnResource) | **Get** /{afId}/provisionedLpis/{provisionedLpiId} | read an active LPI Parameters Provisioning resource for the AF and the provisioned LPI Id



## DeleteAnResource

> DeleteAnResource(ctx, afId, provisionedLpiId).Execute()

Deletes an already existing LPI Parameters Provisioning resource

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
    provisionedLpiId := "provisionedLpiId_example" // string | Identifier of the provisioned LPI parameter resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualLPIParametersProvisioningApi.DeleteAnResource(context.Background(), afId, provisionedLpiId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualLPIParametersProvisioningApi.DeleteAnResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**afId** | **string** | Identifier of the AF | 
**provisionedLpiId** | **string** | Identifier of the provisioned LPI parameter resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAnResourceRequest struct via the builder pattern


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


## FullyUpdateAnResource

> LpiParametersProvision FullyUpdateAnResource(ctx, afId, provisionedLpiId).LpiParametersProvision(lpiParametersProvision).Execute()

Fully updates/replaces an existing LPI Parameters Provisioning resource

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
    provisionedLpiId := "provisionedLpiId_example" // string | Identifier of the provisioned LPI parameter resource
    lpiParametersProvision := *openapiclient.NewLpiParametersProvision(*openapiclient.NewLpi(*openapiclient.NewLocationPrivacyInd()), "SuppFeat_example") // LpiParametersProvision | Parameters to update/replace the existing resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualLPIParametersProvisioningApi.FullyUpdateAnResource(context.Background(), afId, provisionedLpiId).LpiParametersProvision(lpiParametersProvision).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualLPIParametersProvisioningApi.FullyUpdateAnResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FullyUpdateAnResource`: LpiParametersProvision
    fmt.Fprintf(os.Stdout, "Response from `IndividualLPIParametersProvisioningApi.FullyUpdateAnResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**afId** | **string** | Identifier of the AF | 
**provisionedLpiId** | **string** | Identifier of the provisioned LPI parameter resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiFullyUpdateAnResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **lpiParametersProvision** | [**LpiParametersProvision**](LpiParametersProvision.md) | Parameters to update/replace the existing resource | 

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


## PartialUpdateAnResource

> LpiParametersProvision PartialUpdateAnResource(ctx, afId, provisionedLpiId).LpiParametersProvisionPatch(lpiParametersProvisionPatch).Execute()

Partially modifies an existing LPI Parameters Provisioning resource.

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
    provisionedLpiId := "provisionedLpiId_example" // string | Identifier of the provisioned LPI parameter resource
    lpiParametersProvisionPatch := *openapiclient.NewLpiParametersProvisionPatch() // LpiParametersProvisionPatch | Parameters to modify the existing resource.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualLPIParametersProvisioningApi.PartialUpdateAnResource(context.Background(), afId, provisionedLpiId).LpiParametersProvisionPatch(lpiParametersProvisionPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualLPIParametersProvisioningApi.PartialUpdateAnResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PartialUpdateAnResource`: LpiParametersProvision
    fmt.Fprintf(os.Stdout, "Response from `IndividualLPIParametersProvisioningApi.PartialUpdateAnResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**afId** | **string** | Identifier of the AF | 
**provisionedLpiId** | **string** | Identifier of the provisioned LPI parameter resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartialUpdateAnResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **lpiParametersProvisionPatch** | [**LpiParametersProvisionPatch**](LpiParametersProvisionPatch.md) | Parameters to modify the existing resource. | 

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


## ReadAnResource

> LpiParametersProvision ReadAnResource(ctx, afId, provisionedLpiId).Execute()

read an active LPI Parameters Provisioning resource for the AF and the provisioned LPI Id

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
    provisionedLpiId := "provisionedLpiId_example" // string | Identifier of the provisioned LPI parameter resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualLPIParametersProvisioningApi.ReadAnResource(context.Background(), afId, provisionedLpiId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualLPIParametersProvisioningApi.ReadAnResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAnResource`: LpiParametersProvision
    fmt.Fprintf(os.Stdout, "Response from `IndividualLPIParametersProvisioningApi.ReadAnResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**afId** | **string** | Identifier of the AF | 
**provisionedLpiId** | **string** | Identifier of the provisioned LPI parameter resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAnResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**LpiParametersProvision**](LpiParametersProvision.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

