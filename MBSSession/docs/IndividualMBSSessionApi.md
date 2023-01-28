# \IndividualMBSSessionApi

All URIs are relative to *https://example.com/3gpp-mbs-session/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIndMBSSession**](IndividualMBSSessionApi.md#DeleteIndMBSSession) | **Delete** /mbs-sessions/{mbsSessionRef} | Request the Deletion of an existing Individual MBS Session resource.
[**ModifyIndMBSSession**](IndividualMBSSessionApi.md#ModifyIndMBSSession) | **Patch** /mbs-sessions/{mbsSessionRef} | Request the modification of an existing Individual MBS Session resource.



## DeleteIndMBSSession

> DeleteIndMBSSession(ctx, mbsSessionRef).Execute()

Request the Deletion of an existing Individual MBS Session resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/MBSSession"
)

func main() {
    mbsSessionRef := "mbsSessionRef_example" // string | Identifier of the Individual MBS Session resource.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualMBSSessionApi.DeleteIndMBSSession(context.Background(), mbsSessionRef).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualMBSSessionApi.DeleteIndMBSSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mbsSessionRef** | **string** | Identifier of the Individual MBS Session resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIndMBSSessionRequest struct via the builder pattern


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


## ModifyIndMBSSession

> ModifyIndMBSSession(ctx, mbsSessionRef).PatchItem(patchItem).Execute()

Request the modification of an existing Individual MBS Session resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/MBSSession"
)

func main() {
    mbsSessionRef := "mbsSessionRef_example" // string | Identifier of the Individual MBS Session resource.
    patchItem := []openapiclient.PatchItem{*openapiclient.NewPatchItem(*openapiclient.NewPatchOperation(), "Path_example")} // []PatchItem | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualMBSSessionApi.ModifyIndMBSSession(context.Background(), mbsSessionRef).PatchItem(patchItem).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualMBSSessionApi.ModifyIndMBSSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mbsSessionRef** | **string** | Identifier of the Individual MBS Session resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyIndMBSSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchItem** | [**[]PatchItem**](PatchItem.md) |  | 

### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

