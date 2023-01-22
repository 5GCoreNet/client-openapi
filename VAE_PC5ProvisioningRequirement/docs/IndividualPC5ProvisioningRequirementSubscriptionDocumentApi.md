# \IndividualPC5ProvisioningRequirementSubscriptionDocumentApi

All URIs are relative to *https://example.com/vae-pc5-prov-req/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePC5ProvisioningRequirementSubscription**](IndividualPC5ProvisioningRequirementSubscriptionDocumentApi.md#DeletePC5ProvisioningRequirementSubscription) | **Delete** /subscriptions/{subscriptionId} | VAE PC5 Provisioning Requirement Subscription resource delete service Operation
[**ReadPC5ProvisioningRequirementSubscription**](IndividualPC5ProvisioningRequirementSubscriptionDocumentApi.md#ReadPC5ProvisioningRequirementSubscription) | **Get** /subscriptions/{subscriptionId} | VAE PC5 Provisioning Requirement Subscription resource read service Operation



## DeletePC5ProvisioningRequirementSubscription

> DeletePC5ProvisioningRequirementSubscription(ctx, subscriptionId).Execute()

VAE PC5 Provisioning Requirement Subscription resource delete service Operation

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
    subscriptionId := "subscriptionId_example" // string | Unique ID of the PC5 Provisioning Requirement Subscription to be deleted

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualPC5ProvisioningRequirementSubscriptionDocumentApi.DeletePC5ProvisioningRequirementSubscription(context.Background(), subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualPC5ProvisioningRequirementSubscriptionDocumentApi.DeletePC5ProvisioningRequirementSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Unique ID of the PC5 Provisioning Requirement Subscription to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePC5ProvisioningRequirementSubscriptionRequest struct via the builder pattern


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


## ReadPC5ProvisioningRequirementSubscription

> ProvisioningRequirement ReadPC5ProvisioningRequirementSubscription(ctx, subscriptionId).Execute()

VAE PC5 Provisioning Requirement Subscription resource read service Operation

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
    subscriptionId := "subscriptionId_example" // string | Identifier of an PC5 Provisioning Requirement Subscription resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualPC5ProvisioningRequirementSubscriptionDocumentApi.ReadPC5ProvisioningRequirementSubscription(context.Background(), subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualPC5ProvisioningRequirementSubscriptionDocumentApi.ReadPC5ProvisioningRequirementSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadPC5ProvisioningRequirementSubscription`: ProvisioningRequirement
    fmt.Fprintf(os.Stdout, "Response from `IndividualPC5ProvisioningRequirementSubscriptionDocumentApi.ReadPC5ProvisioningRequirementSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Identifier of an PC5 Provisioning Requirement Subscription resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadPC5ProvisioningRequirementSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProvisioningRequirement**](ProvisioningRequirement.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

