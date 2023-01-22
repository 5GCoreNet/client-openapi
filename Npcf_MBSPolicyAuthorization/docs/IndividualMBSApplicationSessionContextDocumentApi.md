# \IndividualMBSApplicationSessionContextDocumentApi

All URIs are relative to *https://example.com/npcf-mbspolicyauth/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMBSAppSessionCtxt**](IndividualMBSApplicationSessionContextDocumentApi.md#DeleteMBSAppSessionCtxt) | **Delete** /contexts/{contextId} | Request the deletion of an existing Individual MBS Application Session Context resource.
[**GetMBSAppSessionCtxt**](IndividualMBSApplicationSessionContextDocumentApi.md#GetMBSAppSessionCtxt) | **Get** /contexts/{contextId} | Read an existing Individual MBS Application Session Context resource.
[**ModifyMBSAppSessionCtxt**](IndividualMBSApplicationSessionContextDocumentApi.md#ModifyMBSAppSessionCtxt) | **Patch** /contexts/{contextId} | Request the modification of an existing Individual MBS Application Session Context resource.



## DeleteMBSAppSessionCtxt

> DeleteMBSAppSessionCtxt(ctx, contextId).Execute()

Request the deletion of an existing Individual MBS Application Session Context resource.

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
    contextId := "contextId_example" // string | Contains the identifier of the Individual MBS Application Session Context resource. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualMBSApplicationSessionContextDocumentApi.DeleteMBSAppSessionCtxt(context.Background(), contextId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualMBSApplicationSessionContextDocumentApi.DeleteMBSAppSessionCtxt``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contextId** | **string** | Contains the identifier of the Individual MBS Application Session Context resource.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMBSAppSessionCtxtRequest struct via the builder pattern


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


## GetMBSAppSessionCtxt

> MbsAppSessionCtxt GetMBSAppSessionCtxt(ctx, contextId).Execute()

Read an existing Individual MBS Application Session Context resource.

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
    contextId := "contextId_example" // string | Contains the identifier of the Individual MBS Application Session Context resource. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualMBSApplicationSessionContextDocumentApi.GetMBSAppSessionCtxt(context.Background(), contextId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualMBSApplicationSessionContextDocumentApi.GetMBSAppSessionCtxt``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMBSAppSessionCtxt`: MbsAppSessionCtxt
    fmt.Fprintf(os.Stdout, "Response from `IndividualMBSApplicationSessionContextDocumentApi.GetMBSAppSessionCtxt`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contextId** | **string** | Contains the identifier of the Individual MBS Application Session Context resource.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMBSAppSessionCtxtRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MbsAppSessionCtxt**](MbsAppSessionCtxt.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyMBSAppSessionCtxt

> MbsAppSessionCtxt ModifyMBSAppSessionCtxt(ctx, contextId).MbsAppSessionCtxtPatch(mbsAppSessionCtxtPatch).Execute()

Request the modification of an existing Individual MBS Application Session Context resource.

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
    contextId := "contextId_example" // string | Contains the identifier of the Individual MBS Application Session Context resource. 
    mbsAppSessionCtxtPatch := *openapiclient.NewMbsAppSessionCtxtPatch() // MbsAppSessionCtxtPatch | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualMBSApplicationSessionContextDocumentApi.ModifyMBSAppSessionCtxt(context.Background(), contextId).MbsAppSessionCtxtPatch(mbsAppSessionCtxtPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualMBSApplicationSessionContextDocumentApi.ModifyMBSAppSessionCtxt``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ModifyMBSAppSessionCtxt`: MbsAppSessionCtxt
    fmt.Fprintf(os.Stdout, "Response from `IndividualMBSApplicationSessionContextDocumentApi.ModifyMBSAppSessionCtxt`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contextId** | **string** | Contains the identifier of the Individual MBS Application Session Context resource.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyMBSAppSessionCtxtRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mbsAppSessionCtxtPatch** | [**MbsAppSessionCtxtPatch**](MbsAppSessionCtxtPatch.md) |  | 

### Return type

[**MbsAppSessionCtxt**](MbsAppSessionCtxt.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

