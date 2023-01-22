# \ASTIConfigurationsCollectionApi

All URIs are relative to *https://example.com/ntsctsf-asti/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ASTIConfiguration**](ASTIConfigurationsCollectionApi.md#ASTIConfiguration) | **Post** /configurations | Creates a new Individual ASTI Configuration resource.



## ASTIConfiguration

> AccessTimeDistributionData ASTIConfiguration(ctx).AccessTimeDistributionData(accessTimeDistributionData).Execute()

Creates a new Individual ASTI Configuration resource.

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
    accessTimeDistributionData := openapiclient.AccessTimeDistributionData{Interface{}: new(interface{})} // AccessTimeDistributionData | Contains the information for the creation the resource.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ASTIConfigurationsCollectionApi.ASTIConfiguration(context.Background()).AccessTimeDistributionData(accessTimeDistributionData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ASTIConfigurationsCollectionApi.ASTIConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ASTIConfiguration`: AccessTimeDistributionData
    fmt.Fprintf(os.Stdout, "Response from `ASTIConfigurationsCollectionApi.ASTIConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiASTIConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessTimeDistributionData** | [**AccessTimeDistributionData**](AccessTimeDistributionData.md) | Contains the information for the creation the resource. | 

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

