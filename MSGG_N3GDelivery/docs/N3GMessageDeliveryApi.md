# \N3GMessageDeliveryApi

All URIs are relative to *https://example.com/msgg-n3gdelivery/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeliverMessagePost**](N3GMessageDeliveryApi.md#DeliverMessagePost) | **Post** /deliver-message | deliver message to NON-3GPP Message Gateway



## DeliverMessagePost

> DeliverMessagePost(ctx).N3gMessageDelivery(n3gMessageDelivery).Execute()

deliver message to NON-3GPP Message Gateway

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
    n3gMessageDelivery := *openapiclient.NewN3gMessageDelivery(*openapiclient.NewAddress(*openapiclient.NewAddressType(), "Addr_example"), *openapiclient.NewAddress(*openapiclient.NewAddressType(), "Addr_example"), "MsgId_example") // N3gMessageDelivery | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.N3GMessageDeliveryApi.DeliverMessagePost(context.Background()).N3gMessageDelivery(n3gMessageDelivery).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `N3GMessageDeliveryApi.DeliverMessagePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeliverMessagePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **n3gMessageDelivery** | [**N3gMessageDelivery**](N3gMessageDelivery.md) |  | 

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

