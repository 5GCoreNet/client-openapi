# \IndividualMBSPOlicyAssociationDocumentApi

All URIs are relative to *https://example.com/npcf-mbspolicycontrol/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateIndMBSPolicy**](IndividualMBSPOlicyAssociationDocumentApi.md#UpdateIndMBSPolicy) | **Post** /mbs-policies/{mbsPolicyId}/update | Request the update of an existing MBS Policy Association.



## UpdateIndMBSPolicy

> MbsPolicyData UpdateIndMBSPolicy(ctx, mbsPolicyId).MbsPolicyCtxtDataUpdate(mbsPolicyCtxtDataUpdate).Execute()

Request the update of an existing MBS Policy Association.

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
    mbsPolicyCtxtDataUpdate := *openapiclient.NewMbsPolicyCtxtDataUpdate() // MbsPolicyCtxtDataUpdate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualMBSPOlicyAssociationDocumentApi.UpdateIndMBSPolicy(context.Background(), mbsPolicyId).MbsPolicyCtxtDataUpdate(mbsPolicyCtxtDataUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualMBSPOlicyAssociationDocumentApi.UpdateIndMBSPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIndMBSPolicy`: MbsPolicyData
    fmt.Fprintf(os.Stdout, "Response from `IndividualMBSPOlicyAssociationDocumentApi.UpdateIndMBSPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mbsPolicyId** | **string** | Contains the identifier of the concerned Individual MBS Policy Association.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIndMBSPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mbsPolicyCtxtDataUpdate** | [**MbsPolicyCtxtDataUpdate**](MbsPolicyCtxtDataUpdate.md) |  | 

### Return type

[**MbsPolicyData**](MbsPolicyData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

