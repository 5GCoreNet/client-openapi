# \DefaultApi

All URIs are relative to *http://example.com/3GPPManagement/fileDataReportingMnS/XXX*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FilesGet**](DefaultApi.md#FilesGet) | **Get** /files | Read information about available files
[**SubscriptionsPost**](DefaultApi.md#SubscriptionsPost) | **Post** /subscriptions | Create a subscription
[**SubscriptionsSubscriptionIdDelete**](DefaultApi.md#SubscriptionsSubscriptionIdDelete) | **Delete** /subscriptions/{subscriptionId} | Delete a subscription



## FilesGet

> []FileInfo FilesGet(ctx).FileDataType(fileDataType).BeginTime(beginTime).EndTime(endTime).Execute()

Read information about available files



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
    fileDataType := openapiclient.FileDataType("Performance") // FileDataType | This parameter selects files based on the file data type.
    beginTime := time.Now() // time.Time | This parameter selects files based on the earliest time they became available (optional)
    endTime := time.Now() // time.Time | This parameter selects files based on the latest time they became available (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.FilesGet(context.Background()).FileDataType(fileDataType).BeginTime(beginTime).EndTime(endTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FilesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FilesGet`: []FileInfo
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FilesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFilesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fileDataType** | [**FileDataType**](FileDataType.md) | This parameter selects files based on the file data type. | 
 **beginTime** | **time.Time** | This parameter selects files based on the earliest time they became available | 
 **endTime** | **time.Time** | This parameter selects files based on the latest time they became available | 

### Return type

[**[]FileInfo**](FileInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
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
    openapiclient "./openapi"
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
    openapiclient "./openapi"
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

