# \SubscriptionsCollectionForMBSSessionsApi

All URIs are relative to *https://example.com/nmbsmf-mbssession/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StatusSubscribe**](SubscriptionsCollectionForMBSSessionsApi.md#StatusSubscribe) | **Post** /mbs-sessions/subscriptions | StatusSubscribe creating a subscription



## StatusSubscribe

> StatusSubscribeRspData StatusSubscribe(ctx).StatusSubscribeReqData(statusSubscribeReqData).Execute()

StatusSubscribe creating a subscription

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nmbsmf_MBSSession"
)

func main() {
    statusSubscribeReqData := *openapiclient.NewStatusSubscribeReqData(*openapiclient.NewMbsSessionSubscription([]openapiclient.MbsSessionEvent{*openapiclient.NewMbsSessionEvent(*openapiclient.NewMbsSessionEventType())}, "NotifyUri_example")) // StatusSubscribeReqData | Data within the StatusSubscribe Request

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionsCollectionForMBSSessionsApi.StatusSubscribe(context.Background()).StatusSubscribeReqData(statusSubscribeReqData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsCollectionForMBSSessionsApi.StatusSubscribe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StatusSubscribe`: StatusSubscribeRspData
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionsCollectionForMBSSessionsApi.StatusSubscribe`: %v\n", resp)
}
```

### Path Parameters



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

