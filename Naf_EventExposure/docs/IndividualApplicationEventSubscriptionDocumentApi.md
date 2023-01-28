# \IndividualApplicationEventSubscriptionDocumentApi

All URIs are relative to *https://example.com/naf-eventexposure/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAfEventExposureSubsc**](IndividualApplicationEventSubscriptionDocumentApi.md#DeleteAfEventExposureSubsc) | **Delete** /subscriptions/{subscriptionId} | Cancels an existing Individual Application Event Subscription 
[**GetAfEventExposureSubsc**](IndividualApplicationEventSubscriptionDocumentApi.md#GetAfEventExposureSubsc) | **Get** /subscriptions/{subscriptionId} | Reads an existing Individual Application Event Subscription
[**PutAfEventExposureSubsc**](IndividualApplicationEventSubscriptionDocumentApi.md#PutAfEventExposureSubsc) | **Put** /subscriptions/{subscriptionId} | Modifies an existing Individual Application Event Subscription 



## DeleteAfEventExposureSubsc

> DeleteAfEventExposureSubsc(ctx, subscriptionId).Execute()

Cancels an existing Individual Application Event Subscription 

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Naf_EventExposure"
)

func main() {
    subscriptionId := "subscriptionId_example" // string | Application Event Subscription ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualApplicationEventSubscriptionDocumentApi.DeleteAfEventExposureSubsc(context.Background(), subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualApplicationEventSubscriptionDocumentApi.DeleteAfEventExposureSubsc``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Application Event Subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAfEventExposureSubscRequest struct via the builder pattern


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


## GetAfEventExposureSubsc

> AfEventExposureSubsc GetAfEventExposureSubsc(ctx, subscriptionId).SuppFeat(suppFeat).Execute()

Reads an existing Individual Application Event Subscription

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Naf_EventExposure"
)

func main() {
    subscriptionId := "subscriptionId_example" // string | Application Event Subscription ID
    suppFeat := "suppFeat_example" // string | Features supported by the NF service consumer (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualApplicationEventSubscriptionDocumentApi.GetAfEventExposureSubsc(context.Background(), subscriptionId).SuppFeat(suppFeat).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualApplicationEventSubscriptionDocumentApi.GetAfEventExposureSubsc``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAfEventExposureSubsc`: AfEventExposureSubsc
    fmt.Fprintf(os.Stdout, "Response from `IndividualApplicationEventSubscriptionDocumentApi.GetAfEventExposureSubsc`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Application Event Subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAfEventExposureSubscRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **suppFeat** | **string** | Features supported by the NF service consumer | 

### Return type

[**AfEventExposureSubsc**](AfEventExposureSubsc.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAfEventExposureSubsc

> AfEventExposureSubsc PutAfEventExposureSubsc(ctx, subscriptionId).AfEventExposureSubsc(afEventExposureSubsc).Execute()

Modifies an existing Individual Application Event Subscription 

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Naf_EventExposure"
)

func main() {
    subscriptionId := "subscriptionId_example" // string | Application Event Subscription ID
    afEventExposureSubsc := *openapiclient.NewAfEventExposureSubsc([]openapiclient.EventsSubs{*openapiclient.NewEventsSubs(*openapiclient.NewAfEvent(), *openapiclient.NewEventFilter())}, *openapiclient.NewReportingInformation(), "NotifUri_example", "NotifId_example") // AfEventExposureSubsc | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualApplicationEventSubscriptionDocumentApi.PutAfEventExposureSubsc(context.Background(), subscriptionId).AfEventExposureSubsc(afEventExposureSubsc).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualApplicationEventSubscriptionDocumentApi.PutAfEventExposureSubsc``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutAfEventExposureSubsc`: AfEventExposureSubsc
    fmt.Fprintf(os.Stdout, "Response from `IndividualApplicationEventSubscriptionDocumentApi.PutAfEventExposureSubsc`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Application Event Subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutAfEventExposureSubscRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **afEventExposureSubsc** | [**AfEventExposureSubsc**](AfEventExposureSubsc.md) |  | 

### Return type

[**AfEventExposureSubsc**](AfEventExposureSubsc.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

