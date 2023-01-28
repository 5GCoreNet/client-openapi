# \DefaultApi

All URIs are relative to *http://example.com/3GPPManagement/FaultSupervisionMnS/XXX*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AlarmsAlarmCountGet**](DefaultApi.md#AlarmsAlarmCountGet) | **Get** /alarms/alarmCount | Get the alarm count per perceived severity
[**AlarmsAlarmIdCommentsPost**](DefaultApi.md#AlarmsAlarmIdCommentsPost) | **Post** /alarms/{alarmId}/comments | Add a comment to a single alarm
[**AlarmsAlarmIdPatch**](DefaultApi.md#AlarmsAlarmIdPatch) | **Patch** /alarms/{alarmId} | Clear, acknowledge or unacknowledge a single alarm
[**AlarmsGet**](DefaultApi.md#AlarmsGet) | **Get** /alarms | Retrieve multiple alarms
[**AlarmsPatch**](DefaultApi.md#AlarmsPatch) | **Patch** /alarms | Clear, acknowledge or unacknowledge multiple alarms
[**SubscriptionsPost**](DefaultApi.md#SubscriptionsPost) | **Post** /subscriptions | Create a subscription
[**SubscriptionsSubscriptionIdDelete**](DefaultApi.md#SubscriptionsSubscriptionIdDelete) | **Delete** /subscriptions/{subscriptionId} | Delete a subscription



## AlarmsAlarmCountGet

> AlarmCount AlarmsAlarmCountGet(ctx).AlarmAckState(alarmAckState).Filter(filter).Execute()

Get the alarm count per perceived severity

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/FaultMnS"
)

func main() {
    alarmAckState := openapiclient.AlarmAckState("ALL_ALARMS") // AlarmAckState |  (optional)
    filter := "filter_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.AlarmsAlarmCountGet(context.Background()).AlarmAckState(alarmAckState).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AlarmsAlarmCountGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlarmsAlarmCountGet`: AlarmCount
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AlarmsAlarmCountGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlarmsAlarmCountGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **alarmAckState** | [**AlarmAckState**](AlarmAckState.md) |  | 
 **filter** | **string** |  | 

### Return type

[**AlarmCount**](AlarmCount.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlarmsAlarmIdCommentsPost

> Comment AlarmsAlarmIdCommentsPost(ctx, alarmId).Comment(comment).Execute()

Add a comment to a single alarm



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/FaultMnS"
)

func main() {
    alarmId := "alarmId_example" // string | Identifies the alarm to which the comment shall be added.
    comment := *openapiclient.NewComment() // Comment | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.AlarmsAlarmIdCommentsPost(context.Background(), alarmId).Comment(comment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AlarmsAlarmIdCommentsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlarmsAlarmIdCommentsPost`: Comment
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AlarmsAlarmIdCommentsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alarmId** | **string** | Identifies the alarm to which the comment shall be added. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlarmsAlarmIdCommentsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **comment** | [**Comment**](Comment.md) |  | 

### Return type

[**Comment**](Comment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlarmsAlarmIdPatch

> AlarmsAlarmIdPatch(ctx, alarmId).AlarmsAlarmIdPatchRequest(alarmsAlarmIdPatchRequest).Execute()

Clear, acknowledge or unacknowledge a single alarm



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/FaultMnS"
)

func main() {
    alarmId := "alarmId_example" // string | Identifies the alarm to be patched.
    alarmsAlarmIdPatchRequest := openapiclient._alarms__alarmId__patch_request{MergePatchAcknowledgeAlarm: openapiclient.NewMergePatchAcknowledgeAlarm("AckUserId_example", openapiclient.AckState("ACKNOWLEDGED"))} // AlarmsAlarmIdPatchRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.AlarmsAlarmIdPatch(context.Background(), alarmId).AlarmsAlarmIdPatchRequest(alarmsAlarmIdPatchRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AlarmsAlarmIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alarmId** | **string** | Identifies the alarm to be patched. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlarmsAlarmIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **alarmsAlarmIdPatchRequest** | [**AlarmsAlarmIdPatchRequest**](AlarmsAlarmIdPatchRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlarmsGet

> map[string]AlarmsGet200ResponseValue AlarmsGet(ctx).AlarmAckState(alarmAckState).BaseObjectInstance(baseObjectInstance).Filter(filter).Execute()

Retrieve multiple alarms



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/FaultMnS"
)

func main() {
    alarmAckState := openapiclient.AlarmAckState("ALL_ALARMS") // AlarmAckState |  (optional)
    baseObjectInstance := "baseObjectInstance_example" // string |  (optional)
    filter := "filter_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.AlarmsGet(context.Background()).AlarmAckState(alarmAckState).BaseObjectInstance(baseObjectInstance).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AlarmsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlarmsGet`: map[string]AlarmsGet200ResponseValue
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AlarmsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlarmsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **alarmAckState** | [**AlarmAckState**](AlarmAckState.md) |  | 
 **baseObjectInstance** | **string** |  | 
 **filter** | **string** |  | 

### Return type

[**map[string]AlarmsGet200ResponseValue**](AlarmsGet200ResponseValue.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlarmsPatch

> AlarmsPatch(ctx).AlarmsGetRequest(alarmsGetRequest).Execute()

Clear, acknowledge or unacknowledge multiple alarms



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/FaultMnS"
)

func main() {
    alarmsGetRequest := openapiclient._alarms_get_request{MapmapOfStringMergePatchAcknowledgeAlarm: new(map[string]MergePatchAcknowledgeAlarm)} // AlarmsGetRequest | Patch documents for acknowledging and unacknowledging, or clearing multiple alarms. The keys in the map are the alarmIds to be patched. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.AlarmsPatch(context.Background()).AlarmsGetRequest(alarmsGetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AlarmsPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlarmsPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **alarmsGetRequest** | [**AlarmsGetRequest**](AlarmsGetRequest.md) | Patch documents for acknowledging and unacknowledging, or clearing multiple alarms. The keys in the map are the alarmIds to be patched. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsPost

> Subscription SubscriptionsPost(ctx).Subscription(subscription).Execute()

Create a subscription



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/FaultMnS"
)

func main() {
    subscription := *openapiclient.NewSubscription() // Subscription | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.SubscriptionsPost(context.Background()).Subscription(subscription).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SubscriptionsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubscriptionsPost`: Subscription
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SubscriptionsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscription** | [**Subscription**](Subscription.md) |  | 

### Return type

[**Subscription**](Subscription.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsSubscriptionIdDelete

> SubscriptionsSubscriptionIdDelete(ctx, subscriptionId).Execute()

Delete a subscription



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/FaultMnS"
)

func main() {
    subscriptionId := "subscriptionId_example" // string | Identifies the subscription to be deleted.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.SubscriptionsSubscriptionIdDelete(context.Background(), subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SubscriptionsSubscriptionIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Identifies the subscription to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsSubscriptionIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

