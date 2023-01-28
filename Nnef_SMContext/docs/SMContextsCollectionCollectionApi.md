# \SMContextsCollectionCollectionApi

All URIs are relative to *https://example.com/nnef-smcontext/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](SMContextsCollectionCollectionApi.md#Create) | **Post** /sm-contexts | Create SM Context



## Create

> SmContextCreatedData Create(ctx).SmContextCreateData(smContextCreateData).Execute()

Create SM Context

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nnef_SMContext"
)

func main() {
    smContextCreateData := *openapiclient.NewSmContextCreateData("Supi_example", int32(123), "Dnn_example", *openapiclient.NewSnssai(int32(123)), "NefId_example", "DlNiddEndPoint_example", "NotificationUri_example") // SmContextCreateData | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SMContextsCollectionCollectionApi.Create(context.Background()).SmContextCreateData(smContextCreateData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SMContextsCollectionCollectionApi.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create`: SmContextCreatedData
    fmt.Fprintf(os.Stdout, "Response from `SMContextsCollectionCollectionApi.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **smContextCreateData** | [**SmContextCreateData**](SmContextCreateData.md) |  | 

### Return type

[**SmContextCreatedData**](SmContextCreatedData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

