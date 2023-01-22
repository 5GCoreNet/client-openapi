# \MBSUserDataIngestSessionStatusSubscriptionsCollectionApi

All URIs are relative to *https://example.com/nmbsf-mbs-ud-ingest/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMBSUserDataIngStatSubsc**](MBSUserDataIngestSessionStatusSubscriptionsCollectionApi.md#CreateMBSUserDataIngStatSubsc) | **Post** /status-subscriptions | Request the creation of a new MBS User Data Ingest Session Status Subscription.
[**RetrieveMBSUserDataIngStatSubscs**](MBSUserDataIngestSessionStatusSubscriptionsCollectionApi.md#RetrieveMBSUserDataIngStatSubscs) | **Get** /status-subscriptions | Retrieve all the active MBS User Data Ingest Session Status Subscription resources managed by the MBSF.



## CreateMBSUserDataIngStatSubsc

> MBSUserDataIngStatSubsc CreateMBSUserDataIngStatSubsc(ctx).MBSUserDataIngStatSubsc(mBSUserDataIngStatSubsc).Execute()

Request the creation of a new MBS User Data Ingest Session Status Subscription.

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
    mBSUserDataIngStatSubsc := *openapiclient.NewMBSUserDataIngStatSubsc("MbsIngSessionId_example", []openapiclient.SubscribedEvent{*openapiclient.NewSubscribedEvent(openapiclient.Event{EventOneOf: penapiclient.Event_oneOf("USER_DATA_ING_SESS_STARTING")})}, "NotifUri_example") // MBSUserDataIngStatSubsc | Contains the parameters to request the creation of a new MBS User Data Ingest Session Status Subscription. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MBSUserDataIngestSessionStatusSubscriptionsCollectionApi.CreateMBSUserDataIngStatSubsc(context.Background()).MBSUserDataIngStatSubsc(mBSUserDataIngStatSubsc).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MBSUserDataIngestSessionStatusSubscriptionsCollectionApi.CreateMBSUserDataIngStatSubsc``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMBSUserDataIngStatSubsc`: MBSUserDataIngStatSubsc
    fmt.Fprintf(os.Stdout, "Response from `MBSUserDataIngestSessionStatusSubscriptionsCollectionApi.CreateMBSUserDataIngStatSubsc`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMBSUserDataIngStatSubscRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mBSUserDataIngStatSubsc** | [**MBSUserDataIngStatSubsc**](MBSUserDataIngStatSubsc.md) | Contains the parameters to request the creation of a new MBS User Data Ingest Session Status Subscription.  | 

### Return type

[**MBSUserDataIngStatSubsc**](MBSUserDataIngStatSubsc.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveMBSUserDataIngStatSubscs

> []MBSUserDataIngStatSubsc RetrieveMBSUserDataIngStatSubscs(ctx).Execute()

Retrieve all the active MBS User Data Ingest Session Status Subscription resources managed by the MBSF.

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MBSUserDataIngestSessionStatusSubscriptionsCollectionApi.RetrieveMBSUserDataIngStatSubscs(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MBSUserDataIngestSessionStatusSubscriptionsCollectionApi.RetrieveMBSUserDataIngStatSubscs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveMBSUserDataIngStatSubscs`: []MBSUserDataIngStatSubsc
    fmt.Fprintf(os.Stdout, "Response from `MBSUserDataIngestSessionStatusSubscriptionsCollectionApi.RetrieveMBSUserDataIngStatSubscs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveMBSUserDataIngStatSubscsRequest struct via the builder pattern


### Return type

[**[]MBSUserDataIngStatSubsc**](MBSUserDataIngStatSubsc.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

