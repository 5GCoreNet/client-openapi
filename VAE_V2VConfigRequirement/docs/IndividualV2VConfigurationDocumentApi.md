# \IndividualV2VConfigurationDocumentApi

All URIs are relative to *https://example.com/vae-v2v-config-req/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteV2VConfiguration**](IndividualV2VConfigurationDocumentApi.md#DeleteV2VConfiguration) | **Delete** /configurations/{configurationId} | VAE V2V Configuration resource delete service Operation
[**ReadV2VConfiguration**](IndividualV2VConfigurationDocumentApi.md#ReadV2VConfiguration) | **Get** /configurations/{configurationId} | VAE V2V Configuration resource read service Operation



## DeleteV2VConfiguration

> DeleteV2VConfiguration(ctx, configurationId).Execute()

VAE V2V Configuration resource delete service Operation

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
    configurationId := "configurationId_example" // string | Unique ID of the V2V Configuration to be deleted

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualV2VConfigurationDocumentApi.DeleteV2VConfiguration(context.Background(), configurationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualV2VConfigurationDocumentApi.DeleteV2VConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configurationId** | **string** | Unique ID of the V2V Configuration to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteV2VConfigurationRequest struct via the builder pattern


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


## ReadV2VConfiguration

> V2vConfigurationData ReadV2VConfiguration(ctx, configurationId).Execute()

VAE V2V Configuration resource read service Operation

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
    configurationId := "configurationId_example" // string | Identifier of a V2V Configuration resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualV2VConfigurationDocumentApi.ReadV2VConfiguration(context.Background(), configurationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualV2VConfigurationDocumentApi.ReadV2VConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadV2VConfiguration`: V2vConfigurationData
    fmt.Fprintf(os.Stdout, "Response from `IndividualV2VConfigurationDocumentApi.ReadV2VConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configurationId** | **string** | Identifier of a V2V Configuration resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadV2VConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V2vConfigurationData**](V2vConfigurationData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

