# \IndividualUERadioCapabilityProvisioningDocumentApi

All URIs are relative to *https://example.com/nucmf-provisioning/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetProvisioning**](IndividualUERadioCapabilityProvisioningDocumentApi.md#GetProvisioning) | **Get** /provisionings/{provisioningId} | Get an Individual UE radio capability provisioning
[**RemoveProvisioning**](IndividualUERadioCapabilityProvisioningDocumentApi.md#RemoveProvisioning) | **Delete** /provisionings/{provisioningId} | Remove an Individual UE radio capability provisioning
[**ReplaceProvisioning**](IndividualUERadioCapabilityProvisioningDocumentApi.md#ReplaceProvisioning) | **Put** /provisionings/{provisioningId} | Replace (PUT) an Individual UE radio capability provisioning
[**UpdateProvisioning**](IndividualUERadioCapabilityProvisioningDocumentApi.md#UpdateProvisioning) | **Patch** /provisionings/{provisioningId} | Update (PATCH) an Individual UE radio capability provisioning



## GetProvisioning

> RacsData GetProvisioning(ctx, provisioningId).Execute()

Get an Individual UE radio capability provisioning

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nucmf_Provisioning"
)

func main() {
    provisioningId := "provisioningId_example" // string | Provisioning ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualUERadioCapabilityProvisioningDocumentApi.GetProvisioning(context.Background(), provisioningId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualUERadioCapabilityProvisioningDocumentApi.GetProvisioning``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProvisioning`: RacsData
    fmt.Fprintf(os.Stdout, "Response from `IndividualUERadioCapabilityProvisioningDocumentApi.GetProvisioning`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provisioningId** | **string** | Provisioning ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProvisioningRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RacsData**](RacsData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveProvisioning

> RemoveProvisioning(ctx, provisioningId).Execute()

Remove an Individual UE radio capability provisioning

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nucmf_Provisioning"
)

func main() {
    provisioningId := "provisioningId_example" // string | Provisioning ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualUERadioCapabilityProvisioningDocumentApi.RemoveProvisioning(context.Background(), provisioningId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualUERadioCapabilityProvisioningDocumentApi.RemoveProvisioning``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provisioningId** | **string** | Provisioning ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveProvisioningRequest struct via the builder pattern


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


## ReplaceProvisioning

> RacsData ReplaceProvisioning(ctx, provisioningId).RacsData(racsData).Execute()

Replace (PUT) an Individual UE radio capability provisioning

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nucmf_Provisioning"
)

func main() {
    provisioningId := "provisioningId_example" // string | Provisioning ID
    racsData := *openapiclient.NewRacsData(map[string]RacsConfiguration{"key": *openapiclient.NewRacsConfiguration("RacsId_example", []string{"ImeiTacs_example"})}) // RacsData | Update an existing parameter provisioning.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualUERadioCapabilityProvisioningDocumentApi.ReplaceProvisioning(context.Background(), provisioningId).RacsData(racsData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualUERadioCapabilityProvisioningDocumentApi.ReplaceProvisioning``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceProvisioning`: RacsData
    fmt.Fprintf(os.Stdout, "Response from `IndividualUERadioCapabilityProvisioningDocumentApi.ReplaceProvisioning`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provisioningId** | **string** | Provisioning ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceProvisioningRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **racsData** | [**RacsData**](RacsData.md) | Update an existing parameter provisioning. | 

### Return type

[**RacsData**](RacsData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProvisioning

> RacsData UpdateProvisioning(ctx, provisioningId).RacsDataPatch(racsDataPatch).Execute()

Update (PATCH) an Individual UE radio capability provisioning

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nucmf_Provisioning"
)

func main() {
    provisioningId := "provisioningId_example" // string | Provisioning ID
    racsDataPatch := *openapiclient.NewRacsDataPatch() // RacsDataPatch | Update an existing parameter provisioning.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualUERadioCapabilityProvisioningDocumentApi.UpdateProvisioning(context.Background(), provisioningId).RacsDataPatch(racsDataPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualUERadioCapabilityProvisioningDocumentApi.UpdateProvisioning``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateProvisioning`: RacsData
    fmt.Fprintf(os.Stdout, "Response from `IndividualUERadioCapabilityProvisioningDocumentApi.UpdateProvisioning`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provisioningId** | **string** | Provisioning ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProvisioningRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **racsDataPatch** | [**RacsDataPatch**](RacsDataPatch.md) | Update an existing parameter provisioning. | 

### Return type

[**RacsData**](RacsData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

