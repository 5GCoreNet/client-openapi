# \ASMessageDeliveryApi

All URIs are relative to *https://example.com/msgs-msgdelivery/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeliverAsMessagePost**](ASMessageDeliveryApi.md#DeliverAsMessagePost) | **Post** /deliver-as-message | AS deliver message to MSGin5G Server



## DeliverAsMessagePost

> MessageDeliveryAck DeliverAsMessagePost(ctx).ASMessageDelivery(aSMessageDelivery).Execute()

AS deliver message to MSGin5G Server

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
    aSMessageDelivery := *openapiclient.NewASMessageDelivery(*openapiclient.NewAddress(*openapiclient.NewAddressType(), "Addr_example"), *openapiclient.NewAddress(*openapiclient.NewAddressType(), "Addr_example"), "MsgId_example", false) // ASMessageDelivery | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ASMessageDeliveryApi.DeliverAsMessagePost(context.Background()).ASMessageDelivery(aSMessageDelivery).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ASMessageDeliveryApi.DeliverAsMessagePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeliverAsMessagePost`: MessageDeliveryAck
    fmt.Fprintf(os.Stdout, "Response from `ASMessageDeliveryApi.DeliverAsMessagePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeliverAsMessagePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aSMessageDelivery** | [**ASMessageDelivery**](ASMessageDelivery.md) |  | 

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

