# \IndividualApplicationSessionContextDocumentApi

All URIs are relative to *https://example.com/npcf-policyauthorization/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAppSession**](IndividualApplicationSessionContextDocumentApi.md#DeleteAppSession) | **Post** /app-sessions/{appSessionId}/delete | Deletes an existing Individual Application Session Context
[**GetAppSession**](IndividualApplicationSessionContextDocumentApi.md#GetAppSession) | **Get** /app-sessions/{appSessionId} | Reads an existing Individual Application Session Context
[**ModAppSession**](IndividualApplicationSessionContextDocumentApi.md#ModAppSession) | **Patch** /app-sessions/{appSessionId} | Modifies an existing Individual Application Session Context



## DeleteAppSession

> AppSessionContext DeleteAppSession(ctx, appSessionId).EventsSubscReqData(eventsSubscReqData).Execute()

Deletes an existing Individual Application Session Context

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
    appSessionId := "appSessionId_example" // string | String identifying the Individual Application Session Context resource.
    eventsSubscReqData := *openapiclient.NewEventsSubscReqData([]openapiclient.AfEventSubscription{*openapiclient.NewAfEventSubscription(*openapiclient.NewAfEvent())}) // EventsSubscReqData | Deletion of the Individual Application Session Context resource, req notification. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualApplicationSessionContextDocumentApi.DeleteAppSession(context.Background(), appSessionId).EventsSubscReqData(eventsSubscReqData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualApplicationSessionContextDocumentApi.DeleteAppSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAppSession`: AppSessionContext
    fmt.Fprintf(os.Stdout, "Response from `IndividualApplicationSessionContextDocumentApi.DeleteAppSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSessionId** | **string** | String identifying the Individual Application Session Context resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAppSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eventsSubscReqData** | [**EventsSubscReqData**](EventsSubscReqData.md) | Deletion of the Individual Application Session Context resource, req notification. | 

### Return type

[**AppSessionContext**](AppSessionContext.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppSession

> AppSessionContext GetAppSession(ctx, appSessionId).Execute()

Reads an existing Individual Application Session Context

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
    appSessionId := "appSessionId_example" // string | String identifying the resource.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualApplicationSessionContextDocumentApi.GetAppSession(context.Background(), appSessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualApplicationSessionContextDocumentApi.GetAppSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAppSession`: AppSessionContext
    fmt.Fprintf(os.Stdout, "Response from `IndividualApplicationSessionContextDocumentApi.GetAppSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appSessionId** | **string** | String identifying the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AppSessionContext**](AppSessionContext.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModAppSession

> AppSessionContext ModAppSession(ctx, appSessionId).AppSessionContextUpdateDataPatch(appSessionContextUpdateDataPatch).Execute()

Modifies an existing Individual Application Session Context

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
    appSessionId := "appSessionId_example" // string | String identifying the resource.
    appSessionContextUpdateDataPatch := *openapiclient.NewAppSessionContextUpdateDataPatch() // AppSessionContextUpdateDataPatch | Modification of the resource.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualApplicationSessionContextDocumentApi.ModAppSession(context.Background(), appSessionId).AppSessionContextUpdateDataPatch(appSessionContextUpdateDataPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualApplicationSessionContextDocumentApi.ModAppSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ModAppSession`: AppSessionContext
    fmt.Fprintf(os.Stdout, "Response from `IndividualApplicationSessionContextDocumentApi.ModAppSession`: %v\n", resp)
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

 **appSessionContextUpdateDataPatch** | [**AppSessionContextUpdateDataPatch**](AppSessionContextUpdateDataPatch.md) | Modification of the resource. | 

### Return type

[**AppSessionContext**](AppSessionContext.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

