# \DefaultApi

All URIs are relative to *http://example.com/3GPPManagement*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NotificationSinkPost**](DefaultApi.md#NotificationSinkPost) | **Post** /notificationSink | Send notifications about performance threshold crossing



## NotificationSinkPost

> NotificationSinkPost(ctx).NotifyThresholdCrossing(notifyThresholdCrossing).Execute()

Send notifications about performance threshold crossing



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    notifyThresholdCrossing := *openapiclient.NewNotifyThresholdCrossing("Href_example", int32(123), openapiclient.NotificationType{AlarmNotificationTypes: penapiclient.AlarmNotificationTypes("notifyNewAlarm")}, time.Now(), "SystemDN_example") // NotifyThresholdCrossing | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.NotificationSinkPost(context.Background()).NotifyThresholdCrossing(notifyThresholdCrossing).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.NotificationSinkPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNotificationSinkPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **notifyThresholdCrossing** | [**NotifyThresholdCrossing**](NotifyThresholdCrossing.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

