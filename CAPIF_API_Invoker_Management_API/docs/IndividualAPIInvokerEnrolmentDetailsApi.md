# \IndividualAPIInvokerEnrolmentDetailsApi

All URIs are relative to *https://example.com/api-invoker-management/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ModifyIndApiInvokeEnrolment**](IndividualAPIInvokerEnrolmentDetailsApi.md#ModifyIndApiInvokeEnrolment) | **Patch** /onboardedInvokers/{onboardingId} | 



## ModifyIndApiInvokeEnrolment

> APIInvokerEnrolmentDetails ModifyIndApiInvokeEnrolment(ctx, onboardingId).APIInvokerEnrolmentDetailsPatch(aPIInvokerEnrolmentDetailsPatch).Execute()





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
    onboardingId := "onboardingId_example" // string | 
    aPIInvokerEnrolmentDetailsPatch := *openapiclient.NewAPIInvokerEnrolmentDetailsPatch() // APIInvokerEnrolmentDetailsPatch | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualAPIInvokerEnrolmentDetailsApi.ModifyIndApiInvokeEnrolment(context.Background(), onboardingId).APIInvokerEnrolmentDetailsPatch(aPIInvokerEnrolmentDetailsPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualAPIInvokerEnrolmentDetailsApi.ModifyIndApiInvokeEnrolment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ModifyIndApiInvokeEnrolment`: APIInvokerEnrolmentDetails
    fmt.Fprintf(os.Stdout, "Response from `IndividualAPIInvokerEnrolmentDetailsApi.ModifyIndApiInvokeEnrolment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**onboardingId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyIndApiInvokeEnrolmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aPIInvokerEnrolmentDetailsPatch** | [**APIInvokerEnrolmentDetailsPatch**](APIInvokerEnrolmentDetailsPatch.md) |  | 

### Return type

[**APIInvokerEnrolmentDetails**](APIInvokerEnrolmentDetails.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

