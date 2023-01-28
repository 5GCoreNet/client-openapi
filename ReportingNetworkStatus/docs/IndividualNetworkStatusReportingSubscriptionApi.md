# \IndividualNetworkStatusReportingSubscriptionApi

All URIs are relative to *https://example.com/3gpp-net-stat-report/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIndNwStatusReportSubscription**](IndividualNetworkStatusReportingSubscriptionApi.md#DeleteIndNwStatusReportSubscription) | **Delete** /{scsAsId}/subscriptions/{subscriptionId} | Delete an existing continuous network status reporting subscription resource.
[**FetchIndNwStatusReportSubscription**](IndividualNetworkStatusReportingSubscriptionApi.md#FetchIndNwStatusReportSubscription) | **Get** /{scsAsId}/subscriptions/{subscriptionId} | Read an active network status reporting subscription resource.
[**ModifyIndNwStatusReportSubscription**](IndividualNetworkStatusReportingSubscriptionApi.md#ModifyIndNwStatusReportSubscription) | **Patch** /{scsAsId}/subscriptions/{subscriptionId} | Modify an existing Individual Network Status Reporting Subscription resource.
[**UpdateIndNwStatusReportSubscription**](IndividualNetworkStatusReportingSubscriptionApi.md#UpdateIndNwStatusReportSubscription) | **Put** /{scsAsId}/subscriptions/{subscriptionId} | Modify an existing subscription resource to update a subscription.



## DeleteIndNwStatusReportSubscription

> DeleteIndNwStatusReportSubscription(ctx, scsAsId, subscriptionId).Execute()

Delete an existing continuous network status reporting subscription resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/ReportingNetworkStatus"
)

func main() {
    scsAsId := "scsAsId_example" // string | Identifier of the SCS/AS
    subscriptionId := "subscriptionId_example" // string | Identifier of the subscription resource of type string

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualNetworkStatusReportingSubscriptionApi.DeleteIndNwStatusReportSubscription(context.Background(), scsAsId, subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualNetworkStatusReportingSubscriptionApi.DeleteIndNwStatusReportSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of the SCS/AS | 
**subscriptionId** | **string** | Identifier of the subscription resource of type string | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIndNwStatusReportSubscriptionRequest struct via the builder pattern


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


## FetchIndNwStatusReportSubscription

> NetworkStatusReportingSubscription FetchIndNwStatusReportSubscription(ctx, scsAsId, subscriptionId).Execute()

Read an active network status reporting subscription resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/ReportingNetworkStatus"
)

func main() {
    scsAsId := "scsAsId_example" // string | Identifier of the SCS/AS
    subscriptionId := "subscriptionId_example" // string | Identifier of the subscription resource of type string

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualNetworkStatusReportingSubscriptionApi.FetchIndNwStatusReportSubscription(context.Background(), scsAsId, subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualNetworkStatusReportingSubscriptionApi.FetchIndNwStatusReportSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchIndNwStatusReportSubscription`: NetworkStatusReportingSubscription
    fmt.Fprintf(os.Stdout, "Response from `IndividualNetworkStatusReportingSubscriptionApi.FetchIndNwStatusReportSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of the SCS/AS | 
**subscriptionId** | **string** | Identifier of the subscription resource of type string | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchIndNwStatusReportSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**NetworkStatusReportingSubscription**](NetworkStatusReportingSubscription.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyIndNwStatusReportSubscription

> NetworkStatusReportingSubscription ModifyIndNwStatusReportSubscription(ctx, scsAsId, subscriptionId).NetStatusRepSubsPatch(netStatusRepSubsPatch).Execute()

Modify an existing Individual Network Status Reporting Subscription resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/ReportingNetworkStatus"
)

func main() {
    scsAsId := "scsAsId_example" // string | Identifier of the SCS/AS
    subscriptionId := "subscriptionId_example" // string | Identifier of the subscription resource of type string
    netStatusRepSubsPatch := *openapiclient.NewNetStatusRepSubsPatch() // NetStatusRepSubsPatch | Contains the parameters to modify an existing Individual Network Status Reporting Subscription resource.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualNetworkStatusReportingSubscriptionApi.ModifyIndNwStatusReportSubscription(context.Background(), scsAsId, subscriptionId).NetStatusRepSubsPatch(netStatusRepSubsPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualNetworkStatusReportingSubscriptionApi.ModifyIndNwStatusReportSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ModifyIndNwStatusReportSubscription`: NetworkStatusReportingSubscription
    fmt.Fprintf(os.Stdout, "Response from `IndividualNetworkStatusReportingSubscriptionApi.ModifyIndNwStatusReportSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of the SCS/AS | 
**subscriptionId** | **string** | Identifier of the subscription resource of type string | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyIndNwStatusReportSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **netStatusRepSubsPatch** | [**NetStatusRepSubsPatch**](NetStatusRepSubsPatch.md) | Contains the parameters to modify an existing Individual Network Status Reporting Subscription resource. | 

### Return type

[**NetworkStatusReportingSubscription**](NetworkStatusReportingSubscription.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIndNwStatusReportSubscription

> NetworkStatusReportingSubscription UpdateIndNwStatusReportSubscription(ctx, scsAsId, subscriptionId).NetworkStatusReportingSubscription(networkStatusReportingSubscription).Execute()

Modify an existing subscription resource to update a subscription.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/ReportingNetworkStatus"
)

func main() {
    scsAsId := "scsAsId_example" // string | Identifier of the SCS/AS
    subscriptionId := "subscriptionId_example" // string | Identifier of the subscription resource of type string
    networkStatusReportingSubscription := *openapiclient.NewNetworkStatusReportingSubscription("NotificationDestination_example", *openapiclient.NewLocationArea()) // NetworkStatusReportingSubscription | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualNetworkStatusReportingSubscriptionApi.UpdateIndNwStatusReportSubscription(context.Background(), scsAsId, subscriptionId).NetworkStatusReportingSubscription(networkStatusReportingSubscription).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualNetworkStatusReportingSubscriptionApi.UpdateIndNwStatusReportSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIndNwStatusReportSubscription`: NetworkStatusReportingSubscription
    fmt.Fprintf(os.Stdout, "Response from `IndividualNetworkStatusReportingSubscriptionApi.UpdateIndNwStatusReportSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of the SCS/AS | 
**subscriptionId** | **string** | Identifier of the subscription resource of type string | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIndNwStatusReportSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **networkStatusReportingSubscription** | [**NetworkStatusReportingSubscription**](NetworkStatusReportingSubscription.md) |  | 

### Return type

[**NetworkStatusReportingSubscription**](NetworkStatusReportingSubscription.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

