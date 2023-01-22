# \IndividualApplicationAMContextApi

All URIs are relative to *https://example.com/3gpp-am-policyauthorization/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAppAmContext**](IndividualApplicationAMContextApi.md#DeleteAppAmContext) | **Delete** /{afId}/app-am-contexts/{appAmContextId} | Deletes an existing Individual Application AM Context
[**GetAppAmContext**](IndividualApplicationAMContextApi.md#GetAppAmContext) | **Get** /{afId}/app-am-contexts/{appAmContextId} | read an existing Individual application AM context
[**ModAppAmContext**](IndividualApplicationAMContextApi.md#ModAppAmContext) | **Patch** /{afId}/app-am-contexts/{appAmContextId} | partial modifies an existing Individual application AM context



## DeleteAppAmContext

> DeleteAppAmContext(ctx, afId, appAmContextId).Execute()

Deletes an existing Individual Application AM Context

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
    appAmContextId := "appAmContextId_example" // string | string identifying the Individual aaplication AM context resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualApplicationAMContextApi.DeleteAppAmContext(context.Background(), afId, appAmContextId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualApplicationAMContextApi.DeleteAppAmContext``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**afId** | **string** | Identifier of the AF | 
**appAmContextId** | **string** | string identifying the Individual aaplication AM context resource | 

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
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppAmContext

> AppAmContextExpData GetAppAmContext(ctx, afId, appAmContextId).Execute()

read an existing Individual application AM context

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
    appAmContextId := "appAmContextId_example" // string | Identifier of the Individual application AM context

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualApplicationAMContextApi.GetAppAmContext(context.Background(), afId, appAmContextId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualApplicationAMContextApi.GetAppAmContext``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAppAmContext`: AppAmContextExpData
    fmt.Fprintf(os.Stdout, "Response from `IndividualApplicationAMContextApi.GetAppAmContext`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**afId** | **string** | Identifier of the AF | 
**appAmContextId** | **string** | Identifier of the Individual application AM context | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppAmContextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AppAmContextExpData**](AppAmContextExpData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModAppAmContext

> AppAmContextExpRespData ModAppAmContext(ctx, afId, appAmContextId).AppAmContextExpUpdateData(appAmContextExpUpdateData).Execute()

partial modifies an existing Individual application AM context

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
    appAmContextId := "appAmContextId_example" // string | Identifier of the application AM context resource
    appAmContextExpUpdateData := *openapiclient.NewAppAmContextExpUpdateData() // AppAmContextExpUpdateData | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualApplicationAMContextApi.ModAppAmContext(context.Background(), afId, appAmContextId).AppAmContextExpUpdateData(appAmContextExpUpdateData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualApplicationAMContextApi.ModAppAmContext``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ModAppAmContext`: AppAmContextExpRespData
    fmt.Fprintf(os.Stdout, "Response from `IndividualApplicationAMContextApi.ModAppAmContext`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**afId** | **string** | Identifier of the AF | 
**appAmContextId** | **string** | Identifier of the application AM context resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiModAppAmContextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **appAmContextExpUpdateData** | [**AppAmContextExpUpdateData**](AppAmContextExpUpdateData.md) |  | 

### Return type

[**AppAmContextExpRespData**](AppAmContextExpRespData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

