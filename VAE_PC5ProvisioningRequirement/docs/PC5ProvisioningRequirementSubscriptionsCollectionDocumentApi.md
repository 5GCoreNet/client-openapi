# \PC5ProvisioningRequirementSubscriptionsCollectionDocumentApi

All URIs are relative to *https://example.com/vae-pc5-prov-req/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](PC5ProvisioningRequirementSubscriptionsCollectionDocumentApi.md#Create) | **Post** /subscriptions | VAE_PC5 Provisioning Requirement resource create service Operation



## Create

> ProvisioningRequirement Create(ctx).ProvisioningRequirement(provisioningRequirement).Execute()

VAE_PC5 Provisioning Requirement resource create service Operation

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
    provisioningRequirement := *openapiclient.NewProvisioningRequirement("NotifUri_example", "ServiceId_example") // ProvisioningRequirement | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PC5ProvisioningRequirementSubscriptionsCollectionDocumentApi.Create(context.Background()).ProvisioningRequirement(provisioningRequirement).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PC5ProvisioningRequirementSubscriptionsCollectionDocumentApi.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create`: ProvisioningRequirement
    fmt.Fprintf(os.Stdout, "Response from `PC5ProvisioningRequirementSubscriptionsCollectionDocumentApi.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **provisioningRequirement** | [**ProvisioningRequirement**](ProvisioningRequirement.md) |  | 

### Return type

[**ProvisioningRequirement**](ProvisioningRequirement.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

