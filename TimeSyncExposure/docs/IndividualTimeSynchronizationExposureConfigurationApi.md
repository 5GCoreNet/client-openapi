# \IndividualTimeSynchronizationExposureConfigurationApi

All URIs are relative to *https://example.com/3gpp-time-sync/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAnConfiguration**](IndividualTimeSynchronizationExposureConfigurationApi.md#DeleteAnConfiguration) | **Delete** /{afId}/subscriptions/{subscriptionId}/configurations/{instanceReference} | Deletes an already existing configuration
[**FullyUpdateAnConfiguration**](IndividualTimeSynchronizationExposureConfigurationApi.md#FullyUpdateAnConfiguration) | **Put** /{afId}/subscriptions/{subscriptionId}/configurations/{instanceReference} | Fully updates/replaces an existing configuration resource



## DeleteAnConfiguration

> DeleteAnConfiguration(ctx, afId, subscriptionId, instanceReference).Execute()

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
    subscriptionId := "subscriptionId_example" // string | Identifier of the subscription resource
    instanceReference := "instanceReference_example" // string | Identifier of the configuration resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualTimeSynchronizationExposureConfigurationApi.DeleteAnConfiguration(context.Background(), afId, subscriptionId, instanceReference).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualTimeSynchronizationExposureConfigurationApi.DeleteAnConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**afId** | **string** | Identifier of the AF | 
**subscriptionId** | **string** | Identifier of the subscription resource | 
**instanceReference** | **string** | Identifier of the configuration resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAnConfigurationRequest struct via the builder pattern


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


## FullyUpdateAnConfiguration

> TimeSyncExposureConfig FullyUpdateAnConfiguration(ctx, afId, subscriptionId, instanceReference).TimeSyncExposureConfig(timeSyncExposureConfig).Execute()

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
    subscriptionId := "subscriptionId_example" // string | Identifier of the subscription resource
    instanceReference := "instanceReference_example" // string | Identifier of the configuration resource
    timeSyncExposureConfig := *openapiclient.NewTimeSyncExposureConfig(int32(123), *openapiclient.NewPtpInstance(*openapiclient.NewInstanceType(), *openapiclient.NewProtocol(), "PtpProfile_example"), int32(123), "ConfigNotifId_example", "ConfigNotifUri_example") // TimeSyncExposureConfig | Parameters to update/replace the existing configuration

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualTimeSynchronizationExposureConfigurationApi.FullyUpdateAnConfiguration(context.Background(), afId, subscriptionId, instanceReference).TimeSyncExposureConfig(timeSyncExposureConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualTimeSynchronizationExposureConfigurationApi.FullyUpdateAnConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FullyUpdateAnConfiguration`: TimeSyncExposureConfig
    fmt.Fprintf(os.Stdout, "Response from `IndividualTimeSynchronizationExposureConfigurationApi.FullyUpdateAnConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**afId** | **string** | Identifier of the AF | 
**subscriptionId** | **string** | Identifier of the subscription resource | 
**instanceReference** | **string** | Identifier of the configuration resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiFullyUpdateAnConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **timeSyncExposureConfig** | [**TimeSyncExposureConfig**](TimeSyncExposureConfig.md) | Parameters to update/replace the existing configuration | 

### Return type

[**TimeSyncExposureConfig**](TimeSyncExposureConfig.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

