# \UEPolicyAssociationsCollectionApi

All URIs are relative to *https://example.com/npcf-ue-policy-control/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIndividualUEPolicyAssociation**](UEPolicyAssociationsCollectionApi.md#CreateIndividualUEPolicyAssociation) | **Post** /policies | Create individual UE policy association.



## CreateIndividualUEPolicyAssociation

> PolicyAssociation CreateIndividualUEPolicyAssociation(ctx).PolicyAssociationRequest(policyAssociationRequest).Execute()

Create individual UE policy association.

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
    policyAssociationRequest := *openapiclient.NewPolicyAssociationRequest("NotificationUri_example", "Supi_example", "SuppFeat_example") // PolicyAssociationRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UEPolicyAssociationsCollectionApi.CreateIndividualUEPolicyAssociation(context.Background()).PolicyAssociationRequest(policyAssociationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UEPolicyAssociationsCollectionApi.CreateIndividualUEPolicyAssociation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIndividualUEPolicyAssociation`: PolicyAssociation
    fmt.Fprintf(os.Stdout, "Response from `UEPolicyAssociationsCollectionApi.CreateIndividualUEPolicyAssociation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIndividualUEPolicyAssociationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **policyAssociationRequest** | [**PolicyAssociationRequest**](PolicyAssociationRequest.md) |  | 

### Return type

[**PolicyAssociation**](PolicyAssociation.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

