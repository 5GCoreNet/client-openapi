# \AMPolicyEventsSubscriptionApi

All URIs are relative to *https://example.com/3gpp-am-policyauthorization/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAmEventsSubsc**](AMPolicyEventsSubscriptionApi.md#DeleteAmEventsSubsc) | **Delete** /{afId}/app-am-contexts/{appAmContextId}/events-subscription | deletes the AM Policy Events Subscription sub-resource
[**UpdateAmEventsSubsc**](AMPolicyEventsSubscriptionApi.md#UpdateAmEventsSubsc) | **Put** /{afId}/app-am-contexts/{appAmContextId}/events-subscription | creates or modifies an AM Policy Events Subscription sub-resource.



## DeleteAmEventsSubsc

> DeleteAmEventsSubsc(ctx, afId, appAmContextId).Execute()

deletes the AM Policy Events Subscription sub-resource

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/AMPolicyAuthorization"
)

func main() {
    afId := "afId_example" // string | Identifier of the AF
    appAmContextId := "appAmContextId_example" // string | string identifying the Individual Application AM Context resource.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AMPolicyEventsSubscriptionApi.DeleteAmEventsSubsc(context.Background(), afId, appAmContextId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AMPolicyEventsSubscriptionApi.DeleteAmEventsSubsc``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**afId** | **string** | Identifier of the AF | 
**appAmContextId** | **string** | string identifying the Individual Application AM Context resource. | 

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
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAmEventsSubsc

> AmEventsSubscRespData UpdateAmEventsSubsc(ctx, afId, appAmContextId).AmEventsSubscData(amEventsSubscData).Execute()

creates or modifies an AM Policy Events Subscription sub-resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/AMPolicyAuthorization"
)

func main() {
    afId := "afId_example" // string | Identifier of the AF
    appAmContextId := "appAmContextId_example" // string | string identifying the AM Policy Events Subscription subresource
    amEventsSubscData := *openapiclient.NewAmEventsSubscData("EventNotifUri_example") // AmEventsSubscData | Creation or modification of an application AM Policy Events Subscription sub-resource. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AMPolicyEventsSubscriptionApi.UpdateAmEventsSubsc(context.Background(), afId, appAmContextId).AmEventsSubscData(amEventsSubscData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AMPolicyEventsSubscriptionApi.UpdateAmEventsSubsc``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAmEventsSubsc`: AmEventsSubscRespData
    fmt.Fprintf(os.Stdout, "Response from `AMPolicyEventsSubscriptionApi.UpdateAmEventsSubsc`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**afId** | **string** | Identifier of the AF | 
**appAmContextId** | **string** | string identifying the AM Policy Events Subscription subresource | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAmEventsSubscRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **amEventsSubscData** | [**AmEventsSubscData**](AmEventsSubscData.md) | Creation or modification of an application AM Policy Events Subscription sub-resource.  | 

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

