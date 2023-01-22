# \IndividualPFDManagementTransactionApi

All URIs are relative to *https://example.com/3gpp-pfd-management/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIndPFDManagementTransaction**](IndividualPFDManagementTransactionApi.md#DeleteIndPFDManagementTransaction) | **Delete** /{scsAsId}/transactions/{transactionId} | Delete PFDs for a given SCS/AS and a transaction for one or more external Application Identifier(s).
[**FetchIndPFDManagementTransaction**](IndividualPFDManagementTransactionApi.md#FetchIndPFDManagementTransaction) | **Get** /{scsAsId}/transactions/{transactionId} | Read all PFDs for a given SCS/AS and a transaction for one or more external Application Identifier(s).
[**ModifyIndPFDManagementTransaction**](IndividualPFDManagementTransactionApi.md#ModifyIndPFDManagementTransaction) | **Patch** /{scsAsId}/transactions/{transactionId} | Modify an existing PFD Management Transaction resource.
[**UpdateIndPFDManagementTransaction**](IndividualPFDManagementTransactionApi.md#UpdateIndPFDManagementTransaction) | **Put** /{scsAsId}/transactions/{transactionId} | Update PFDs for a given SCS/AS and a transaction for one or more external Application Identifier(s).



## DeleteIndPFDManagementTransaction

> DeleteIndPFDManagementTransaction(ctx, scsAsId, transactionId).Execute()

Delete PFDs for a given SCS/AS and a transaction for one or more external Application Identifier(s).

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
    scsAsId := "scsAsId_example" // string | Identifier of the SCS/AS as defined in clause 5.2.4 of 3GPP TS 29.122.
    transactionId := "transactionId_example" // string | Transaction ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualPFDManagementTransactionApi.DeleteIndPFDManagementTransaction(context.Background(), scsAsId, transactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualPFDManagementTransactionApi.DeleteIndPFDManagementTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of the SCS/AS as defined in clause 5.2.4 of 3GPP TS 29.122. | 
**transactionId** | **string** | Transaction ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIndPFDManagementTransactionRequest struct via the builder pattern


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


## FetchIndPFDManagementTransaction

> PfdManagement FetchIndPFDManagementTransaction(ctx, scsAsId, transactionId).Execute()

Read all PFDs for a given SCS/AS and a transaction for one or more external Application Identifier(s).

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
    scsAsId := "scsAsId_example" // string | Identifier of the SCS/AS as defined in clause 5.2.4 of 3GPP TS 29.122.
    transactionId := "transactionId_example" // string | Transaction ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualPFDManagementTransactionApi.FetchIndPFDManagementTransaction(context.Background(), scsAsId, transactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualPFDManagementTransactionApi.FetchIndPFDManagementTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchIndPFDManagementTransaction`: PfdManagement
    fmt.Fprintf(os.Stdout, "Response from `IndividualPFDManagementTransactionApi.FetchIndPFDManagementTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of the SCS/AS as defined in clause 5.2.4 of 3GPP TS 29.122. | 
**transactionId** | **string** | Transaction ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchIndPFDManagementTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PfdManagement**](PfdManagement.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyIndPFDManagementTransaction

> PfdManagement ModifyIndPFDManagementTransaction(ctx, scsAsId, transactionId).PfdManagementPatch(pfdManagementPatch).Execute()

Modify an existing PFD Management Transaction resource.

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
    scsAsId := "scsAsId_example" // string | Identifier of the SCS/AS as defined in clause 5.2.4 of 3GPP TS 29.122.
    transactionId := "transactionId_example" // string | Transaction ID
    pfdManagementPatch := *openapiclient.NewPfdManagementPatch() // PfdManagementPatch | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualPFDManagementTransactionApi.ModifyIndPFDManagementTransaction(context.Background(), scsAsId, transactionId).PfdManagementPatch(pfdManagementPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualPFDManagementTransactionApi.ModifyIndPFDManagementTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ModifyIndPFDManagementTransaction`: PfdManagement
    fmt.Fprintf(os.Stdout, "Response from `IndividualPFDManagementTransactionApi.ModifyIndPFDManagementTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of the SCS/AS as defined in clause 5.2.4 of 3GPP TS 29.122. | 
**transactionId** | **string** | Transaction ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyIndPFDManagementTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pfdManagementPatch** | [**PfdManagementPatch**](PfdManagementPatch.md) |  | 

### Return type

[**PfdManagement**](PfdManagement.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIndPFDManagementTransaction

> PfdManagement UpdateIndPFDManagementTransaction(ctx, scsAsId, transactionId).PfdManagement(pfdManagement).Execute()

Update PFDs for a given SCS/AS and a transaction for one or more external Application Identifier(s).

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
    scsAsId := "scsAsId_example" // string | Identifier of the SCS/AS as defined in clause 5.2.4 of 3GPP TS 29.122.
    transactionId := "transactionId_example" // string | Transaction ID
    pfdManagement := *openapiclient.NewPfdManagement(map[string]PfdData{"key": *openapiclient.NewPfdData("ExternalAppId_example", map[string]Pfd{"key": *openapiclient.NewPfd("PfdId_example")})}) // PfdManagement | Change information in PFD management transaction.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualPFDManagementTransactionApi.UpdateIndPFDManagementTransaction(context.Background(), scsAsId, transactionId).PfdManagement(pfdManagement).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualPFDManagementTransactionApi.UpdateIndPFDManagementTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIndPFDManagementTransaction`: PfdManagement
    fmt.Fprintf(os.Stdout, "Response from `IndividualPFDManagementTransactionApi.UpdateIndPFDManagementTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of the SCS/AS as defined in clause 5.2.4 of 3GPP TS 29.122. | 
**transactionId** | **string** | Transaction ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIndPFDManagementTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pfdManagement** | [**PfdManagement**](PfdManagement.md) | Change information in PFD management transaction. | 

### Return type

[**PfdManagement**](PfdManagement.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

