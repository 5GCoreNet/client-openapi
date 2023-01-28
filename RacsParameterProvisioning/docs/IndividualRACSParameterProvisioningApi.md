# \IndividualRACSParameterProvisioningApi

All URIs are relative to *https://example.com/3gpp-racs-pp/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIndRACSParameterProvisioning**](IndividualRACSParameterProvisioningApi.md#DeleteIndRACSParameterProvisioning) | **Delete** /{scsAsId}/provisionings/{provisioningId} | Delete a RACS parameter provisioning.
[**FetchIndRACSParameterProvisioning**](IndividualRACSParameterProvisioningApi.md#FetchIndRACSParameterProvisioning) | **Get** /{scsAsId}/provisionings/{provisioningId} | Read an existing RACS parameter provisioning.
[**ModifyIndRACSParameterProvisioning**](IndividualRACSParameterProvisioningApi.md#ModifyIndRACSParameterProvisioning) | **Patch** /{scsAsId}/provisionings/{provisioningId} | Modify some properties in an existing RACS parameter provisioning.
[**UpdateIndRACSParameterProvisioning**](IndividualRACSParameterProvisioningApi.md#UpdateIndRACSParameterProvisioning) | **Put** /{scsAsId}/provisionings/{provisioningId} | Modify all properties in an existing RACS parameter provisioning.



## DeleteIndRACSParameterProvisioning

> DeleteIndRACSParameterProvisioning(ctx, scsAsId, provisioningId).Execute()

Delete a RACS parameter provisioning.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/RacsParameterProvisioning"
)

func main() {
    scsAsId := "scsAsId_example" // string | Identifier of the SCS/AS as defined in clause 5.2.4 of 3GPP TS 29.122.
    provisioningId := "provisioningId_example" // string | Provisioning ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualRACSParameterProvisioningApi.DeleteIndRACSParameterProvisioning(context.Background(), scsAsId, provisioningId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualRACSParameterProvisioningApi.DeleteIndRACSParameterProvisioning``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of the SCS/AS as defined in clause 5.2.4 of 3GPP TS 29.122. | 
**provisioningId** | **string** | Provisioning ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIndRACSParameterProvisioningRequest struct via the builder pattern


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


## FetchIndRACSParameterProvisioning

> RacsProvisioningData FetchIndRACSParameterProvisioning(ctx, scsAsId, provisioningId).Execute()

Read an existing RACS parameter provisioning.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/RacsParameterProvisioning"
)

func main() {
    scsAsId := "scsAsId_example" // string | Identifier of the SCS/AS as defined in clause 5.2.4 of 3GPP TS 29.122.
    provisioningId := "provisioningId_example" // string | Provisioning ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualRACSParameterProvisioningApi.FetchIndRACSParameterProvisioning(context.Background(), scsAsId, provisioningId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualRACSParameterProvisioningApi.FetchIndRACSParameterProvisioning``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchIndRACSParameterProvisioning`: RacsProvisioningData
    fmt.Fprintf(os.Stdout, "Response from `IndividualRACSParameterProvisioningApi.FetchIndRACSParameterProvisioning`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of the SCS/AS as defined in clause 5.2.4 of 3GPP TS 29.122. | 
**provisioningId** | **string** | Provisioning ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchIndRACSParameterProvisioningRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RacsProvisioningData**](RacsProvisioningData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyIndRACSParameterProvisioning

> RacsProvisioningData ModifyIndRACSParameterProvisioning(ctx, scsAsId, provisioningId).RacsProvisioningDataPatch(racsProvisioningDataPatch).Execute()

Modify some properties in an existing RACS parameter provisioning.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/RacsParameterProvisioning"
)

func main() {
    scsAsId := "scsAsId_example" // string | Identifier of the SCS/AS as defined in clause 5.2.4 of 3GPP TS 29.122.
    provisioningId := "provisioningId_example" // string | Provisioning ID
    racsProvisioningDataPatch := *openapiclient.NewRacsProvisioningDataPatch() // RacsProvisioningDataPatch | update an existing parameter provisioning.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualRACSParameterProvisioningApi.ModifyIndRACSParameterProvisioning(context.Background(), scsAsId, provisioningId).RacsProvisioningDataPatch(racsProvisioningDataPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualRACSParameterProvisioningApi.ModifyIndRACSParameterProvisioning``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ModifyIndRACSParameterProvisioning`: RacsProvisioningData
    fmt.Fprintf(os.Stdout, "Response from `IndividualRACSParameterProvisioningApi.ModifyIndRACSParameterProvisioning`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of the SCS/AS as defined in clause 5.2.4 of 3GPP TS 29.122. | 
**provisioningId** | **string** | Provisioning ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyIndRACSParameterProvisioningRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **racsProvisioningDataPatch** | [**RacsProvisioningDataPatch**](RacsProvisioningDataPatch.md) | update an existing parameter provisioning. | 

### Return type

[**RacsProvisioningData**](RacsProvisioningData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIndRACSParameterProvisioning

> RacsProvisioningData UpdateIndRACSParameterProvisioning(ctx, scsAsId, provisioningId).RacsProvisioningData(racsProvisioningData).Execute()

Modify all properties in an existing RACS parameter provisioning.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/RacsParameterProvisioning"
)

func main() {
    scsAsId := "scsAsId_example" // string | Identifier of the SCS/AS as defined in clause 5.2.4 of 3GPP TS 29.122.
    provisioningId := "provisioningId_example" // string | Provisioning ID
    racsProvisioningData := *openapiclient.NewRacsProvisioningData(map[string]RacsConfiguration{"key": *openapiclient.NewRacsConfiguration("RacsId_example", []string{"ImeiTacs_example"})}) // RacsProvisioningData | update an existing parameter provisioning.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualRACSParameterProvisioningApi.UpdateIndRACSParameterProvisioning(context.Background(), scsAsId, provisioningId).RacsProvisioningData(racsProvisioningData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualRACSParameterProvisioningApi.UpdateIndRACSParameterProvisioning``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIndRACSParameterProvisioning`: RacsProvisioningData
    fmt.Fprintf(os.Stdout, "Response from `IndividualRACSParameterProvisioningApi.UpdateIndRACSParameterProvisioning`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of the SCS/AS as defined in clause 5.2.4 of 3GPP TS 29.122. | 
**provisioningId** | **string** | Provisioning ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIndRACSParameterProvisioningRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **racsProvisioningData** | [**RacsProvisioningData**](RacsProvisioningData.md) | update an existing parameter provisioning. | 

### Return type

[**RacsProvisioningData**](RacsProvisioningData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

