# \BDTSubscriptionApi

All URIs are relative to *https://example.com/3gpp-bdt/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBDTSubscription**](BDTSubscriptionApi.md#CreateBDTSubscription) | **Post** /{scsAsId}/subscriptions | Creates a new background data transfer subscription resource.
[**FetchAllActiveBDTSubscriptions**](BDTSubscriptionApi.md#FetchAllActiveBDTSubscriptions) | **Get** /{scsAsId}/subscriptions | Fetch all active background data transfer subscription resources for a given SCS/AS.



## CreateBDTSubscription

> Bdt CreateBDTSubscription(ctx, scsAsId).Bdt(bdt).Execute()

Creates a new background data transfer subscription resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/5GCoreNet/client-openapi/ResourceManagementOfBdt"
)

func main() {
    scsAsId := "scsAsId_example" // string | String identifying the SCS/AS.
    bdt := *openapiclient.NewBdt(*openapiclient.NewUsageThreshold(), int32(123), *openapiclient.NewTimeWindow(time.Now(), time.Now())) // Bdt | Contains the data to create a BDT subscription.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BDTSubscriptionApi.CreateBDTSubscription(context.Background(), scsAsId).Bdt(bdt).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BDTSubscriptionApi.CreateBDTSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBDTSubscription`: Bdt
    fmt.Fprintf(os.Stdout, "Response from `BDTSubscriptionApi.CreateBDTSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | String identifying the SCS/AS. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBDTSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bdt** | [**Bdt**](Bdt.md) | Contains the data to create a BDT subscription. | 

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


## FetchAllActiveBDTSubscriptions

> []Bdt FetchAllActiveBDTSubscriptions(ctx, scsAsId).Execute()

Fetch all active background data transfer subscription resources for a given SCS/AS.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/ResourceManagementOfBdt"
)

func main() {
    scsAsId := "scsAsId_example" // string | String identifying the SCS/AS.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BDTSubscriptionApi.FetchAllActiveBDTSubscriptions(context.Background(), scsAsId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BDTSubscriptionApi.FetchAllActiveBDTSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchAllActiveBDTSubscriptions`: []Bdt
    fmt.Fprintf(os.Stdout, "Response from `BDTSubscriptionApi.FetchAllActiveBDTSubscriptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | String identifying the SCS/AS. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchAllActiveBDTSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Bdt**](Bdt.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

