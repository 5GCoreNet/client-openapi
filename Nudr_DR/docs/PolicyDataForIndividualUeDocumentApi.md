# \PolicyDataForIndividualUeDocumentApi

All URIs are relative to *https://example.com/nudr-dr/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReadPolicyData**](PolicyDataForIndividualUeDocumentApi.md#ReadPolicyData) | **Get** /policy-data/ues/{ueId} | Retrieve the policy data for a subscriber



## ReadPolicyData

> PolicyDataForIndividualUe ReadPolicyData(ctx, ueId).SuppFeat(suppFeat).DataSubsetNames(dataSubsetNames).Execute()

Retrieve the policy data for a subscriber

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
    ueId := "ueId_example" // string | 
    suppFeat := "suppFeat_example" // string | Supported Features (optional)
    dataSubsetNames := []openapiclient.PolicyDataSubset{*openapiclient.NewPolicyDataSubset()} // []PolicyDataSubset | List of policy data subset names (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyDataForIndividualUeDocumentApi.ReadPolicyData(context.Background(), ueId).SuppFeat(suppFeat).DataSubsetNames(dataSubsetNames).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyDataForIndividualUeDocumentApi.ReadPolicyData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadPolicyData`: PolicyDataForIndividualUe
    fmt.Fprintf(os.Stdout, "Response from `PolicyDataForIndividualUeDocumentApi.ReadPolicyData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadPolicyDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **suppFeat** | **string** | Supported Features | 
 **dataSubsetNames** | [**[]PolicyDataSubset**](PolicyDataSubset.md) | List of policy data subset names | 

### Return type

[**PolicyDataForIndividualUe**](PolicyDataForIndividualUe.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

