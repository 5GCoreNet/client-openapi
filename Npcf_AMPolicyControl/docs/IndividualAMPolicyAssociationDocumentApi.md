# \IndividualAMPolicyAssociationDocumentApi

All URIs are relative to *https://example.com/npcf-am-policy-control/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIndividualAMPolicyAssociation**](IndividualAMPolicyAssociationDocumentApi.md#DeleteIndividualAMPolicyAssociation) | **Delete** /policies/{polAssoId} | Delete individual AM policy association.
[**ReadIndividualAMPolicyAssociation**](IndividualAMPolicyAssociationDocumentApi.md#ReadIndividualAMPolicyAssociation) | **Get** /policies/{polAssoId} | Read individual AM policy association.
[**ReportObservedEventTriggersForIndividualAMPolicyAssociation**](IndividualAMPolicyAssociationDocumentApi.md#ReportObservedEventTriggersForIndividualAMPolicyAssociation) | **Post** /policies/{polAssoId}/update | Report observed event triggers and obtain updated policies for an individual AM policy association. 



## DeleteIndividualAMPolicyAssociation

> DeleteIndividualAMPolicyAssociation(ctx, polAssoId).Execute()

Delete individual AM policy association.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Npcf_AMPolicyControl"
)

func main() {
    polAssoId := "polAssoId_example" // string | Identifier of a policy association

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualAMPolicyAssociationDocumentApi.DeleteIndividualAMPolicyAssociation(context.Background(), polAssoId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualAMPolicyAssociationDocumentApi.DeleteIndividualAMPolicyAssociation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**polAssoId** | **string** | Identifier of a policy association | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIndividualAMPolicyAssociationRequest struct via the builder pattern


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


## ReadIndividualAMPolicyAssociation

> PolicyAssociation ReadIndividualAMPolicyAssociation(ctx, polAssoId).Execute()

Read individual AM policy association.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Npcf_AMPolicyControl"
)

func main() {
    polAssoId := "polAssoId_example" // string | Identifier of a policy association

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualAMPolicyAssociationDocumentApi.ReadIndividualAMPolicyAssociation(context.Background(), polAssoId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualAMPolicyAssociationDocumentApi.ReadIndividualAMPolicyAssociation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadIndividualAMPolicyAssociation`: PolicyAssociation
    fmt.Fprintf(os.Stdout, "Response from `IndividualAMPolicyAssociationDocumentApi.ReadIndividualAMPolicyAssociation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**polAssoId** | **string** | Identifier of a policy association | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadIndividualAMPolicyAssociationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PolicyAssociation**](PolicyAssociation.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportObservedEventTriggersForIndividualAMPolicyAssociation

> PolicyUpdate ReportObservedEventTriggersForIndividualAMPolicyAssociation(ctx, polAssoId).PolicyAssociationUpdateRequest(policyAssociationUpdateRequest).Execute()

Report observed event triggers and obtain updated policies for an individual AM policy association. 

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Npcf_AMPolicyControl"
)

func main() {
    polAssoId := "polAssoId_example" // string | Identifier of a policy association
    policyAssociationUpdateRequest := *openapiclient.NewPolicyAssociationUpdateRequest() // PolicyAssociationUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualAMPolicyAssociationDocumentApi.ReportObservedEventTriggersForIndividualAMPolicyAssociation(context.Background(), polAssoId).PolicyAssociationUpdateRequest(policyAssociationUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualAMPolicyAssociationDocumentApi.ReportObservedEventTriggersForIndividualAMPolicyAssociation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportObservedEventTriggersForIndividualAMPolicyAssociation`: PolicyUpdate
    fmt.Fprintf(os.Stdout, "Response from `IndividualAMPolicyAssociationDocumentApi.ReportObservedEventTriggersForIndividualAMPolicyAssociation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**polAssoId** | **string** | Identifier of a policy association | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportObservedEventTriggersForIndividualAMPolicyAssociationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **policyAssociationUpdateRequest** | [**PolicyAssociationUpdateRequest**](PolicyAssociationUpdateRequest.md) |  | 

### Return type

[**PolicyUpdate**](PolicyUpdate.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

