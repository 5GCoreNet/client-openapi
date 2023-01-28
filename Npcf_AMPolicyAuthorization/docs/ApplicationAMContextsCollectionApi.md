# \ApplicationAMContextsCollectionApi

All URIs are relative to *https://example.com/npcf-am-policyauthorization/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostAppAmContexts**](ApplicationAMContextsCollectionApi.md#PostAppAmContexts) | **Post** /app-am-contexts | Creates a new Individual Application AM Context resource



## PostAppAmContexts

> AppAmContextRespData PostAppAmContexts(ctx).AppAmContextData(appAmContextData).Execute()

Creates a new Individual Application AM Context resource

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Npcf_AMPolicyAuthorization"
)

func main() {
    appAmContextData := *openapiclient.NewAppAmContextData("Supi_example", "TermNotifUri_example") // AppAmContextData | Contains the information for the creation the resource.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationAMContextsCollectionApi.PostAppAmContexts(context.Background()).AppAmContextData(appAmContextData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAMContextsCollectionApi.PostAppAmContexts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostAppAmContexts`: AppAmContextRespData
    fmt.Fprintf(os.Stdout, "Response from `ApplicationAMContextsCollectionApi.PostAppAmContexts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAppAmContextsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appAmContextData** | [**AppAmContextData**](AppAmContextData.md) | Contains the information for the creation the resource. | 

### Return type

[**AppAmContextRespData**](AppAmContextRespData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

