# \SubscriberDataDocumentApi

All URIs are relative to *https://example.com/nhss-gba-sdm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSubscriberData**](SubscriberDataDocumentApi.md#GetSubscriberData) | **Get** /{ueId}/subscriber-data | Retrieve the subscriber data of a user



## GetSubscriberData

> GbaSubscriberData GetSubscriberData(ctx, ueId).SupportedFeatures(supportedFeatures).Execute()

Retrieve the subscriber data of a user

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nhss_gbaSDM"
)

func main() {
    ueId := *openapiclient.NewUeId() // UeId | UE Identity
    supportedFeatures := "supportedFeatures_example" // string | Supported Features (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriberDataDocumentApi.GetSubscriberData(context.Background(), ueId).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriberDataDocumentApi.GetSubscriberData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSubscriberData`: GbaSubscriberData
    fmt.Fprintf(os.Stdout, "Response from `SubscriberDataDocumentApi.GetSubscriberData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | [**UeId**](.md) | UE Identity | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubscriberDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **supportedFeatures** | **string** | Supported Features | 

### Return type

[**GbaSubscriberData**](GbaSubscriberData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

