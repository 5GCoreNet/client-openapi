# \MBSSessionSubscriptionsApi

All URIs are relative to *https://example.com/3gpp-mbs-session/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMBSSessionsSubsc**](MBSSessionSubscriptionsApi.md#CreateMBSSessionsSubsc) | **Post** /mbs-sessions/subscriptions | Request the creation of a new Individual MBS Session subscription resource.
[**ReadMBSSessionsSubscs**](MBSSessionSubscriptionsApi.md#ReadMBSSessionsSubscs) | **Get** /mbs-sessions/subscriptions | Retrieve all the active MBS Sessions subscriptions.



## CreateMBSSessionsSubsc

> MbsSessionSubsc CreateMBSSessionsSubsc(ctx).MbsSessionSubsc(mbsSessionSubsc).Execute()

Request the creation of a new Individual MBS Session subscription resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/MBSSession"
)

func main() {
    mbsSessionSubsc := *openapiclient.NewMbsSessionSubsc("AfId_example", *openapiclient.NewMbsSessionSubscription([]openapiclient.MbsSessionEvent{*openapiclient.NewMbsSessionEvent(*openapiclient.NewMbsSessionEventType())}, "NotifyUri_example")) // MbsSessionSubsc | Request the creation of a new MBS Session subscription resource.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MBSSessionSubscriptionsApi.CreateMBSSessionsSubsc(context.Background()).MbsSessionSubsc(mbsSessionSubsc).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MBSSessionSubscriptionsApi.CreateMBSSessionsSubsc``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMBSSessionsSubsc`: MbsSessionSubsc
    fmt.Fprintf(os.Stdout, "Response from `MBSSessionSubscriptionsApi.CreateMBSSessionsSubsc`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMBSSessionsSubscRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mbsSessionSubsc** | [**MbsSessionSubsc**](MbsSessionSubsc.md) | Request the creation of a new MBS Session subscription resource. | 

### Return type

[**MbsSessionSubsc**](MbsSessionSubsc.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadMBSSessionsSubscs

> []MbsSessionSubsc ReadMBSSessionsSubscs(ctx).Execute()

Retrieve all the active MBS Sessions subscriptions.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/MBSSession"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MBSSessionSubscriptionsApi.ReadMBSSessionsSubscs(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MBSSessionSubscriptionsApi.ReadMBSSessionsSubscs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadMBSSessionsSubscs`: []MbsSessionSubsc
    fmt.Fprintf(os.Stdout, "Response from `MBSSessionSubscriptionsApi.ReadMBSSessionsSubscs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiReadMBSSessionsSubscsRequest struct via the builder pattern


### Return type

[**[]MbsSessionSubsc**](MbsSessionSubsc.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

