# \SubscriptionsCollectionForMBSContextsApi

All URIs are relative to *https://example.com/nmbsmf-mbssession/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContextStatusSubscribe**](SubscriptionsCollectionForMBSContextsApi.md#ContextStatusSubscribe) | **Post** /mbs-sessions/contexts/subscriptions | ContextStatusSubscribe creating a subscription



## ContextStatusSubscribe

> ContextStatusSubscribeRspData ContextStatusSubscribe(ctx).ContextStatusSubscribeReqData(contextStatusSubscribeReqData).Execute()

ContextStatusSubscribe creating a subscription

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
    contextStatusSubscribeReqData := *openapiclient.NewContextStatusSubscribeReqData(*openapiclient.NewContextStatusSubscription("NfcInstanceId_example", *openapiclient.NewMbsSessionId(), []openapiclient.ContextStatusEvent{*openapiclient.NewContextStatusEvent(*openapiclient.NewContextStatusEventType())}, "NotifyUri_example")) // ContextStatusSubscribeReqData | Data within the ContextStatusSubscribe Request

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionsCollectionForMBSContextsApi.ContextStatusSubscribe(context.Background()).ContextStatusSubscribeReqData(contextStatusSubscribeReqData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsCollectionForMBSContextsApi.ContextStatusSubscribe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContextStatusSubscribe`: ContextStatusSubscribeRspData
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionsCollectionForMBSContextsApi.ContextStatusSubscribe`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContextStatusSubscribeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contextStatusSubscribeReqData** | [**ContextStatusSubscribeReqData**](ContextStatusSubscribeReqData.md) | Data within the ContextStatusSubscribe Request | 

### Return type

[**ContextStatusSubscribeRspData**](ContextStatusSubscribeRspData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

