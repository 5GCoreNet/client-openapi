# \PriorityInfoRetrievalApi

All URIs are relative to *https://example.com/nhss-ims-sdm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPriorityInfo**](PriorityInfoRetrievalApi.md#GetPriorityInfo) | **Get** /{imsUeId}/ims-data/profile-data/priority-levels | Retrieve the service priority levels associated to the user



## GetPriorityInfo

> PriorityLevels GetPriorityInfo(ctx, imsUeId).SupportedFeatures(supportedFeatures).Execute()

Retrieve the service priority levels associated to the user

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
    imsUeId := "imsUeId_example" // string | IMS Identity
    supportedFeatures := "supportedFeatures_example" // string | Supported Features (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PriorityInfoRetrievalApi.GetPriorityInfo(context.Background(), imsUeId).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PriorityInfoRetrievalApi.GetPriorityInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPriorityInfo`: PriorityLevels
    fmt.Fprintf(os.Stdout, "Response from `PriorityInfoRetrievalApi.GetPriorityInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imsUeId** | **string** | IMS Identity | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPriorityInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **supportedFeatures** | **string** | Supported Features | 

### Return type

[**PriorityLevels**](PriorityLevels.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

