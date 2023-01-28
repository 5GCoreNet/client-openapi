# \IndividualMBSUserDataIngestSessionStatusSubscriptionDocumentApi

All URIs are relative to *https://example.com/3gpp-mbs-ud-ingest/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIndMBSUserDataIngStatSubsc**](IndividualMBSUserDataIngestSessionStatusSubscriptionDocumentApi.md#DeleteIndMBSUserDataIngStatSubsc) | **Delete** /status-subscriptions/{subscriptionId} | Deletes an existing Individual MBS User Data Ingest Session Status Subscription resource.
[**ModifyIndMBSUserDataIngStatSubsc**](IndividualMBSUserDataIngestSessionStatusSubscriptionDocumentApi.md#ModifyIndMBSUserDataIngStatSubsc) | **Patch** /status-subscriptions/{subscriptionId} | Request the modification of an existing Individual MBS User Data Ingest Session Status Subscription resource.
[**RetrieveIndMBSUserDataIngStatSubsc**](IndividualMBSUserDataIngestSessionStatusSubscriptionDocumentApi.md#RetrieveIndMBSUserDataIngStatSubsc) | **Get** /status-subscriptions/{subscriptionId} | Retrieve an existing Individual MBS User Data Ingest Session Status Subscription resource.
[**UpdateIndMBSUserDataIngStatSubsc**](IndividualMBSUserDataIngestSessionStatusSubscriptionDocumentApi.md#UpdateIndMBSUserDataIngStatSubsc) | **Put** /status-subscriptions/{subscriptionId} | Request the update of an existing Individual MBS User Data Ingest Session Status Subscription resource.



## DeleteIndMBSUserDataIngStatSubsc

> DeleteIndMBSUserDataIngStatSubsc(ctx, subscriptionId).Execute()

Deletes an existing Individual MBS User Data Ingest Session Status Subscription resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/MBSUserDataIngestSession"
)

func main() {
    subscriptionId := "subscriptionId_example" // string | Identifier of the Individual MBS User Data Ingest Session Status Subscription resource. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualMBSUserDataIngestSessionStatusSubscriptionDocumentApi.DeleteIndMBSUserDataIngStatSubsc(context.Background(), subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualMBSUserDataIngestSessionStatusSubscriptionDocumentApi.DeleteIndMBSUserDataIngStatSubsc``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Identifier of the Individual MBS User Data Ingest Session Status Subscription resource.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIndMBSUserDataIngStatSubscRequest struct via the builder pattern


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


## ModifyIndMBSUserDataIngStatSubsc

> MBSUserDataIngStatSubsc ModifyIndMBSUserDataIngStatSubsc(ctx, subscriptionId).MBSUserDataIngStatSubscPatch(mBSUserDataIngStatSubscPatch).Execute()

Request the modification of an existing Individual MBS User Data Ingest Session Status Subscription resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/MBSUserDataIngestSession"
)

func main() {
    subscriptionId := "subscriptionId_example" // string | Identifier of the Individual MBS User Data Ingest Session Status Subscription resource. 
    mBSUserDataIngStatSubscPatch := *openapiclient.NewMBSUserDataIngStatSubscPatch() // MBSUserDataIngStatSubscPatch | Contains the parameters to request the modification of the Individual MBS User Data Ingest Session Status Subscription resource. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualMBSUserDataIngestSessionStatusSubscriptionDocumentApi.ModifyIndMBSUserDataIngStatSubsc(context.Background(), subscriptionId).MBSUserDataIngStatSubscPatch(mBSUserDataIngStatSubscPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualMBSUserDataIngestSessionStatusSubscriptionDocumentApi.ModifyIndMBSUserDataIngStatSubsc``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ModifyIndMBSUserDataIngStatSubsc`: MBSUserDataIngStatSubsc
    fmt.Fprintf(os.Stdout, "Response from `IndividualMBSUserDataIngestSessionStatusSubscriptionDocumentApi.ModifyIndMBSUserDataIngStatSubsc`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Identifier of the Individual MBS User Data Ingest Session Status Subscription resource.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyIndMBSUserDataIngStatSubscRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mBSUserDataIngStatSubscPatch** | [**MBSUserDataIngStatSubscPatch**](MBSUserDataIngStatSubscPatch.md) | Contains the parameters to request the modification of the Individual MBS User Data Ingest Session Status Subscription resource.  | 

### Return type

[**MBSUserDataIngStatSubsc**](MBSUserDataIngStatSubsc.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveIndMBSUserDataIngStatSubsc

> MBSUserDataIngStatSubsc RetrieveIndMBSUserDataIngStatSubsc(ctx, subscriptionId).Execute()

Retrieve an existing Individual MBS User Data Ingest Session Status Subscription resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/MBSUserDataIngestSession"
)

func main() {
    subscriptionId := "subscriptionId_example" // string | Identifier of the Individual MBS User Data Ingest Session Status Subscription resource. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualMBSUserDataIngestSessionStatusSubscriptionDocumentApi.RetrieveIndMBSUserDataIngStatSubsc(context.Background(), subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualMBSUserDataIngestSessionStatusSubscriptionDocumentApi.RetrieveIndMBSUserDataIngStatSubsc``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveIndMBSUserDataIngStatSubsc`: MBSUserDataIngStatSubsc
    fmt.Fprintf(os.Stdout, "Response from `IndividualMBSUserDataIngestSessionStatusSubscriptionDocumentApi.RetrieveIndMBSUserDataIngStatSubsc`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Identifier of the Individual MBS User Data Ingest Session Status Subscription resource.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveIndMBSUserDataIngStatSubscRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MBSUserDataIngStatSubsc**](MBSUserDataIngStatSubsc.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIndMBSUserDataIngStatSubsc

> MBSUserDataIngStatSubsc UpdateIndMBSUserDataIngStatSubsc(ctx, subscriptionId).MBSUserDataIngStatSubsc(mBSUserDataIngStatSubsc).Execute()

Request the update of an existing Individual MBS User Data Ingest Session Status Subscription resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/MBSUserDataIngestSession"
)

func main() {
    subscriptionId := "subscriptionId_example" // string | Identifier of the Individual MBS User Data Ingest Session Status Subscription resource. 
    mBSUserDataIngStatSubsc := *openapiclient.NewMBSUserDataIngStatSubsc("MbsIngSessionId_example", []openapiclient.SubscribedEvent{*openapiclient.NewSubscribedEvent(openapiclient.Event{EventOneOf: penapiclient.Event_oneOf("USER_DATA_ING_SESS_STARTING")})}, "NotifUri_example") // MBSUserDataIngStatSubsc | Contains the updated representation of the Individual MBS User Data Ingest Session Status Subscription resource. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualMBSUserDataIngestSessionStatusSubscriptionDocumentApi.UpdateIndMBSUserDataIngStatSubsc(context.Background(), subscriptionId).MBSUserDataIngStatSubsc(mBSUserDataIngStatSubsc).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualMBSUserDataIngestSessionStatusSubscriptionDocumentApi.UpdateIndMBSUserDataIngStatSubsc``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIndMBSUserDataIngStatSubsc`: MBSUserDataIngStatSubsc
    fmt.Fprintf(os.Stdout, "Response from `IndividualMBSUserDataIngestSessionStatusSubscriptionDocumentApi.UpdateIndMBSUserDataIngStatSubsc`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Identifier of the Individual MBS User Data Ingest Session Status Subscription resource.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIndMBSUserDataIngStatSubscRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mBSUserDataIngStatSubsc** | [**MBSUserDataIngStatSubsc**](MBSUserDataIngStatSubsc.md) | Contains the updated representation of the Individual MBS User Data Ingest Session Status Subscription resource.  | 

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

