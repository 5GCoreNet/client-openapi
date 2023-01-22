# \IndividualPDUSessionApi

All URIs are relative to *https://example.com/nsmf-nidd/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Deliver**](IndividualPDUSessionApi.md#Deliver) | **Post** /pdu-sessions/{pduSessionRef}/deliver | Delivery Service Operation



## Deliver

> Deliver(ctx, pduSessionRef).JsonData(jsonData).BinaryMtData(binaryMtData).Execute()

Delivery Service Operation

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
    pduSessionRef := "pduSessionRef_example" // string | PDU session reference
    jsonData := *openapiclient.NewDeliverReqData(*openapiclient.NewRefToBinaryData("ContentId_example")) // DeliverReqData |  (optional)
    binaryMtData := os.NewFile(1234, "some_file") // *os.File |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualPDUSessionApi.Deliver(context.Background(), pduSessionRef).JsonData(jsonData).BinaryMtData(binaryMtData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualPDUSessionApi.Deliver``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pduSessionRef** | **string** | PDU session reference | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliverRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jsonData** | [**DeliverReqData**](DeliverReqData.md) |  | 
 **binaryMtData** | ***os.File** |  | 

### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: multipart/related
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

