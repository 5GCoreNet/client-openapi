# \CPProvisioningSubscriptionsApi

All URIs are relative to *https://example.com/3gpp-cp-parameter-provisioning/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCPProvisioningSubscription**](CPProvisioningSubscriptionsApi.md#CreateCPProvisioningSubscription) | **Post** /{scsAsId}/subscriptions | Create a new subscription resource of provisioning CP parameter set(s).
[**FetchAllCPProvisioningSubscriptions**](CPProvisioningSubscriptionsApi.md#FetchAllCPProvisioningSubscriptions) | **Get** /{scsAsId}/subscriptions | Read all active CP parameter provisioning subscription resources for a given SCS/AS.



## CreateCPProvisioningSubscription

> CpInfo CreateCPProvisioningSubscription(ctx, scsAsId).CpInfo(cpInfo).Execute()

Create a new subscription resource of provisioning CP parameter set(s).

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/CpProvisioning"
)

func main() {
    scsAsId := "scsAsId_example" // string | Identifier of the SCS/AS as defined in clause 5.2.4 of 3GPP TS 29.122.
    cpInfo := openapiclient.CpInfo{Interface{}: new(interface{})} // CpInfo | create new subscriptions for a given SCS/AS and the provisioning CP parameter sets.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CPProvisioningSubscriptionsApi.CreateCPProvisioningSubscription(context.Background(), scsAsId).CpInfo(cpInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CPProvisioningSubscriptionsApi.CreateCPProvisioningSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCPProvisioningSubscription`: CpInfo
    fmt.Fprintf(os.Stdout, "Response from `CPProvisioningSubscriptionsApi.CreateCPProvisioningSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of the SCS/AS as defined in clause 5.2.4 of 3GPP TS 29.122. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCPProvisioningSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cpInfo** | [**CpInfo**](CpInfo.md) | create new subscriptions for a given SCS/AS and the provisioning CP parameter sets. | 

### Return type

[**CpInfo**](CpInfo.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchAllCPProvisioningSubscriptions

> []CpInfo FetchAllCPProvisioningSubscriptions(ctx, scsAsId).Execute()

Read all active CP parameter provisioning subscription resources for a given SCS/AS.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/CpProvisioning"
)

func main() {
    scsAsId := "scsAsId_example" // string | Identifier of the SCS/AS as defined in clause 5.2.4 of 3GPP TS 29.122.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CPProvisioningSubscriptionsApi.FetchAllCPProvisioningSubscriptions(context.Background(), scsAsId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CPProvisioningSubscriptionsApi.FetchAllCPProvisioningSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchAllCPProvisioningSubscriptions`: []CpInfo
    fmt.Fprintf(os.Stdout, "Response from `CPProvisioningSubscriptionsApi.FetchAllCPProvisioningSubscriptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of the SCS/AS as defined in clause 5.2.4 of 3GPP TS 29.122. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchAllCPProvisioningSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]CpInfo**](CpInfo.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

