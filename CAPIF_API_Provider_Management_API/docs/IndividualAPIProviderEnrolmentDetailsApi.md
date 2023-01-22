# \IndividualAPIProviderEnrolmentDetailsApi

All URIs are relative to *https://example.com/api-provider-management/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ModifyIndApiProviderEnrolment**](IndividualAPIProviderEnrolmentDetailsApi.md#ModifyIndApiProviderEnrolment) | **Patch** /registrations/{registrationId} | 



## ModifyIndApiProviderEnrolment

> APIProviderEnrolmentDetails ModifyIndApiProviderEnrolment(ctx, registrationId).APIProviderEnrolmentDetailsPatch(aPIProviderEnrolmentDetailsPatch).Execute()





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
    registrationId := "registrationId_example" // string | 
    aPIProviderEnrolmentDetailsPatch := *openapiclient.NewAPIProviderEnrolmentDetailsPatch() // APIProviderEnrolmentDetailsPatch | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualAPIProviderEnrolmentDetailsApi.ModifyIndApiProviderEnrolment(context.Background(), registrationId).APIProviderEnrolmentDetailsPatch(aPIProviderEnrolmentDetailsPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualAPIProviderEnrolmentDetailsApi.ModifyIndApiProviderEnrolment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ModifyIndApiProviderEnrolment`: APIProviderEnrolmentDetails
    fmt.Fprintf(os.Stdout, "Response from `IndividualAPIProviderEnrolmentDetailsApi.ModifyIndApiProviderEnrolment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registrationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyIndApiProviderEnrolmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aPIProviderEnrolmentDetailsPatch** | [**APIProviderEnrolmentDetailsPatch**](APIProviderEnrolmentDetailsPatch.md) |  | 

### Return type

[**APIProviderEnrolmentDetails**](APIProviderEnrolmentDetails.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

