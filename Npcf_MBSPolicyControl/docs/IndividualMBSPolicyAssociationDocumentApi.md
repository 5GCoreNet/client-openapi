# \IndividualMBSPolicyAssociationDocumentApi

All URIs are relative to *https://example.com/npcf-mbspolicycontrol/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIndMBSPolicy**](IndividualMBSPolicyAssociationDocumentApi.md#DeleteIndMBSPolicy) | **Delete** /mbs-policies/{mbsPolicyId} | Deletes an existing Individual MBS Policy resource.
[**GetIndMBSPolicy**](IndividualMBSPolicyAssociationDocumentApi.md#GetIndMBSPolicy) | **Get** /mbs-policies/{mbsPolicyId} | Read an Individual MBS Policy.



## DeleteIndMBSPolicy

> DeleteIndMBSPolicy(ctx, mbsPolicyId).Execute()

Deletes an existing Individual MBS Policy resource.

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
    mbsPolicyId := "mbsPolicyId_example" // string | Contains the identifier of the concerned Individual MBS Policy Association. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualMBSPolicyAssociationDocumentApi.DeleteIndMBSPolicy(context.Background(), mbsPolicyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualMBSPolicyAssociationDocumentApi.DeleteIndMBSPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mbsPolicyId** | **string** | Contains the identifier of the concerned Individual MBS Policy Association.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIndMBSPolicyRequest struct via the builder pattern


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


## GetIndMBSPolicy

> MbsPolicyData GetIndMBSPolicy(ctx, mbsPolicyId).Execute()

Read an Individual MBS Policy.

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
    mbsPolicyId := "mbsPolicyId_example" // string | Contains the identifier of the concerned Individual MBS Policy Association. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualMBSPolicyAssociationDocumentApi.GetIndMBSPolicy(context.Background(), mbsPolicyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualMBSPolicyAssociationDocumentApi.GetIndMBSPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIndMBSPolicy`: MbsPolicyData
    fmt.Fprintf(os.Stdout, "Response from `IndividualMBSPolicyAssociationDocumentApi.GetIndMBSPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mbsPolicyId** | **string** | Contains the identifier of the concerned Individual MBS Policy Association.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIndMBSPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MbsPolicyData**](MbsPolicyData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

