# \IndividualApplicationPFDManagementApi

All URIs are relative to *https://example.com/3gpp-pfd-management/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIndApplicationPFDManagement**](IndividualApplicationPFDManagementApi.md#DeleteIndApplicationPFDManagement) | **Delete** /{scsAsId}/transactions/{transactionId}/applications/{appId} | Delete PFDs at individual application level.
[**FetchIndApplicationPFDManagement**](IndividualApplicationPFDManagementApi.md#FetchIndApplicationPFDManagement) | **Get** /{scsAsId}/transactions/{transactionId}/applications/{appId} | Read PFDs at individual application level.
[**ModifyIndApplicationPFDManagement**](IndividualApplicationPFDManagementApi.md#ModifyIndApplicationPFDManagement) | **Patch** /{scsAsId}/transactions/{transactionId}/applications/{appId} | Update PFDs at individual application level.
[**UpdateIndApplicationPFDManagement**](IndividualApplicationPFDManagementApi.md#UpdateIndApplicationPFDManagement) | **Put** /{scsAsId}/transactions/{transactionId}/applications/{appId} | Update PFDs at individual application level.



## DeleteIndApplicationPFDManagement

> DeleteIndApplicationPFDManagement(ctx, scsAsId, transactionId, appId).Execute()

Delete PFDs at individual application level.

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
    appId := "appId_example" // string | Identifier of the application

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualApplicationPFDManagementApi.DeleteIndApplicationPFDManagement(context.Background(), scsAsId, transactionId, appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualApplicationPFDManagementApi.DeleteIndApplicationPFDManagement``: %v\n", err)
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
**appId** | **string** | Identifier of the application | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIndApplicationPFDManagementRequest struct via the builder pattern


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


## FetchIndApplicationPFDManagement

> PfdData FetchIndApplicationPFDManagement(ctx, scsAsId, transactionId, appId).Execute()

Read PFDs at individual application level.

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
    appId := "appId_example" // string | Identifier of the application

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualApplicationPFDManagementApi.FetchIndApplicationPFDManagement(context.Background(), scsAsId, transactionId, appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualApplicationPFDManagementApi.FetchIndApplicationPFDManagement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchIndApplicationPFDManagement`: PfdData
    fmt.Fprintf(os.Stdout, "Response from `IndividualApplicationPFDManagementApi.FetchIndApplicationPFDManagement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of the SCS/AS as defined in clause 5.2.4 of 3GPP TS 29.122. | 
**transactionId** | **string** | Transaction ID | 
**appId** | **string** | Identifier of the application | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchIndApplicationPFDManagementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**PfdData**](PfdData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyIndApplicationPFDManagement

> PfdData ModifyIndApplicationPFDManagement(ctx, scsAsId, transactionId, appId).PfdData(pfdData).Execute()

Update PFDs at individual application level.

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
    appId := "appId_example" // string | Identifier of the application
    pfdData := *openapiclient.NewPfdData("ExternalAppId_example", map[string]Pfd{"key": *openapiclient.NewPfd("PfdId_example")}) // PfdData | Change information in PFD management transaction.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualApplicationPFDManagementApi.ModifyIndApplicationPFDManagement(context.Background(), scsAsId, transactionId, appId).PfdData(pfdData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualApplicationPFDManagementApi.ModifyIndApplicationPFDManagement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ModifyIndApplicationPFDManagement`: PfdData
    fmt.Fprintf(os.Stdout, "Response from `IndividualApplicationPFDManagementApi.ModifyIndApplicationPFDManagement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of the SCS/AS as defined in clause 5.2.4 of 3GPP TS 29.122. | 
**transactionId** | **string** | Transaction ID | 
**appId** | **string** | Identifier of the application | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyIndApplicationPFDManagementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pfdData** | [**PfdData**](PfdData.md) | Change information in PFD management transaction. | 

### Return type

[**PfdData**](PfdData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIndApplicationPFDManagement

> PfdData UpdateIndApplicationPFDManagement(ctx, scsAsId, transactionId, appId).PfdData(pfdData).Execute()

Update PFDs at individual application level.

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
    appId := "appId_example" // string | Identifier of the application
    pfdData := *openapiclient.NewPfdData("ExternalAppId_example", map[string]Pfd{"key": *openapiclient.NewPfd("PfdId_example")}) // PfdData | Change information in application.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualApplicationPFDManagementApi.UpdateIndApplicationPFDManagement(context.Background(), scsAsId, transactionId, appId).PfdData(pfdData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualApplicationPFDManagementApi.UpdateIndApplicationPFDManagement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIndApplicationPFDManagement`: PfdData
    fmt.Fprintf(os.Stdout, "Response from `IndividualApplicationPFDManagementApi.UpdateIndApplicationPFDManagement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of the SCS/AS as defined in clause 5.2.4 of 3GPP TS 29.122. | 
**transactionId** | **string** | Transaction ID | 
**appId** | **string** | Identifier of the application | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIndApplicationPFDManagementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pfdData** | [**PfdData**](PfdData.md) | Change information in application. | 

### Return type

[**PfdData**](PfdData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

