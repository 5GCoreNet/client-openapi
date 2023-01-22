# \IndividualBDTPolicyDocumentApi

All URIs are relative to *https://example.com/npcf-bdtpolicycontrol/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBDTPolicy**](IndividualBDTPolicyDocumentApi.md#GetBDTPolicy) | **Get** /bdtpolicies/{bdtPolicyId} | Read an Individual BDT policy
[**UpdateBDTPolicy**](IndividualBDTPolicyDocumentApi.md#UpdateBDTPolicy) | **Patch** /bdtpolicies/{bdtPolicyId} | Update an Individual BDT policy



## GetBDTPolicy

> BdtPolicy GetBDTPolicy(ctx, bdtPolicyId).Execute()

Read an Individual BDT policy

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
    bdtPolicyId := "bdtPolicyId_example" // string | String identifying the individual BDT policy resource in the PCF.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualBDTPolicyDocumentApi.GetBDTPolicy(context.Background(), bdtPolicyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualBDTPolicyDocumentApi.GetBDTPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBDTPolicy`: BdtPolicy
    fmt.Fprintf(os.Stdout, "Response from `IndividualBDTPolicyDocumentApi.GetBDTPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bdtPolicyId** | **string** | String identifying the individual BDT policy resource in the PCF. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBDTPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BdtPolicy**](BdtPolicy.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBDTPolicy

> BdtPolicy UpdateBDTPolicy(ctx, bdtPolicyId).PatchBdtPolicy(patchBdtPolicy).Execute()

Update an Individual BDT policy

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
    bdtPolicyId := "bdtPolicyId_example" // string | String identifying the individual BDT policy resource in the PCF.
    patchBdtPolicy := *openapiclient.NewPatchBdtPolicy() // PatchBdtPolicy | Contains modification instruction to be performed on the BdtPolicy data structure to select a transfer policy and in addition, may indicate whether the BDT warning notification is enabled or disabled. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualBDTPolicyDocumentApi.UpdateBDTPolicy(context.Background(), bdtPolicyId).PatchBdtPolicy(patchBdtPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualBDTPolicyDocumentApi.UpdateBDTPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBDTPolicy`: BdtPolicy
    fmt.Fprintf(os.Stdout, "Response from `IndividualBDTPolicyDocumentApi.UpdateBDTPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bdtPolicyId** | **string** | String identifying the individual BDT policy resource in the PCF. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBDTPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchBdtPolicy** | [**PatchBdtPolicy**](PatchBdtPolicy.md) | Contains modification instruction to be performed on the BdtPolicy data structure to select a transfer policy and in addition, may indicate whether the BDT warning notification is enabled or disabled.  | 

### Return type

[**BdtPolicy**](BdtPolicy.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

