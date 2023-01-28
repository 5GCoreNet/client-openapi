# \IndividualEASDeploymentInformationApi

All URIs are relative to *https://example.com/3gpp-eas-deployment/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAnDeployment**](IndividualEASDeploymentInformationApi.md#DeleteAnDeployment) | **Delete** /{afId}/eas-deployment-info/{easDeployInfoId} | Deletes an already existing EAS Deployment information resource
[**FullyUpdateAnDeployment**](IndividualEASDeploymentInformationApi.md#FullyUpdateAnDeployment) | **Put** /{afId}/eas-deployment-info/{easDeployInfoId} | Fully updates/replaces an existing resource
[**ReadAnDeployment**](IndividualEASDeploymentInformationApi.md#ReadAnDeployment) | **Get** /{afId}/eas-deployment-info/{easDeployInfoId} | Read an active Individual EAS Deployment Information resource for the AF



## DeleteAnDeployment

> DeleteAnDeployment(ctx, afId, easDeployInfoId).Execute()

Deletes an already existing EAS Deployment information resource

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
    easDeployInfoId := "easDeployInfoId_example" // string | Identifier of the EAS Deployment information resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualEASDeploymentInformationApi.DeleteAnDeployment(context.Background(), afId, easDeployInfoId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualEASDeploymentInformationApi.DeleteAnDeployment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**afId** | **string** | Identifier of the AF | 
**easDeployInfoId** | **string** | Identifier of the EAS Deployment information resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAnDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FullyUpdateAnDeployment

> EasDeployInfo FullyUpdateAnDeployment(ctx, afId, easDeployInfoId).EasDeployInfo(easDeployInfo).Execute()

Fully updates/replaces an existing resource

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
    easDeployInfoId := "easDeployInfoId_example" // string | Identifier of the EAS Deployment information resource
    easDeployInfo := *openapiclient.NewEasDeployInfo([]openapiclient.FqdnPatternMatchingRule{openapiclient.FqdnPatternMatchingRule{Interface{}: new(interface{})}}) // EasDeployInfo | Parameters to update/replace the existing resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualEASDeploymentInformationApi.FullyUpdateAnDeployment(context.Background(), afId, easDeployInfoId).EasDeployInfo(easDeployInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualEASDeploymentInformationApi.FullyUpdateAnDeployment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FullyUpdateAnDeployment`: EasDeployInfo
    fmt.Fprintf(os.Stdout, "Response from `IndividualEASDeploymentInformationApi.FullyUpdateAnDeployment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**afId** | **string** | Identifier of the AF | 
**easDeployInfoId** | **string** | Identifier of the EAS Deployment information resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiFullyUpdateAnDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **easDeployInfo** | [**EasDeployInfo**](EasDeployInfo.md) | Parameters to update/replace the existing resource | 

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


## ReadAnDeployment

> EasDeployInfo ReadAnDeployment(ctx, afId, easDeployInfoId).Execute()

Read an active Individual EAS Deployment Information resource for the AF

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
    easDeployInfoId := "easDeployInfoId_example" // string | Identifier of an EAS Deployment Information.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualEASDeploymentInformationApi.ReadAnDeployment(context.Background(), afId, easDeployInfoId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualEASDeploymentInformationApi.ReadAnDeployment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAnDeployment`: EasDeployInfo
    fmt.Fprintf(os.Stdout, "Response from `IndividualEASDeploymentInformationApi.ReadAnDeployment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**afId** | **string** | Identifier of the AF | 
**easDeployInfoId** | **string** | Identifier of an EAS Deployment Information. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAnDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EasDeployInfo**](EasDeployInfo.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

