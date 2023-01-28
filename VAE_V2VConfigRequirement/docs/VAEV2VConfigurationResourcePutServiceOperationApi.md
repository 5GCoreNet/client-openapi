# \VAEV2VConfigurationResourcePutServiceOperationApi

All URIs are relative to *https://example.com/vae-v2v-config-req/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateV2VConfiguration**](VAEV2VConfigurationResourcePutServiceOperationApi.md#UpdateV2VConfiguration) | **Put** /configurations/{configurationId} | Updates/replaces an existing configuration resource



## UpdateV2VConfiguration

> V2vConfigurationData UpdateV2VConfiguration(ctx, configurationId).V2vConfigurationData(v2vConfigurationData).Execute()

Updates/replaces an existing configuration resource

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
    configurationId := "configurationId_example" // string | Identifier of a V2V Configuration resource
    v2vConfigurationData := *openapiclient.NewV2vConfigurationData() // V2vConfigurationData | Parameters to update/replace the existing configuration

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VAEV2VConfigurationResourcePutServiceOperationApi.UpdateV2VConfiguration(context.Background(), configurationId).V2vConfigurationData(v2vConfigurationData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VAEV2VConfigurationResourcePutServiceOperationApi.UpdateV2VConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateV2VConfiguration`: V2vConfigurationData
    fmt.Fprintf(os.Stdout, "Response from `VAEV2VConfigurationResourcePutServiceOperationApi.UpdateV2VConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configurationId** | **string** | Identifier of a V2V Configuration resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateV2VConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v2vConfigurationData** | [**V2vConfigurationData**](V2vConfigurationData.md) | Parameters to update/replace the existing configuration | 

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

