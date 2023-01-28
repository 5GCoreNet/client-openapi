# \MMESGSNDeregistrationApi

All URIs are relative to *https://example.com/nhss-uecm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeregisterSN**](MMESGSNDeregistrationApi.md#DeregisterSN) | **Post** /deregister-sn | MME/SGSN Deregistration



## DeregisterSN

> DeregisterSN(ctx).DeregistrationRequest(deregistrationRequest).Execute()

MME/SGSN Deregistration

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nhss_UECM"
)

func main() {
    deregistrationRequest := *openapiclient.NewDeregistrationRequest("Imsi_example", *openapiclient.NewDeregistrationReason()) // DeregistrationRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MMESGSNDeregistrationApi.DeregisterSN(context.Background()).DeregistrationRequest(deregistrationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MMESGSNDeregistrationApi.DeregisterSN``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeregisterSNRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deregistrationRequest** | [**DeregistrationRequest**](DeregistrationRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

