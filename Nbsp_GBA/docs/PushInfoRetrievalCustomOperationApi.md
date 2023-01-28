# \PushInfoRetrievalCustomOperationApi

All URIs are relative to *https://example.com/nbsp-gba/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PushInfoRetrieval**](PushInfoRetrievalCustomOperationApi.md#PushInfoRetrieval) | **Post** /push-info-retrieval | Retrieve Push Info from GBA BSF from Push-NAF



## PushInfoRetrieval

> PushInfoResponse PushInfoRetrieval(ctx).PushInfoRequest(pushInfoRequest).Execute()

Retrieve Push Info from GBA BSF from Push-NAF

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/5GCoreNet/client-openapi/Nbsp_GBA"
)

func main() {
    pushInfoRequest := *openapiclient.NewPushInfoRequest("UeId_example", *openapiclient.NewUeIdType(), "UiccAppLabel_example", *openapiclient.NewNafId("NafFqdn_example", "UaSecProtId_example"), "PtId_example", *openapiclient.NewUiccOrMe(), time.Now()) // PushInfoRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PushInfoRetrievalCustomOperationApi.PushInfoRetrieval(context.Background()).PushInfoRequest(pushInfoRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PushInfoRetrievalCustomOperationApi.PushInfoRetrieval``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PushInfoRetrieval`: PushInfoResponse
    fmt.Fprintf(os.Stdout, "Response from `PushInfoRetrievalCustomOperationApi.PushInfoRetrieval`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPushInfoRetrievalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pushInfoRequest** | [**PushInfoRequest**](PushInfoRequest.md) |  | 

### Return type

[**PushInfoResponse**](PushInfoResponse.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

