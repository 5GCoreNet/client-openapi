# \ApplicationAMContextsApi

All URIs are relative to *https://example.com/3gpp-am-policyauthorization/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostAppAmContexts**](ApplicationAMContextsApi.md#PostAppAmContexts) | **Post** /{afId}/app-am-contexts | Creates a new Individual application AM Context resource



## PostAppAmContexts

> AppAmContextExpRespData PostAppAmContexts(ctx, afId).AppAmContextExpData(appAmContextExpData).Execute()

Creates a new Individual application AM Context resource

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/AMPolicyAuthorization"
)

func main() {
    afId := "afId_example" // string | Identifier of the AF
    appAmContextExpData := *openapiclient.NewAppAmContextExpData("Gpsi_example") // AppAmContextExpData | new resource creation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationAMContextsApi.PostAppAmContexts(context.Background(), afId).AppAmContextExpData(appAmContextExpData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAMContextsApi.PostAppAmContexts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostAppAmContexts`: AppAmContextExpRespData
    fmt.Fprintf(os.Stdout, "Response from `ApplicationAMContextsApi.PostAppAmContexts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**afId** | **string** | Identifier of the AF | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAppAmContextsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appAmContextExpData** | [**AppAmContextExpData**](AppAmContextExpData.md) | new resource creation | 

### Return type

[**AppAmContextExpRespData**](AppAmContextExpRespData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

