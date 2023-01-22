# \IndividualMBSSessionSubscriptionApi

All URIs are relative to *https://example.com/3gpp-mbs-session/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIndMBSSessionsSubsc**](IndividualMBSSessionSubscriptionApi.md#DeleteIndMBSSessionsSubsc) | **Delete** /mbs-sessions/subscriptions/{subscriptionId} | Request the deletion of an existing Individual MBS Session subscription resource.
[**ReadIndMBSSessionsSubsc**](IndividualMBSSessionSubscriptionApi.md#ReadIndMBSSessionsSubsc) | **Get** /mbs-sessions/subscriptions/{subscriptionId} | Retrieve an existing Individual MBS Session Subscription resource.



## DeleteIndMBSSessionsSubsc

> DeleteIndMBSSessionsSubsc(ctx, subscriptionId).Execute()

Request the deletion of an existing Individual MBS Session subscription resource.

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
    subscriptionId := "subscriptionId_example" // string | Identifier of the Individual MBS Session Subscription resource.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualMBSSessionSubscriptionApi.DeleteIndMBSSessionsSubsc(context.Background(), subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualMBSSessionSubscriptionApi.DeleteIndMBSSessionsSubsc``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Identifier of the Individual MBS Session Subscription resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIndMBSSessionsSubscRequest struct via the builder pattern


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


## ReadIndMBSSessionsSubsc

> MbsSessionSubsc ReadIndMBSSessionsSubsc(ctx, subscriptionId).Execute()

Retrieve an existing Individual MBS Session Subscription resource.

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
    subscriptionId := "subscriptionId_example" // string | Identifier of the Individual MBS Session Subscription resource.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualMBSSessionSubscriptionApi.ReadIndMBSSessionsSubsc(context.Background(), subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualMBSSessionSubscriptionApi.ReadIndMBSSessionsSubsc``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadIndMBSSessionsSubsc`: MbsSessionSubsc
    fmt.Fprintf(os.Stdout, "Response from `IndividualMBSSessionSubscriptionApi.ReadIndMBSSessionsSubsc`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Identifier of the Individual MBS Session Subscription resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadIndMBSSessionsSubscRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MbsSessionSubsc**](MbsSessionSubsc.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

