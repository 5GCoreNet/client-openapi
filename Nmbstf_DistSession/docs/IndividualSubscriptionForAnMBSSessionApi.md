# \IndividualSubscriptionForAnMBSSessionApi

All URIs are relative to *https://example.com/nmbstf-distsession/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StatusUnSubscribe**](IndividualSubscriptionForAnMBSSessionApi.md#StatusUnSubscribe) | **Delete** /dist-sessions/{distSessionRef}/subscriptions/{subscriptionId} | StatusUnSubscribe to unsubscribe from the Status Subscription



## StatusUnSubscribe

> StatusUnSubscribe(ctx, subscriptionId, distSessionRef).Execute()

StatusUnSubscribe to unsubscribe from the Status Subscription

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
    subscriptionId := "subscriptionId_example" // string | Unique ID of the subscription
    distSessionRef := "distSessionRef_example" // string | Unique ID of the MBS distribution session

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualSubscriptionForAnMBSSessionApi.StatusUnSubscribe(context.Background(), subscriptionId, distSessionRef).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualSubscriptionForAnMBSSessionApi.StatusUnSubscribe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Unique ID of the subscription | 
**distSessionRef** | **string** | Unique ID of the MBS distribution session | 

### Other Parameters

Other parameters are passed through a pointer to a apiStatusUnSubscribeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

