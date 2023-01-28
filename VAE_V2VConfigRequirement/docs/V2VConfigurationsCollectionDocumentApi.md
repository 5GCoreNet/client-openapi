# \V2VConfigurationsCollectionDocumentApi

All URIs are relative to *https://example.com/vae-v2v-config-req/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](V2VConfigurationsCollectionDocumentApi.md#Create) | **Post** /configurations | VAE V2V Configuration resource create service Operation



## Create

> V2vConfigurationData Create(ctx).V2vConfigurationData(v2vConfigurationData).Execute()

VAE V2V Configuration resource create service Operation

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/VAE_V2VConfigRequirement"
)

func main() {
    v2vConfigurationData := *openapiclient.NewV2vConfigurationData() // V2vConfigurationData | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.V2VConfigurationsCollectionDocumentApi.Create(context.Background()).V2vConfigurationData(v2vConfigurationData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `V2VConfigurationsCollectionDocumentApi.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create`: V2vConfigurationData
    fmt.Fprintf(os.Stdout, "Response from `V2VConfigurationsCollectionDocumentApi.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v2vConfigurationData** | [**V2vConfigurationData**](V2vConfigurationData.md) |  | 

### Return type

[**V2vConfigurationData**](V2vConfigurationData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

