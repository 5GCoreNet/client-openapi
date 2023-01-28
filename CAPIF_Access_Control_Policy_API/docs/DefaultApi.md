# \DefaultApi

All URIs are relative to *https://example.com/access-control-policy/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccessControlPolicyListServiceApiIdGet**](DefaultApi.md#AccessControlPolicyListServiceApiIdGet) | **Get** /accessControlPolicyList/{serviceApiId} | 



## AccessControlPolicyListServiceApiIdGet

> AccessControlPolicyList AccessControlPolicyListServiceApiIdGet(ctx, serviceApiId).AefId(aefId).ApiInvokerId(apiInvokerId).SupportedFeatures(supportedFeatures).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/CAPIF_Access_Control_Policy_API"
)

func main() {
    serviceApiId := "serviceApiId_example" // string | Identifier of a published service API
    aefId := "aefId_example" // string | Identifier of the AEF
    apiInvokerId := "apiInvokerId_example" // string | Identifier of the API invoker (optional)
    supportedFeatures := "supportedFeatures_example" // string | To filter irrelevant responses related to unsupported features (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.AccessControlPolicyListServiceApiIdGet(context.Background(), serviceApiId).AefId(aefId).ApiInvokerId(apiInvokerId).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AccessControlPolicyListServiceApiIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccessControlPolicyListServiceApiIdGet`: AccessControlPolicyList
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AccessControlPolicyListServiceApiIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceApiId** | **string** | Identifier of a published service API | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccessControlPolicyListServiceApiIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aefId** | **string** | Identifier of the AEF | 
 **apiInvokerId** | **string** | Identifier of the API invoker | 
 **supportedFeatures** | **string** | To filter irrelevant responses related to unsupported features | 

### Return type

[**AccessControlPolicyList**](AccessControlPolicyList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

