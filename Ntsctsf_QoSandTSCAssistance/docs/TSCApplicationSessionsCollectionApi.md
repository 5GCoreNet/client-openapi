# \TSCApplicationSessionsCollectionApi

All URIs are relative to *https://example.com/ntsctsf-qos-tscai/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostTSCAppSessions**](TSCApplicationSessionsCollectionApi.md#PostTSCAppSessions) | **Post** /tsc-app-sessions | Creates a new Individual TSC Application Session Context resource



## PostTSCAppSessions

> TscAppSessionContextData PostTSCAppSessions(ctx).TscAppSessionContextData(tscAppSessionContextData).Execute()

Creates a new Individual TSC Application Session Context resource

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
    tscAppSessionContextData := openapiclient.TscAppSessionContextData{Interface{}: new(interface{})} // TscAppSessionContextData | Contains the information for the creation the resource.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TSCApplicationSessionsCollectionApi.PostTSCAppSessions(context.Background()).TscAppSessionContextData(tscAppSessionContextData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TSCApplicationSessionsCollectionApi.PostTSCAppSessions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostTSCAppSessions`: TscAppSessionContextData
    fmt.Fprintf(os.Stdout, "Response from `TSCApplicationSessionsCollectionApi.PostTSCAppSessions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostTSCAppSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tscAppSessionContextData** | [**TscAppSessionContextData**](TscAppSessionContextData.md) | Contains the information for the creation the resource. | 

### Return type

[**TscAppSessionContextData**](TscAppSessionContextData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

