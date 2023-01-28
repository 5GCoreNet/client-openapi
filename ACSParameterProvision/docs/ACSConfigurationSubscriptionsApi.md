# \ACSConfigurationSubscriptionsApi

All URIs are relative to *https://example.com/3gpp-acs-pp/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAnSubscription**](ACSConfigurationSubscriptionsApi.md#CreateAnSubscription) | **Post** /{afId}/subscriptions | Creates a new subscription resource
[**ReadAllSubscriptions**](ACSConfigurationSubscriptionsApi.md#ReadAllSubscriptions) | **Get** /{afId}/subscriptions | read all of the active subscriptions for the AF



## CreateAnSubscription

> AcsConfigurationData CreateAnSubscription(ctx, afId).AcsConfigurationData(acsConfigurationData).Execute()

Creates a new subscription resource

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/ACSParameterProvision"
)

func main() {
    afId := "afId_example" // string | Identifier of the AF
    acsConfigurationData := *openapiclient.NewAcsConfigurationData(*openapiclient.NewAcsInfo(), "SuppFeat_example") // AcsConfigurationData | new subscription creation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ACSConfigurationSubscriptionsApi.CreateAnSubscription(context.Background(), afId).AcsConfigurationData(acsConfigurationData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ACSConfigurationSubscriptionsApi.CreateAnSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAnSubscription`: AcsConfigurationData
    fmt.Fprintf(os.Stdout, "Response from `ACSConfigurationSubscriptionsApi.CreateAnSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**afId** | **string** | Identifier of the AF | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAnSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acsConfigurationData** | [**AcsConfigurationData**](AcsConfigurationData.md) | new subscription creation | 

### Return type

[**AcsConfigurationData**](AcsConfigurationData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadAllSubscriptions

> []AcsConfigurationData ReadAllSubscriptions(ctx, afId).Execute()

read all of the active subscriptions for the AF

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/ACSParameterProvision"
)

func main() {
    afId := "afId_example" // string | Identifier of the AF

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ACSConfigurationSubscriptionsApi.ReadAllSubscriptions(context.Background(), afId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ACSConfigurationSubscriptionsApi.ReadAllSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllSubscriptions`: []AcsConfigurationData
    fmt.Fprintf(os.Stdout, "Response from `ACSConfigurationSubscriptionsApi.ReadAllSubscriptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**afId** | **string** | Identifier of the AF | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAllSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]AcsConfigurationData**](AcsConfigurationData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

