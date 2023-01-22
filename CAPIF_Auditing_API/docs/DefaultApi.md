# \DefaultApi

All URIs are relative to *https://example.com/logs/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiInvocationLogsGet**](DefaultApi.md#ApiInvocationLogsGet) | **Get** /apiInvocationLogs | 



## ApiInvocationLogsGet

> InvocationLog ApiInvocationLogsGet(ctx).AefId(aefId).ApiInvokerId(apiInvokerId).TimeRangeStart(timeRangeStart).TimeRangeEnd(timeRangeEnd).ApiId(apiId).ApiName(apiName).ApiVersion(apiVersion).Protocol(protocol).Operation(operation).Result(result).ResourceName(resourceName).SrcInterface(srcInterface).DestInterface(destInterface).SupportedFeatures(supportedFeatures).Execute()





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
    aefId := "aefId_example" // string | String identifying the API exposing function. (optional)
    apiInvokerId := "apiInvokerId_example" // string | String identifying the API invoker which invoked the service API. (optional)
    timeRangeStart := time.Now() // time.Time | Start time of the invocation time range. (optional)
    timeRangeEnd := time.Now() // time.Time | End time of the invocation time range. (optional)
    apiId := "apiId_example" // string | String identifying the API invoked. (optional)
    apiName := "apiName_example" // string | API name, it is set as {apiName} part of the URI structure as defined in clause 5.2.4 of 3GPP TS 29.122.  (optional)
    apiVersion := "apiVersion_example" // string | Version of the API which was invoked. (optional)
    protocol := *openapiclient.NewProtocol() // Protocol | Protocol invoked. (optional)
    operation := *openapiclient.NewOperation() // Operation | Operation that was invoked on the API. (optional)
    result := "result_example" // string | Result or output of the invocation. (optional)
    resourceName := "resourceName_example" // string | Name of the specific resource invoked. (optional)
    srcInterface := map[string][]openapiclient.InterfaceDescription{ ... } // InterfaceDescription | Interface description of the API invoker. (optional)
    destInterface := map[string][]openapiclient.InterfaceDescription{ ... } // InterfaceDescription | Interface description of the API invoked. (optional)
    supportedFeatures := "supportedFeatures_example" // string | To filter irrelevant responses related to unsupported features (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ApiInvocationLogsGet(context.Background()).AefId(aefId).ApiInvokerId(apiInvokerId).TimeRangeStart(timeRangeStart).TimeRangeEnd(timeRangeEnd).ApiId(apiId).ApiName(apiName).ApiVersion(apiVersion).Protocol(protocol).Operation(operation).Result(result).ResourceName(resourceName).SrcInterface(srcInterface).DestInterface(destInterface).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApiInvocationLogsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiInvocationLogsGet`: InvocationLog
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApiInvocationLogsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiInvocationLogsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aefId** | **string** | String identifying the API exposing function. | 
 **apiInvokerId** | **string** | String identifying the API invoker which invoked the service API. | 
 **timeRangeStart** | **time.Time** | Start time of the invocation time range. | 
 **timeRangeEnd** | **time.Time** | End time of the invocation time range. | 
 **apiId** | **string** | String identifying the API invoked. | 
 **apiName** | **string** | API name, it is set as {apiName} part of the URI structure as defined in clause 5.2.4 of 3GPP TS 29.122.  | 
 **apiVersion** | **string** | Version of the API which was invoked. | 
 **protocol** | [**Protocol**](Protocol.md) | Protocol invoked. | 
 **operation** | [**Operation**](Operation.md) | Operation that was invoked on the API. | 
 **result** | **string** | Result or output of the invocation. | 
 **resourceName** | **string** | Name of the specific resource invoked. | 
 **srcInterface** | [**InterfaceDescription**](InterfaceDescription.md) | Interface description of the API invoker. | 
 **destInterface** | [**InterfaceDescription**](InterfaceDescription.md) | Interface description of the API invoked. | 
 **supportedFeatures** | **string** | To filter irrelevant responses related to unsupported features | 

### Return type

[**InvocationLog**](InvocationLog.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

