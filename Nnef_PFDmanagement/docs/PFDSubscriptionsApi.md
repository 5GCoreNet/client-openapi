# \PFDSubscriptionsApi

All URIs are relative to *https://example.com/nnef-pfdmanagement/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NnefPFDmanagementCreateSubscr**](PFDSubscriptionsApi.md#NnefPFDmanagementCreateSubscr) | **Post** /subscriptions | Subscribe the notification of PFD changes.



## NnefPFDmanagementCreateSubscr

> PfdSubscription NnefPFDmanagementCreateSubscr(ctx).PfdSubscription(pfdSubscription).Execute()

Subscribe the notification of PFD changes.

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
    pfdSubscription := *openapiclient.NewPfdSubscription("NotifyUri_example", "SupportedFeatures_example") // PfdSubscription | a PfdSubscription resource to be created.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PFDSubscriptionsApi.NnefPFDmanagementCreateSubscr(context.Background()).PfdSubscription(pfdSubscription).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PFDSubscriptionsApi.NnefPFDmanagementCreateSubscr``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NnefPFDmanagementCreateSubscr`: PfdSubscription
    fmt.Fprintf(os.Stdout, "Response from `PFDSubscriptionsApi.NnefPFDmanagementCreateSubscr`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNnefPFDmanagementCreateSubscrRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pfdSubscription** | [**PfdSubscription**](PfdSubscription.md) | a PfdSubscription resource to be created. | 

### Return type

[**PfdSubscription**](PfdSubscription.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

