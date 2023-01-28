# \IPTVConfigurationsApi

All URIs are relative to *https://example.com/3gpp-iptvconfiguration/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNewSubscription**](IPTVConfigurationsApi.md#CreateNewSubscription) | **Post** /{afId}/configurations | Creates a new configuration resource
[**ReadAllSubscriptions**](IPTVConfigurationsApi.md#ReadAllSubscriptions) | **Get** /{afId}/configurations | read all of the active configurations for the AF



## CreateNewSubscription

> IptvConfigData CreateNewSubscription(ctx, afId).IptvConfigData(iptvConfigData).Execute()

Creates a new configuration resource

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/IPTVConfiguration"
)

func main() {
    afId := "afId_example" // string | Identifier of the AF
    iptvConfigData := *openapiclient.NewIptvConfigData("AfAppId_example", map[string]MulticastAccessControl{"key": *openapiclient.NewMulticastAccessControl(*openapiclient.NewAccessRightStatus())}, "SuppFeat_example") // IptvConfigData | new configuration creation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IPTVConfigurationsApi.CreateNewSubscription(context.Background(), afId).IptvConfigData(iptvConfigData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPTVConfigurationsApi.CreateNewSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNewSubscription`: IptvConfigData
    fmt.Fprintf(os.Stdout, "Response from `IPTVConfigurationsApi.CreateNewSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**afId** | **string** | Identifier of the AF | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNewSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iptvConfigData** | [**IptvConfigData**](IptvConfigData.md) | new configuration creation | 

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


## ReadAllSubscriptions

> []IptvConfigData ReadAllSubscriptions(ctx, afId).Execute()

read all of the active configurations for the AF

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/IPTVConfiguration"
)

func main() {
    afId := "afId_example" // string | Identifier of the AF

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IPTVConfigurationsApi.ReadAllSubscriptions(context.Background(), afId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPTVConfigurationsApi.ReadAllSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllSubscriptions`: []IptvConfigData
    fmt.Fprintf(os.Stdout, "Response from `IPTVConfigurationsApi.ReadAllSubscriptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**afId** | **string** | Identifier of the AF | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAllSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]IptvConfigData**](IptvConfigData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

