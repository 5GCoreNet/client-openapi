# \VAEPC5ProvisioningRequirementSubscriptionResourcePutServiceOperationApi

All URIs are relative to *https://example.com/vae-pc5-prov-req/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdatePC5ProvisioningRequirementSubscription**](VAEPC5ProvisioningRequirementSubscriptionResourcePutServiceOperationApi.md#UpdatePC5ProvisioningRequirementSubscription) | **Put** /subscriptions/{subscriptionId} | Updates/replaces an existing subscription resource



## UpdatePC5ProvisioningRequirementSubscription

> ProvisioningRequirement UpdatePC5ProvisioningRequirementSubscription(ctx, subscriptionId).ProvisioningRequirement(provisioningRequirement).Execute()

Updates/replaces an existing subscription resource

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/VAE_PC5ProvisioningRequirement"
)

func main() {
    subscriptionId := "subscriptionId_example" // string | Identifier of an PC5 Provisioning Requirement Subscription resource
    provisioningRequirement := *openapiclient.NewProvisioningRequirement("NotifUri_example", "ServiceId_example") // ProvisioningRequirement | Parameters to update/replace the existing subscription

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VAEPC5ProvisioningRequirementSubscriptionResourcePutServiceOperationApi.UpdatePC5ProvisioningRequirementSubscription(context.Background(), subscriptionId).ProvisioningRequirement(provisioningRequirement).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VAEPC5ProvisioningRequirementSubscriptionResourcePutServiceOperationApi.UpdatePC5ProvisioningRequirementSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePC5ProvisioningRequirementSubscription`: ProvisioningRequirement
    fmt.Fprintf(os.Stdout, "Response from `VAEPC5ProvisioningRequirementSubscriptionResourcePutServiceOperationApi.UpdatePC5ProvisioningRequirementSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Identifier of an PC5 Provisioning Requirement Subscription resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePC5ProvisioningRequirementSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **provisioningRequirement** | [**ProvisioningRequirement**](ProvisioningRequirement.md) | Parameters to update/replace the existing subscription | 

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

