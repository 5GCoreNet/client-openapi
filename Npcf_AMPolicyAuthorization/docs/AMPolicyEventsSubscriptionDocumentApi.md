# \AMPolicyEventsSubscriptionDocumentApi

All URIs are relative to *https://example.com/npcf-am-policyauthorization/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAmEventsSubsc**](AMPolicyEventsSubscriptionDocumentApi.md#DeleteAmEventsSubsc) | **Delete** /app-am-contexts/{appAmContextId}/events-subscription | deletes the AM Policy Events Subscription subresource
[**UpdateAmEventsSubsc**](AMPolicyEventsSubscriptionDocumentApi.md#UpdateAmEventsSubsc) | **Put** /app-am-contexts/{appAmContextId}/events-subscription | creates or modifies an AM Policy Events Subscription subresource.



## DeleteAmEventsSubsc

> DeleteAmEventsSubsc(ctx, appAmContextId).Execute()

deletes the AM Policy Events Subscription subresource

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Npcf_AMPolicyAuthorization"
)

func main() {
    appAmContextId := "appAmContextId_example" // string | String identifying the Individual Application AM Context resource.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AMPolicyEventsSubscriptionDocumentApi.DeleteAmEventsSubsc(context.Background(), appAmContextId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AMPolicyEventsSubscriptionDocumentApi.DeleteAmEventsSubsc``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appAmContextId** | **string** | String identifying the Individual Application AM Context resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAmEventsSubscRequest struct via the builder pattern


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


## UpdateAmEventsSubsc

> AmEventsSubscRespData UpdateAmEventsSubsc(ctx, appAmContextId).AmEventsSubscData(amEventsSubscData).Execute()

creates or modifies an AM Policy Events Subscription subresource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Npcf_AMPolicyAuthorization"
)

func main() {
    appAmContextId := "appAmContextId_example" // string | String identifying the AM Policy Events Subscription subresource.
    amEventsSubscData := *openapiclient.NewAmEventsSubscData("EventNotifUri_example") // AmEventsSubscData | Creation or modification of an AM Policy Events Subscription subresource.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AMPolicyEventsSubscriptionDocumentApi.UpdateAmEventsSubsc(context.Background(), appAmContextId).AmEventsSubscData(amEventsSubscData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AMPolicyEventsSubscriptionDocumentApi.UpdateAmEventsSubsc``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAmEventsSubsc`: AmEventsSubscRespData
    fmt.Fprintf(os.Stdout, "Response from `AMPolicyEventsSubscriptionDocumentApi.UpdateAmEventsSubsc`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appAmContextId** | **string** | String identifying the AM Policy Events Subscription subresource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAmEventsSubscRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **amEventsSubscData** | [**AmEventsSubscData**](AmEventsSubscData.md) | Creation or modification of an AM Policy Events Subscription subresource. | 

### Return type

[**AmEventsSubscRespData**](AmEventsSubscRespData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

