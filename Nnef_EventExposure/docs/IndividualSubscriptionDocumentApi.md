# \IndividualSubscriptionDocumentApi

All URIs are relative to *https://example.com/nnef-eventexposure/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIndividualSubcription**](IndividualSubscriptionDocumentApi.md#DeleteIndividualSubcription) | **Delete** /subscriptions/{subscriptionId} | unsubscribe from notifications
[**GetIndividualSubcription**](IndividualSubscriptionDocumentApi.md#GetIndividualSubcription) | **Get** /subscriptions/{subscriptionId} | retrieve subscription
[**ReplaceIndividualSubcription**](IndividualSubscriptionDocumentApi.md#ReplaceIndividualSubcription) | **Put** /subscriptions/{subscriptionId} | update subscription



## DeleteIndividualSubcription

> DeleteIndividualSubcription(ctx, subscriptionId).Execute()

unsubscribe from notifications

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
    subscriptionId := "subscriptionId_example" // string | Event Subscription ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualSubscriptionDocumentApi.DeleteIndividualSubcription(context.Background(), subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualSubscriptionDocumentApi.DeleteIndividualSubcription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Event Subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIndividualSubcriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIndividualSubcription

> NefEventExposureSubsc GetIndividualSubcription(ctx, subscriptionId).SuppFeat(suppFeat).Execute()

retrieve subscription

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
    subscriptionId := "subscriptionId_example" // string | Event Subscription ID
    suppFeat := "suppFeat_example" // string | Features supported by the NF service consumer (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualSubscriptionDocumentApi.GetIndividualSubcription(context.Background(), subscriptionId).SuppFeat(suppFeat).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualSubscriptionDocumentApi.GetIndividualSubcription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIndividualSubcription`: NefEventExposureSubsc
    fmt.Fprintf(os.Stdout, "Response from `IndividualSubscriptionDocumentApi.GetIndividualSubcription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Event Subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIndividualSubcriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **suppFeat** | **string** | Features supported by the NF service consumer | 

### Return type

[**NefEventExposureSubsc**](NefEventExposureSubsc.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceIndividualSubcription

> NefEventExposureSubsc ReplaceIndividualSubcription(ctx, subscriptionId).NefEventExposureSubsc(nefEventExposureSubsc).Execute()

update subscription

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
    subscriptionId := "subscriptionId_example" // string | Event Subscription ID
    nefEventExposureSubsc := *openapiclient.NewNefEventExposureSubsc([]openapiclient.NefEventSubs{*openapiclient.NewNefEventSubs(*openapiclient.NewNefEvent())}, "NotifUri_example", "NotifId_example") // NefEventExposureSubsc | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualSubscriptionDocumentApi.ReplaceIndividualSubcription(context.Background(), subscriptionId).NefEventExposureSubsc(nefEventExposureSubsc).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualSubscriptionDocumentApi.ReplaceIndividualSubcription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceIndividualSubcription`: NefEventExposureSubsc
    fmt.Fprintf(os.Stdout, "Response from `IndividualSubscriptionDocumentApi.ReplaceIndividualSubcription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Event Subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceIndividualSubcriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nefEventExposureSubsc** | [**NefEventExposureSubsc**](NefEventExposureSubsc.md) |  | 

### Return type

[**NefEventExposureSubsc**](NefEventExposureSubsc.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

