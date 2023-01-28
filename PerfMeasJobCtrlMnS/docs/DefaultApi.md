# \DefaultApi

All URIs are relative to *http://example.com/3GPPManagement/PerfMeasJobCtrlMnS/XXX*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MeasJobsGet**](DefaultApi.md#MeasJobsGet) | **Get** /measJobs | Read resources of measurement jobs
[**MeasJobsJobIdDelete**](DefaultApi.md#MeasJobsJobIdDelete) | **Delete** /measJobs/{jobId} | Delete a single measurement job
[**MeasJobsJobIdGet**](DefaultApi.md#MeasJobsJobIdGet) | **Get** /measJobs/{jobId} | Read resource of a single measurement job
[**MeasJobsPost**](DefaultApi.md#MeasJobsPost) | **Post** /measJobs | Create a measurement job



## MeasJobsGet

> MeasJobsRetrievalResponseType MeasJobsGet(ctx).JobIdList(jobIdList).Execute()

Read resources of measurement jobs



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/PerfMeasJobCtrlMnS"
)

func main() {
    jobIdList := []string{"Inner_example"} // []string | This parameter identifies the list of jobId to select the resources from the collection resources identified with the path component of the URI.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.MeasJobsGet(context.Background()).JobIdList(jobIdList).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.MeasJobsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeasJobsGet`: MeasJobsRetrievalResponseType
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.MeasJobsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeasJobsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jobIdList** | **[]string** | This parameter identifies the list of jobId to select the resources from the collection resources identified with the path component of the URI. | 

### Return type

[**MeasJobsRetrievalResponseType**](MeasJobsRetrievalResponseType.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeasJobsJobIdDelete

> MeasJobsJobIdDelete(ctx, jobId).Execute()

Delete a single measurement job



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/PerfMeasJobCtrlMnS"
)

func main() {
    jobId := "jobId_example" // string | Identifies the measurement job to be deleted.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.MeasJobsJobIdDelete(context.Background(), jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.MeasJobsJobIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Identifies the measurement job to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeasJobsJobIdDeleteRequest struct via the builder pattern


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


## MeasJobsJobIdGet

> MeasJobsRetrievalResponseType MeasJobsJobIdGet(ctx, jobId).Execute()

Read resource of a single measurement job



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/PerfMeasJobCtrlMnS"
)

func main() {
    jobId := "jobId_example" // string | Identifies the measurement job to be read.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.MeasJobsJobIdGet(context.Background(), jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.MeasJobsJobIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeasJobsJobIdGet`: MeasJobsRetrievalResponseType
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.MeasJobsJobIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Identifies the measurement job to be read. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMeasJobsJobIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MeasJobsRetrievalResponseType**](MeasJobsRetrievalResponseType.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MeasJobsPost

> MeasJobCreationResponseType MeasJobsPost(ctx).MeasJobCreationRequestType(measJobCreationRequestType).Execute()

Create a measurement job



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/PerfMeasJobCtrlMnS"
)

func main() {
    measJobCreationRequestType := *openapiclient.NewMeasJobCreationRequestType() // MeasJobCreationRequestType | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.MeasJobsPost(context.Background()).MeasJobCreationRequestType(measJobCreationRequestType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.MeasJobsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MeasJobsPost`: MeasJobCreationResponseType
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.MeasJobsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMeasJobsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **measJobCreationRequestType** | [**MeasJobCreationRequestType**](MeasJobCreationRequestType.md) |  | 

### Return type

[**MeasJobCreationResponseType**](MeasJobCreationResponseType.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

