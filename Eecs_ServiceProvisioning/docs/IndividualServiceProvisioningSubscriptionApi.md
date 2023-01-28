# \IndividualServiceProvisioningSubscriptionApi

All URIs are relative to *https://example.com/eecs-serviceprovisioning/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubscriptionsSubscriptionIdDelete**](IndividualServiceProvisioningSubscriptionApi.md#SubscriptionsSubscriptionIdDelete) | **Delete** /subscriptions/{subscriptionId} | 
[**SubscriptionsSubscriptionIdPatch**](IndividualServiceProvisioningSubscriptionApi.md#SubscriptionsSubscriptionIdPatch) | **Patch** /subscriptions/{subscriptionId} | 
[**SubscriptionsSubscriptionIdPut**](IndividualServiceProvisioningSubscriptionApi.md#SubscriptionsSubscriptionIdPut) | **Put** /subscriptions/{subscriptionId} | 



## SubscriptionsSubscriptionIdDelete

> SubscriptionsSubscriptionIdDelete(ctx, subscriptionId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Eecs_ServiceProvisioning"
)

func main() {
    subscriptionId := "subscriptionId_example" // string | Identifies an individual service provisioning subscription

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualServiceProvisioningSubscriptionApi.SubscriptionsSubscriptionIdDelete(context.Background(), subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualServiceProvisioningSubscriptionApi.SubscriptionsSubscriptionIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Identifies an individual service provisioning subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsSubscriptionIdDeleteRequest struct via the builder pattern


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


## SubscriptionsSubscriptionIdPatch

> ECSServProvSubscription SubscriptionsSubscriptionIdPatch(ctx, subscriptionId).ECSServProvSubscriptionPatch(eCSServProvSubscriptionPatch).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Eecs_ServiceProvisioning"
)

func main() {
    subscriptionId := "subscriptionId_example" // string | Identifies an individual service provisioning subscription
    eCSServProvSubscriptionPatch := *openapiclient.NewECSServProvSubscriptionPatch() // ECSServProvSubscriptionPatch | Parameters to replace the existing subscription

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualServiceProvisioningSubscriptionApi.SubscriptionsSubscriptionIdPatch(context.Background(), subscriptionId).ECSServProvSubscriptionPatch(eCSServProvSubscriptionPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualServiceProvisioningSubscriptionApi.SubscriptionsSubscriptionIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubscriptionsSubscriptionIdPatch`: ECSServProvSubscription
    fmt.Fprintf(os.Stdout, "Response from `IndividualServiceProvisioningSubscriptionApi.SubscriptionsSubscriptionIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Identifies an individual service provisioning subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsSubscriptionIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eCSServProvSubscriptionPatch** | [**ECSServProvSubscriptionPatch**](ECSServProvSubscriptionPatch.md) | Parameters to replace the existing subscription | 

### Return type

[**ECSServProvSubscription**](ECSServProvSubscription.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsSubscriptionIdPut

> ECSServProvSubscription SubscriptionsSubscriptionIdPut(ctx, subscriptionId).ECSServProvSubscription(eCSServProvSubscription).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Eecs_ServiceProvisioning"
)

func main() {
    subscriptionId := "subscriptionId_example" // string | Identifies an individual service provisioning subscription
    eCSServProvSubscription := *openapiclient.NewECSServProvSubscription("EecId_example") // ECSServProvSubscription | Parameters to replace the existing subscription

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualServiceProvisioningSubscriptionApi.SubscriptionsSubscriptionIdPut(context.Background(), subscriptionId).ECSServProvSubscription(eCSServProvSubscription).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualServiceProvisioningSubscriptionApi.SubscriptionsSubscriptionIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubscriptionsSubscriptionIdPut`: ECSServProvSubscription
    fmt.Fprintf(os.Stdout, "Response from `IndividualServiceProvisioningSubscriptionApi.SubscriptionsSubscriptionIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Identifies an individual service provisioning subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsSubscriptionIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eCSServProvSubscription** | [**ECSServProvSubscription**](ECSServProvSubscription.md) | Parameters to replace the existing subscription | 

### Return type

[**ECSServProvSubscription**](ECSServProvSubscription.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

