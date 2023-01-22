# \IndividualUEPolicyAssociationDocumentApi

All URIs are relative to *https://example.com/npcf-ue-policy-control/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIndividualUEPolicyAssociation**](IndividualUEPolicyAssociationDocumentApi.md#DeleteIndividualUEPolicyAssociation) | **Delete** /policies/{polAssoId} | Delete individual UE policy association.
[**ReadIndividualUEPolicyAssociation**](IndividualUEPolicyAssociationDocumentApi.md#ReadIndividualUEPolicyAssociation) | **Get** /policies/{polAssoId} | Read individual UE policy association.
[**ReportObservedEventTriggersForIndividualUEPolicyAssociation**](IndividualUEPolicyAssociationDocumentApi.md#ReportObservedEventTriggersForIndividualUEPolicyAssociation) | **Post** /policies/{polAssoId}/update | Report observed event triggers and possibly obtain updated policies for an individual UE policy association. 



## DeleteIndividualUEPolicyAssociation

> DeleteIndividualUEPolicyAssociation(ctx, polAssoId).Execute()

Delete individual UE policy association.

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
    polAssoId := "polAssoId_example" // string | Identifier of a policy association

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualUEPolicyAssociationDocumentApi.DeleteIndividualUEPolicyAssociation(context.Background(), polAssoId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualUEPolicyAssociationDocumentApi.DeleteIndividualUEPolicyAssociation``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteIndividualUEPolicyAssociationRequest struct via the builder pattern


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


## ReadIndividualUEPolicyAssociation

> PolicyAssociation ReadIndividualUEPolicyAssociation(ctx, polAssoId).Execute()

Read individual UE policy association.

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
    polAssoId := "polAssoId_example" // string | Identifier of a policy association

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualUEPolicyAssociationDocumentApi.ReadIndividualUEPolicyAssociation(context.Background(), polAssoId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualUEPolicyAssociationDocumentApi.ReadIndividualUEPolicyAssociation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadIndividualUEPolicyAssociation`: PolicyAssociation
    fmt.Fprintf(os.Stdout, "Response from `IndividualUEPolicyAssociationDocumentApi.ReadIndividualUEPolicyAssociation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**polAssoId** | **string** | Identifier of a policy association | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadIndividualUEPolicyAssociationRequest struct via the builder pattern


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


## ReportObservedEventTriggersForIndividualUEPolicyAssociation

> PolicyUpdate ReportObservedEventTriggersForIndividualUEPolicyAssociation(ctx, polAssoId).PolicyAssociationUpdateRequest(policyAssociationUpdateRequest).Execute()

Report observed event triggers and possibly obtain updated policies for an individual UE policy association. 

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
    polAssoId := "polAssoId_example" // string | Identifier of a policy association
    policyAssociationUpdateRequest := *openapiclient.NewPolicyAssociationUpdateRequest() // PolicyAssociationUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualUEPolicyAssociationDocumentApi.ReportObservedEventTriggersForIndividualUEPolicyAssociation(context.Background(), polAssoId).PolicyAssociationUpdateRequest(policyAssociationUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualUEPolicyAssociationDocumentApi.ReportObservedEventTriggersForIndividualUEPolicyAssociation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportObservedEventTriggersForIndividualUEPolicyAssociation`: PolicyUpdate
    fmt.Fprintf(os.Stdout, "Response from `IndividualUEPolicyAssociationDocumentApi.ReportObservedEventTriggersForIndividualUEPolicyAssociation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**polAssoId** | **string** | Identifier of a policy association | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportObservedEventTriggersForIndividualUEPolicyAssociationRequest struct via the builder pattern


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

