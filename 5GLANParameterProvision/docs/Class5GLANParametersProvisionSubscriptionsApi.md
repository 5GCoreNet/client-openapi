# \Class5GLANParametersProvisionSubscriptionsApi

All URIs are relative to *https://example.com/3gpp-5glan-pp/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAnSubscription**](Class5GLANParametersProvisionSubscriptionsApi.md#CreateAnSubscription) | **Post** /{afId}/subscriptions | Creates a new subscription resource
[**RealAllSubscriptions**](Class5GLANParametersProvisionSubscriptionsApi.md#RealAllSubscriptions) | **Get** /{afId}/subscriptions | read all of the active subscriptions for the AF



## CreateAnSubscription

> Model5GLanParametersProvision CreateAnSubscription(ctx, afId).Model5GLanParametersProvision(model5GLanParametersProvision).Execute()

Creates a new subscription resource

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
    afId := "afId_example" // string | Identifier of the AF
    model5GLanParametersProvision := *openapiclient.NewModel5GLanParametersProvision(*openapiclient.NewModel5GLanParameters("ExterGroupId_example", map[string]string{"key": "Inner_example"}, "Dnn_example", *openapiclient.NewSnssai(int32(123)), *openapiclient.NewPduSessionType(), map[string]AppDescriptor{"key": *openapiclient.NewAppDescriptor("OsId_example", map[string]string{"key": "Inner_example"})}), "SuppFeat_example") // Model5GLanParametersProvision | new subscription creation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Class5GLANParametersProvisionSubscriptionsApi.CreateAnSubscription(context.Background(), afId).Model5GLanParametersProvision(model5GLanParametersProvision).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Class5GLANParametersProvisionSubscriptionsApi.CreateAnSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAnSubscription`: Model5GLanParametersProvision
    fmt.Fprintf(os.Stdout, "Response from `Class5GLANParametersProvisionSubscriptionsApi.CreateAnSubscription`: %v\n", resp)
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

 **model5GLanParametersProvision** | [**Model5GLanParametersProvision**](Model5GLanParametersProvision.md) | new subscription creation | 

### Return type

[**Model5GLanParametersProvision**](Model5GLanParametersProvision.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RealAllSubscriptions

> []Model5GLanParametersProvision RealAllSubscriptions(ctx, afId).Execute()

read all of the active subscriptions for the AF

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
    afId := "afId_example" // string | Identifier of the AF

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Class5GLANParametersProvisionSubscriptionsApi.RealAllSubscriptions(context.Background(), afId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Class5GLANParametersProvisionSubscriptionsApi.RealAllSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RealAllSubscriptions`: []Model5GLanParametersProvision
    fmt.Fprintf(os.Stdout, "Response from `Class5GLANParametersProvisionSubscriptionsApi.RealAllSubscriptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**afId** | **string** | Identifier of the AF | 

### Other Parameters

Other parameters are passed through a pointer to a apiRealAllSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Model5GLanParametersProvision**](Model5GLanParametersProvision.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

