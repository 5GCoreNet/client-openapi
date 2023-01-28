# \IndividualVALGroupDocumentDocumentApi

All URIs are relative to *https://example.com/ss-gm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIndValGroupDoc**](IndividualVALGroupDocumentDocumentApi.md#DeleteIndValGroupDoc) | **Delete** /group-documents/{groupDocId} | 
[**ModifyIndValGroupDoc**](IndividualVALGroupDocumentDocumentApi.md#ModifyIndValGroupDoc) | **Patch** /group-documents/{groupDocId} | 
[**RetrieveIndValGroupDoc**](IndividualVALGroupDocumentDocumentApi.md#RetrieveIndValGroupDoc) | **Get** /group-documents/{groupDocId} | 
[**UpdateIndValGroupDoc**](IndividualVALGroupDocumentDocumentApi.md#UpdateIndValGroupDoc) | **Put** /group-documents/{groupDocId} | 



## DeleteIndValGroupDoc

> DeleteIndValGroupDoc(ctx, groupDocId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/SS_GroupManagement"
)

func main() {
    groupDocId := "groupDocId_example" // string | String identifying an individual VAL group document resource.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualVALGroupDocumentDocumentApi.DeleteIndValGroupDoc(context.Background(), groupDocId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualVALGroupDocumentDocumentApi.DeleteIndValGroupDoc``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupDocId** | **string** | String identifying an individual VAL group document resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIndValGroupDocRequest struct via the builder pattern


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


## ModifyIndValGroupDoc

> VALGroupDocument ModifyIndValGroupDoc(ctx, groupDocId).VALGroupDocumentPatch(vALGroupDocumentPatch).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/SS_GroupManagement"
)

func main() {
    groupDocId := "groupDocId_example" // string | Identifier of an individual VAL group document.
    vALGroupDocumentPatch := *openapiclient.NewVALGroupDocumentPatch() // VALGroupDocumentPatch | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualVALGroupDocumentDocumentApi.ModifyIndValGroupDoc(context.Background(), groupDocId).VALGroupDocumentPatch(vALGroupDocumentPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualVALGroupDocumentDocumentApi.ModifyIndValGroupDoc``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ModifyIndValGroupDoc`: VALGroupDocument
    fmt.Fprintf(os.Stdout, "Response from `IndividualVALGroupDocumentDocumentApi.ModifyIndValGroupDoc`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupDocId** | **string** | Identifier of an individual VAL group document. | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyIndValGroupDocRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vALGroupDocumentPatch** | [**VALGroupDocumentPatch**](VALGroupDocumentPatch.md) |  | 

### Return type

[**VALGroupDocument**](VALGroupDocument.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveIndValGroupDoc

> VALGroupDocument RetrieveIndValGroupDoc(ctx, groupDocId).GroupMembers(groupMembers).GroupConfiguration(groupConfiguration).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/SS_GroupManagement"
)

func main() {
    groupDocId := "groupDocId_example" // string | String identifying an individual VAL group document resource.
    groupMembers := true // bool | When set to true indicates the group management server to send the members list information of the VAL group.  (optional)
    groupConfiguration := true // bool | When set to true indicates the group management server to send the group configuration information of the VAL group.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualVALGroupDocumentDocumentApi.RetrieveIndValGroupDoc(context.Background(), groupDocId).GroupMembers(groupMembers).GroupConfiguration(groupConfiguration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualVALGroupDocumentDocumentApi.RetrieveIndValGroupDoc``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveIndValGroupDoc`: VALGroupDocument
    fmt.Fprintf(os.Stdout, "Response from `IndividualVALGroupDocumentDocumentApi.RetrieveIndValGroupDoc`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupDocId** | **string** | String identifying an individual VAL group document resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveIndValGroupDocRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **groupMembers** | **bool** | When set to true indicates the group management server to send the members list information of the VAL group.  | 
 **groupConfiguration** | **bool** | When set to true indicates the group management server to send the group configuration information of the VAL group.  | 

### Return type

[**VALGroupDocument**](VALGroupDocument.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIndValGroupDoc

> VALGroupDocument UpdateIndValGroupDoc(ctx, groupDocId).VALGroupDocument(vALGroupDocument).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/SS_GroupManagement"
)

func main() {
    groupDocId := "groupDocId_example" // string | String identifying an individual VAL group document resource
    vALGroupDocument := *openapiclient.NewVALGroupDocument("ValGroupId_example") // VALGroupDocument | VAL group document to be updated in Group management server.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualVALGroupDocumentDocumentApi.UpdateIndValGroupDoc(context.Background(), groupDocId).VALGroupDocument(vALGroupDocument).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualVALGroupDocumentDocumentApi.UpdateIndValGroupDoc``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIndValGroupDoc`: VALGroupDocument
    fmt.Fprintf(os.Stdout, "Response from `IndividualVALGroupDocumentDocumentApi.UpdateIndValGroupDoc`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupDocId** | **string** | String identifying an individual VAL group document resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIndValGroupDocRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vALGroupDocument** | [**VALGroupDocument**](VALGroupDocument.md) | VAL group document to be updated in Group management server. | 

### Return type

[**VALGroupDocument**](VALGroupDocument.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

