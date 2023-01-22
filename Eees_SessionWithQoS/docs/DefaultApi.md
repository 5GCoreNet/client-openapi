# \DefaultApi

All URIs are relative to *https://example.com/eees-session-with-qos/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SessionsGet**](DefaultApi.md#SessionsGet) | **Get** /sessions | 
[**SessionsPost**](DefaultApi.md#SessionsPost) | **Post** /sessions | 
[**SessionsSessionIdDelete**](DefaultApi.md#SessionsSessionIdDelete) | **Delete** /sessions/{sessionId} | 
[**SessionsSessionIdGet**](DefaultApi.md#SessionsSessionIdGet) | **Get** /sessions/{sessionId} | 
[**SessionsSessionIdPatch**](DefaultApi.md#SessionsSessionIdPatch) | **Patch** /sessions/{sessionId} | 
[**SessionsSessionIdPut**](DefaultApi.md#SessionsSessionIdPut) | **Put** /sessions/{sessionId} | 



## SessionsGet

> []SessionWithQoS SessionsGet(ctx).EasId(easId).Execute()





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
    easId := "easId_example" // string | Identifier of the EAS which querying the status of subscriptions.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.SessionsGet(context.Background()).EasId(easId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SessionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SessionsGet`: []SessionWithQoS
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SessionsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSessionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **easId** | **string** | Identifier of the EAS which querying the status of subscriptions. | 

### Return type

[**[]SessionWithQoS**](SessionWithQoS.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SessionsPost

> SessionWithQoS SessionsPost(ctx).SessionWithQoS(sessionWithQoS).Execute()





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
    sessionWithQoS := *openapiclient.NewSessionWithQoS("EasId_example", []string{"IpFlows_example"}) // SessionWithQoS | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.SessionsPost(context.Background()).SessionWithQoS(sessionWithQoS).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SessionsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SessionsPost`: SessionWithQoS
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SessionsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSessionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionWithQoS** | [**SessionWithQoS**](SessionWithQoS.md) |  | 

### Return type

[**SessionWithQoS**](SessionWithQoS.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SessionsSessionIdDelete

> SessionsSessionIdDelete(ctx, sessionId).Execute()





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
    sessionId := "sessionId_example" // string | session Id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.SessionsSessionIdDelete(context.Background(), sessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SessionsSessionIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | session Id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSessionsSessionIdDeleteRequest struct via the builder pattern


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


## SessionsSessionIdGet

> SessionWithQoS SessionsSessionIdGet(ctx, sessionId).Execute()





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
    sessionId := "sessionId_example" // string | Session Id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.SessionsSessionIdGet(context.Background(), sessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SessionsSessionIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SessionsSessionIdGet`: SessionWithQoS
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SessionsSessionIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Session Id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSessionsSessionIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SessionWithQoS**](SessionWithQoS.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SessionsSessionIdPatch

> SessionWithQoS SessionsSessionIdPatch(ctx, sessionId).SessionWithQoSPatch(sessionWithQoSPatch).Execute()



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
    sessionId := "sessionId_example" // string | session Id.
    sessionWithQoSPatch := *openapiclient.NewSessionWithQoSPatch() // SessionWithQoSPatch | Partial update an existing Individual Session with QoS resource identified by a sessionId. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.SessionsSessionIdPatch(context.Background(), sessionId).SessionWithQoSPatch(sessionWithQoSPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SessionsSessionIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SessionsSessionIdPatch`: SessionWithQoS
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SessionsSessionIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | session Id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSessionsSessionIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sessionWithQoSPatch** | [**SessionWithQoSPatch**](SessionWithQoSPatch.md) | Partial update an existing Individual Session with QoS resource identified by a sessionId.  | 

### Return type

[**SessionWithQoS**](SessionWithQoS.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SessionsSessionIdPut

> SessionWithQoS SessionsSessionIdPut(ctx, sessionId).SessionWithQoS(sessionWithQoS).Execute()





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
    sessionId := "sessionId_example" // string | Session Id.
    sessionWithQoS := *openapiclient.NewSessionWithQoS("EasId_example", []string{"IpFlows_example"}) // SessionWithQoS | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.SessionsSessionIdPut(context.Background(), sessionId).SessionWithQoS(sessionWithQoS).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SessionsSessionIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SessionsSessionIdPut`: SessionWithQoS
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SessionsSessionIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Session Id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSessionsSessionIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sessionWithQoS** | [**SessionWithQoS**](SessionWithQoS.md) |  | 

### Return type

[**SessionWithQoS**](SessionWithQoS.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

