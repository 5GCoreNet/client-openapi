# \IndividualACSConfigurationSubscriptionApi

All URIs are relative to *https://example.com/3gpp-acs-pp/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAnSubscription**](IndividualACSConfigurationSubscriptionApi.md#DeleteAnSubscription) | **Delete** /{afId}/subscriptions/{subscriptionId} | Deletes an already existing subscription
[**FullyUpdateAnSubscription**](IndividualACSConfigurationSubscriptionApi.md#FullyUpdateAnSubscription) | **Put** /{afId}/subscriptions/{subscriptionId} | Fully updates/replaces an existing subscription resource
[**PartialUpdateAnSubscription**](IndividualACSConfigurationSubscriptionApi.md#PartialUpdateAnSubscription) | **Patch** /{afId}/subscriptions/{subscriptionId} | Partial modifies an existing subscription resource.
[**ReadAnSubscription**](IndividualACSConfigurationSubscriptionApi.md#ReadAnSubscription) | **Get** /{afId}/subscriptions/{subscriptionId} | read an active subscription for the AF and the subscription Id



## DeleteAnSubscription

> DeleteAnSubscription(ctx, afId, subscriptionId).Execute()

Deletes an already existing subscription

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
    subscriptionId := "subscriptionId_example" // string | Identifier of the subscription resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualACSConfigurationSubscriptionApi.DeleteAnSubscription(context.Background(), afId, subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualACSConfigurationSubscriptionApi.DeleteAnSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**afId** | **string** | Identifier of the AF | 
**subscriptionId** | **string** | Identifier of the subscription resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAnSubscriptionRequest struct via the builder pattern


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


## FullyUpdateAnSubscription

> AcsConfigurationData FullyUpdateAnSubscription(ctx, afId, subscriptionId).AcsConfigurationData(acsConfigurationData).Execute()

Fully updates/replaces an existing subscription resource

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
    subscriptionId := "subscriptionId_example" // string | Identifier of the subscription resource
    acsConfigurationData := *openapiclient.NewAcsConfigurationData(*openapiclient.NewAcsInfo(), "SuppFeat_example") // AcsConfigurationData | Parameters to update/replace the existing subscription

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualACSConfigurationSubscriptionApi.FullyUpdateAnSubscription(context.Background(), afId, subscriptionId).AcsConfigurationData(acsConfigurationData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualACSConfigurationSubscriptionApi.FullyUpdateAnSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FullyUpdateAnSubscription`: AcsConfigurationData
    fmt.Fprintf(os.Stdout, "Response from `IndividualACSConfigurationSubscriptionApi.FullyUpdateAnSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**afId** | **string** | Identifier of the AF | 
**subscriptionId** | **string** | Identifier of the subscription resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiFullyUpdateAnSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **acsConfigurationData** | [**AcsConfigurationData**](AcsConfigurationData.md) | Parameters to update/replace the existing subscription | 

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


## PartialUpdateAnSubscription

> AcsConfigurationData PartialUpdateAnSubscription(ctx, afId, subscriptionId).AcsConfigurationDataPatch(acsConfigurationDataPatch).Execute()

Partial modifies an existing subscription resource.

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
    subscriptionId := "subscriptionId_example" // string | Identifier of the subscription resource
    acsConfigurationDataPatch := *openapiclient.NewAcsConfigurationDataPatch() // AcsConfigurationDataPatch | Parameters to modify the existing subscription.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualACSConfigurationSubscriptionApi.PartialUpdateAnSubscription(context.Background(), afId, subscriptionId).AcsConfigurationDataPatch(acsConfigurationDataPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualACSConfigurationSubscriptionApi.PartialUpdateAnSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PartialUpdateAnSubscription`: AcsConfigurationData
    fmt.Fprintf(os.Stdout, "Response from `IndividualACSConfigurationSubscriptionApi.PartialUpdateAnSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**afId** | **string** | Identifier of the AF | 
**subscriptionId** | **string** | Identifier of the subscription resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartialUpdateAnSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **acsConfigurationDataPatch** | [**AcsConfigurationDataPatch**](AcsConfigurationDataPatch.md) | Parameters to modify the existing subscription. | 

### Return type

[**AcsConfigurationData**](AcsConfigurationData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadAnSubscription

> AcsConfigurationData ReadAnSubscription(ctx, afId, subscriptionId).Execute()

read an active subscription for the AF and the subscription Id

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
    subscriptionId := "subscriptionId_example" // string | Identifier of the subscription resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualACSConfigurationSubscriptionApi.ReadAnSubscription(context.Background(), afId, subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualACSConfigurationSubscriptionApi.ReadAnSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAnSubscription`: AcsConfigurationData
    fmt.Fprintf(os.Stdout, "Response from `IndividualACSConfigurationSubscriptionApi.ReadAnSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**afId** | **string** | Identifier of the AF | 
**subscriptionId** | **string** | Identifier of the subscription resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAnSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AcsConfigurationData**](AcsConfigurationData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

