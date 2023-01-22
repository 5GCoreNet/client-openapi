# \IndividualCPProvisioningSubscriptionApi

All URIs are relative to *https://example.com/3gpp-cp-parameter-provisioning/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIndCPProvisioningSubscription**](IndividualCPProvisioningSubscriptionApi.md#DeleteIndCPProvisioningSubscription) | **Delete** /{scsAsId}/subscriptions/{subscriptionId} | Delete a CP parameter provisioning subscription resource.
[**FetchIndCPProvisioningSubscription**](IndividualCPProvisioningSubscriptionApi.md#FetchIndCPProvisioningSubscription) | **Get** /{scsAsId}/subscriptions/{subscriptionId} | Read a CP parameter provisioning subscription resource.
[**UpdateIndCPProvisioningSubscription**](IndividualCPProvisioningSubscriptionApi.md#UpdateIndCPProvisioningSubscription) | **Put** /{scsAsId}/subscriptions/{subscriptionId} | Modify a CP parameter provisioning subscription resource.



## DeleteIndCPProvisioningSubscription

> DeleteIndCPProvisioningSubscription(ctx, scsAsId, subscriptionId).Execute()

Delete a CP parameter provisioning subscription resource.

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
    scsAsId := "scsAsId_example" // string | Identifier of the SCS/AS as defined in clause 5.2.4 of 3GPP TS 29.122.
    subscriptionId := "subscriptionId_example" // string | Subscription ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualCPProvisioningSubscriptionApi.DeleteIndCPProvisioningSubscription(context.Background(), scsAsId, subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualCPProvisioningSubscriptionApi.DeleteIndCPProvisioningSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of the SCS/AS as defined in clause 5.2.4 of 3GPP TS 29.122. | 
**subscriptionId** | **string** | Subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIndCPProvisioningSubscriptionRequest struct via the builder pattern


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


## FetchIndCPProvisioningSubscription

> CpInfo FetchIndCPProvisioningSubscription(ctx, scsAsId, subscriptionId).Execute()

Read a CP parameter provisioning subscription resource.

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
    scsAsId := "scsAsId_example" // string | Identifier of the SCS/AS as defined in clause 5.2.4 of 3GPP TS 29.122.
    subscriptionId := "subscriptionId_example" // string | Subscription ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualCPProvisioningSubscriptionApi.FetchIndCPProvisioningSubscription(context.Background(), scsAsId, subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualCPProvisioningSubscriptionApi.FetchIndCPProvisioningSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchIndCPProvisioningSubscription`: CpInfo
    fmt.Fprintf(os.Stdout, "Response from `IndividualCPProvisioningSubscriptionApi.FetchIndCPProvisioningSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of the SCS/AS as defined in clause 5.2.4 of 3GPP TS 29.122. | 
**subscriptionId** | **string** | Subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchIndCPProvisioningSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CpInfo**](CpInfo.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIndCPProvisioningSubscription

> CpInfo UpdateIndCPProvisioningSubscription(ctx, scsAsId, subscriptionId).CpInfo(cpInfo).Execute()

Modify a CP parameter provisioning subscription resource.

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
    scsAsId := "scsAsId_example" // string | Identifier of the SCS/AS as defined in clause 5.2.4 of 3GPP TS 29.122.
    subscriptionId := "subscriptionId_example" // string | Subscription ID
    cpInfo := openapiclient.CpInfo{Interface{}: new(interface{})} // CpInfo | Modify a CP parameter provisioning subscription resource.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualCPProvisioningSubscriptionApi.UpdateIndCPProvisioningSubscription(context.Background(), scsAsId, subscriptionId).CpInfo(cpInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualCPProvisioningSubscriptionApi.UpdateIndCPProvisioningSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIndCPProvisioningSubscription`: CpInfo
    fmt.Fprintf(os.Stdout, "Response from `IndividualCPProvisioningSubscriptionApi.UpdateIndCPProvisioningSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of the SCS/AS as defined in clause 5.2.4 of 3GPP TS 29.122. | 
**subscriptionId** | **string** | Subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIndCPProvisioningSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cpInfo** | [**CpInfo**](CpInfo.md) | Modify a CP parameter provisioning subscription resource. | 

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

