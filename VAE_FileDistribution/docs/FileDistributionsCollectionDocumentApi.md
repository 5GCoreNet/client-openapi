# \FileDistributionsCollectionDocumentApi

All URIs are relative to *https://example.com/vae-file-distribution/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFileDistributions**](FileDistributionsCollectionDocumentApi.md#CreateFileDistributions) | **Post** /file-distributions | VAE File Distributions resource create service Operation



## CreateFileDistributions

> FileDistributionData CreateFileDistributions(ctx).FileDistributionData(fileDistributionData).Execute()

VAE File Distributions resource create service Operation

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/5GCoreNet/client-openapi/VAE_FileDistribution"
)

func main() {
    fileDistributionData := *openapiclient.NewFileDistributionData([]openapiclient.FileList{*openapiclient.NewFileList("FileUri_example", "FileDisplayUri_example", time.Now(), time.Now(), *openapiclient.NewFileStatus(), time.Now(), int32(123))}, *openapiclient.NewGeographicArea(*openapiclient.NewSupportedGADShapes(), *openapiclient.NewGeographicalCoordinates(float64(123), float64(123)), float32(123), *openapiclient.NewUncertaintyEllipse(float32(123), float32(123), int32(123)), int32(123), []openapiclient.GeographicalCoordinates{*openapiclient.NewGeographicalCoordinates(float64(123), float64(123))}, float64(123), float32(123), int32(123), float32(123), int32(123), int32(123)), "MaxBitrate_example", int32(123)) // FileDistributionData | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FileDistributionsCollectionDocumentApi.CreateFileDistributions(context.Background()).FileDistributionData(fileDistributionData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FileDistributionsCollectionDocumentApi.CreateFileDistributions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFileDistributions`: FileDistributionData
    fmt.Fprintf(os.Stdout, "Response from `FileDistributionsCollectionDocumentApi.CreateFileDistributions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFileDistributionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fileDistributionData** | [**FileDistributionData**](FileDistributionData.md) |  | 

### Return type

[**FileDistributionData**](FileDistributionData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

