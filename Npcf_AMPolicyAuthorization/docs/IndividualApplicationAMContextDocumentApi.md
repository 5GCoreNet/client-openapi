# \IndividualApplicationAMContextDocumentApi

All URIs are relative to *https://example.com/npcf-am-policyauthorization/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAppAmContext**](IndividualApplicationAMContextDocumentApi.md#DeleteAppAmContext) | **Delete** /app-am-contexts/{appAmContextId} | Deletes an existing Individual Application AM Context
[**GetAppAmContext**](IndividualApplicationAMContextDocumentApi.md#GetAppAmContext) | **Get** /app-am-contexts/{appAmContextId} | Reads an existing Individual Application AM Context
[**ModAppAmContext**](IndividualApplicationAMContextDocumentApi.md#ModAppAmContext) | **Patch** /app-am-contexts/{appAmContextId} | Modifies an existing Individual Application AM Context



## DeleteAppAmContext

> DeleteAppAmContext(ctx, appAmContextId).Execute()

Deletes an existing Individual Application AM Context

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Npcf_AMPolicyAuthorization"
)

func main() {
    appAmContextId := "appAmContextId_example" // string | String identifying the Individual Application AM Context resource.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualApplicationAMContextDocumentApi.DeleteAppAmContext(context.Background(), appAmContextId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualApplicationAMContextDocumentApi.DeleteAppAmContext``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appAmContextId** | **string** | String identifying the Individual Application AM Context resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAppAmContextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppAmContext

> AppAmContextData GetAppAmContext(ctx, appAmContextId).Execute()

Reads an existing Individual Application AM Context

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Npcf_AMPolicyAuthorization"
)

func main() {
    appAmContextId := "appAmContextId_example" // string | String identifying the resource.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualApplicationAMContextDocumentApi.GetAppAmContext(context.Background(), appAmContextId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualApplicationAMContextDocumentApi.GetAppAmContext``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAppAmContext`: AppAmContextData
    fmt.Fprintf(os.Stdout, "Response from `IndividualApplicationAMContextDocumentApi.GetAppAmContext`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appAmContextId** | **string** | String identifying the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppAmContextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AppAmContextData**](AppAmContextData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModAppAmContext

> AppAmContextRespData ModAppAmContext(ctx, appAmContextId).AppAmContextUpdateData(appAmContextUpdateData).Execute()

Modifies an existing Individual Application AM Context

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Npcf_AMPolicyAuthorization"
)

func main() {
    appAmContextId := "appAmContextId_example" // string | String identifying the resource.
    appAmContextUpdateData := *openapiclient.NewAppAmContextUpdateData() // AppAmContextUpdateData | Modification of the resource.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualApplicationAMContextDocumentApi.ModAppAmContext(context.Background(), appAmContextId).AppAmContextUpdateData(appAmContextUpdateData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualApplicationAMContextDocumentApi.ModAppAmContext``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ModAppAmContext`: AppAmContextRespData
    fmt.Fprintf(os.Stdout, "Response from `IndividualApplicationAMContextDocumentApi.ModAppAmContext`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appAmContextId** | **string** | String identifying the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiModAppAmContextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appAmContextUpdateData** | [**AppAmContextUpdateData**](AppAmContextUpdateData.md) | Modification of the resource. | 

### Return type

[**AppAmContextRespData**](AppAmContextRespData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

