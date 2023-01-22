# \IndividualMediaStreamingEventExposureSubscriptionDocumentApi

All URIs are relative to *https://example.com/3gpp-ms-event-exposure/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIndivMSEventExposureSubsc**](IndividualMediaStreamingEventExposureSubscriptionDocumentApi.md#DeleteIndivMSEventExposureSubsc) | **Delete** /subscriptions/{subscriptionId} | Request the deletion of an existing Individual Media Streaming Event Exposure Subscription resource.
[**RetrieveIndivMSEventExposureSubsc**](IndividualMediaStreamingEventExposureSubscriptionDocumentApi.md#RetrieveIndivMSEventExposureSubsc) | **Get** /subscriptions/{subscriptionId} | Retrieve an existing Individual Media Streaming Event Exposure Subscription resource.
[**UpdateIndivMSEventExposureSubsc**](IndividualMediaStreamingEventExposureSubscriptionDocumentApi.md#UpdateIndivMSEventExposureSubsc) | **Put** /subscriptions/{subscriptionId} | Request the update of an existing Individual Media Streaming Event Exposure Subscription resource.



## DeleteIndivMSEventExposureSubsc

> DeleteIndivMSEventExposureSubsc(ctx, subscriptionId).Execute()

Request the deletion of an existing Individual Media Streaming Event Exposure Subscription resource.

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
    subscriptionId := "subscriptionId_example" // string | Identifier of the Individual Media Streaming Event Exposure Subscription resource.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualMediaStreamingEventExposureSubscriptionDocumentApi.DeleteIndivMSEventExposureSubsc(context.Background(), subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualMediaStreamingEventExposureSubscriptionDocumentApi.DeleteIndivMSEventExposureSubsc``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Identifier of the Individual Media Streaming Event Exposure Subscription resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIndivMSEventExposureSubscRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveIndivMSEventExposureSubsc

> AfEventExposureSubsc RetrieveIndivMSEventExposureSubsc(ctx, subscriptionId).Execute()

Retrieve an existing Individual Media Streaming Event Exposure Subscription resource.

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
    subscriptionId := "subscriptionId_example" // string | Identifier of the Individual Media Streaming Event Exposure Subscription resource.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualMediaStreamingEventExposureSubscriptionDocumentApi.RetrieveIndivMSEventExposureSubsc(context.Background(), subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualMediaStreamingEventExposureSubscriptionDocumentApi.RetrieveIndivMSEventExposureSubsc``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveIndivMSEventExposureSubsc`: AfEventExposureSubsc
    fmt.Fprintf(os.Stdout, "Response from `IndividualMediaStreamingEventExposureSubscriptionDocumentApi.RetrieveIndivMSEventExposureSubsc`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Identifier of the Individual Media Streaming Event Exposure Subscription resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveIndivMSEventExposureSubscRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## UpdateIndivMSEventExposureSubsc

> AfEventExposureSubsc UpdateIndivMSEventExposureSubsc(ctx, subscriptionId).AfEventExposureSubsc(afEventExposureSubsc).Execute()

Request the update of an existing Individual Media Streaming Event Exposure Subscription resource.

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
    subscriptionId := "subscriptionId_example" // string | Identifier of the Individual Media Streaming Event Exposure Subscription resource.
    afEventExposureSubsc := *openapiclient.NewAfEventExposureSubsc([]openapiclient.EventsSubs{*openapiclient.NewEventsSubs(*openapiclient.NewAfEvent(), *openapiclient.NewEventFilter())}, *openapiclient.NewReportingInformation(), "NotifUri_example", "NotifId_example") // AfEventExposureSubsc | Contains the updated representation of the Individual Media Streaming Event Exposure  Subscription resource. resource. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualMediaStreamingEventExposureSubscriptionDocumentApi.UpdateIndivMSEventExposureSubsc(context.Background(), subscriptionId).AfEventExposureSubsc(afEventExposureSubsc).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualMediaStreamingEventExposureSubscriptionDocumentApi.UpdateIndivMSEventExposureSubsc``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIndivMSEventExposureSubsc`: AfEventExposureSubsc
    fmt.Fprintf(os.Stdout, "Response from `IndividualMediaStreamingEventExposureSubscriptionDocumentApi.UpdateIndivMSEventExposureSubsc`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Identifier of the Individual Media Streaming Event Exposure Subscription resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIndivMSEventExposureSubscRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **afEventExposureSubsc** | [**AfEventExposureSubsc**](AfEventExposureSubsc.md) | Contains the updated representation of the Individual Media Streaming Event Exposure  Subscription resource. resource.  | 

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

