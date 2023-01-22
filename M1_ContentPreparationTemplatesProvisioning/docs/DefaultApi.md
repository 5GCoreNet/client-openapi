# \DefaultApi

All URIs are relative to *https://example.com/3gpp-m1/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateContentPreparationTemplate**](DefaultApi.md#CreateContentPreparationTemplate) | **Post** /provisioning-sessions/{provisioningSessionId}/content-preparation-templates | Create (and optionally upload) a new Content Preparation Template for the specified Provisioning Session
[**DestroyContentPreparationTemplate**](DefaultApi.md#DestroyContentPreparationTemplate) | **Delete** /provisioning-sessions/{provisioningSessionId}/content-preparation-templates/{contentPreparationTemplateId} | Destroy the specified Content Preparation Template of the specified Provisioning Session
[**PatchContentPreparationTemplate**](DefaultApi.md#PatchContentPreparationTemplate) | **Patch** /provisioning-sessions/{provisioningSessionId}/content-preparation-templates/{contentPreparationTemplateId} | Patch the specified Content Preparation Template for the specified Provisioning Session
[**RetrieveContentPreparationTemplate**](DefaultApi.md#RetrieveContentPreparationTemplate) | **Get** /provisioning-sessions/{provisioningSessionId}/content-preparation-templates/{contentPreparationTemplateId} | Retrieve the specified Content Preparation Template of the specified Provisioning Session
[**UpdateContentPreparationTemplate**](DefaultApi.md#UpdateContentPreparationTemplate) | **Put** /provisioning-sessions/{provisioningSessionId}/content-preparation-templates/{contentPreparationTemplateId} | Update the specified Content Preparation Template for the specified Provisioning Session



## CreateContentPreparationTemplate

> CreateContentPreparationTemplate(ctx, provisioningSessionId).Body(body).Execute()

Create (and optionally upload) a new Content Preparation Template for the specified Provisioning Session

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
    provisioningSessionId := "provisioningSessionId_example" // string | The resource identifier of an existing Provisioning Session.
    body := "body_example" // string | A Content Preparation Template of any type

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateContentPreparationTemplate(context.Background(), provisioningSessionId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateContentPreparationTemplate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCreateContentPreparationTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** | A Content Preparation Template of any type | 

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


## DestroyContentPreparationTemplate

> DestroyContentPreparationTemplate(ctx, provisioningSessionId, contentPreparationTemplateId).Execute()

Destroy the specified Content Preparation Template of the specified Provisioning Session

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
    provisioningSessionId := "provisioningSessionId_example" // string | The resource identifier of an existing Provisioning Session.
    contentPreparationTemplateId := "contentPreparationTemplateId_example" // string | The resource identifier of an existing Content Preparation Template.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DestroyContentPreparationTemplate(context.Background(), provisioningSessionId, contentPreparationTemplateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DestroyContentPreparationTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provisioningSessionId** | **string** | The resource identifier of an existing Provisioning Session. | 
**contentPreparationTemplateId** | **string** | The resource identifier of an existing Content Preparation Template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDestroyContentPreparationTemplateRequest struct via the builder pattern


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


## PatchContentPreparationTemplate

> string PatchContentPreparationTemplate(ctx, provisioningSessionId, contentPreparationTemplateId).Body(body).Execute()

Patch the specified Content Preparation Template for the specified Provisioning Session

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
    provisioningSessionId := "provisioningSessionId_example" // string | The resource identifier of an existing Provisioning Session.
    contentPreparationTemplateId := "contentPreparationTemplateId_example" // string | The resource identifier of an existing Content Preparation Template.
    body := "body_example" // string | A Content Preparation Template patch of any type

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.PatchContentPreparationTemplate(context.Background(), provisioningSessionId, contentPreparationTemplateId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.PatchContentPreparationTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchContentPreparationTemplate`: string
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.PatchContentPreparationTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provisioningSessionId** | **string** | The resource identifier of an existing Provisioning Session. | 
**contentPreparationTemplateId** | **string** | The resource identifier of an existing Content Preparation Template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchContentPreparationTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **string** | A Content Preparation Template patch of any type | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/merge-patch+json, application/json-patch+json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveContentPreparationTemplate

> string RetrieveContentPreparationTemplate(ctx, provisioningSessionId, contentPreparationTemplateId).Execute()

Retrieve the specified Content Preparation Template of the specified Provisioning Session

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
    provisioningSessionId := "provisioningSessionId_example" // string | The resource identifier of an existing Provisioning Session.
    contentPreparationTemplateId := "contentPreparationTemplateId_example" // string | The resource identifier of an existing Content Preparation Template.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.RetrieveContentPreparationTemplate(context.Background(), provisioningSessionId, contentPreparationTemplateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RetrieveContentPreparationTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveContentPreparationTemplate`: string
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.RetrieveContentPreparationTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provisioningSessionId** | **string** | The resource identifier of an existing Provisioning Session. | 
**contentPreparationTemplateId** | **string** | The resource identifier of an existing Content Preparation Template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveContentPreparationTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateContentPreparationTemplate

> UpdateContentPreparationTemplate(ctx, provisioningSessionId, contentPreparationTemplateId).Body(body).Execute()

Update the specified Content Preparation Template for the specified Provisioning Session

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
    provisioningSessionId := "provisioningSessionId_example" // string | The resource identifier of an existing Provisioning Session.
    contentPreparationTemplateId := "contentPreparationTemplateId_example" // string | The resource identifier of an existing Content Preparation Template.
    body := "body_example" // string | A Content Preparation Template of any type

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateContentPreparationTemplate(context.Background(), provisioningSessionId, contentPreparationTemplateId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateContentPreparationTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provisioningSessionId** | **string** | The resource identifier of an existing Provisioning Session. | 
**contentPreparationTemplateId** | **string** | The resource identifier of an existing Content Preparation Template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateContentPreparationTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **string** | A Content Preparation Template of any type | 

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

