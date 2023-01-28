# \IndividualASTIConfigurationDocumentApi

All URIs are relative to *https://example.com/ntsctsf-asti/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIndividualASTIConfiguration**](IndividualASTIConfigurationDocumentApi.md#DeleteIndividualASTIConfiguration) | **Delete** /configurations/{configId} | Delete an Individual ASTI Configuration
[**ModifyIndividualASTIConfiguration**](IndividualASTIConfigurationDocumentApi.md#ModifyIndividualASTIConfiguration) | **Put** /configurations/{configId} | Modifies an existing Individual ASTI Configuration resource.



## DeleteIndividualASTIConfiguration

> DeleteIndividualASTIConfiguration(ctx, configId).Execute()

Delete an Individual ASTI Configuration

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Ntsctsf_ASTI"
)

func main() {
    configId := "configId_example" // string | String identifying an Individual ASTI Configuration.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualASTIConfigurationDocumentApi.DeleteIndividualASTIConfiguration(context.Background(), configId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualASTIConfigurationDocumentApi.DeleteIndividualASTIConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configId** | **string** | String identifying an Individual ASTI Configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIndividualASTIConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyIndividualASTIConfiguration

> AccessTimeDistributionData ModifyIndividualASTIConfiguration(ctx, configId).AccessTimeDistributionData(accessTimeDistributionData).Execute()

Modifies an existing Individual ASTI Configuration resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Ntsctsf_ASTI"
)

func main() {
    configId := "configId_example" // string | String identifying an Individual ASTI Configuration.
    accessTimeDistributionData := openapiclient.AccessTimeDistributionData{Interface{}: new(interface{})} // AccessTimeDistributionData | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualASTIConfigurationDocumentApi.ModifyIndividualASTIConfiguration(context.Background(), configId).AccessTimeDistributionData(accessTimeDistributionData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualASTIConfigurationDocumentApi.ModifyIndividualASTIConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ModifyIndividualASTIConfiguration`: AccessTimeDistributionData
    fmt.Fprintf(os.Stdout, "Response from `IndividualASTIConfigurationDocumentApi.ModifyIndividualASTIConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configId** | **string** | String identifying an Individual ASTI Configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyIndividualASTIConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessTimeDistributionData** | [**AccessTimeDistributionData**](AccessTimeDistributionData.md) |  | 

### Return type

[**AccessTimeDistributionData**](AccessTimeDistributionData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

