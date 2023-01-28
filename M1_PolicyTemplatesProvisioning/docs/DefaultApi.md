# \DefaultApi

All URIs are relative to *https://example.com/3gpp-m1/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePolicyTemplate**](DefaultApi.md#CreatePolicyTemplate) | **Post** /provisioning-sessions/{provisioningSessionId}/policy-templates | Create (and optionally upload) a new Policy Template
[**DestroyPolicyTemplate**](DefaultApi.md#DestroyPolicyTemplate) | **Delete** /provisioning-sessions/{provisioningSessionId}/policy-templates/{policyTemplateId} | 
[**PatchPolicyTemplate**](DefaultApi.md#PatchPolicyTemplate) | **Patch** /provisioning-sessions/{provisioningSessionId}/policy-templates/{policyTemplateId} | Patch the Policy Template for the specified Provisioning Session
[**RetrievePolicyTemplate**](DefaultApi.md#RetrievePolicyTemplate) | **Get** /provisioning-sessions/{provisioningSessionId}/policy-templates/{policyTemplateId} | Retrieve a representation of an existing Policy Template in the specified Provisioning Session
[**UpdatePolicyTemplate**](DefaultApi.md#UpdatePolicyTemplate) | **Put** /provisioning-sessions/{provisioningSessionId}/policy-templates/{policyTemplateId} | Update a Policy Template for the specified Provisioning Session



## CreatePolicyTemplate

> CreatePolicyTemplate(ctx, provisioningSessionId).PolicyTemplate(policyTemplate).Execute()

Create (and optionally upload) a new Policy Template

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/M1_PolicyTemplatesProvisioning"
)

func main() {
    provisioningSessionId := "provisioningSessionId_example" // string | The resource identifier of an existing Provisioning Session.
    policyTemplate := *openapiclient.NewPolicyTemplate("PolicyTemplateId_example", *openapiclient.NewPolicyTemplateState(), "ApiEndPoint_example", *openapiclient.NewPolicyTemplateApiType(), "ExternalReference_example", *openapiclient.NewPolicyTemplateApplicationSessionContext()) // PolicyTemplate | A JSON representation of a Policy Template

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreatePolicyTemplate(context.Background(), provisioningSessionId).PolicyTemplate(policyTemplate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreatePolicyTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provisioningSessionId** | **string** | The resource identifier of an existing Provisioning Session. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePolicyTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **policyTemplate** | [**PolicyTemplate**](PolicyTemplate.md) | A JSON representation of a Policy Template | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyPolicyTemplate

> DestroyPolicyTemplate(ctx, provisioningSessionId, policyTemplateId).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/M1_PolicyTemplatesProvisioning"
)

func main() {
    provisioningSessionId := "provisioningSessionId_example" // string | A unique identifier of the Provisioning Session.
    policyTemplateId := "policyTemplateId_example" // string | A resource identifier of a Policy Template.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DestroyPolicyTemplate(context.Background(), provisioningSessionId, policyTemplateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DestroyPolicyTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provisioningSessionId** | **string** | A unique identifier of the Provisioning Session. | 
**policyTemplateId** | **string** | A resource identifier of a Policy Template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDestroyPolicyTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchPolicyTemplate

> PolicyTemplate PatchPolicyTemplate(ctx, provisioningSessionId, policyTemplateId).PolicyTemplate(policyTemplate).Execute()

Patch the Policy Template for the specified Provisioning Session

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/M1_PolicyTemplatesProvisioning"
)

func main() {
    provisioningSessionId := "provisioningSessionId_example" // string | A unique identifier of the Provisioning Session.
    policyTemplateId := "policyTemplateId_example" // string | A resource identifier of a Policy Template.
    policyTemplate := *openapiclient.NewPolicyTemplate("PolicyTemplateId_example", *openapiclient.NewPolicyTemplateState(), "ApiEndPoint_example", *openapiclient.NewPolicyTemplateApiType(), "ExternalReference_example", *openapiclient.NewPolicyTemplateApplicationSessionContext()) // PolicyTemplate | A JSON representation of a Policy Template

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.PatchPolicyTemplate(context.Background(), provisioningSessionId, policyTemplateId).PolicyTemplate(policyTemplate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.PatchPolicyTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchPolicyTemplate`: PolicyTemplate
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.PatchPolicyTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provisioningSessionId** | **string** | A unique identifier of the Provisioning Session. | 
**policyTemplateId** | **string** | A resource identifier of a Policy Template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchPolicyTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **policyTemplate** | [**PolicyTemplate**](PolicyTemplate.md) | A JSON representation of a Policy Template | 

### Return type

[**PolicyTemplate**](PolicyTemplate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/merge-patch+json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrievePolicyTemplate

> PolicyTemplate RetrievePolicyTemplate(ctx, provisioningSessionId, policyTemplateId).Execute()

Retrieve a representation of an existing Policy Template in the specified Provisioning Session

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/M1_PolicyTemplatesProvisioning"
)

func main() {
    provisioningSessionId := "provisioningSessionId_example" // string | A unique identifier of the Provisioning Session.
    policyTemplateId := "policyTemplateId_example" // string | A resource identifier of a Policy Template.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.RetrievePolicyTemplate(context.Background(), provisioningSessionId, policyTemplateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RetrievePolicyTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrievePolicyTemplate`: PolicyTemplate
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.RetrievePolicyTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provisioningSessionId** | **string** | A unique identifier of the Provisioning Session. | 
**policyTemplateId** | **string** | A resource identifier of a Policy Template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrievePolicyTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PolicyTemplate**](PolicyTemplate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePolicyTemplate

> UpdatePolicyTemplate(ctx, provisioningSessionId, policyTemplateId).PolicyTemplate(policyTemplate).Execute()

Update a Policy Template for the specified Provisioning Session

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/M1_PolicyTemplatesProvisioning"
)

func main() {
    provisioningSessionId := "provisioningSessionId_example" // string | A unique identifier of the Provisioning Session.
    policyTemplateId := "policyTemplateId_example" // string | A resource identifier of a Policy Template.
    policyTemplate := *openapiclient.NewPolicyTemplate("PolicyTemplateId_example", *openapiclient.NewPolicyTemplateState(), "ApiEndPoint_example", *openapiclient.NewPolicyTemplateApiType(), "ExternalReference_example", *openapiclient.NewPolicyTemplateApplicationSessionContext()) // PolicyTemplate | A JSON representation of a Policy Template

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdatePolicyTemplate(context.Background(), provisioningSessionId, policyTemplateId).PolicyTemplate(policyTemplate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdatePolicyTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provisioningSessionId** | **string** | A unique identifier of the Provisioning Session. | 
**policyTemplateId** | **string** | A resource identifier of a Policy Template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePolicyTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **policyTemplate** | [**PolicyTemplate**](PolicyTemplate.md) | A JSON representation of a Policy Template | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

