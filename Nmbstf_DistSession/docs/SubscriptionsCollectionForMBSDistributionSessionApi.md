# \SubscriptionsCollectionForMBSDistributionSessionApi

All URIs are relative to *https://example.com/nmbstf-distsession/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StatusSubscribe**](SubscriptionsCollectionForMBSDistributionSessionApi.md#StatusSubscribe) | **Post** /dist-sessions/{distSessionRef}/subscriptions | StatusSubscribe creating a subscription



## StatusSubscribe

> StatusSubscribeRspData StatusSubscribe(ctx, distSessionRef).StatusSubscribeReqData(statusSubscribeReqData).Execute()

StatusSubscribe creating a subscription

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
    distSessionRef := "distSessionRef_example" // string | Unique ID of the MBS distribution session
    statusSubscribeReqData := *openapiclient.NewStatusSubscribeReqData(*openapiclient.NewDistSessionSubscription([]openapiclient.DistSessionEventType{*openapiclient.NewDistSessionEventType()}, "NotifyUri_example")) // StatusSubscribeReqData | Data within the StatusSubscribe Request

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionsCollectionForMBSDistributionSessionApi.StatusSubscribe(context.Background(), distSessionRef).StatusSubscribeReqData(statusSubscribeReqData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsCollectionForMBSDistributionSessionApi.StatusSubscribe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StatusSubscribe`: StatusSubscribeRspData
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionsCollectionForMBSDistributionSessionApi.StatusSubscribe`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**distSessionRef** | **string** | Unique ID of the MBS distribution session | 

### Other Parameters

Other parameters are passed through a pointer to a apiStatusSubscribeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **statusSubscribeReqData** | [**StatusSubscribeReqData**](StatusSubscribeReqData.md) | Data within the StatusSubscribe Request | 

### Return type

[**StatusSubscribeRspData**](StatusSubscribeRspData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

