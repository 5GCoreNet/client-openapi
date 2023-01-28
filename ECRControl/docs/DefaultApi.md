# \DefaultApi

All URIs are relative to *https://example.com/3gpp-ecr-control/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConfigurePost**](DefaultApi.md#ConfigurePost) | **Post** /configure | Configure the enhanced converage restriction for a UE.
[**QueryPost**](DefaultApi.md#QueryPost) | **Post** /query | Query the status of enhanced converage restriction for a UE.



## ConfigurePost

> ECRData ConfigurePost(ctx).ECRControl(eCRControl).Execute()

Configure the enhanced converage restriction for a UE.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/ECRControl"
)

func main() {
    eCRControl := openapiclient.ECRControl{Interface{}: new(interface{})} // ECRControl | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ConfigurePost(context.Background()).ECRControl(eCRControl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ConfigurePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConfigurePost`: ECRData
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ConfigurePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConfigurePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eCRControl** | [**ECRControl**](ECRControl.md) |  | 

### Return type

[**ECRData**](ECRData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryPost

> ECRData QueryPost(ctx).ECRControl(eCRControl).Execute()

Query the status of enhanced converage restriction for a UE.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/ECRControl"
)

func main() {
    eCRControl := openapiclient.ECRControl{Interface{}: new(interface{})} // ECRControl | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.QueryPost(context.Background()).ECRControl(eCRControl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.QueryPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryPost`: ECRData
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.QueryPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eCRControl** | [**ECRControl**](ECRControl.md) |  | 

### Return type

[**ECRData**](ECRData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

