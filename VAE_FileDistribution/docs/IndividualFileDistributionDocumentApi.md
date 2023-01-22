# \IndividualFileDistributionDocumentApi

All URIs are relative to *https://example.com/vae-file-distribution/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFileDistribution**](IndividualFileDistributionDocumentApi.md#DeleteFileDistribution) | **Delete** /file-distributions/{distributionId} | VAE File Distribution resource delete service Operation
[**ReadIndividualFileDistribution**](IndividualFileDistributionDocumentApi.md#ReadIndividualFileDistribution) | **Get** /file-distributions/{distributionId} | Get an existing individual file distribution resource



## DeleteFileDistribution

> DeleteFileDistribution(ctx, distributionId).Execute()

VAE File Distribution resource delete service Operation

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
    distributionId := "distributionId_example" // string | Unique ID of the file distribution to be deleted

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualFileDistributionDocumentApi.DeleteFileDistribution(context.Background(), distributionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualFileDistributionDocumentApi.DeleteFileDistribution``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**distributionId** | **string** | Unique ID of the file distribution to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFileDistributionRequest struct via the builder pattern


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


## ReadIndividualFileDistribution

> FileDistributionData ReadIndividualFileDistribution(ctx, distributionId).Execute()

Get an existing individual file distribution resource

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
    distributionId := "distributionId_example" // string | Identifier of a file distribution resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualFileDistributionDocumentApi.ReadIndividualFileDistribution(context.Background(), distributionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualFileDistributionDocumentApi.ReadIndividualFileDistribution``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadIndividualFileDistribution`: FileDistributionData
    fmt.Fprintf(os.Stdout, "Response from `IndividualFileDistributionDocumentApi.ReadIndividualFileDistribution`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**distributionId** | **string** | Identifier of a file distribution resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadIndividualFileDistributionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FileDistributionData**](FileDistributionData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

