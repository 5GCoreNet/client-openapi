# \IndividualCPSetProvisioningApi

All URIs are relative to *https://example.com/3gpp-cp-parameter-provisioning/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIndCPSetProvisioning**](IndividualCPSetProvisioningApi.md#DeleteIndCPSetProvisioning) | **Delete** /{scsAsId}/subscriptions/{subscriptionId}/cpSets/{setId} | Delete CP at individual CP set(s) level associated with a CP parameter set Id.
[**FetchIndCPSetProvisioning**](IndividualCPSetProvisioningApi.md#FetchIndCPSetProvisioning) | **Get** /{scsAsId}/subscriptions/{subscriptionId}/cpSets/{setId} | Read CP at individual CP set(s) level associated with a CP parameter set Id.
[**UpdateIndCPSetProvisioning**](IndividualCPSetProvisioningApi.md#UpdateIndCPSetProvisioning) | **Put** /{scsAsId}/subscriptions/{subscriptionId}/cpSets/{setId} | Update CP at individual CP set(s) level associated with a CP parameter set Id.



## DeleteIndCPSetProvisioning

> DeleteIndCPSetProvisioning(ctx, scsAsId, subscriptionId, setId).Execute()

Delete CP at individual CP set(s) level associated with a CP parameter set Id.

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
    setId := "setId_example" // string | Identifier of the CP parameter set

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualCPSetProvisioningApi.DeleteIndCPSetProvisioning(context.Background(), scsAsId, subscriptionId, setId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualCPSetProvisioningApi.DeleteIndCPSetProvisioning``: %v\n", err)
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
**setId** | **string** | Identifier of the CP parameter set | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIndCPSetProvisioningRequest struct via the builder pattern


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


## FetchIndCPSetProvisioning

> CpParameterSet FetchIndCPSetProvisioning(ctx, scsAsId, subscriptionId, setId).Execute()

Read CP at individual CP set(s) level associated with a CP parameter set Id.

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
    setId := "setId_example" // string | Identifier of the CP parameter set

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualCPSetProvisioningApi.FetchIndCPSetProvisioning(context.Background(), scsAsId, subscriptionId, setId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualCPSetProvisioningApi.FetchIndCPSetProvisioning``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchIndCPSetProvisioning`: CpParameterSet
    fmt.Fprintf(os.Stdout, "Response from `IndividualCPSetProvisioningApi.FetchIndCPSetProvisioning`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of the SCS/AS as defined in clause 5.2.4 of 3GPP TS 29.122. | 
**subscriptionId** | **string** | Subscription ID | 
**setId** | **string** | Identifier of the CP parameter set | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchIndCPSetProvisioningRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**CpParameterSet**](CpParameterSet.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIndCPSetProvisioning

> CpParameterSet UpdateIndCPSetProvisioning(ctx, scsAsId, subscriptionId, setId).CpParameterSet(cpParameterSet).Execute()

Update CP at individual CP set(s) level associated with a CP parameter set Id.

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
    setId := "setId_example" // string | Identifier of the CP parameter set
    cpParameterSet := *openapiclient.NewCpParameterSet("SetId_example") // CpParameterSet | Change information for a CP parameter set.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualCPSetProvisioningApi.UpdateIndCPSetProvisioning(context.Background(), scsAsId, subscriptionId, setId).CpParameterSet(cpParameterSet).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualCPSetProvisioningApi.UpdateIndCPSetProvisioning``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIndCPSetProvisioning`: CpParameterSet
    fmt.Fprintf(os.Stdout, "Response from `IndividualCPSetProvisioningApi.UpdateIndCPSetProvisioning`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of the SCS/AS as defined in clause 5.2.4 of 3GPP TS 29.122. | 
**subscriptionId** | **string** | Subscription ID | 
**setId** | **string** | Identifier of the CP parameter set | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIndCPSetProvisioningRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **cpParameterSet** | [**CpParameterSet**](CpParameterSet.md) | Change information for a CP parameter set. | 

### Return type

[**CpParameterSet**](CpParameterSet.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

