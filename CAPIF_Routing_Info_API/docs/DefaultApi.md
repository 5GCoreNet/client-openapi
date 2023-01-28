# \DefaultApi

All URIs are relative to *https://example.com/capif-routing-info/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ServiceApisServiceApiIdGet**](DefaultApi.md#ServiceApisServiceApiIdGet) | **Get** /service-apis/{serviceApiId} | 



## ServiceApisServiceApiIdGet

> RoutingInfo ServiceApisServiceApiIdGet(ctx, serviceApiId).AefId(aefId).SuppFeat(suppFeat).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/CAPIF_Routing_Info_API"
)

func main() {
    serviceApiId := "serviceApiId_example" // string | Identifier of a published service API
    aefId := "aefId_example" // string | Identifier of the AEF
    suppFeat := "suppFeat_example" // string | To filter irrelevant responses related to unsupported features (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ServiceApisServiceApiIdGet(context.Background(), serviceApiId).AefId(aefId).SuppFeat(suppFeat).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ServiceApisServiceApiIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServiceApisServiceApiIdGet`: RoutingInfo
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ServiceApisServiceApiIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceApiId** | **string** | Identifier of a published service API | 

### Other Parameters

Other parameters are passed through a pointer to a apiServiceApisServiceApiIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aefId** | **string** | Identifier of the AEF | 
 **suppFeat** | **string** | To filter irrelevant responses related to unsupported features | 

### Return type

[**RoutingInfo**](RoutingInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

