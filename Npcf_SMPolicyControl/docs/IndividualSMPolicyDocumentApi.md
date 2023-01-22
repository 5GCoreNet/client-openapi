# \IndividualSMPolicyDocumentApi

All URIs are relative to *https://example.com/npcf-smpolicycontrol/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSMPolicy**](IndividualSMPolicyDocumentApi.md#DeleteSMPolicy) | **Post** /sm-policies/{smPolicyId}/delete | Delete an existing Individual SM Policy
[**GetSMPolicy**](IndividualSMPolicyDocumentApi.md#GetSMPolicy) | **Get** /sm-policies/{smPolicyId} | Read an Individual SM Policy
[**UpdateSMPolicy**](IndividualSMPolicyDocumentApi.md#UpdateSMPolicy) | **Post** /sm-policies/{smPolicyId}/update | Update an existing Individual SM Policy



## DeleteSMPolicy

> DeleteSMPolicy(ctx, smPolicyId).SmPolicyDeleteData(smPolicyDeleteData).Execute()

Delete an existing Individual SM Policy

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
    smPolicyId := "smPolicyId_example" // string | Identifier of a policy association
    smPolicyDeleteData := *openapiclient.NewSmPolicyDeleteData() // SmPolicyDeleteData | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualSMPolicyDocumentApi.DeleteSMPolicy(context.Background(), smPolicyId).SmPolicyDeleteData(smPolicyDeleteData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualSMPolicyDocumentApi.DeleteSMPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**smPolicyId** | **string** | Identifier of a policy association | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSMPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **smPolicyDeleteData** | [**SmPolicyDeleteData**](SmPolicyDeleteData.md) |  | 

### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSMPolicy

> SmPolicyControl GetSMPolicy(ctx, smPolicyId).Execute()

Read an Individual SM Policy

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
    smPolicyId := "smPolicyId_example" // string | Identifier of a policy association

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualSMPolicyDocumentApi.GetSMPolicy(context.Background(), smPolicyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualSMPolicyDocumentApi.GetSMPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSMPolicy`: SmPolicyControl
    fmt.Fprintf(os.Stdout, "Response from `IndividualSMPolicyDocumentApi.GetSMPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**smPolicyId** | **string** | Identifier of a policy association | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSMPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SmPolicyControl**](SmPolicyControl.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSMPolicy

> SmPolicyDecision UpdateSMPolicy(ctx, smPolicyId).SmPolicyUpdateContextData(smPolicyUpdateContextData).Execute()

Update an existing Individual SM Policy

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
    smPolicyId := "smPolicyId_example" // string | Identifier of a policy association
    smPolicyUpdateContextData := *openapiclient.NewSmPolicyUpdateContextData() // SmPolicyUpdateContextData | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualSMPolicyDocumentApi.UpdateSMPolicy(context.Background(), smPolicyId).SmPolicyUpdateContextData(smPolicyUpdateContextData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualSMPolicyDocumentApi.UpdateSMPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSMPolicy`: SmPolicyDecision
    fmt.Fprintf(os.Stdout, "Response from `IndividualSMPolicyDocumentApi.UpdateSMPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**smPolicyId** | **string** | Identifier of a policy association | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSMPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **smPolicyUpdateContextData** | [**SmPolicyUpdateContextData**](SmPolicyUpdateContextData.md) |  | 

### Return type

[**SmPolicyDecision**](SmPolicyDecision.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

