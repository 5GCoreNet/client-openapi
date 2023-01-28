# \EASDeploymentInformationCollectionApi

All URIs are relative to *https://example.com/3gpp-eas-deployment/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAnDeployment**](EASDeploymentInformationCollectionApi.md#CreateAnDeployment) | **Post** /{afId}/eas-deployment-info | Create a new Individual EAS Deployment information resource.
[**ReadAllDeployment**](EASDeploymentInformationCollectionApi.md#ReadAllDeployment) | **Get** /{afId}/eas-deployment-info | Read all EAS Deployment information for a given AF



## CreateAnDeployment

> EasDeployInfo CreateAnDeployment(ctx, afId).EasDeployInfo(easDeployInfo).Execute()

Create a new Individual EAS Deployment information resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/EASDeployment"
)

func main() {
    afId := "afId_example" // string | Identifier of the AF
    easDeployInfo := *openapiclient.NewEasDeployInfo([]openapiclient.FqdnPatternMatchingRule{openapiclient.FqdnPatternMatchingRule{Interface{}: new(interface{})}}) // EasDeployInfo | new resource creation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EASDeploymentInformationCollectionApi.CreateAnDeployment(context.Background(), afId).EasDeployInfo(easDeployInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EASDeploymentInformationCollectionApi.CreateAnDeployment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAnDeployment`: EasDeployInfo
    fmt.Fprintf(os.Stdout, "Response from `EASDeploymentInformationCollectionApi.CreateAnDeployment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**afId** | **string** | Identifier of the AF | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAnDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **easDeployInfo** | [**EasDeployInfo**](EasDeployInfo.md) | new resource creation | 

### Return type

[**EasDeployInfo**](EasDeployInfo.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadAllDeployment

> []EasDeployInfo ReadAllDeployment(ctx, afId).Execute()

Read all EAS Deployment information for a given AF

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/EASDeployment"
)

func main() {
    afId := "afId_example" // string | Identifier of the AF

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EASDeploymentInformationCollectionApi.ReadAllDeployment(context.Background(), afId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EASDeploymentInformationCollectionApi.ReadAllDeployment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllDeployment`: []EasDeployInfo
    fmt.Fprintf(os.Stdout, "Response from `EASDeploymentInformationCollectionApi.ReadAllDeployment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**afId** | **string** | Identifier of the AF | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAllDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]EasDeployInfo**](EasDeployInfo.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

