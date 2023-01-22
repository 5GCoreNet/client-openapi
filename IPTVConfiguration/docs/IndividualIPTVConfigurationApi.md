# \IndividualIPTVConfigurationApi

All URIs are relative to *https://example.com/3gpp-iptvconfiguration/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAnSubscription**](IndividualIPTVConfigurationApi.md#DeleteAnSubscription) | **Delete** /{afId}/configurations/{configurationId} | Deletes an already existing configuration
[**FullyUpdateAnSubscription**](IndividualIPTVConfigurationApi.md#FullyUpdateAnSubscription) | **Put** /{afId}/configurations/{configurationId} | Fully updates/replaces an existing configuration resource
[**PartialUpdateAnSubscription**](IndividualIPTVConfigurationApi.md#PartialUpdateAnSubscription) | **Patch** /{afId}/configurations/{configurationId} | Partial updates an existing configuration resource
[**ReadAnSubscription**](IndividualIPTVConfigurationApi.md#ReadAnSubscription) | **Get** /{afId}/configurations/{configurationId} | read an active configuration for the AF and the configuration Id



## DeleteAnSubscription

> DeleteAnSubscription(ctx, afId, configurationId).Execute()

Deletes an already existing configuration

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
    afId := "afId_example" // string | Identifier of the AF
    configurationId := "configurationId_example" // string | Identifier of the configuration resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualIPTVConfigurationApi.DeleteAnSubscription(context.Background(), afId, configurationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualIPTVConfigurationApi.DeleteAnSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**afId** | **string** | Identifier of the AF | 
**configurationId** | **string** | Identifier of the configuration resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAnSubscriptionRequest struct via the builder pattern


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


## FullyUpdateAnSubscription

> IptvConfigData FullyUpdateAnSubscription(ctx, afId, configurationId).IptvConfigData(iptvConfigData).Execute()

Fully updates/replaces an existing configuration resource

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
    afId := "afId_example" // string | Identifier of the AF
    configurationId := "configurationId_example" // string | Identifier of the configuration resource
    iptvConfigData := *openapiclient.NewIptvConfigData("AfAppId_example", map[string]MulticastAccessControl{"key": *openapiclient.NewMulticastAccessControl(*openapiclient.NewAccessRightStatus())}, "SuppFeat_example") // IptvConfigData | Parameters to update/replace the existing configuration

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualIPTVConfigurationApi.FullyUpdateAnSubscription(context.Background(), afId, configurationId).IptvConfigData(iptvConfigData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualIPTVConfigurationApi.FullyUpdateAnSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FullyUpdateAnSubscription`: IptvConfigData
    fmt.Fprintf(os.Stdout, "Response from `IndividualIPTVConfigurationApi.FullyUpdateAnSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**afId** | **string** | Identifier of the AF | 
**configurationId** | **string** | Identifier of the configuration resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiFullyUpdateAnSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **iptvConfigData** | [**IptvConfigData**](IptvConfigData.md) | Parameters to update/replace the existing configuration | 

### Return type

[**IptvConfigData**](IptvConfigData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartialUpdateAnSubscription

> IptvConfigData PartialUpdateAnSubscription(ctx, afId, configurationId).IptvConfigDataPatch(iptvConfigDataPatch).Execute()

Partial updates an existing configuration resource

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
    afId := "afId_example" // string | Identifier of the AF
    configurationId := "configurationId_example" // string | Identifier of the configuration resource
    iptvConfigDataPatch := *openapiclient.NewIptvConfigDataPatch() // IptvConfigDataPatch | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualIPTVConfigurationApi.PartialUpdateAnSubscription(context.Background(), afId, configurationId).IptvConfigDataPatch(iptvConfigDataPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualIPTVConfigurationApi.PartialUpdateAnSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PartialUpdateAnSubscription`: IptvConfigData
    fmt.Fprintf(os.Stdout, "Response from `IndividualIPTVConfigurationApi.PartialUpdateAnSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**afId** | **string** | Identifier of the AF | 
**configurationId** | **string** | Identifier of the configuration resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartialUpdateAnSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **iptvConfigDataPatch** | [**IptvConfigDataPatch**](IptvConfigDataPatch.md) |  | 

### Return type

[**IptvConfigData**](IptvConfigData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadAnSubscription

> IptvConfigData ReadAnSubscription(ctx, afId, configurationId).Execute()

read an active configuration for the AF and the configuration Id

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
    afId := "afId_example" // string | Identifier of the AF
    configurationId := "configurationId_example" // string | Identifier of the configuration resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualIPTVConfigurationApi.ReadAnSubscription(context.Background(), afId, configurationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualIPTVConfigurationApi.ReadAnSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAnSubscription`: IptvConfigData
    fmt.Fprintf(os.Stdout, "Response from `IndividualIPTVConfigurationApi.ReadAnSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**afId** | **string** | Identifier of the AF | 
**configurationId** | **string** | Identifier of the configuration resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAnSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**IptvConfigData**](IptvConfigData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

