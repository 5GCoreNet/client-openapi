# \PDUSessionsCollectionApi

All URIs are relative to *https://example.com/nsmf-pdusession/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostPduSessions**](PDUSessionsCollectionApi.md#PostPduSessions) | **Post** /pdu-sessions | Create



## PostPduSessions

> PduSessionCreatedData PostPduSessions(ctx).PduSessionCreateData(pduSessionCreateData).Execute()

Create

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
    pduSessionCreateData := openapiclient.PduSessionCreateData{Interface{}: new(interface{})} // PduSessionCreateData | representation of the PDU session to be created in the H-SMF or SMF

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PDUSessionsCollectionApi.PostPduSessions(context.Background()).PduSessionCreateData(pduSessionCreateData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PDUSessionsCollectionApi.PostPduSessions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostPduSessions`: PduSessionCreatedData
    fmt.Fprintf(os.Stdout, "Response from `PDUSessionsCollectionApi.PostPduSessions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostPduSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pduSessionCreateData** | [**PduSessionCreateData**](PduSessionCreateData.md) | representation of the PDU session to be created in the H-SMF or SMF | 

### Return type

[**PduSessionCreatedData**](PduSessionCreatedData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json, multipart/related
- **Accept**: application/json, multipart/related, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

