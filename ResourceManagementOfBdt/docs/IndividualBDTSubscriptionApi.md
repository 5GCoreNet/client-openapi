# \IndividualBDTSubscriptionApi

All URIs are relative to *https://example.com/3gpp-bdt/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteBDTSubscription**](IndividualBDTSubscriptionApi.md#DeleteBDTSubscription) | **Delete** /{scsAsId}/subscriptions/{subscriptionId} | Delete a background data transfer resource.
[**FetchIndBDTSubscription**](IndividualBDTSubscriptionApi.md#FetchIndBDTSubscription) | **Get** /{scsAsId}/subscriptions/{subscriptionId} | Read a background data transfer subscription resource.
[**ModifyBDTSubscription**](IndividualBDTSubscriptionApi.md#ModifyBDTSubscription) | **Patch** /{scsAsId}/subscriptions/{subscriptionId} | Modify a background data transfer subscription resource to select one of the transfer policies offered by the SCEF.
[**UpdateBDTSubscription**](IndividualBDTSubscriptionApi.md#UpdateBDTSubscription) | **Put** /{scsAsId}/subscriptions/{subscriptionId} | Update a background data transfer subscription resource for negotiation of background data transfer policy.



## DeleteBDTSubscription

> DeleteBDTSubscription(ctx, scsAsId, subscriptionId).Execute()

Delete a background data transfer resource.

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
    scsAsId := "scsAsId_example" // string | String identifying the SCS/AS.
    subscriptionId := "subscriptionId_example" // string | String identifying the individual BDT policy resource in the SCEF.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualBDTSubscriptionApi.DeleteBDTSubscription(context.Background(), scsAsId, subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualBDTSubscriptionApi.DeleteBDTSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | String identifying the SCS/AS. | 
**subscriptionId** | **string** | String identifying the individual BDT policy resource in the SCEF. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBDTSubscriptionRequest struct via the builder pattern


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


## FetchIndBDTSubscription

> Bdt FetchIndBDTSubscription(ctx, scsAsId, subscriptionId).Execute()

Read a background data transfer subscription resource.

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
    scsAsId := "scsAsId_example" // string | String identifying the SCS/AS.
    subscriptionId := "subscriptionId_example" // string | String identifying the individual BDT policy resource in the SCEF.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualBDTSubscriptionApi.FetchIndBDTSubscription(context.Background(), scsAsId, subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualBDTSubscriptionApi.FetchIndBDTSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchIndBDTSubscription`: Bdt
    fmt.Fprintf(os.Stdout, "Response from `IndividualBDTSubscriptionApi.FetchIndBDTSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | String identifying the SCS/AS. | 
**subscriptionId** | **string** | String identifying the individual BDT policy resource in the SCEF. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchIndBDTSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Bdt**](Bdt.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyBDTSubscription

> Bdt ModifyBDTSubscription(ctx, scsAsId, subscriptionId).BdtPatch(bdtPatch).Execute()

Modify a background data transfer subscription resource to select one of the transfer policies offered by the SCEF.

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
    scsAsId := "scsAsId_example" // string | String identifying the SCS/AS.
    subscriptionId := "subscriptionId_example" // string | String identifying the individual BDT policy resource in the SCEF.
    bdtPatch := *openapiclient.NewBdtPatch(int32(123)) // BdtPatch | Contains information to be performed on the Bdt data structure to select a transfer policy.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualBDTSubscriptionApi.ModifyBDTSubscription(context.Background(), scsAsId, subscriptionId).BdtPatch(bdtPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualBDTSubscriptionApi.ModifyBDTSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ModifyBDTSubscription`: Bdt
    fmt.Fprintf(os.Stdout, "Response from `IndividualBDTSubscriptionApi.ModifyBDTSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | String identifying the SCS/AS. | 
**subscriptionId** | **string** | String identifying the individual BDT policy resource in the SCEF. | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyBDTSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **bdtPatch** | [**BdtPatch**](BdtPatch.md) | Contains information to be performed on the Bdt data structure to select a transfer policy. | 

### Return type

[**Bdt**](Bdt.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBDTSubscription

> Bdt UpdateBDTSubscription(ctx, scsAsId, subscriptionId).Bdt(bdt).Execute()

Update a background data transfer subscription resource for negotiation of background data transfer policy.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    scsAsId := "scsAsId_example" // string | String identifying the SCS/AS.
    subscriptionId := "subscriptionId_example" // string | String identifying the individual BDT policy resource in the SCEF.
    bdt := *openapiclient.NewBdt(*openapiclient.NewUsageThreshold(), int32(123), *openapiclient.NewTimeWindow(time.Now(), time.Now())) // Bdt | Parameters to update/replace the existing BDT subscription

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualBDTSubscriptionApi.UpdateBDTSubscription(context.Background(), scsAsId, subscriptionId).Bdt(bdt).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualBDTSubscriptionApi.UpdateBDTSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBDTSubscription`: Bdt
    fmt.Fprintf(os.Stdout, "Response from `IndividualBDTSubscriptionApi.UpdateBDTSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | String identifying the SCS/AS. | 
**subscriptionId** | **string** | String identifying the individual BDT policy resource in the SCEF. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBDTSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **bdt** | [**Bdt**](Bdt.md) | Parameters to update/replace the existing BDT subscription | 

### Return type

[**Bdt**](Bdt.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

