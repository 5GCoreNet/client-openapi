# \ApplicationSessionsCollectionApi

All URIs are relative to *https://example.com/npcf-policyauthorization/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostAppSessions**](ApplicationSessionsCollectionApi.md#PostAppSessions) | **Post** /app-sessions | Creates a new Individual Application Session Context resource



## PostAppSessions

> AppSessionContext PostAppSessions(ctx).AppSessionContext(appSessionContext).Execute()

Creates a new Individual Application Session Context resource

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Npcf_PolicyAuthorization"
)

func main() {
    appSessionContext := *openapiclient.NewAppSessionContext() // AppSessionContext | Contains the information for the creation the resource.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationSessionsCollectionApi.PostAppSessions(context.Background()).AppSessionContext(appSessionContext).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationSessionsCollectionApi.PostAppSessions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostAppSessions`: AppSessionContext
    fmt.Fprintf(os.Stdout, "Response from `ApplicationSessionsCollectionApi.PostAppSessions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAppSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appSessionContext** | [**AppSessionContext**](AppSessionContext.md) | Contains the information for the creation the resource. | 

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

