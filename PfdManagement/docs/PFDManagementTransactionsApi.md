# \PFDManagementTransactionsApi

All URIs are relative to *https://example.com/3gpp-pfd-management/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePFDManagementTransaction**](PFDManagementTransactionsApi.md#CreatePFDManagementTransaction) | **Post** /{scsAsId}/transactions | Create PFDs for a given SCS/AS and one or more external Application Identifier(s).
[**FetchAllPFDManagementTransactions**](PFDManagementTransactionsApi.md#FetchAllPFDManagementTransactions) | **Get** /{scsAsId}/transactions | Read all or queried PFDs for a given SCS/AS.



## CreatePFDManagementTransaction

> PfdManagement CreatePFDManagementTransaction(ctx, scsAsId).PfdManagement(pfdManagement).Execute()

Create PFDs for a given SCS/AS and one or more external Application Identifier(s).

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
    pfdManagement := *openapiclient.NewPfdManagement(map[string]PfdData{"key": *openapiclient.NewPfdData("ExternalAppId_example", map[string]Pfd{"key": *openapiclient.NewPfd("PfdId_example")})}) // PfdManagement | Create a new transaction for PFD management.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PFDManagementTransactionsApi.CreatePFDManagementTransaction(context.Background(), scsAsId).PfdManagement(pfdManagement).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PFDManagementTransactionsApi.CreatePFDManagementTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePFDManagementTransaction`: PfdManagement
    fmt.Fprintf(os.Stdout, "Response from `PFDManagementTransactionsApi.CreatePFDManagementTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of the SCS/AS as defined in clause 5.2.4 of 3GPP TS 29.122. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePFDManagementTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pfdManagement** | [**PfdManagement**](PfdManagement.md) | Create a new transaction for PFD management. | 

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


## FetchAllPFDManagementTransactions

> []PfdManagement FetchAllPFDManagementTransactions(ctx, scsAsId).ExternalAppIds(externalAppIds).Execute()

Read all or queried PFDs for a given SCS/AS.

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
    externalAppIds := []string{"Inner_example"} // []string | The external application identifier(s) of the requested PFD data. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PFDManagementTransactionsApi.FetchAllPFDManagementTransactions(context.Background(), scsAsId).ExternalAppIds(externalAppIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PFDManagementTransactionsApi.FetchAllPFDManagementTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchAllPFDManagementTransactions`: []PfdManagement
    fmt.Fprintf(os.Stdout, "Response from `PFDManagementTransactionsApi.FetchAllPFDManagementTransactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of the SCS/AS as defined in clause 5.2.4 of 3GPP TS 29.122. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchAllPFDManagementTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **externalAppIds** | **[]string** | The external application identifier(s) of the requested PFD data. | 

### Return type

[**[]PfdManagement**](PfdManagement.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

