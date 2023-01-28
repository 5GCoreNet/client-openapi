# \IndividualTSCApplicationSessionContextDocumentApi

All URIs are relative to *https://example.com/ntsctsf-qos-tscai/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteTSCAppSession**](IndividualTSCApplicationSessionContextDocumentApi.md#DeleteTSCAppSession) | **Post** /tsc-app-sessions/{appSessionId}/delete | Deletes an existing Individual TSC Application Session Context
[**GetTSCAppSession**](IndividualTSCApplicationSessionContextDocumentApi.md#GetTSCAppSession) | **Get** /tsc-app-sessions/{appSessionId} | Reads an existing Individual TSC Application Session Context
[**ModAppSession**](IndividualTSCApplicationSessionContextDocumentApi.md#ModAppSession) | **Patch** /tsc-app-sessions/{appSessionId} | Modifies an existing Individual TSC Application Session Context



## DeleteTSCAppSession

> EventsNotification DeleteTSCAppSession(ctx, appSessionId).EventsSubscReqData(eventsSubscReqData).Execute()

Deletes an existing Individual TSC Application Session Context

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Ntsctsf_QoSandTSCAssistance"
)

func main() {
    appSessionId := "appSessionId_example" // string | String identifying the Individual TSC Application Session Context resource.
    eventsSubscReqData := *openapiclient.NewEventsSubscReqData([]openapiclient.TscEvent{*openapiclient.NewTscEvent()}, "NotifUri_example", "NotifCorreId_example") // EventsSubscReqData | Deletion of the Individual TSC Application Session Context resource, request notification.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualTSCApplicationSessionContextDocumentApi.DeleteTSCAppSession(context.Background(), appSessionId).EventsSubscReqData(eventsSubscReqData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualTSCApplicationSessionContextDocumentApi.DeleteTSCAppSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteTSCAppSession`: EventsNotification
    fmt.Fprintf(os.Stdout, "Response from `IndividualTSCApplicationSessionContextDocumentApi.DeleteTSCAppSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSessionId** | **string** | String identifying the Individual TSC Application Session Context resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTSCAppSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eventsSubscReqData** | [**EventsSubscReqData**](EventsSubscReqData.md) | Deletion of the Individual TSC Application Session Context resource, request notification.  | 

### Return type

[**EventsNotification**](EventsNotification.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTSCAppSession

> TscAppSessionContextData GetTSCAppSession(ctx, appSessionId).Execute()

Reads an existing Individual TSC Application Session Context

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Ntsctsf_QoSandTSCAssistance"
)

func main() {
    appSessionId := "appSessionId_example" // string | String identifying the resource.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualTSCApplicationSessionContextDocumentApi.GetTSCAppSession(context.Background(), appSessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualTSCApplicationSessionContextDocumentApi.GetTSCAppSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTSCAppSession`: TscAppSessionContextData
    fmt.Fprintf(os.Stdout, "Response from `IndividualTSCApplicationSessionContextDocumentApi.GetTSCAppSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSessionId** | **string** | String identifying the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTSCAppSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TscAppSessionContextData**](TscAppSessionContextData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModAppSession

> TscAppSessionContextData ModAppSession(ctx, appSessionId).TscAppSessionContextUpdateData(tscAppSessionContextUpdateData).Execute()

Modifies an existing Individual TSC Application Session Context

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Ntsctsf_QoSandTSCAssistance"
)

func main() {
    appSessionId := "appSessionId_example" // string | String identifying the resource.
    tscAppSessionContextUpdateData := *openapiclient.NewTscAppSessionContextUpdateData() // TscAppSessionContextUpdateData | Modification of the resource.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualTSCApplicationSessionContextDocumentApi.ModAppSession(context.Background(), appSessionId).TscAppSessionContextUpdateData(tscAppSessionContextUpdateData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualTSCApplicationSessionContextDocumentApi.ModAppSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ModAppSession`: TscAppSessionContextData
    fmt.Fprintf(os.Stdout, "Response from `IndividualTSCApplicationSessionContextDocumentApi.ModAppSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSessionId** | **string** | String identifying the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiModAppSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tscAppSessionContextUpdateData** | [**TscAppSessionContextUpdateData**](TscAppSessionContextUpdateData.md) | Modification of the resource. | 

### Return type

[**TscAppSessionContextData**](TscAppSessionContextData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

