# \NSSAIACKDocumentApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QueryNssaiAck**](NSSAIACKDocumentApi.md#QueryNssaiAck) | **Get** /subscription-data/{ueId}/ue-update-confirmation-data/subscribed-snssais | Retrieves the UPU acknowledgement information of a UE



## QueryNssaiAck

> NssaiAckData QueryNssaiAck(ctx, ueId).SupportedFeatures(supportedFeatures).Execute()

Retrieves the UPU acknowledgement information of a UE

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Subscription_Data"
)

func main() {
    ueId := "ueId_example" // string | UE id
    supportedFeatures := "supportedFeatures_example" // string | Supported Features (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NSSAIACKDocumentApi.QueryNssaiAck(context.Background(), ueId).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NSSAIACKDocumentApi.QueryNssaiAck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryNssaiAck`: NssaiAckData
    fmt.Fprintf(os.Stdout, "Response from `NSSAIACKDocumentApi.QueryNssaiAck`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | UE id | 

### Other Parameters

Other parameters are passed through a pointer to a apiQueryNssaiAckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **supportedFeatures** | **string** | Supported Features | 

### Return type

[**NssaiAckData**](NssaiAckData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

