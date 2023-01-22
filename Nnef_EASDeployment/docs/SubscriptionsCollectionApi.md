# \SubscriptionsCollectionApi

All URIs are relative to *https://example.com/nnef-eas-deployment/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIndividualSubcription**](SubscriptionsCollectionApi.md#CreateIndividualSubcription) | **Post** /subscriptions | subscribe to notifications



## CreateIndividualSubcription

> EasDeploySubData CreateIndividualSubcription(ctx).EasDeploySubData(easDeploySubData).Execute()

subscribe to notifications

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
    easDeploySubData := *openapiclient.NewEasDeploySubData(*openapiclient.NewEasEvent(), "NotifId_example", "NotifUri_example") // EasDeploySubData | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionsCollectionApi.CreateIndividualSubcription(context.Background()).EasDeploySubData(easDeploySubData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsCollectionApi.CreateIndividualSubcription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIndividualSubcription`: EasDeploySubData
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionsCollectionApi.CreateIndividualSubcription`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIndividualSubcriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **easDeploySubData** | [**EasDeploySubData**](EasDeploySubData.md) |  | 

### Return type

[**EasDeploySubData**](EasDeploySubData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

