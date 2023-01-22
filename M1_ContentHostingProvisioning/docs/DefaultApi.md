# \DefaultApi

All URIs are relative to *https://example.com/3gpp-m1/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateContentHostingConfiguration**](DefaultApi.md#CreateContentHostingConfiguration) | **Post** /provisioning-sessions/{provisioningSessionId}/content-hosting-configuration | Create (and optionally upload) the Content Hosting Configuration for the specified Provisioning Session
[**DestroyContentHostingConfiguration**](DefaultApi.md#DestroyContentHostingConfiguration) | **Delete** /provisioning-sessions/{provisioningSessionId}/content-hosting-configuration | Destroy the current Content Hosting Configuration of the specified Provisioning Session
[**PatchContentHostingConfiguration**](DefaultApi.md#PatchContentHostingConfiguration) | **Patch** /provisioning-sessions/{provisioningSessionId}/content-hosting-configuration | Patch the Content Hosting Configuration for the specified Provisioning Session
[**PurgeContentHostingCache**](DefaultApi.md#PurgeContentHostingCache) | **Post** /provisioning-sessions/{provisioningSessionId}/content-hosting-configuration/purge | Purge the content of the cache for the Content Hosting Configuration of the specified Provisioning Session
[**RetrieveContentHostingConfiguration**](DefaultApi.md#RetrieveContentHostingConfiguration) | **Get** /provisioning-sessions/{provisioningSessionId}/content-hosting-configuration | Retrieve the Content Hosting Configuration of the specified Provisioning Session
[**UpdateContentHostingConfiguration**](DefaultApi.md#UpdateContentHostingConfiguration) | **Put** /provisioning-sessions/{provisioningSessionId}/content-hosting-configuration | Update the Content Hosting Configuration for the specified Provisioning Session



## CreateContentHostingConfiguration

> CreateContentHostingConfiguration(ctx, provisioningSessionId).ContentHostingConfiguration(contentHostingConfiguration).Execute()

Create (and optionally upload) the Content Hosting Configuration for the specified Provisioning Session

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
    provisioningSessionId := "provisioningSessionId_example" // string | The resource identifier of an existing Provisioning Session.
    contentHostingConfiguration := *openapiclient.NewContentHostingConfiguration("Name_example", *openapiclient.NewIngestConfiguration(), []openapiclient.DistributionConfiguration{*openapiclient.NewDistributionConfiguration("CanonicalDomainName_example", "DomainNameAlias_example")}) // ContentHostingConfiguration | A JSON representation of a Content Hosting Configuration

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateContentHostingConfiguration(context.Background(), provisioningSessionId).ContentHostingConfiguration(contentHostingConfiguration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateContentHostingConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provisioningSessionId** | **string** | The resource identifier of an existing Provisioning Session. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateContentHostingConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentHostingConfiguration** | [**ContentHostingConfiguration**](ContentHostingConfiguration.md) | A JSON representation of a Content Hosting Configuration | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyContentHostingConfiguration

> DestroyContentHostingConfiguration(ctx, provisioningSessionId).Execute()

Destroy the current Content Hosting Configuration of the specified Provisioning Session

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
    provisioningSessionId := "provisioningSessionId_example" // string | The resource identifier of an existing Provisioning Session.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DestroyContentHostingConfiguration(context.Background(), provisioningSessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DestroyContentHostingConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provisioningSessionId** | **string** | The resource identifier of an existing Provisioning Session. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDestroyContentHostingConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchContentHostingConfiguration

> ContentHostingConfiguration PatchContentHostingConfiguration(ctx, provisioningSessionId).ContentHostingConfiguration(contentHostingConfiguration).Execute()

Patch the Content Hosting Configuration for the specified Provisioning Session

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
    provisioningSessionId := "provisioningSessionId_example" // string | The resource identifier of an existing Provisioning Session.
    contentHostingConfiguration := *openapiclient.NewContentHostingConfiguration("Name_example", *openapiclient.NewIngestConfiguration(), []openapiclient.DistributionConfiguration{*openapiclient.NewDistributionConfiguration("CanonicalDomainName_example", "DomainNameAlias_example")}) // ContentHostingConfiguration | A JSON representation of a Content Hosting Configuration

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.PatchContentHostingConfiguration(context.Background(), provisioningSessionId).ContentHostingConfiguration(contentHostingConfiguration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.PatchContentHostingConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchContentHostingConfiguration`: ContentHostingConfiguration
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.PatchContentHostingConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provisioningSessionId** | **string** | The resource identifier of an existing Provisioning Session. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchContentHostingConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentHostingConfiguration** | [**ContentHostingConfiguration**](ContentHostingConfiguration.md) | A JSON representation of a Content Hosting Configuration | 

### Return type

[**ContentHostingConfiguration**](ContentHostingConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/merge-patch+json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PurgeContentHostingCache

> PurgeContentHostingCache(ctx, provisioningSessionId).Pattern(pattern).Value(value).Execute()

Purge the content of the cache for the Content Hosting Configuration of the specified Provisioning Session

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
    provisioningSessionId := "provisioningSessionId_example" // string | A unique identifier of the Provisioning
    pattern := "pattern_example" // string | Keyword (optional)
    value := "value_example" // string | The regular expression (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.PurgeContentHostingCache(context.Background(), provisioningSessionId).Pattern(pattern).Value(value).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.PurgeContentHostingCache``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provisioningSessionId** | **string** | A unique identifier of the Provisioning | 

### Other Parameters

Other parameters are passed through a pointer to a apiPurgeContentHostingCacheRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pattern** | **string** | Keyword | 
 **value** | **string** | The regular expression | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveContentHostingConfiguration

> ContentHostingConfiguration RetrieveContentHostingConfiguration(ctx, provisioningSessionId).Execute()

Retrieve the Content Hosting Configuration of the specified Provisioning Session

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
    provisioningSessionId := "provisioningSessionId_example" // string | The resource identifier of an existing Provisioning Session.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.RetrieveContentHostingConfiguration(context.Background(), provisioningSessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RetrieveContentHostingConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveContentHostingConfiguration`: ContentHostingConfiguration
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.RetrieveContentHostingConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provisioningSessionId** | **string** | The resource identifier of an existing Provisioning Session. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveContentHostingConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ContentHostingConfiguration**](ContentHostingConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateContentHostingConfiguration

> UpdateContentHostingConfiguration(ctx, provisioningSessionId).ContentHostingConfiguration(contentHostingConfiguration).Execute()

Update the Content Hosting Configuration for the specified Provisioning Session

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
    provisioningSessionId := "provisioningSessionId_example" // string | The resource identifier of an existing Provisioning Session.
    contentHostingConfiguration := *openapiclient.NewContentHostingConfiguration("Name_example", *openapiclient.NewIngestConfiguration(), []openapiclient.DistributionConfiguration{*openapiclient.NewDistributionConfiguration("CanonicalDomainName_example", "DomainNameAlias_example")}) // ContentHostingConfiguration | A JSON representation of a Content Hosting Configuration

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateContentHostingConfiguration(context.Background(), provisioningSessionId).ContentHostingConfiguration(contentHostingConfiguration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateContentHostingConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provisioningSessionId** | **string** | The resource identifier of an existing Provisioning Session. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateContentHostingConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentHostingConfiguration** | [**ContentHostingConfiguration**](ContentHostingConfiguration.md) | A JSON representation of a Content Hosting Configuration | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

