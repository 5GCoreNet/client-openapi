# \DefaultApi

All URIs are relative to *https://example.com/api-invoker-management/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OnboardedInvokersOnboardingIdDelete**](DefaultApi.md#OnboardedInvokersOnboardingIdDelete) | **Delete** /onboardedInvokers/{onboardingId} | 
[**OnboardedInvokersOnboardingIdPut**](DefaultApi.md#OnboardedInvokersOnboardingIdPut) | **Put** /onboardedInvokers/{onboardingId} | 
[**OnboardedInvokersPost**](DefaultApi.md#OnboardedInvokersPost) | **Post** /onboardedInvokers | 



## OnboardedInvokersOnboardingIdDelete

> OnboardedInvokersOnboardingIdDelete(ctx, onboardingId).Execute()





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
    onboardingId := "onboardingId_example" // string | String identifying an individual on-boarded API invoker resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.OnboardedInvokersOnboardingIdDelete(context.Background(), onboardingId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OnboardedInvokersOnboardingIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**onboardingId** | **string** | String identifying an individual on-boarded API invoker resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiOnboardedInvokersOnboardingIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OnboardedInvokersOnboardingIdPut

> APIInvokerEnrolmentDetails OnboardedInvokersOnboardingIdPut(ctx, onboardingId).APIInvokerEnrolmentDetails(aPIInvokerEnrolmentDetails).Execute()





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
    onboardingId := "onboardingId_example" // string | String identifying an individual on-boarded API invoker resource
    aPIInvokerEnrolmentDetails := *openapiclient.NewAPIInvokerEnrolmentDetails(*openapiclient.NewOnboardingInformation("ApiInvokerPublicKey_example"), "NotificationDestination_example") // APIInvokerEnrolmentDetails | representation of the API invoker details to be updated in CAPIF core function

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.OnboardedInvokersOnboardingIdPut(context.Background(), onboardingId).APIInvokerEnrolmentDetails(aPIInvokerEnrolmentDetails).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OnboardedInvokersOnboardingIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OnboardedInvokersOnboardingIdPut`: APIInvokerEnrolmentDetails
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OnboardedInvokersOnboardingIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**onboardingId** | **string** | String identifying an individual on-boarded API invoker resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiOnboardedInvokersOnboardingIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aPIInvokerEnrolmentDetails** | [**APIInvokerEnrolmentDetails**](APIInvokerEnrolmentDetails.md) | representation of the API invoker details to be updated in CAPIF core function | 

### Return type

[**APIInvokerEnrolmentDetails**](APIInvokerEnrolmentDetails.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OnboardedInvokersPost

> APIInvokerEnrolmentDetails OnboardedInvokersPost(ctx).APIInvokerEnrolmentDetails(aPIInvokerEnrolmentDetails).Execute()





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
    aPIInvokerEnrolmentDetails := *openapiclient.NewAPIInvokerEnrolmentDetails(*openapiclient.NewOnboardingInformation("ApiInvokerPublicKey_example"), "NotificationDestination_example") // APIInvokerEnrolmentDetails | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.OnboardedInvokersPost(context.Background()).APIInvokerEnrolmentDetails(aPIInvokerEnrolmentDetails).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OnboardedInvokersPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OnboardedInvokersPost`: APIInvokerEnrolmentDetails
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OnboardedInvokersPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOnboardedInvokersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aPIInvokerEnrolmentDetails** | [**APIInvokerEnrolmentDetails**](APIInvokerEnrolmentDetails.md) |  | 

### Return type

[**APIInvokerEnrolmentDetails**](APIInvokerEnrolmentDetails.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

