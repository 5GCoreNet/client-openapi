# \AMPolicyAssociationsCollectionApi

All URIs are relative to *https://example.com/npcf-am-policy-control/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIndividualAMPolicyAssociation**](AMPolicyAssociationsCollectionApi.md#CreateIndividualAMPolicyAssociation) | **Post** /policies | Create individual AM policy association.



## CreateIndividualAMPolicyAssociation

> PolicyAssociation CreateIndividualAMPolicyAssociation(ctx).PolicyAssociationRequest(policyAssociationRequest).Execute()

Create individual AM policy association.

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
    resp, r, err := apiClient.AMPolicyAssociationsCollectionApi.CreateIndividualAMPolicyAssociation(context.Background()).PolicyAssociationRequest(policyAssociationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AMPolicyAssociationsCollectionApi.CreateIndividualAMPolicyAssociation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIndividualAMPolicyAssociation`: PolicyAssociation
    fmt.Fprintf(os.Stdout, "Response from `AMPolicyAssociationsCollectionApi.CreateIndividualAMPolicyAssociation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIndividualAMPolicyAssociationRequest struct via the builder pattern


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

