# \L3GMessageDeliveryApi

All URIs are relative to *https://example.com/msgg-l3gdelivery/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeliverMessagePost**](L3GMessageDeliveryApi.md#DeliverMessagePost) | **Post** /deliver-message | deliver message to Legacy 3GPP Message Gateway



## DeliverMessagePost

> DeliverMessagePost(ctx).L3gMessageDelivery(l3gMessageDelivery).Execute()

deliver message to Legacy 3GPP Message Gateway

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
    l3gMessageDelivery := *openapiclient.NewL3gMessageDelivery(*openapiclient.NewAddress(*openapiclient.NewAddressType(), "Addr_example"), *openapiclient.NewAddress(*openapiclient.NewAddressType(), "Addr_example"), "MsgId_example") // L3gMessageDelivery | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.L3GMessageDeliveryApi.DeliverMessagePost(context.Background()).L3gMessageDelivery(l3gMessageDelivery).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `L3GMessageDeliveryApi.DeliverMessagePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeliverMessagePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **l3gMessageDelivery** | [**L3gMessageDelivery**](L3gMessageDelivery.md) |  | 

### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

