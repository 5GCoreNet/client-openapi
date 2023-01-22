# \UEMessageDeliveryApi

All URIs are relative to *https://example.com/msgs-msgdelivery/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeliverUeMessagePost**](UEMessageDeliveryApi.md#DeliverUeMessagePost) | **Post** /deliver-ue-message | UE deliver message to MSGin5G Server



## DeliverUeMessagePost

> MessageDeliveryAck DeliverUeMessagePost(ctx).UEMessageDelivery(uEMessageDelivery).Execute()

UE deliver message to MSGin5G Server

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
    uEMessageDelivery := *openapiclient.NewUEMessageDelivery(*openapiclient.NewAddress(*openapiclient.NewAddressType(), "Addr_example"), *openapiclient.NewAddress(*openapiclient.NewAddressType(), "Addr_example"), "MsgId_example", false) // UEMessageDelivery | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UEMessageDeliveryApi.DeliverUeMessagePost(context.Background()).UEMessageDelivery(uEMessageDelivery).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UEMessageDeliveryApi.DeliverUeMessagePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliverUeMessagePost`: MessageDeliveryAck
    fmt.Fprintf(os.Stdout, "Response from `UEMessageDeliveryApi.DeliverUeMessagePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeliverUeMessagePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uEMessageDelivery** | [**UEMessageDelivery**](UEMessageDelivery.md) |  | 

### Return type

[**MessageDeliveryAck**](MessageDeliveryAck.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

